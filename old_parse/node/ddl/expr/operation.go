package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type OperationArity int

const (
	OperationArityUnary OperationArity = iota + 1
	OperationArityBinary
	OperationArityTernary
)

type Operator int

const (
	OperatorUnspecified Operator = iota
	OperatorUnaryIsNull
	OperatorUnaryIsNotNull
	OperatorUnaryIsTrue
	OperatorUnaryIsNotTrue
	OperatorUnaryIsFalse
	OperatorUnaryIsNotFalse
	OperatorUnaryIsUnknown
	OperatorUnaryIsNotUnknown
	OperatorUnaryMinus
	OperatorUnaryPlus
	OperatorUnaryNot

	OperatorBinaryMul
	OperatorBinaryDiv
	OperatorBinaryConcat
	OperatorBinaryAdd
	OperatorBinarySub
	OperatorBinaryBitShl
	OperatorBinaryBitShr
	OperatorBinaryBitAnd
	OperatorBinaryBitXor
	OperatorBinaryBitOr
	OperatorBinaryEQ
	OperatorBinaryLT
	OperatorBinaryGT
	OperatorBinaryLEQ
	OperatorBinaryGEQ
	OperatorBinaryNEQ
	OperatorBinaryLike
	OperatorBinaryNotLike
	OperatorBinaryIn
	OperatorBinaryNotIn
	OperatorBinaryNot
	OperatorBinaryAnd
	OperatorBinaryOr

	OperatorTernaryBetween
)

type OperationNode interface {
	node.Node
	Expr() ExprNode
	Arity() OperationArity
	Operator() Operator
	Operands() []OperandNode
}
