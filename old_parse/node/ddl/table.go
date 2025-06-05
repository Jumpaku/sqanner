package ddl

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
)

type InterleaveNode interface {
	node.Node
	Parent() node.IdentifierNode
	HasOnDelete() bool
	OnDelete() OnDeleteAction
}
type PrimaryKeyColumnNode interface {
	node.Node
	Name() node.IdentifierNode
	Order() node.Order
}

type RowDeletionPolicyDaysNode interface {
	node.Node
	Value() int64
}
type RowDeletionPolicyNode interface {
	node.Node
	Column() node.IdentifierNode
	Days() RowDeletionPolicyDaysNode
}
type CreateTableNode interface {
	node.Node
	Table() string
	IfNotExists() bool
	Columns() []ColumnDefNode
	PrimaryKey() []PrimaryKeyColumnNode
	HasInterleave() bool
	Interleave() InterleaveNode
	Constraints() []ConstraintNode
	HasRowDeletionPolicy() bool
}
