// Code generated by "stringer -type=BrushState -output state_string.go -linecomment"; DO NOT EDIT.

package goralb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StateUnknown-0]
	_ = x[StateIdle-2]
	_ = x[StateRun-3]
	_ = x[StateCharge-4]
}

const (
	_BrushState_name_0 = "Unknown"
	_BrushState_name_1 = "IdleRunCharge"
)

var (
	_BrushState_index_1 = [...]uint8{0, 4, 7, 13}
)

func (i BrushState) String() string {
	switch {
	case i == 0:
		return _BrushState_name_0
	case 2 <= i && i <= 4:
		i -= 2
		return _BrushState_name_1[_BrushState_index_1[i]:_BrushState_index_1[i+1]]
	default:
		return "BrushState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}