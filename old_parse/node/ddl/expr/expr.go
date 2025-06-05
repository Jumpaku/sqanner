package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type ExprKind int

const (
	ExprKindUnspecified ExprKind = iota
	ExprKindOperand
	ExprKindOperation
)

type ExprNode interface {
	node.Node
	ExprKind() ExprKind
	Operand() OperandNode
	Operation() OperationNode
}
