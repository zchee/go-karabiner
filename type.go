// Copyright 2017 The go-karabiner Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package karabiner

type Karabiner struct {
	Global   Global    `json:"global"`
	Profiles []Profile `json:"profiles"`
}

type Global struct {
	CheckForUpdatesOnStartup bool `json:"check_for_updates_on_startup"`
	ShowInMenuBar            bool `json:"show_in_menu_bar"`
	ShowProfileNameInMenuBar bool `json:"show_profile_name_in_menu_bar"`
}

type Profile struct {
	ComplexModifications ComplexModifications `json:"complex_modifications,omitempty"`
	Devices              []Device             `json:"devices"`
	FnFunctionKeys       FnFunctionKeys       `json:"fn_function_keys"`
	Name                 string               `json:"name"`
	Selected             bool                 `json:"selected"`
	SimpleModifications  map[KeyCode]KeyCode  `json:"simple_modifications,omitempty"`
	VirtualHidKeyboard   VirtualHidKeyboard   `json:"virtual_hid_keyboard"`
}

type ComplexModifications struct {
	Parameters Parameters `json:"parameters"`
	Rules      []Rule     `json:"rules"`
}

type Parameters struct {
	BasicToIfAloneTimeoutMilliseconds int `json:"basic.to_if_alone_timeout_milliseconds"`
}

type Rule struct {
	Description  string        `json:"description"`
	Manipulators []Manipulator `json:"manipulators"`
}

type Manipulator struct {
	From      From        `json:"from"`
	To        []To        `json:"to"`
	ToIfAlone []ToIfAlone `json:"to_if_alone,omitempty"`
	Type      string      `json:"type"`
}

type From struct {
	KeyCode   string    `json:"key_code"`
	Modifiers Modifiers `json:"modifiers,omitempty"`
}

type Modifiers struct {
	Optional []string `json:"optional"`
}

type To struct {
	KeyCode string `json:"key_code"`
}

type ToIfAlone struct {
	KeyCode string `json:"key_code"`
}

type Device struct {
	DisableBuiltInKeyboardIfExists bool                `json:"disable_built_in_keyboard_if_exists"`
	FnFunctionKeys                 FnFunctionKeys      `json:"fn_function_keys"`
	Identifiers                    Identifiers         `json:"identifiers"`
	Ignore                         bool                `json:"ignore"`
	SimpleModifications            map[KeyCode]KeyCode `json:"simple_modifications,omitempty"`
}

type Identifiers struct {
	IsKeyboard       bool `json:"is_keyboard"`
	IsPointingDevice bool `json:"is_pointing_device"`
	ProductID        int  `json:"product_id"`
	VendorID         int  `json:"vendor_id"`
}

type FnFunctionKeys struct {
	F1  string `json:"f1"`
	F2  string `json:"f2"`
	F3  string `json:"f3"`
	F4  string `json:"f4"`
	F5  string `json:"f5"`
	F6  string `json:"f6"`
	F7  string `json:"f7"`
	F8  string `json:"f8"`
	F9  string `json:"f9"`
	F10 string `json:"f10"`
	F11 string `json:"f11"`
	F12 string `json:"f12"`
}

type VirtualHidKeyboard struct {
	CapsLockDelayMilliseconds int          `json:"caps_lock_delay_milliseconds"`
	KeyboardType              KeyboardType `json:"keyboard_type"`
}

type KeyboardType string

const (
	NONE KeyboardType = "none"
	ANSI KeyboardType = "ansi"
	ISO  KeyboardType = "iso"
	JIS  KeyboardType = "jis"
)

// KeyCode represents a karabiner keycode.
// This conevert table based by:
//  github.com/tekezo/Karabiner-Elements/src/share/types.hpp@8068628
type KeyCode string

const (
	A                                      KeyCode = "a"                                           // key_code(kHIDUsage_KeyboardA)
	B                                      KeyCode = "b"                                           // key_code(kHIDUsage_KeyboardB)
	C                                      KeyCode = "c"                                           // key_code(kHIDUsage_KeyboardC)
	D                                      KeyCode = "d"                                           // key_code(kHIDUsage_KeyboardD)
	E                                      KeyCode = "e"                                           // key_code(kHIDUsage_KeyboardE)
	F                                      KeyCode = "f"                                           // key_code(kHIDUsage_KeyboardF)
	G                                      KeyCode = "g"                                           // key_code(kHIDUsage_KeyboardG)
	H                                      KeyCode = "h"                                           // key_code(kHIDUsage_KeyboardH)
	I                                      KeyCode = "i"                                           // key_code(kHIDUsage_KeyboardI)
	J                                      KeyCode = "j"                                           // key_code(kHIDUsage_KeyboardJ)
	K                                      KeyCode = "k"                                           // key_code(kHIDUsage_KeyboardK)
	L                                      KeyCode = "l"                                           // key_code(kHIDUsage_KeyboardL)
	M                                      KeyCode = "m"                                           // key_code(kHIDUsage_KeyboardM)
	N                                      KeyCode = "n"                                           // key_code(kHIDUsage_KeyboardN)
	O                                      KeyCode = "o"                                           // key_code(kHIDUsage_KeyboardO)
	P                                      KeyCode = "p"                                           // key_code(kHIDUsage_KeyboardP)
	Q                                      KeyCode = "q"                                           // key_code(kHIDUsage_KeyboardQ)
	R                                      KeyCode = "r"                                           // key_code(kHIDUsage_KeyboardR)
	S                                      KeyCode = "s"                                           // key_code(kHIDUsage_KeyboardS)
	T                                      KeyCode = "t"                                           // key_code(kHIDUsage_KeyboardT)
	U                                      KeyCode = "u"                                           // key_code(kHIDUsage_KeyboardU)
	V                                      KeyCode = "v"                                           // key_code(kHIDUsage_KeyboardV)
	W                                      KeyCode = "w"                                           // key_code(kHIDUsage_KeyboardW)
	X                                      KeyCode = "x"                                           // key_code(kHIDUsage_KeyboardX)
	Y                                      KeyCode = "y"                                           // key_code(kHIDUsage_KeyboardY)
	Z                                      KeyCode = "z"                                           // key_code(kHIDUsage_KeyboardZ)
	One                                    KeyCode = "1"                                           // key_code(kHIDUsage_Keyboard1)
	Two                                    KeyCode = "2"                                           // key_code(kHIDUsage_Keyboard2)
	Three                                  KeyCode = "3"                                           // key_code(kHIDUsage_Keyboard3)
	Four                                   KeyCode = "4"                                           // key_code(kHIDUsage_Keyboard4)
	Five                                   KeyCode = "5"                                           // key_code(kHIDUsage_Keyboard5)
	Six                                    KeyCode = "6"                                           // key_code(kHIDUsage_Keyboard6)
	Seven                                  KeyCode = "7"                                           // key_code(kHIDUsage_Keyboard7)
	Eight                                  KeyCode = "8"                                           // key_code(kHIDUsage_Keyboard8)
	Nine                                   KeyCode = "9"                                           // key_code(kHIDUsage_Keyboard9)
	Ten                                    KeyCode = "0"                                           // key_code(kHIDUsage_Keyboard0)
	ReturnOrEnter                          KeyCode = "return_or_enter"                             // key_code(kHIDUsage_KeyboardReturnOrEnter)
	Escape                                 KeyCode = "escape"                                      // key_code(kHIDUsage_KeyboardEscape)
	DeleteOrBackspace                      KeyCode = "delete_or_backspace"                         // key_code(kHIDUsage_KeyboardDeleteOrBackspace)
	Tab                                    KeyCode = "tab"                                         // key_code(kHIDUsage_KeyboardTab)
	Spacebar                               KeyCode = "spacebar"                                    // key_code(kHIDUsage_KeyboardSpacebar)
	Hypen                                  KeyCode = "hyphen"                                      // key_code(kHIDUsage_KeyboardHyphen)
	EqualSign                              KeyCode = "equal_sign"                                  // key_code(kHIDUsage_KeyboardEqualSign)
	OpenBracket                            KeyCode = "open_bracket"                                // key_code(kHIDUsage_KeyboardOpenBracket)
	CloseBracket                           KeyCode = "close_bracket"                               // key_code(kHIDUsage_KeyboardCloseBracket)
	Backslash                              KeyCode = "backslash"                                   // key_code(kHIDUsage_KeyboardBackslash)
	NonUSPound                             KeyCode = "non_us_pound"                                // key_code(kHIDUsage_KeyboardNonUSPound)
	Semicolon                              KeyCode = "semicolon"                                   // key_code(kHIDUsage_KeyboardSemicolon)
	Quote                                  KeyCode = "quote"                                       // key_code(kHIDUsage_KeyboardQuote)
	GraveAccentAndTilde                    KeyCode = "grave_accent_and_tilde"                      // key_code(kHIDUsage_KeyboardGraveAccentAndTilde)
	Comma                                  KeyCode = "comma"                                       // key_code(kHIDUsage_KeyboardComma)
	Period                                 KeyCode = "period"                                      // key_code(kHIDUsage_KeyboardPeriod)
	Slash                                  KeyCode = "slash"                                       // key_code(kHIDUsage_KeyboardSlash)
	CapsLock                               KeyCode = "caps_lock"                                   // key_code(kHIDUsage_KeyboardCapsLock)
	F1                                     KeyCode = "f1"                                          // key_code(kHIDUsage_KeyboardF1)
	F2                                     KeyCode = "f2"                                          // key_code(kHIDUsage_KeyboardF2)
	F3                                     KeyCode = "f3"                                          // key_code(kHIDUsage_KeyboardF3)
	F4                                     KeyCode = "f4"                                          // key_code(kHIDUsage_KeyboardF4)
	F5                                     KeyCode = "f5"                                          // key_code(kHIDUsage_KeyboardF5)
	F6                                     KeyCode = "f6"                                          // key_code(kHIDUsage_KeyboardF6)
	F7                                     KeyCode = "f7"                                          // key_code(kHIDUsage_KeyboardF7)
	F8                                     KeyCode = "f8"                                          // key_code(kHIDUsage_KeyboardF8)
	F9                                     KeyCode = "f9"                                          // key_code(kHIDUsage_KeyboardF9)
	F10                                    KeyCode = "f10"                                         // key_code(kHIDUsage_KeyboardF10)
	F11                                    KeyCode = "f11"                                         // key_code(kHIDUsage_KeyboardF11)
	F12                                    KeyCode = "f12"                                         // key_code(kHIDUsage_KeyboardF12)
	PrintScreen                            KeyCode = "print_screen"                                // key_code(kHIDUsage_KeyboardPrintScreen)
	ScrollLock                             KeyCode = "scroll_lock"                                 // key_code(kHIDUsage_KeyboardScrollLock)
	Pause                                  KeyCode = "pause"                                       // key_code(kHIDUsage_KeyboardPause)
	Insert                                 KeyCode = "insert"                                      // key_code(kHIDUsage_KeyboardInsert)
	Home                                   KeyCode = "home"                                        // key_code(kHIDUsage_KeyboardHome)
	PageUp                                 KeyCode = "page_up"                                     // key_code(kHIDUsage_KeyboardPageUp)
	DeleteForward                          KeyCode = "delete_forward"                              // key_code(kHIDUsage_KeyboardDeleteForward)
	End                                    KeyCode = "end"                                         // key_code(kHIDUsage_KeyboardEnd)
	PageDown                               KeyCode = "page_down"                                   // key_code(kHIDUsage_KeyboardPageDown)
	RightArrow                             KeyCode = "right_arrow"                                 // key_code(kHIDUsage_KeyboardRightArrow)
	LeftArrow                              KeyCode = "left_arrow"                                  // key_code(kHIDUsage_KeyboardLeftArrow)
	DownArrow                              KeyCode = "down_arrow"                                  // key_code(kHIDUsage_KeyboardDownArrow)
	UpArrow                                KeyCode = "up_arrow"                                    // key_code(kHIDUsage_KeyboardUpArrow)
	KeypadNumLock                          KeyCode = "keypad_num_lock"                             // key_code(kHIDUsage_KeypadNumLock)
	KeypadSlash                            KeyCode = "keypad_slash"                                // key_code(kHIDUsage_KeypadSlash)
	KeypadAsterisk                         KeyCode = "keypad_asterisk"                             // key_code(kHIDUsage_KeypadAsterisk)
	KeypadHyphen                           KeyCode = "keypad_hyphen"                               // key_code(kHIDUsage_KeypadHyphen)
	KeypadPlus                             KeyCode = "keypad_plus"                                 // key_code(kHIDUsage_KeypadPlus)
	KeypadEnter                            KeyCode = "keypad_enter"                                // key_code(kHIDUsage_KeypadEnter)
	Keypad1                                KeyCode = "keypad_1"                                    // key_code(kHIDUsage_Keypad1)
	Keypad2                                KeyCode = "keypad_2"                                    // key_code(kHIDUsage_Keypad2)
	Keypad3                                KeyCode = "keypad_3"                                    // key_code(kHIDUsage_Keypad3)
	Keypad4                                KeyCode = "keypad_4"                                    // key_code(kHIDUsage_Keypad4)
	Keypad5                                KeyCode = "keypad_5"                                    // key_code(kHIDUsage_Keypad5)
	Keypad6                                KeyCode = "keypad_6"                                    // key_code(kHIDUsage_Keypad6)
	Keypad7                                KeyCode = "keypad_7"                                    // key_code(kHIDUsage_Keypad7)
	Keypad8                                KeyCode = "keypad_8"                                    // key_code(kHIDUsage_Keypad8)
	Keypad9                                KeyCode = "keypad_9"                                    // key_code(kHIDUsage_Keypad9)
	Keypad0                                KeyCode = "keypad_0"                                    // key_code(kHIDUsage_Keypad0)
	KeypadPeriod                           KeyCode = "keypad_period"                               // key_code(kHIDUsage_KeypadPeriod)
	NonUSBackslash                         KeyCode = "non_us_backslash"                            // key_code(kHIDUsage_KeyboardNonUSBackslash)
	Application                            KeyCode = "application"                                 // key_code(kHIDUsage_KeyboardApplication)
	Power                                  KeyCode = "power"                                       // key_code(kHIDUsage_KeyboardPower)
	KeypadEqualSign                        KeyCode = "keypad_equal_sign"                           // key_code(kHIDUsage_KeypadEqualSign)
	F13                                    KeyCode = "f13"                                         // key_code(kHIDUsage_KeyboardF13)
	F14                                    KeyCode = "f14"                                         // key_code(kHIDUsage_KeyboardF14)
	F15                                    KeyCode = "f15"                                         // key_code(kHIDUsage_KeyboardF15)
	F16                                    KeyCode = "f16"                                         // key_code(kHIDUsage_KeyboardF16)
	F17                                    KeyCode = "f17"                                         // key_code(kHIDUsage_KeyboardF17)
	F18                                    KeyCode = "f18"                                         // key_code(kHIDUsage_KeyboardF18)
	F19                                    KeyCode = "f19"                                         // key_code(kHIDUsage_KeyboardF19)
	F20                                    KeyCode = "f20"                                         // key_code(kHIDUsage_KeyboardF20)
	F21                                    KeyCode = "f21"                                         // key_code(kHIDUsage_KeyboardF21)
	F22                                    KeyCode = "f22"                                         // key_code(kHIDUsage_KeyboardF22)
	F23                                    KeyCode = "f23"                                         // key_code(kHIDUsage_KeyboardF23)
	F24                                    KeyCode = "f24"                                         // key_code(kHIDUsage_KeyboardF24)
	Execute                                KeyCode = "execute"                                     // key_code(kHIDUsage_KeyboardExecute)
	Help                                   KeyCode = "help"                                        // key_code(kHIDUsage_KeyboardHelp)
	Menu                                   KeyCode = "menu"                                        // key_code(kHIDUsage_KeyboardMenu)
	Select                                 KeyCode = "select"                                      // key_code(kHIDUsage_KeyboardSelect)
	Stop                                   KeyCode = "stop"                                        // key_code(kHIDUsage_KeyboardStop)
	Again                                  KeyCode = "again"                                       // key_code(kHIDUsage_KeyboardAgain)
	Undo                                   KeyCode = "undo"                                        // key_code(kHIDUsage_KeyboardUndo)
	Cut                                    KeyCode = "cut"                                         // key_code(kHIDUsage_KeyboardCut)
	Copy                                   KeyCode = "copy"                                        // key_code(kHIDUsage_KeyboardCopy)
	Paste                                  KeyCode = "paste"                                       // key_code(kHIDUsage_KeyboardPaste)
	Find                                   KeyCode = "find"                                        // key_code(kHIDUsage_KeyboardFind)
	Mute                                   KeyCode = "mute"                                        // key_code(kHIDUsage_KeyboardMute)
	VolumeDecrement                        KeyCode = "volume_decrement"                            // key_code(kHIDUsage_KeyboardVolumeDown)
	VolumeIncrement                        KeyCode = "volume_increment"                            // key_code(kHIDUsage_KeyboardVolumeUp)
	LockingCapsLock                        KeyCode = "locking_caps_lock"                           // key_code(kHIDUsage_KeyboardLockingCapsLock)
	LockingNumLock                         KeyCode = "locking_num_lock"                            // key_code(kHIDUsage_KeyboardLockingNumLock)
	LockingScrollLock                      KeyCode = "locking_scroll_lock"                         // key_code(kHIDUsage_KeyboardLockingScrollLock)
	KeypadComma                            KeyCode = "keypad_comma"                                // key_code(kHIDUsage_KeypadComma)
	KeypadEqualSignAS400                   KeyCode = "keypad_equal_sign_as400"                     // key_code(kHIDUsage_KeypadEqualSignAS400)
	International1                         KeyCode = "international1"                              // key_code(kHIDUsage_KeyboardInternational1)
	International2                         KeyCode = "international2"                              // key_code(kHIDUsage_KeyboardInternational2)
	International3                         KeyCode = "international3"                              // key_code(kHIDUsage_KeyboardInternational3)
	International4                         KeyCode = "international4"                              // key_code(kHIDUsage_KeyboardInternational4)
	International5                         KeyCode = "international5"                              // key_code(kHIDUsage_KeyboardInternational5)
	International6                         KeyCode = "international6"                              // key_code(kHIDUsage_KeyboardInternational6)
	International7                         KeyCode = "international7"                              // key_code(kHIDUsage_KeyboardInternational7)
	International8                         KeyCode = "international8"                              // key_code(kHIDUsage_KeyboardInternational8)
	International9                         KeyCode = "international9"                              // key_code(kHIDUsage_KeyboardInternational9)
	Lang1                                  KeyCode = "lang1"                                       // key_code(kHIDUsage_KeyboardLANG1)
	Lang2                                  KeyCode = "lang2"                                       // key_code(kHIDUsage_KeyboardLANG2)
	Lang3                                  KeyCode = "lang3"                                       // key_code(kHIDUsage_KeyboardLANG3)
	Lang4                                  KeyCode = "lang4"                                       // key_code(kHIDUsage_KeyboardLANG4)
	Lang5                                  KeyCode = "lang5"                                       // key_code(kHIDUsage_KeyboardLANG5)
	Lang6                                  KeyCode = "lang6"                                       // key_code(kHIDUsage_KeyboardLANG6)
	Lang7                                  KeyCode = "lang7"                                       // key_code(kHIDUsage_KeyboardLANG7)
	Lang8                                  KeyCode = "lang8"                                       // key_code(kHIDUsage_KeyboardLANG8)
	Lang9                                  KeyCode = "lang9"                                       // key_code(kHIDUsage_KeyboardLANG9)
	AlternateErase                         KeyCode = "alternate_erase"                             // key_code(kHIDUsage_KeyboardAlternateErase)
	SysReqOrAttention                      KeyCode = "sys_req_or_attention"                        // key_code(kHIDUsage_KeyboardSysReqOrAttention)
	Cancel                                 KeyCode = "cancel"                                      // key_code(kHIDUsage_KeyboardCancel)
	Clear                                  KeyCode = "clear"                                       // key_code(kHIDUsage_KeyboardClear)
	Prior                                  KeyCode = "prior"                                       // key_code(kHIDUsage_KeyboardPrior)
	Return                                 KeyCode = "return"                                      // key_code(kHIDUsage_KeyboardReturn)
	Separator                              KeyCode = "separator"                                   // key_code(kHIDUsage_KeyboardSeparator)
	Out                                    KeyCode = "out"                                         // key_code(kHIDUsage_KeyboardOut)
	Oper                                   KeyCode = "oper"                                        // key_code(kHIDUsage_KeyboardOper)
	ClearOrAgain                           KeyCode = "clear_or_again"                              // key_code(kHIDUsage_KeyboardClearOrAgain)
	CrSelOrProps                           KeyCode = "cr_sel_or_props"                             // key_code(kHIDUsage_KeyboardCrSelOrProps)
	ExSel                                  KeyCode = "ex_sel"                                      // key_code(kHIDUsage_KeyboardExSel)
	LeftControl                            KeyCode = "left_control"                                // key_code(kHIDUsage_KeyboardLeftControl)
	LeftShift                              KeyCode = "left_shift"                                  // key_code(kHIDUsage_KeyboardLeftShift)
	LeftAlt                                KeyCode = "left_alt"                                    // key_code(kHIDUsage_KeyboardLeftAlt)
	LeftGUI                                KeyCode = "left_gui"                                    // key_code(kHIDUsage_KeyboardLeftGUI)
	RightControl                           KeyCode = "right_control"                               // key_code(kHIDUsage_KeyboardRightControl)
	RightShift                             KeyCode = "right_shift"                                 // key_code(kHIDUsage_KeyboardRightShift)
	RightAlt                               KeyCode = "right_alt"                                   // key_code(kHIDUsage_KeyboardRightAlt)
	RightGUI                               KeyCode = "right_gui"                                   // key_code(kHIDUsage_KeyboardRightGUI)
	VkNone                                 KeyCode = "vk_none"                                     // key_code::vk_none
	Fn                                     KeyCode = "fn"                                          // key_code::fn
	DisplayBrightnessDecrement             KeyCode = "display_brightness_decrement"                // key_code::display_brightness_decrement
	DisplayBrightnessIncrement             KeyCode = "display_brightness_increment"                // key_code::display_brightness_increment
	MissionControl                         KeyCode = "mission_control"                             // key_code::mission_control
	Launchpad                              KeyCode = "launchpad"                                   // key_code::launchpad
	Dashboard                              KeyCode = "dashboard"                                   // key_code::dashboard
	IlluminationDecrement                  KeyCode = "illumination_decrement"                      // key_code::illumination_decrement
	IlluminationIncrement                  KeyCode = "illumination_increment"                      // key_code::illumination_increment
	Rewind                                 KeyCode = "rewind"                                      // key_code::rewind
	PlayOrPause                            KeyCode = "play_or_pause"                               // key_code::play_or_pause
	Fastforward                            KeyCode = "fastforward"                                 // key_code::fastforward
	Eject                                  KeyCode = "eject"                                       // key_code::eject
	AppleDisplayBrightnessDecrement        KeyCode = "apple_display_brightness_decrement"          // key_code::apple_display_brightness_decrement
	AppleDisplayBrightnessIncrement        KeyCode = "apple_display_brightness_increment"          // key_code::apple_display_brightness_increment
	AppleTopCaseDisplayBrightnessDecrement KeyCode = "apple_top_case_display_brightness_decrement" // key_code::apple_top_case_display_brightness_decrement
	AppleTopCaseDisplayBrightnessIncrement KeyCode = "apple_top_case_display_brightness_increment" // key_code::apple_top_case_display_brightness_increment
	LeftOption                             KeyCode = "left_option"                                 // key_code(kHIDUsage_KeyboardLeftAlt)
	LeftCommand                            KeyCode = "left_command"                                // key_code(kHIDUsage_KeyboardLeftGUI)
	RightOption                            KeyCode = "right_option"                                // key_code(kHIDUsage_KeyboardRightAlt)
	RightCommand                           KeyCode = "right_command"                               // key_code(kHIDUsage_KeyboardRightGUI)
	JapaneseEisuu                          KeyCode = "japanese_eisuu"                              // key_code(kHIDUsage_KeyboardLANG2)
	JapaneseKana                           KeyCode = "japanese_kana"                               // key_code(kHIDUsage_KeyboardLANG1)
	JapanesePCNfer                         KeyCode = "japanese_pc_nfer"                            // key_code(kHIDUsage_KeyboardInternational5)
	JapanesePCXfer                         KeyCode = "japanese_pc_xfer"                            // key_code(kHIDUsage_KeyboardInternational4)
	JapanesePCKatakana                     KeyCode = "japanese_pc_katakana"                        // key_code(kHIDUsage_KeyboardInternational2)
	VkConsumerBrightnessDown               KeyCode = "vk_consumer_brightness_down"                 // key_code::display_brightness_decrement
	VkCcnsumerBrightnessUp                 KeyCode = "vk_consumer_brightness_up"                   // key_code::display_brightness_increment
	VkMissionControl                       KeyCode = "vk_mission_control"                          // key_code::mission_control
	VkLaunchpad                            KeyCode = "vk_launchpad"                                // key_code::launchpad
	VkDashboard                            KeyCode = "vk_dashboard"                                // key_code::dashboard
	VkConsumerIlluminationDown             KeyCode = "vk_consumer_illumination_down"               // key_code::illumination_decrement
	VkConsumerIlluminationUp               KeyCode = "vk_consumer_illumination_up"                 // key_code::illumination_increment
	VkConsumerPrevious                     KeyCode = "vk_consumer_previous"                        // key_code::rewind
	VkConsumerPlay                         KeyCode = "vk_consumer_play"                            // key_code::play_or_pause
	VkConsumerNext                         KeyCode = "vk_consumer_next"                            // key_code::fastforward
	VolumeDown                             KeyCode = "volume_down"                                 // key_code(kHIDUsage_KeyboardVolumeDown)
	VolumeUp                               KeyCode = "volume_up"                                   // key_code(kHIDUsage_KeyboardVolumeUp)
)
