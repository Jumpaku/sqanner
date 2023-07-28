package node

import (
	"github.com/Jumpaku/sqanner/tokenize"
)

type TypeCode int

const (
	TypeUnspecified TypeCode = iota
	TypeArray
	TypeBool
	TypeBytes
	TypeDate
	TypeJSON
	TypeInt64
	TypeNumeric
	TypeFloat64
	TypeString
	TypeStruct
	TypeTimestamp
)

type TypeNode interface {
	Node
	TypeCode() TypeCode
	IsScalar() bool
	HasSize() bool
	Size() TypeSizeNode
	IsArray() bool
	IsStruct() bool
	ScalarName() string
	ArrayElementType() TypeNode
	StructFields() []StructTypeFieldNode
}

type TypeSizeNode interface {
	TerminalNode
	Max() bool
	Size() int
}

type StructTypeFieldNode interface {
	Node
	Name() IdentifierNode
	Type() TypeNode
}
type StructTypeNode interface {
	Node
	Fields() []StructTypeFieldNode
}

func TypeSize(begin int, token tokenize.Token, size int) TypeSizeNode {

	return typeSize{
		nodeBase: nodeBase{kind: NodeTypeSize, begin: begin, tokens: []tokenize.Token{token}},
		size:     size,
	}
}

func TypeSizeMax(begin int, token tokenize.Token) TypeSizeNode {
	return typeSize{
		nodeBase: nodeBase{kind: NodeTypeSize, begin: begin, tokens: []tokenize.Token{token}},
		max:      true,
	}
}

func ArrayType(begin int, tokens []tokenize.Token, element TypeNode) TypeNode {
	return anyType{
		nodeBase:     nodeBase{kind: NodeTypeArray, begin: begin, tokens: tokens},
		code:         TypeArray,
		arrayElement: element,
	}
}

func BoolType(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeBool,
	}
}

func BytesType(begin int, tokens []tokenize.Token, size TypeSizeNode) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: tokens},
		code:     TypeBytes,
		size:     size,
	}
}

func DateType(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeDate,
	}
}

func JSONType(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeJSON,
	}
}

func Int64Type(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeInt64,
	}
}

func NumericType(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeNumeric,
	}
}

func Float64Type(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeFloat64,
	}
}

func StringType(begin int, tokens []tokenize.Token, size TypeSizeNode) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: tokens},
		code:     TypeString,
		size:     size,
	}
}

func StructType(begin int, tokens []tokenize.Token, fields []StructTypeFieldNode) TypeNode {
	return anyType{
		nodeBase:     nodeBase{kind: NodeTypeStruct, begin: begin, tokens: tokens},
		code:         TypeStruct,
		structFields: fields,
	}
}

func TimestampType(begin int, token tokenize.Token) TypeNode {
	return anyType{
		nodeBase: nodeBase{kind: NodeTypeScalar, begin: begin, tokens: []tokenize.Token{token}},
		code:     TypeTimestamp,
	}
}

func StructTypeField(begin int, tokens []tokenize.Token, fieldName IdentifierNode, fieldType TypeNode) StructTypeFieldNode {
	return structTypeField{
		nodeBase:  nodeBase{kind: NodeTypeStructField, begin: begin, tokens: tokens},
		fieldName: fieldName,
		fieldType: fieldType,
	}
}

type anyType struct {
	nodeBase
	code         TypeCode
	size         TypeSizeNode
	arrayElement TypeNode
	structFields []StructTypeFieldNode
}

func (n anyType) IsScalar() bool {
	return n.code != TypeArray && n.code != TypeStruct
}

func (n anyType) HasSize() bool {
	return n.code == TypeBytes || n.code == TypeString
}

func (n anyType) Size() TypeSizeNode {
	return n.size
}

func (n anyType) IsArray() bool {
	return n.code == TypeArray
}

func (n anyType) IsStruct() bool {
	return n.code == TypeStruct
}

func (n anyType) ScalarName() string {
	switch n.code {
	default:
		return ""
	case TypeBool:
		return "BOOL"
	case TypeBytes:
		return "BYTES"
	case TypeDate:
		return "DATE"
	case TypeJSON:
		return "JSON"
	case TypeInt64:
		return "INT64"
	case TypeNumeric:
		return "NUMERIC"
	case TypeFloat64:
		return "FLOAT64"
	case TypeString:
		return "STRING"
	case TypeTimestamp:
		return "TIMESTAMP"
	}
}

func (n anyType) ArrayElementType() TypeNode {
	return n.arrayElement
}

func (n anyType) StructFields() []StructTypeFieldNode {
	return n.structFields
}

var _ TypeNode = anyType{}

func (n anyType) TypeCode() TypeCode {
	return n.code
}

type typeSize struct {
	nodeBase
	size int
	max  bool
}

func (n typeSize) Token() tokenize.Token {
	return n.Tokens()[0]
}

func (n typeSize) Max() bool {
	return n.max
}

func (n typeSize) Size() int {
	return n.size
}

var _ TypeSizeNode = typeSize{}

type scalarType struct {
	nodeBase
	name string
	size int
	max  bool
}

func (n scalarType) Max() bool {
	return n.max
}

func (n scalarType) Size() int {
	return n.size
}

func (n scalarType) Name() string {
	return n.name
}

func (n scalarType) Token() tokenize.Token {
	return n.tokens[0]
}

var _ Node = scalarType{}
var _ TerminalNode = scalarType{}

type structType struct {
	nodeBase
	fields []StructTypeFieldNode
}

var _ Node = structType{}
var _ StructTypeNode = structType{}

func (n structType) Fields() []StructTypeFieldNode {
	return n.fields
}

type structTypeField struct {
	nodeBase
	fieldName IdentifierNode
	fieldType TypeNode
}

var _ Node = structTypeField{}
var _ StructTypeFieldNode = structTypeField{}

func (n structTypeField) Name() IdentifierNode {
	return n.fieldName
}

func (n structTypeField) Type() TypeNode {
	return n.fieldType
}
