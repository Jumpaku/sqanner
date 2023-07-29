package node

import "github.com/Jumpaku/go-assert"

type NodeKind int

const (
	NodeUnspecified NodeKind = iota

	NodeIdentifier
	NodePath
	NodeStructTypeField
	NodeType
	NodeTypeSize
)

func (k NodeKind) String() string {
	switch k {
	default:
		assert.State(false, `NodeKind=%d`, k)
	case NodeUnspecified:
		return "NodeUnspecified"
	case NodeIdentifier:
		return "NodeIdentifier"
	case NodePath:
		return "NodePath"
	case NodeStructTypeField:
		return "NodeStructTypeField"
	case NodeType:
		return "NodeType"
	case NodeTypeSize:
		return "NodeTypeSize"
	}

	return assert.Unexpected1[string](`unexpected NodeKind value: %d`, k)
}
