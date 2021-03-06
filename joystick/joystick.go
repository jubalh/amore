// The joystick Package handles any joystick or gamepad events on the gl context,
// it can be used for feedback and input.
package joystick

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// vibration is a struct that keeps track of vibration patterns and lengths for
// communication with sdl
type vibration struct {
	Left, Right float32
	Effect      sdl.HapticEffect
	Data        [4]uint16
	ID          int
	Endtime     uint32
}

// Joystick is an instance of any joystick or gamepad that is connected to the
// program.
type Joystick struct {
	id         int
	stick      *sdl.Joystick
	controller *sdl.GameController
	haptic     *sdl.Haptic
	vibration  *vibration
}

// open connects the joystick to the sdl instance, it will return true if successful
// and false if it could not get the joystick.
func (joystick *Joystick) open() bool {
	joystick.close()

	joystick.stick = sdl.JoystickOpen(sdl.JoystickID(joystick.id))

	if joystick.stick != nil && sdl.IsGameController(joystick.id) {
		joystick.controller = sdl.GameControllerOpen(joystick.id)
	}

	return joystick.IsConnected()
}

// IsGamepad gets whether the Joystick is recognized as a gamepad.
func (joystick *Joystick) IsGamepad() bool {
	return joystick.controller != nil
}

// GetID Gets the joystick's unique identifier.
func (joystick *Joystick) GetID() int {
	return joystick.id
}

// GetName the name of the joystick as it is identified on the os.
func (joystick *Joystick) GetName() string {
	// Prefer the Joystick name for consistency.
	name := joystick.stick.Name()
	if name == "" && joystick.controller != nil {
		name = joystick.controller.Name()
	}

	return name
}

// IsConnected gets whether the Joystick is connected.
func (joystick *Joystick) IsConnected() bool {
	return joystick.stick != nil && joystick.stick.GetAttached()
}

// GetGUID gets a stable GUID unique to the type of the physical joystick.
func (joystick *Joystick) GetGUID() string {
	return sdl.JoystickGetGUIDString(joystick.stick.GetGUID())
}

// getHandle gets the sdl instance id for this joystick
func (joystick *Joystick) getHandle() *sdl.Joystick {
	return joystick.stick
}

// close disconnects this joystick from the sdl instance.
func (joystick *Joystick) close() {
	if joystick.controller != nil {
		joystick.controller.Close()
	}
	if joystick.stick != nil {
		joystick.stick.Close()
	}
	joystick.stick = nil
	joystick.controller = nil
}

// IsDown checks if a button on the Joystick is pressed.
func (joystick *Joystick) IsDown(button int) bool {
	if joystick.IsConnected() == false {
		return false
	}
	return joystick.stick.GetButton(button) == 1
}

// GetAxisCount gets the number of axes on the joystick.
func (joystick *Joystick) GetAxisCount() int {
	if joystick.IsConnected() == false {
		return 0
	}
	return joystick.stick.NumAxes()
}

// GetButtonCount gets the number of buttons on the joystick.
func (joystick *Joystick) GetButtonCount() int {
	if joystick.IsConnected() == false {
		return 0
	}
	return joystick.stick.NumButtons()
}

// GetHatCount gets the number of hats on the joystick.
func (joystick *Joystick) GetHatCount() int {
	if joystick.IsConnected() == false {
		return 0
	}
	return joystick.stick.NumHats()
}

// GetAxis gets the direction of an axis. Values are clamped.
func (joystick *Joystick) GetAxis(axisindex int) float32 {
	if joystick.IsConnected() == false || axisindex < 0 || axisindex >= joystick.GetAxisCount() {
		return 0.0
	}

	return clampval(float32(joystick.stick.GetAxis(axisindex)) / 32768.0)
}

// GetAxes gets the direction of each axis. Values are clamped.
func (joystick *Joystick) GetAxes() []float32 {
	count := joystick.GetAxisCount()
	axes := []float32{}

	if joystick.IsConnected() == false || count <= 0 {
		return axes
	}

	for i := 0; i < count; i++ {
		axes = append(axes, clampval(float32(joystick.stick.GetAxis(i))/32768.0))
	}

	return axes
}

// GetHat gets the direction of a hat.
func (joystick *Joystick) GetHat(hatindex int) byte {
	if joystick.IsConnected() == false || hatindex < 0 || hatindex >= joystick.GetHatCount() {
		return 0
	}

	return joystick.stick.GetHat(hatindex)
}

// GetGamepadAxis gets the direction of a virtual gamepad axis. Values are clamped.
func (joystick *Joystick) GetGamepadAxis(axis GameControllerAxis) float32 {
	if joystick.IsConnected() == false || joystick.IsGamepad() == false {
		return 0.0
	}

	value := joystick.controller.GetAxis(sdl.GameControllerAxis(axis))

	return clampval(float32(value) / 32768.0)
}

// IsGamePadDown checks if a virtual gamepad button on the Joystick is pressed.
func (joystick *Joystick) IsGamepadDown(button GameControllerButton) bool {
	if joystick.IsConnected() == false || joystick.IsGamepad() == false {
		return false
	}

	return joystick.controller.GetButton(sdl.GameControllerButton(button)) == 1
}

// IsVibrationSupported gets whether the Joystick supports vibration.
func (joystick *Joystick) IsVibrationSupported() bool {
	if joystick.checkCreateHaptic() == false {
		return false
	}

	features := joystick.haptic.Query()

	if (features & sdl.HAPTIC_LEFTRIGHT) != 0 {
		return true
	}

	// Some gamepad drivers only support left/right motors via a custom effect.
	if joystick.IsGamepad() && (features&sdl.HAPTIC_CUSTOM) != 0 {
		return true
	}

	// Test for simple sine wave support as a last resort.
	if (features & sdl.HAPTIC_SINE) != 0 {
		return true
	}

	return false
}

// checkCreateHaptic will check the controller and the system if it supports
// haptic feedback and will enable it if it is able to do so. It will return
// if it was successful or not
func (joystick *Joystick) checkCreateHaptic() bool {
	if joystick.IsConnected() == false {
		return false
	}

	if sdl.WasInit(sdl.INIT_HAPTIC) == 0 && sdl.InitSubSystem(sdl.INIT_HAPTIC) != nil {
		return false
	}

	if joystick.haptic != nil && sdl.HapticIndex(joystick.haptic) != -1 {
		return true
	}

	if joystick.haptic != nil {
		joystick.haptic.Close()
	}

	joystick.haptic = sdl.HapticOpenFromJoystick(joystick.stick)
	joystick.vibration = &vibration{
		ID:      -1,
		Left:    0.0,
		Right:   0.0,
		Endtime: sdl.HAPTIC_INFINITY,
		Effect:  sdl.HapticEffect{},
	}

	return joystick.haptic != nil
}

// runVibrationEffect will initiate the joysticks haptic vibration and will return
// if it was successful in running the vibration.
func (joystick *Joystick) runVibrationEffect() bool {
	if joystick.vibration.ID != -1 {
		if joystick.haptic.UpdateEffect(joystick.vibration.ID, &joystick.vibration.Effect) == 0 {
			if joystick.haptic.RunEffect(joystick.vibration.ID, 1) == 0 {
				return true
			}
		}

		// If the effect fails to update, we should destroy and re-create it.
		joystick.haptic.DestroyEffect(joystick.vibration.ID)
		joystick.vibration.ID = -1
	}

	joystick.vibration.ID = joystick.haptic.NewEffect(&joystick.vibration.Effect)

	if joystick.vibration.ID != -1 && joystick.haptic.RunEffect(joystick.vibration.ID, 1) == 0 {
		return true
	}

	return false
}

// SetVibration sets the vibration motor speeds on a Joystick with rumble support.
// if passed no arguments it will stop the vibration. If passed one value it will
// set the vibration for both left and right with a duration of INFINITY. If passed
// two parameters it will set left and right intensities with a duration of INFINITY.
// If passed 3 parameters it will set both intensities and the diration as the third
// parameter, all other parameters will be ignored.
func (joystick *Joystick) SetVibration(args ...float32) bool {
	if !joystick.checkCreateHaptic() {
		return false
	}

	// no arguments given means stop the vibration
	if len(args) == 0 {
		return joystick.stopVibration()
	}

	left := float32(math.Min(math.Max(float64(args[0]), 0.0), 1.0))
	right := left

	if len(args) > 1 {
		right = float32(math.Min(math.Max(float64(right), 0.0), 1.0))
	}

	if left == 0.0 && right == 0.0 {
		return joystick.stopVibration()
	}

	length := sdl.HAPTIC_INFINITY
	if len(args) > 2 && args[2] >= 0.0 {
		maxduration := math.MaxUint32 / 1000.0
		length = int(math.Min(float64(args[2]), float64(maxduration)) * 1000)
	}

	success := false
	features := joystick.haptic.Query()
	axes := joystick.haptic.NumAxes()

	if (features & sdl.HAPTIC_LEFTRIGHT) != 0 {
		joystick.vibration.Effect.SetType(sdl.HAPTIC_LEFTRIGHT)
		lr := joystick.vibration.Effect.LeftRight()
		lr.Length = uint32(length)
		lr.LargeMagnitude = uint16(left * math.MaxUint16)
		lr.SmallMagnitude = uint16(right * math.MaxUint16)
		success = joystick.runVibrationEffect()
	}

	// Some gamepad drivers only give support for controlling individual motors
	// through a custom FF effect.
	if !success && joystick.IsGamepad() && (features&sdl.HAPTIC_CUSTOM) != 0 && axes == 2 {
		// NOTE: this may cause issues with drivers which support custom effects
		// but aren't similar to https://github.com/d235j/360Controller .

		// Custom effect data is clamped to 0x7FFF in SDL.
		joystick.vibration.Data[0] = uint16(left * 0x7FFF)
		joystick.vibration.Data[2] = uint16(left * 0x7FFF)
		joystick.vibration.Data[1] = uint16(right * 0x7FFF)
		joystick.vibration.Data[3] = uint16(right * 0x7FFF)

		joystick.vibration.Effect.SetType(sdl.HAPTIC_CUSTOM)
		custom := joystick.vibration.Effect.Custom()
		custom.Length = uint32(length)
		custom.Channels = 2
		custom.Period = 10
		custom.Samples = 2
		custom.Data = &joystick.vibration.Data[0]

		success = joystick.runVibrationEffect()
	}

	//// Fall back to a simple sine wave if all else fails. This only supports a
	//// single strength value.
	if !success && (features&sdl.HAPTIC_SINE) != 0 {
		joystick.vibration.Effect.SetType(sdl.HAPTIC_SINE)
		periodic := joystick.vibration.Effect.Periodic()
		periodic.Length = uint32(length)
		periodic.Period = 10
		strength := math.Max(float64(left), float64(right))
		periodic.Magnitude = int16(strength * 0x7FFF)
		success = joystick.runVibrationEffect()
	}

	if success {
		joystick.vibration.Left = left
		joystick.vibration.Right = right
		if length == sdl.HAPTIC_INFINITY {
			joystick.vibration.Endtime = sdl.HAPTIC_INFINITY
		} else {
			joystick.vibration.Endtime = sdl.GetTicks() + uint32(length)
		}
	} else {
		joystick.vibration.Left = 0.0
		joystick.vibration.Right = 0.0
		joystick.vibration.Endtime = sdl.HAPTIC_INFINITY
	}

	return success
}

// stopVibration will pause the vibration in the controller.
func (joystick *Joystick) stopVibration() bool {
	success := true

	if sdl.WasInit(sdl.INIT_HAPTIC) == 0 && joystick.haptic != nil && sdl.HapticIndex(joystick.haptic) != -1 {
		success = (joystick.haptic.StopEffect(joystick.vibration.ID) == 0)
	}

	if success {
		joystick.vibration.Left = 0.0
		joystick.vibration.Right = 0.0
	}

	return success
}

// GetVibration will return the current intensity of both the left and right motors
// currently vibrating. It will return 0s if they are not vibrating.
func (joystick *Joystick) GetVibration() (float32, float32) {
	if joystick.vibration.Endtime != sdl.HAPTIC_INFINITY {
		// With some drivers, the effect physically stops at the right time, but
		// SDL_HapticGetEffectStatus still thinks it's playing. So we explicitly
		// stop it once it's done, just to be sure.
		if joystick.vibration.Endtime-sdl.GetTicks() <= 0 {
			joystick.stopVibration()
			joystick.vibration.Endtime = sdl.HAPTIC_INFINITY
		}
	}

	// Check if the haptic effect has stopped playing.
	if joystick.haptic == nil || joystick.vibration.ID == -1 || joystick.haptic.GetEffectStatus(joystick.vibration.ID) != 1 {
		joystick.vibration.Left = 0.0
		joystick.vibration.Right = 0.0
	}

	return joystick.vibration.Left, joystick.vibration.Right
}
