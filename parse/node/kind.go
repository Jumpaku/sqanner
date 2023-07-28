package node

type NodeKind int

const (
	NodeUnspecified NodeKind = iota

	NodePath
	NodeIdentifier

	NodeType
	NodeTypeArray
	NodeTypeScalar
	NodeTypeSize
	NodeTypeStruct
	NodeTypeStructField
)
