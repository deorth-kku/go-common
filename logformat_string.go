// Code generated by "stringer -type=LogFormat -linecomment --trimprefix formatEnd"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DefaultFormat-0]
	_ = x[TextFormat-1]
	_ = x[JsonFormat-2]
	_ = x[formatEnd-3]
}

const _LogFormat_name = "DEFAULTTEXTJSON"

var _LogFormat_index = [...]uint8{0, 7, 11, 15, 15}

func (i LogFormat) String() string {
	if i >= LogFormat(len(_LogFormat_index)-1) {
		return "LogFormat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LogFormat_name[_LogFormat_index[i]:_LogFormat_index[i+1]]
}
