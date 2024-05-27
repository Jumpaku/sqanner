package ddl

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
	"github.com/Jumpaku/sqanner/old_parse/node/ddl/expr"
)

type ColumnOption int

const (
	ColumnOptionAllowCommitTimestamp ColumnOption = iota
)

type AllowCommitTimestampValueNode interface {
	node.Node
	IsNull() bool
	Value() bool
}
type ColumnOptionNode interface {
	node.Node
	Option() ColumnOption
	AllowCommitTimestamp() AllowCommitTimestampValueNode
}
type ColumnDefNode interface {
	node.Node
	Name() node.IdentifierNode
	Type() node.TypeNode
	NotNull() bool
	HasDefault() bool
	Default() expr.ExprNode
	Stored() bool
	AsStored() expr.ExprNode
	Options() []ColumnOptionNode
}
