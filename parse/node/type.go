package node

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
	ScalarName() string
	ScalarSized() bool
	ScalarSize() TypeSizeNode
	IsArray() bool
	ArrayElement() TypeNode
	IsStruct() bool
	StructFields() []StructTypeFieldNode
}

func ArrayType(element TypeNode) NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase:     nodeBase{kind: NodeType, begin: begin, end: end},
			code:         TypeArray,
			arrayElement: element,
		}
	}
}

func StructType(fields []StructTypeFieldNode) NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase:     nodeBase{kind: NodeType, begin: begin, end: end},
			code:         TypeStruct,
			structFields: fields,
		}
	}
}

func BoolType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeBool,
		}
	}
}

func BytesTypeSized(size TypeSizeNode) NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeBytes,
			sized:    true,
			size:     size,
		}
	}
}

func BytesType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeBytes,
		}
	}
}

func DateType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeDate,
		}
	}
}

func JSONType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeJSON,
		}
	}
}

func Int64Type() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeInt64,
		}
	}
}

func NumericType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeNumeric,
		}
	}
}

func Float64Type() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeFloat64,
		}
	}
}

func StringTypeSized(size TypeSizeNode) NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeString,
			sized:    true,
			size:     size,
		}
	}
}

func StringType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeString,
		}
	}
}

func TimestampType() NewNodeFunc[TypeNode] {
	return func(begin int, end int) TypeNode {
		return anyType{
			nodeBase: nodeBase{kind: NodeType, begin: begin, end: end},
			code:     TypeTimestamp,
		}
	}
}

type anyType struct {
	nodeBase
	code         TypeCode
	size         TypeSizeNode
	sized        bool
	arrayElement TypeNode
	structFields []StructTypeFieldNode
}

var _ TypeNode = anyType{}

func (n anyType) Children() []Node {
	switch n.TypeCode() {
	default:
		return nil
	case TypeBytes, TypeString:
		if n.sized {
			return []Node{n.size}
		}
		return nil
	case TypeArray:
		return []Node{n.arrayElement}
	case TypeStruct:
		children := []Node{}
		for _, structField := range n.structFields {
			children = append(children, structField)
		}
		return children
	}
}

func (n anyType) IsScalar() bool {
	return n.code != TypeArray && n.code != TypeStruct
}

func (n anyType) ScalarSized() bool {
	return n.sized
}

func (n anyType) ScalarSize() TypeSizeNode {
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

func (n anyType) ArrayElement() TypeNode {
	return n.arrayElement
}

func (n anyType) StructFields() []StructTypeFieldNode {
	return n.structFields
}

func (n anyType) TypeCode() TypeCode {
	return n.code
}
