// Copyright 2017 The go-karabiner Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import "fmt"

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
	ComplexModifications ComplexModifications `json:"complex_modifications"`
	Devices              []Device             `json:"devices"`
	FnFunctionKeys       FnFunctionKeys       `json:"fn_function_keys"`
	Name                 string               `json:"name"`
	Selected             bool                 `json:"selected"`
	SimpleModifications  SimpleModifications  `json:"simple_modifications"`
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
	ToIfAlone []ToIfAlone `json:"to_if_alone"`
	Type      string      `json:"type"`
}

type From struct {
	KeyCode   string    `json:"key_code"`
	Modifiers Modifiers `json:"modifiers"`
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
	SimpleModifications            SimpleModifications `json:"simple_modifications"`
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

type SimpleModifications struct {
	CapsLock    string `json:"caps_lock"`
	RightOption string `json:"right_option"`
}

type VirtualHidKeyboard struct {
	CapsLockDelayMilliseconds int          `json:"caps_lock_delay_milliseconds"`
	KeyboardType              KeyboardType `json:"keyboard_type"`
}

type KeyboardType int

const (
	NONE KeyboardType = iota
	ANSI
	ISO
	JIS
)

func (k KeyboardType) String() string {
	switch k {
	case ANSI:
		return "ANSI"
	case ISO:
		return "ISO"
	case JIS:
		return "JIS"
	default:
		return fmt.Sprintf("KeyCode(%d)", k)
	}
}
