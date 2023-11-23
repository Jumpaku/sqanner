package node

type StructTypeFieldNode interface {
	Node
	Named() bool
	Name() IdentifierNode
	Type() TypeNode
}

func StructTypeFieldUnnamed(fieldType TypeNode) NewNodeFunc[StructTypeFieldNode] {
	return func(begin, end int) StructTypeFieldNode {
		return structTypeField{
			NodeBase:  NodeBase{kind: KindStructTypeField, begin: begin, end: end},
			fieldType: fieldType,
		}
	}
}
func StructTypeField(fieldName IdentifierNode, fieldType TypeNode) NewNodeFunc[StructTypeFieldNode] {
	return func(begin, end int) StructTypeFieldNode {
		return structTypeField{
			NodeBase:  NodeBase{kind: KindStructTypeField, begin: begin, end: end},
			fieldName: fieldName,
			fieldType: fieldType,
		}
	}
}

type structTypeField struct {
	NodeBase
	fieldName IdentifierNode
	fieldType TypeNode
}

var _ Node = structTypeField{}
var _ StructTypeFieldNode = structTypeField{}

func (n structTypeField) Children() []Node {
	if n.Named() {
		return []Node{n.fieldName, n.fieldType}
	}
	return []Node{n.fieldType}
}

func (n structTypeField) Name() IdentifierNode {
	return n.fieldName
}

func (n structTypeField) Named() bool {
	return n.fieldName != nil
}

func (n structTypeField) Type() TypeNode {
	return n.fieldType
}
