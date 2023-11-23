package ddl

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
	"github.com/Jumpaku/sqanner/old_parse/node/ddl/expr"
)

//go:generate go run "golang.org/x/tools/cmd/stringer" -type ConstraintKind constraint.go
type ConstraintKind int

const (
	ConstraintKindUnspecified ConstraintKind = iota
	ConstraintKindCheck
	ConstraintKindForeignKey
)

type ConstraintNode interface {
	node.Node
	ConstraintKind() ConstraintKind
	Check() CheckConstraintNode
	ForeignKey() ForeignKeyConstraintNode
}
type CheckConstraintNode interface {
	node.Node
	HasName() bool
	Name() node.IdentifierNode
	Condition() expr.ExprNode
}
type ForeignKeyConstraintNode interface {
	node.Node
	HasName() bool
	Name() node.IdentifierNode
	Columns() []node.IdentifierNode
	ReferencedTable() node.IdentifierNode
	ReferencedColumns() []node.IdentifierNode
	OnDelete() OnDeleteAction
}
