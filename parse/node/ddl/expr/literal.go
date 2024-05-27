package expr

import (
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

type LiteralKind int

const (
	LiteralKindUnspecified LiteralKind = iota
	LiteralKindBool
	LiteralKindInt64
	LiteralKindFloat64
	LiteralKindString
	LiteralKindBytes
)

type LiteralNode interface {
	node.Node
	LiteralKind() LiteralKind
	// StringValue returns a string value.
	// For LiteralKindBool, `TRUE` or `FALSE` are returned.
	// For LiteralKindString, a string in form of `"<string value>"` is returned, in which special characters are escaped.
	// For LiteralKindBytes, a string in form of `B"<string value>"` is returned, in which special characters are escaped.
	Value() string
	// UnquotedStringValue returns an unquoted and unescaped string value of this string literal.
	// This function is valid only if LiteralKind() returns one of LiteralKindString or LiteralKindBytes.
	UnquotedStringValue() string
}

type literalNode struct {
	node.NodeBase
	literalKind LiteralKind
	value       string
}

var _ LiteralNode = (*literalNode)(nil)

func AcceptLiteral(s *parse.ParseState, kind LiteralKind, value string) *literalNode {
	return &literalNode{
		NodeBase:    node.NewNodeBase(s.Begin(), s.End()),
		literalKind: kind,
		value:       value,
	}
}
func AcceptBoolLiteral(s *parse.ParseState, value string) *literalNode {
	return AcceptLiteral(s, LiteralKindBool, value)
}

func AcceptFloat64Literal(s *parse.ParseState, value string) *literalNode {
	return AcceptLiteral(s, LiteralKindFloat64, value)
}

func AcceptInt64Literal(s *parse.ParseState, value string) *literalNode {
	return AcceptLiteral(s, LiteralKindInt64, value)
}

func AcceptStringLiteral(s *parse.ParseState, value string) *literalNode {
	return AcceptLiteral(s, LiteralKindString, value)
}

func AcceptBytesLiteral(s *parse.ParseState, value string) *literalNode {
	return AcceptLiteral(s, LiteralKindBytes, value)
}

func (n *literalNode) Children() []node.Node {
	return nil
}

func (n *literalNode) LiteralKind() LiteralKind {
	return n.literalKind
}

func (n *literalNode) Value() string {
	return n.value
}

func (n *literalNode) UnquotedStringValue() string {
	switch n.LiteralKind() {
	case LiteralKindBytes, LiteralKindString:
		v, _ := unquote(n.value)
		return v
	default:
		return n.Value()
	}
}
