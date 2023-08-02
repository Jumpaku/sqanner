package node

type StructTypeFieldNode interface {
	Node
	Name() IdentifierNode
	Type() TypeNode
}

func StructTypeField(fieldName IdentifierNode, fieldType TypeNode) NewNodeFunc[StructTypeFieldNode] {
	return func(begin, end int) StructTypeFieldNode {
		return structTypeField{
			nodeBase:  nodeBase{kind: NodeStructTypeField, begin: begin, end: end},
			fieldName: fieldName,
			fieldType: fieldType,
		}
	}
}

type structTypeField struct {
	nodeBase
	fieldName IdentifierNode
	fieldType TypeNode
}

var _ Node = structTypeField{}
var _ StructTypeFieldNode = structTypeField{}

func (n structTypeField) Children() []Node {
	return []Node{n.fieldName, n.fieldType}
}

func (n structTypeField) Name() IdentifierNode {
	return n.fieldName
}

func (n structTypeField) Type() TypeNode {
	return n.fieldType
}
