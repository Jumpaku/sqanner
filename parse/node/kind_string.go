// Code generated by "stringer -type Kind kind.go"; DO NOT EDIT.

package node

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KindUnspecified-0]
	_ = x[KindIdentifier-1]
	_ = x[KindKeyword-2]
	_ = x[KindPath-3]
	_ = x[KindStructTypeField-4]
	_ = x[KindType-5]
	_ = x[KindTypeSize-6]
}

const _Kind_name = "KindUnspecifiedKindIdentifierKindKeywordKindPathKindStructTypeFieldKindTypeKindTypeSize"

var _Kind_index = [...]uint8{0, 15, 29, 40, 48, 67, 75, 87}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}