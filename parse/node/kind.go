package node

type NodeKind int

const (
	NodeUnspecified NodeKind = iota

	NodePath
	NodeIdentifier

	NodeType
	NodeTypeSize
	NodeStructField
)
