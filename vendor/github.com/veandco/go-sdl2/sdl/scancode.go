package sdl

// #include "sdl_wrapper.h"
import "C"

const (
	SCANCODE_UNKNOWN = 0

	/**
	 *  \name Usage page 0x07
	 *
	 *  These values are from usage page 0x07 (USB keyboard page).
	 */
	/*@{*/

	SCANCODE_A = C.SDL_SCANCODE_A
	SCANCODE_B = C.SDL_SCANCODE_B
	SCANCODE_C = C.SDL_SCANCODE_C
	SCANCODE_D = C.SDL_SCANCODE_D
	SCANCODE_E = C.SDL_SCANCODE_E
	SCANCODE_F = C.SDL_SCANCODE_F
	SCANCODE_G = C.SDL_SCANCODE_G
	SCANCODE_H = C.SDL_SCANCODE_H
	SCANCODE_I = C.SDL_SCANCODE_I
	SCANCODE_J = C.SDL_SCANCODE_J
	SCANCODE_K = C.SDL_SCANCODE_K
	SCANCODE_L = C.SDL_SCANCODE_L
	SCANCODE_M = C.SDL_SCANCODE_M
	SCANCODE_N = C.SDL_SCANCODE_N
	SCANCODE_O = C.SDL_SCANCODE_O
	SCANCODE_P = C.SDL_SCANCODE_P
	SCANCODE_Q = C.SDL_SCANCODE_Q
	SCANCODE_R = C.SDL_SCANCODE_R
	SCANCODE_S = C.SDL_SCANCODE_S
	SCANCODE_T = C.SDL_SCANCODE_T
	SCANCODE_U = C.SDL_SCANCODE_U
	SCANCODE_V = C.SDL_SCANCODE_V
	SCANCODE_W = C.SDL_SCANCODE_W
	SCANCODE_X = C.SDL_SCANCODE_X
	SCANCODE_Y = C.SDL_SCANCODE_Y
	SCANCODE_Z = C.SDL_SCANCODE_Z

	SCANCODE_1 = C.SDL_SCANCODE_1
	SCANCODE_2 = C.SDL_SCANCODE_2
	SCANCODE_3 = C.SDL_SCANCODE_3
	SCANCODE_4 = C.SDL_SCANCODE_4
	SCANCODE_5 = C.SDL_SCANCODE_5
	SCANCODE_6 = C.SDL_SCANCODE_6
	SCANCODE_7 = C.SDL_SCANCODE_7
	SCANCODE_8 = C.SDL_SCANCODE_8
	SCANCODE_9 = C.SDL_SCANCODE_9
	SCANCODE_0 = C.SDL_SCANCODE_B

	SCANCODE_RETURN    = C.SDL_SCANCODE_RETURN
	SCANCODE_ESCAPE    = C.SDL_SCANCODE_ESCAPE
	SCANCODE_BACKSPACE = C.SDL_SCANCODE_BACKSPACE
	SCANCODE_TAB       = C.SDL_SCANCODE_TAB
	SCANCODE_SPACE     = C.SDL_SCANCODE_SPACE

	SCANCODE_MINUS        = C.SDL_SCANCODE_MINUS
	SCANCODE_EQUALS       = C.SDL_SCANCODE_EQUALS
	SCANCODE_LEFTBRACKET  = C.SDL_SCANCODE_LEFTBRACKET
	SCANCODE_RIGHTBRACKET = C.SDL_SCANCODE_RIGHTBRACKET
	SCANCODE_BACKSLASH    = C.SDL_SCANCODE_BACKSLASH
	SCANCODE_NONUSHASH    = C.SDL_SCANCODE_NONUSHASH
	SCANCODE_SEMICOLON    = C.SDL_SCANCODE_SEMICOLON
	SCANCODE_APOSTROPHE   = C.SDL_SCANCODE_APOSTROPHE
	SCANCODE_GRAVE        = C.SDL_SCANCODE_GRAVE
	SCANCODE_COMMA        = C.SDL_SCANCODE_COMMA
	SCANCODE_PERIOD       = C.SDL_SCANCODE_PERIOD
	SCANCODE_SLASH        = C.SDL_SCANCODE_SLASH
	SCANCODE_CAPSLOCK     = C.SDL_SCANCODE_CAPSLOCK
	SCANCODE_F1           = C.SDL_SCANCODE_F1
	SCANCODE_F2           = C.SDL_SCANCODE_F2
	SCANCODE_F3           = C.SDL_SCANCODE_F3
	SCANCODE_F4           = C.SDL_SCANCODE_F4
	SCANCODE_F5           = C.SDL_SCANCODE_F5
	SCANCODE_F6           = C.SDL_SCANCODE_F6
	SCANCODE_F7           = C.SDL_SCANCODE_F7
	SCANCODE_F8           = C.SDL_SCANCODE_F8
	SCANCODE_F9           = C.SDL_SCANCODE_F9
	SCANCODE_F10          = C.SDL_SCANCODE_F10
	SCANCODE_F11          = C.SDL_SCANCODE_F11
	SCANCODE_F12          = C.SDL_SCANCODE_F12
	SCANCODE_PRINTSCREEN  = C.SDL_SCANCODE_PRINTSCREEN
	SCANCODE_SCROLLLOCK   = C.SDL_SCANCODE_SCROLLLOCK
	SCANCODE_PAUSE        = C.SDL_SCANCODE_PAUSE
	SCANCODE_INSERT       = C.SDL_SCANCODE_INSERT
	SCANCODE_HOME         = C.SDL_SCANCODE_HOME
	SCANCODE_PAGEUP       = C.SDL_SCANCODE_PAGEUP
	SCANCODE_DELETE       = C.SDL_SCANCODE_DELETE
	SCANCODE_END          = C.SDL_SCANCODE_END
	SCANCODE_PAGEDOWN     = C.SDL_SCANCODE_PAGEDOWN
	SCANCODE_RIGHT        = C.SDL_SCANCODE_RIGHT
	SCANCODE_LEFT         = C.SDL_SCANCODE_LEFT
	SCANCODE_DOWN         = C.SDL_SCANCODE_DOWN
	SCANCODE_UP           = C.SDL_SCANCODE_UP

	SCANCODE_NUMLOCKCLEAR = C.SDL_SCANCODE_NUMLOCKCLEAR
	SCANCODE_KP_DIVIDE    = C.SDL_SCANCODE_KP_DIVIDE
	SCANCODE_KP_MULTIPLY  = C.SDL_SCANCODE_KP_MULTIPLY
	SCANCODE_KP_MINUS     = C.SDL_SCANCODE_KP_MINUS
	SCANCODE_KP_PLUS      = C.SDL_SCANCODE_KP_PLUS
	SCANCODE_KP_ENTER     = C.SDL_SCANCODE_KP_ENTER
	SCANCODE_KP_1         = C.SDL_SCANCODE_KP_1
	SCANCODE_KP_2         = C.SDL_SCANCODE_KP_2
	SCANCODE_KP_3         = C.SDL_SCANCODE_KP_3
	SCANCODE_KP_4         = C.SDL_SCANCODE_KP_4
	SCANCODE_KP_5         = C.SDL_SCANCODE_KP_5
	SCANCODE_KP_6         = C.SDL_SCANCODE_KP_6
	SCANCODE_KP_7         = C.SDL_SCANCODE_KP_7
	SCANCODE_KP_8         = C.SDL_SCANCODE_KP_8
	SCANCODE_KP_9         = C.SDL_SCANCODE_KP_9
	SCANCODE_KP_0         = C.SDL_SCANCODE_KP_0
	SCANCODE_KP_PERIOD    = C.SDL_SCANCODE_KP_PERIOD

	SCANCODE_NONUSBACKSLASH = C.SDL_SCANCODE_NONUSBACKSLASH
	SCANCODE_APPLICATION    = C.SDL_SCANCODE_APPLICATION
	SCANCODE_POWER          = C.SDL_SCANCODE_POWER
	SCANCODE_KP_EQUALS      = C.SDL_SCANCODE_KP_EQUALS
	SCANCODE_F13            = C.SDL_SCANCODE_F13
	SCANCODE_F14            = C.SDL_SCANCODE_F14
	SCANCODE_F15            = C.SDL_SCANCODE_F15
	SCANCODE_F16            = C.SDL_SCANCODE_F16
	SCANCODE_F17            = C.SDL_SCANCODE_F17
	SCANCODE_F18            = C.SDL_SCANCODE_F18
	SCANCODE_F19            = C.SDL_SCANCODE_F19
	SCANCODE_F20            = C.SDL_SCANCODE_F20
	SCANCODE_F21            = C.SDL_SCANCODE_F21
	SCANCODE_F22            = C.SDL_SCANCODE_F22
	SCANCODE_F23            = C.SDL_SCANCODE_F23
	SCANCODE_F24            = C.SDL_SCANCODE_F24
	SCANCODE_EXECUTE        = C.SDL_SCANCODE_EXECUTE
	SCANCODE_HELP           = C.SDL_SCANCODE_HELP
	SCANCODE_MENU           = C.SDL_SCANCODE_MENU
	SCANCODE_SELECT         = C.SDL_SCANCODE_SELECT
	SCANCODE_STOP           = C.SDL_SCANCODE_STOP
	SCANCODE_AGAIN          = C.SDL_SCANCODE_AGAIN
	SCANCODE_UNDO           = C.SDL_SCANCODE_UNDO
	SCANCODE_CUT            = C.SDL_SCANCODE_CUT
	SCANCODE_COPY           = C.SDL_SCANCODE_COPY
	SCANCODE_PASTE          = C.SDL_SCANCODE_PASTE
	SCANCODE_FIND           = C.SDL_SCANCODE_FIND
	SCANCODE_MUTE           = C.SDL_SCANCODE_MUTE
	SCANCODE_VOLUMEUP       = C.SDL_SCANCODE_VOLUMEUP
	SCANCODE_VOLUMEDOWN     = C.SDL_SCANCODE_VOLUMEDOWN
	SCANCODE_KP_COMMA       = C.SDL_SCANCODE_KP_COMMA
	SCANCODE_KP_EQUALSAS400 = C.SDL_SCANCODE_KP_EQUALSAS400

	SCANCODE_INTERNATIONAL1 = C.SDL_SCANCODE_INTERNATIONAL1
	SCANCODE_INTERNATIONAL2 = C.SDL_SCANCODE_INTERNATIONAL2
	SCANCODE_INTERNATIONAL3 = C.SDL_SCANCODE_INTERNATIONAL3
	SCANCODE_INTERNATIONAL4 = C.SDL_SCANCODE_INTERNATIONAL4
	SCANCODE_INTERNATIONAL5 = C.SDL_SCANCODE_INTERNATIONAL5
	SCANCODE_INTERNATIONAL6 = C.SDL_SCANCODE_INTERNATIONAL6
	SCANCODE_INTERNATIONAL7 = C.SDL_SCANCODE_INTERNATIONAL7
	SCANCODE_INTERNATIONAL8 = C.SDL_SCANCODE_INTERNATIONAL8
	SCANCODE_INTERNATIONAL9 = C.SDL_SCANCODE_INTERNATIONAL9
	SCANCODE_LANG1          = C.SDL_SCANCODE_LANG1
	SCANCODE_LANG2          = C.SDL_SCANCODE_LANG2
	SCANCODE_LANG3          = C.SDL_SCANCODE_LANG3
	SCANCODE_LANG4          = C.SDL_SCANCODE_LANG4
	SCANCODE_LANG5          = C.SDL_SCANCODE_LANG5
	SCANCODE_LANG6          = C.SDL_SCANCODE_LANG6
	SCANCODE_LANG7          = C.SDL_SCANCODE_LANG7
	SCANCODE_LANG8          = C.SDL_SCANCODE_LANG8
	SCANCODE_LANG9          = C.SDL_SCANCODE_LANG9

	SCANCODE_ALTERASE   = C.SDL_SCANCODE_ALTERASE
	SCANCODE_SYSREQ     = C.SDL_SCANCODE_SYSREQ
	SCANCODE_CANCEL     = C.SDL_SCANCODE_CANCEL
	SCANCODE_CLEAR      = C.SDL_SCANCODE_CLEAR
	SCANCODE_PRIOR      = C.SDL_SCANCODE_PRIOR
	SCANCODE_RETURN2    = C.SDL_SCANCODE_RETURN2
	SCANCODE_SEPARATOR  = C.SDL_SCANCODE_SEPARATOR
	SCANCODE_OUT        = C.SDL_SCANCODE_OUT
	SCANCODE_OPER       = C.SDL_SCANCODE_OPER
	SCANCODE_CLEARAGAIN = C.SDL_SCANCODE_CLEARAGAIN
	SCANCODE_CRSEL      = C.SDL_SCANCODE_CRSEL
	SCANCODE_EXSEL      = C.SDL_SCANCODE_EXSEL

	SCANCODE_KP_00              = C.SDL_SCANCODE_KP_00
	SCANCODE_KP_000             = C.SDL_SCANCODE_KP_000
	SCANCODE_THOUSANDSSEPARATOR = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	SCANCODE_DECIMALSEPARATOR   = C.SDL_SCANCODE_DECIMALSEPARATOR
	SCANCODE_CURRENCYUNIT       = C.SDL_SCANCODE_CURRENCYUNIT
	SCANCODE_CURRENCYSUBUNIT    = C.SDL_SCANCODE_CURRENCYSUBUNIT
	SCANCODE_KP_LEFTPAREN       = C.SDL_SCANCODE_KP_LEFTPAREN
	SCANCODE_KP_RIGHTPAREN      = C.SDL_SCANCODE_KP_RIGHTPAREN
	SCANCODE_KP_LEFTBRACE       = C.SDL_SCANCODE_KP_LEFTBRACE
	SCANCODE_KP_RIGHTBRACE      = C.SDL_SCANCODE_KP_RIGHTBRACE
	SCANCODE_KP_TAB             = C.SDL_SCANCODE_KP_TAB
	SCANCODE_KP_BACKSPACE       = C.SDL_SCANCODE_KP_BACKSPACE
	SCANCODE_KP_A               = C.SDL_SCANCODE_KP_A
	SCANCODE_KP_B               = C.SDL_SCANCODE_KP_B
	SCANCODE_KP_C               = C.SDL_SCANCODE_KP_C
	SCANCODE_KP_D               = C.SDL_SCANCODE_KP_D
	SCANCODE_KP_E               = C.SDL_SCANCODE_KP_E
	SCANCODE_KP_F               = C.SDL_SCANCODE_KP_F
	SCANCODE_KP_XOR             = C.SDL_SCANCODE_KP_XOR
	SCANCODE_KP_POWER           = C.SDL_SCANCODE_KP_POWER
	SCANCODE_KP_PERCENT         = C.SDL_SCANCODE_KP_PERCENT
	SCANCODE_KP_LESS            = C.SDL_SCANCODE_KP_LESS
	SCANCODE_KP_GREATER         = C.SDL_SCANCODE_KP_GREATER
	SCANCODE_KP_AMPERSAND       = C.SDL_SCANCODE_KP_AMPERSAND
	SCANCODE_KP_DBLAMPERSAND    = C.SDL_SCANCODE_KP_DBLAMPERSAND
	SCANCODE_KP_VERTICALBAR     = C.SDL_SCANCODE_KP_VERTICALBAR
	SCANCODE_KP_DBLVERTICALBAR  = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	SCANCODE_KP_COLON           = C.SDL_SCANCODE_KP_COLON
	SCANCODE_KP_HASH            = C.SDL_SCANCODE_KP_HASH
	SCANCODE_KP_SPACE           = C.SDL_SCANCODE_KP_SPACE
	SCANCODE_KP_AT              = C.SDL_SCANCODE_KP_AT
	SCANCODE_KP_EXCLAM          = C.SDL_SCANCODE_KP_EXCLAM
	SCANCODE_KP_MEMSTORE        = C.SDL_SCANCODE_KP_MEMSTORE
	SCANCODE_KP_MEMRECALL       = C.SDL_SCANCODE_KP_MEMRECALL
	SCANCODE_KP_MEMCLEAR        = C.SDL_SCANCODE_KP_MEMCLEAR
	SCANCODE_KP_MEMADD          = C.SDL_SCANCODE_KP_MEMADD
	SCANCODE_KP_MEMSUBTRACT     = C.SDL_SCANCODE_KP_MEMSUBTRACT
	SCANCODE_KP_MEMMULTIPLY     = C.SDL_SCANCODE_KP_MEMMULTIPLY
	SCANCODE_KP_MEMDIVIDE       = C.SDL_SCANCODE_KP_MEMDIVIDE
	SCANCODE_KP_PLUSMINUS       = C.SDL_SCANCODE_KP_PLUSMINUS
	SCANCODE_KP_CLEAR           = C.SDL_SCANCODE_KP_CLEAR
	SCANCODE_KP_CLEARENTRY      = C.SDL_SCANCODE_KP_CLEARENTRY
	SCANCODE_KP_BINARY          = C.SDL_SCANCODE_KP_BINARY
	SCANCODE_KP_OCTAL           = C.SDL_SCANCODE_KP_OCTAL
	SCANCODE_KP_DECIMAL         = C.SDL_SCANCODE_KP_DECIMAL
	SCANCODE_KP_HEXADECIMAL     = C.SDL_SCANCODE_KP_HEXADECIMAL

	SCANCODE_LCTRL          = C.SDL_SCANCODE_LCTRL
	SCANCODE_LSHIFT         = C.SDL_SCANCODE_LSHIFT
	SCANCODE_LALT           = C.SDL_SCANCODE_LALT
	SCANCODE_LGUI           = C.SDL_SCANCODE_LGUI
	SCANCODE_RCTRL          = C.SDL_SCANCODE_RCTRL
	SCANCODE_RSHIFT         = C.SDL_SCANCODE_RSHIFT
	SCANCODE_RALT           = C.SDL_SCANCODE_RALT
	SCANCODE_RGUI           = C.SDL_SCANCODE_RGUI
	SCANCODE_MODE           = C.SDL_SCANCODE_MODE
	SCANCODE_AUDIONEXT      = C.SDL_SCANCODE_AUDIONEXT
	SCANCODE_AUDIOPREV      = C.SDL_SCANCODE_AUDIOPREV
	SCANCODE_AUDIOSTOP      = C.SDL_SCANCODE_AUDIOSTOP
	SCANCODE_AUDIOPLAY      = C.SDL_SCANCODE_AUDIOPLAY
	SCANCODE_AUDIOMUTE      = C.SDL_SCANCODE_AUDIOMUTE
	SCANCODE_MEDIASELECT    = C.SDL_SCANCODE_MEDIASELECT
	SCANCODE_WWW            = C.SDL_SCANCODE_WWW
	SCANCODE_MAIL           = C.SDL_SCANCODE_MAIL
	SCANCODE_CALCULATOR     = C.SDL_SCANCODE_CALCULATOR
	SCANCODE_COMPUTER       = C.SDL_SCANCODE_COMPUTER
	SCANCODE_AC_SEARCH      = C.SDL_SCANCODE_AC_SEARCH
	SCANCODE_AC_HOME        = C.SDL_SCANCODE_AC_HOME
	SCANCODE_AC_BACK        = C.SDL_SCANCODE_AC_BACK
	SCANCODE_AC_FORWARD     = C.SDL_SCANCODE_AC_FORWARD
	SCANCODE_AC_STOP        = C.SDL_SCANCODE_AC_STOP
	SCANCODE_AC_REFRESH     = C.SDL_SCANCODE_AC_REFRESH
	SCANCODE_AC_BOOKMARKS   = C.SDL_SCANCODE_AC_BOOKMARKS
	SCANCODE_BRIGHTNESSDOWN = C.SDL_SCANCODE_BRIGHTNESSDOWN
	SCANCODE_BRIGHTNESSUP   = C.SDL_SCANCODE_BRIGHTNESSUP
	SCANCODE_DISPLAYSWITCH  = C.SDL_SCANCODE_DISPLAYSWITCH
	SCANCODE_KBDILLUMTOGGLE = C.SDL_SCANCODE_KBDILLUMTOGGLE
	SCANCODE_KBDILLUMDOWN   = C.SDL_SCANCODE_KBDILLUMDOWN
	SCANCODE_KBDILLUMUP     = C.SDL_SCANCODE_KBDILLUMUP
	SCANCODE_EJECT          = C.SDL_SCANCODE_EJECT
	SCANCODE_SLEEP          = C.SDL_SCANCODE_SLEEP
	SCANCODE_APP1           = C.SDL_SCANCODE_APP1
	SCANCODE_APP2           = C.SDL_SCANCODE_APP2
	NUM_SCANCODES           = C.SDL_NUM_SCANCODES
)

// Scancode (https://wiki.libsdl.org/SDL_Scancode)
type Scancode uint32

func (code Scancode) c() C.SDL_Scancode {
	return C.SDL_Scancode(code)
}
