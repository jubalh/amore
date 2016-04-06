// The mouse Package handles the mouse events from the gl context
package mouse

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/tanema/amore/window"
)

//Checks whether a certain button is down.
func IsDown(button Button) bool {
	_, _, state := sdl.GetMouseState()

	if (uint32(button) & state) == 1 {
		return true
	}

	return false
}

//Returns the current x-position of the mouse.
func GetX() float32 {
	x, _ := GetPosition()
	return x
}

//Returns the current y-position of the mouse.
func GetY() float32 {
	_, y := GetPosition()
	return y
}

//Sets the current X position of the mouse.
func SetX(x float32) {
	_, y := GetPosition()
	SetPosition(x, y)
}

//Sets the current Y position of the mouse.
func SetY(y float32) {
	x, _ := GetPosition()
	SetPosition(x, y)
}

//Returns the current position of the mouse.
func GetPosition() (float32, float32) {
	return window.GetMousePosition()
}

//Sets the current position of the mouse.
func SetPosition(x, y float32) {
	window.SetMousePosition(x, y)
}

//Gets whether relative mode is enabled for the mouse.
func GetRelativeMode() bool {
	return sdl.GetRelativeMouseMode() != false
}

//	Sets whether relative mode is enabled for the mouse.
func SetRelativeMode(isvisible bool) {
	sdl.SetRelativeMouseMode(isvisible)
}

//Checks if the mouse is grabbed.
func IsGrabbed() bool {
	return window.IsMouseGrabbed()
}

//Grabs the mouse and confines it to the window.
func SetGrabbed(enabled bool) {
	window.SetMouseGrab(enabled)
}

//Checks if the cursor is visible.
func IsVisible() bool {
	return sdl.ShowCursor(sdl.QUERY) == sdl.ENABLE
}

//Sets the current visibility of the cursor.
func SetVisible(isvisible bool) {
	state := sdl.ENABLE
	if isvisible == false {
		state = sdl.DISABLE
	}
	sdl.ShowCursor(state)
}
