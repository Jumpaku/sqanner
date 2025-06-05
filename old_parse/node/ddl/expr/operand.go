package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type OperandKind int

const (
	OperandKindUnspecified OperandKind = iota
	OperandKindEnclosedExpr
	OperandKindPrimitive
	OperandKindCast
	OperandKindConditional
	OperandKindFieldAccess
	OperandKindArraySubscript
	OperandKindJSONSubscript
	OperandKindFunctionCall
)

type CastNode interface {
	node.Node
	Expr() ExprNode
	Type() node.TypeNode
}

type OperandNode interface {
	node.Node
	OperandKind() OperandKind
	EnclosedExpr() ExprNode
	Primitive() PrimitiveNode
	Cast() CastNode
	Conditional() ConditionalNode
	FieldAccess() FieldAccessNode
	ArraySubscript() ArraySubscriptNode
	JSONSubscript() JSONSubscriptNode
	FunctionCall() FunctionCallNode
}
