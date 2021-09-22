// Code generated by "stringer -type=DcdcModet"; DO NOT EDIT.

package dcdcusb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Dumb-0]
	_ = x[Automotive-2]
	_ = x[Script-1]
	_ = x[UPS-3]
	_ = x[Unknown-255]
}

const (
	_DcdcModet_name_0 = "DumbScriptAutomotiveUPS"
	_DcdcModet_name_1 = "Unknown"
)

var (
	_DcdcModet_index_0 = [...]uint8{0, 4, 10, 20, 23}
)

func (i DcdcModet) String() string {
	switch {
	case 0 <= i && i <= 3:
		return _DcdcModet_name_0[_DcdcModet_index_0[i]:_DcdcModet_index_0[i+1]]
	case i == 255:
		return _DcdcModet_name_1
	default:
		return "DcdcModet(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
