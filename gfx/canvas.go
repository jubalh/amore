package gfx

import (
	"fmt"
	"image"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/tanema/amore/gfx/gl"
)

// Canvas is an off-screen render target.
type Canvas struct {
	*Texture
	fbo              gl.Framebuffer
	depth_stencil    gl.Renderbuffer
	status           uint32
	attachedCanvases []*Canvas
	width, height    int32
	systemViewport   []int32
}

// NewCanvas creates a pointer to a new canvas with the privided width and height
func NewCanvas(width, height int32) *Canvas {
	new_canvas := &Canvas{
		width:  width,
		height: height,
	}
	registerVolatile(new_canvas)
	return new_canvas
}

// loadVolatile will create the framebuffer and return true if successful
func (canvas *Canvas) loadVolatile() bool {
	canvas.status = gl.FRAMEBUFFER_COMPLETE

	// glTexImage2D is guaranteed to error in this case.
	if canvas.width > maxTextureSize || canvas.height > maxTextureSize {
		canvas.status = gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT
		return false
	}

	canvas.Texture = newTexture(canvas.width, canvas.height, false)
	//NULL means reserve texture memory, but texels are undefined
	gl.TexImage2D(gl.TEXTURE_2D, 0, int(canvas.width), int(canvas.height), gl.RGBA, gl.UNSIGNED_BYTE, nil)
	if gl.GetError() != gl.NO_ERROR {
		canvas.Texture.Release()
		canvas.status = gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT
		return false
	}

	canvas.fbo, canvas.status = newFBO(canvas.getHandle())

	if canvas.status != gl.FRAMEBUFFER_COMPLETE {
		if canvas.fbo.Valid() {
			gl.DeleteFramebuffer(canvas.fbo)
			canvas.fbo = gl.Framebuffer{}
		}
		return false
	}

	return true
}

// unLoadVolatile will release the texture, framebuffer and depth buffer
func (canvas *Canvas) unLoadVolatile() {
	if gl_state.currentCanvas == canvas {
		canvas.stopGrab(false)
	}
	gl.DeleteFramebuffer(canvas.fbo)
	gl.DeleteRenderbuffer(canvas.depth_stencil)

	canvas.fbo = gl.Framebuffer{}
	canvas.depth_stencil = gl.Renderbuffer{}

	canvas.attachedCanvases = []*Canvas{}
	canvas.Texture.Release()
}

// Release will release all the gl objects associates with the canvas and clean
// up the memory
func (canvas *Canvas) Release() {
	releaseVolatile(canvas)
}

// startGrab will bind this canvas to grab all drawing operations
// multiple canvases can only be passed in on non mobile platforms
func (canvas *Canvas) startGrab(canvases ...*Canvas) error {
	if gl_state.currentCanvas == canvas {
		return nil // already grabbing
	}

	if canvases != nil && len(canvases) > 0 {
		// Whether the new canvas list is different from the old one.
		// A more thorough check is done below.
		if maxRenderTargets < 4 {
			return fmt.Errorf("Multi-canvas rendering is not supported on this system.")
		}

		if int32(len(canvases)+1) > maxRenderTargets {
			return fmt.Errorf("This system can't simultaneously render to %v canvases.", len(canvases)+1)
		}

		for i := 0; i < len(canvases); i++ {
			if canvases[i].width != canvas.width || canvases[i].height != canvas.height {
				return fmt.Errorf("All canvases must have the same dimensions.")
			}
		}
	}

	// cleanup after previous Canvas
	if gl_state.currentCanvas != nil {
		canvas.systemViewport = gl_state.currentCanvas.systemViewport
		gl_state.currentCanvas.stopGrab(true)
	} else {
		canvas.systemViewport = GetViewport()
	}

	// indicate we are using this Canvas.
	gl_state.currentCanvas = canvas
	// bind the framebuffer object.
	gl.BindFramebuffer(gl.FRAMEBUFFER, canvas.fbo)
	SetViewport(0, 0, canvas.width, canvas.height)
	// Set up the projection matrix
	gl_state.projectionStack.Push()
	gl_state.projectionStack.Load(mgl32.Ortho(0.0, float32(screen_width), 0.0, float32(screen_height), -1, 1))

	canvas.attacheExtra(canvases)
	canvas.attachedCanvases = canvases

	return nil
}

// stopGrab will bind the context back to the default framebuffer and set back
// all the settings
func (canvas *Canvas) stopGrab(switchingToOtherCanvas bool) error {
	// i am not grabbing. leave me alone
	if gl_state.currentCanvas != canvas {
		return nil
	}
	gl_state.projectionStack.Pop()
	if !switchingToOtherCanvas {
		// bind system framebuffer.
		gl_state.currentCanvas = nil
		gl.BindFramebuffer(gl.FRAMEBUFFER, getDefaultFBO())
		SetViewport(canvas.systemViewport[0], canvas.systemViewport[1], canvas.systemViewport[2], canvas.systemViewport[3])
	}
	return nil
}

// NewImageData will create an image from the canvas data. It will return an error
// only if the dimensions given are invalid
func (canvas *Canvas) NewImageData(x, y, w, h int32) (image.Image, error) {
	if x < 0 || y < 0 || w <= 0 || h <= 0 || (x+w) > canvas.width || (y+h) > canvas.height {
		return nil, fmt.Errorf("Invalid ImageData rectangle dimensions.")
	}

	prev_canvas := GetCanvas()
	SetCanvas(canvas)

	screenshot := image.NewRGBA(image.Rect(int(x), int(y), int(w), int(h)))
	stride := int32(screenshot.Stride)
	pixels := make([]byte, len(screenshot.Pix))
	gl.ReadPixels(pixels, int(x), int(y), int(w), int(h), gl.RGBA, gl.UNSIGNED_BYTE)

	for y := int32(0); y < h; y++ {
		i := (h - 1 - y) * stride
		copy(screenshot.Pix[y*stride:], pixels[i:i+w*4])
	}

	SetCanvas(prev_canvas...)

	// The new ImageData now owns the pixel data, so we don't delete it here.
	return screenshot, nil
}

// checkCreateStencil if a stencil is set on a canvas then we need to create
// some buffers to handle this.
func (canvas *Canvas) checkCreateStencil() bool {
	// Do nothing if we've already created the stencil buffer.
	if canvas.depth_stencil.Valid() {
		return true
	}

	if gl_state.currentCanvas != canvas {
		gl.BindFramebuffer(gl.FRAMEBUFFER, canvas.fbo)
	}

	format := gl.STENCIL_INDEX8
	attachment := gl.STENCIL_ATTACHMENT

	canvas.depth_stencil = gl.CreateRenderbuffer()
	gl.BindRenderbuffer(gl.RENDERBUFFER, canvas.depth_stencil)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.Enum(format), int(canvas.width), int(canvas.height))

	// Attach the stencil buffer to the framebuffer object.
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.Enum(attachment), gl.RENDERBUFFER, canvas.depth_stencil)
	gl.BindRenderbuffer(gl.RENDERBUFFER, gl.Renderbuffer{})

	success := (gl.CheckFramebufferStatus(gl.FRAMEBUFFER) == gl.FRAMEBUFFER_COMPLETE)

	// We don't want the stencil buffer filled with garbage.
	if success {
		gl.Clear(gl.STENCIL_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	} else {
		gl.DeleteRenderbuffer(canvas.depth_stencil)
		canvas.depth_stencil = gl.Renderbuffer{}
	}

	if gl_state.currentCanvas != nil && gl_state.currentCanvas != canvas {
		gl.BindFramebuffer(gl.FRAMEBUFFER, gl_state.currentCanvas.fbo)
	} else if gl_state.currentCanvas == nil {
		gl.BindFramebuffer(gl.FRAMEBUFFER, getDefaultFBO())
	}

	return success
}

// newFBO will generate a new Frame Buffer Object for use with the canvas
func newFBO(texture gl.Texture) (gl.Framebuffer, uint32) {
	// get currently bound fbo to reset to it later
	current_fbo := gl.GetBoundFramebuffer()

	framebuffer := gl.CreateFramebuffer()
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
	if texture.Valid() {
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)
		// Initialize the texture to transparent black.
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	}
	status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)

	// unbind framebuffer
	gl.BindFramebuffer(gl.FRAMEBUFFER, current_fbo)

	return framebuffer, uint32(status)
}
