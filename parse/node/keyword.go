package node

type KeywordNode interface {
	Node
	Value() string
	AsIdentifier() IdentifierNode
}

func Keyword(value string) NewNodeFunc[KeywordNode] {
	return func(begin, end int) KeywordNode {
		return keyword{
			nodeBase: nodeBase{kind: NodeKeyword, begin: begin, end: end},
			value:    value,
		}
	}
}

type keyword struct {
	nodeBase
	value string
}

var (
	_ Node        = keyword{}
	_ KeywordNode = keyword{}
)

func (n keyword) Children() []Node {
	return nil
}

func (n keyword) Value() string {
	return n.value
}

func (n keyword) AsIdentifier() IdentifierNode {
	return Identifier(n.Value())(n.Begin(), n.End())
}
