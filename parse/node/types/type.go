package types

import (
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

type TypeNode interface {
	node.Node
	TypeCode() TypeCode
	IsScalar() bool
	ScalarName() string
	ScalarSized() bool
	ScalarSize() TypeSizeNode
	IsArray() bool
	ArrayElement() TypeNode
	IsStruct() bool
	StructFields() []StructFieldNode
}

func AcceptArray(s *parse.ParseState, element TypeNode) TypeNode {
	return &typeNode{
		NodeBase:     node.NewNodeBase(s.Begin(), s.End()),
		code:         TypeCodeArray,
		arrayElement: element,
	}
}

func AcceptStruct(s *parse.ParseState, fields []StructFieldNode) TypeNode {
	return &typeNode{
		NodeBase:     node.NewNodeBase(s.Begin(), s.End()),
		code:         TypeCodeStruct,
		structFields: fields,
	}
}

func AcceptBool(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeBool,
	}
}

func AcceptBytesSized(s *parse.ParseState, size TypeSizeNode) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeBytes,
		sized:    true,
		size:     size,
	}
}

func AcceptBytes(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeBytes,
	}
}

func AcceptDate(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeDate,
	}
}

func AcceptJSON(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeJSON,
	}
}

func AcceptInt64(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeInt64,
	}
}

func AcceptNumeric(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeNumeric,
	}
}

func AcceptFloat64(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeFloat64,
	}
}

func AcceptStringSized(s *parse.ParseState, size TypeSizeNode) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeString,
		sized:    true,
		size:     size,
	}
}

func AcceptString(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeString,
	}
}

func AcceptTimestamp(s *parse.ParseState) TypeNode {
	return &typeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		code:     TypeCodeTimestamp,
	}
}

func NewArray(element TypeNode) TypeNode {
	return &typeNode{code: TypeCodeArray, arrayElement: element}
}

func NewStruct(fields []StructFieldNode) TypeNode {
	return &typeNode{code: TypeCodeStruct, structFields: fields}
}

func NewBool() TypeNode {
	return &typeNode{code: TypeCodeBool}
}

func NewBytesSized(size TypeSizeNode) TypeNode {
	return &typeNode{
		code:  TypeCodeBytes,
		sized: true,
		size:  size,
	}
}

func NewBytes() TypeNode {
	return &typeNode{code: TypeCodeBytes}
}

func NewDate() TypeNode {
	return &typeNode{code: TypeCodeDate}
}

func NewJSON() TypeNode {
	return &typeNode{code: TypeCodeJSON}
}

func NewInt64() TypeNode {
	return &typeNode{code: TypeCodeInt64}
}

func NewNumeric() TypeNode {
	return &typeNode{code: TypeCodeNumeric}
}

func NewFloat64() TypeNode {
	return &typeNode{code: TypeCodeFloat64}
}

func NewStringSized(size TypeSizeNode) TypeNode {
	return &typeNode{
		code:  TypeCodeString,
		sized: true,
		size:  size,
	}
}

func NewString() TypeNode {
	return &typeNode{code: TypeCodeString}
}

func NewTimestamp() TypeNode {
	return &typeNode{code: TypeCodeTimestamp}
}

type typeNode struct {
	node.NodeBase
	code         TypeCode
	size         TypeSizeNode
	sized        bool
	arrayElement TypeNode
	structFields []StructFieldNode
}

var _ TypeNode = (*typeNode)(nil)

func (n *typeNode) Children() []node.Node {
	switch n.TypeCode() {
	default:
		return nil
	case TypeCodeBytes, TypeCodeString:
		if n.sized {
			return []node.Node{n.size}
		}
		return nil
	case TypeCodeArray:
		return []node.Node{n.arrayElement}
	case TypeCodeStruct:
		children := []node.Node{}
		for _, structField := range n.structFields {
			children = append(children, structField)
		}
		return children
	}
}

func (n *typeNode) TypeCode() TypeCode {
	return n.code
}

func (n *typeNode) IsScalar() bool {
	return n.code != TypeCodeArray && n.code != TypeCodeStruct
}

func (n *typeNode) ScalarName() string {
	switch n.code {
	default:
		return ""
	case TypeCodeBool:
		return "BOOL"
	case TypeCodeBytes:
		return "BYTES"
	case TypeCodeDate:
		return "DATE"
	case TypeCodeJSON:
		return "JSON"
	case TypeCodeInt64:
		return "INT64"
	case TypeCodeNumeric:
		return "NUMERIC"
	case TypeCodeFloat64:
		return "FLOAT64"
	case TypeCodeString:
		return "STRING"
	case TypeCodeTimestamp:
		return "TIMESTAMP"
	}
}

func (n *typeNode) ScalarSized() bool {
	return n.sized
}

func (n *typeNode) ScalarSize() TypeSizeNode {
	return n.size
}

func (n *typeNode) IsArray() bool {
	return n.code == TypeCodeArray
}

func (n *typeNode) ArrayElement() TypeNode {
	return n.arrayElement
}

func (n *typeNode) IsStruct() bool {
	return n.code == TypeCodeStruct
}

func (n *typeNode) StructFields() []StructFieldNode {
	return n.structFields
}
