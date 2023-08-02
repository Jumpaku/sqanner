package node

import (
	"github.com/Jumpaku/go-assert"
	"strings"
)

type KeywordCode int

const (
	KeywordUnspecified KeywordCode = iota
	KeywordAll
	KeywordAnd
	KeywordAny
	KeywordArray
	KeywordAs
	KeywordAsc
	KeywordAssertRowsModified
	KeywordAt
	KeywordBetween
	KeywordBy
	KeywordCase
	KeywordCast
	KeywordCollate
	KeywordContains
	KeywordCreate
	KeywordCross
	KeywordCube
	KeywordCurrent
	KeywordDefault
	KeywordDefine
	KeywordDesc
	KeywordDistinct
	KeywordElse
	KeywordEnd
	KeywordEnum
	KeywordEscape
	KeywordExcept
	KeywordExclude
	KeywordExists
	KeywordExtract
	KeywordFalse
	KeywordFetch
	KeywordFollowing
	KeywordFor
	KeywordFrom
	KeywordFull
	KeywordGroup
	KeywordGrouping
	KeywordGroups
	KeywordHash
	KeywordHaving
	KeywordIf
	KeywordIgnore
	KeywordIn
	KeywordInner
	KeywordIntersect
	KeywordInterval
	KeywordInto
	KeywordIs
	KeywordJoin
	KeywordLateral
	KeywordLeft
	KeywordLike
	KeywordLimit
	KeywordLookup
	KeywordMerge
	KeywordNatural
	KeywordNew
	KeywordNo
	KeywordNot
	KeywordNull
	KeywordNulls
	KeywordOf
	KeywordOn
	KeywordOr
	KeywordOrder
	KeywordOuter
	KeywordOver
	KeywordPartition
	KeywordPreceding
	KeywordProto
	KeywordRange
	KeywordRecursive
	KeywordRespect
	KeywordRight
	KeywordRollup
	KeywordRows
	KeywordSelect
	KeywordSet
	KeywordSome
	KeywordStruct
	KeywordTablesample
	KeywordThen
	KeywordTo
	KeywordTreat
	KeywordTrue
	KeywordUnbound
	KeywordUnion
	KeywordUnnest
	KeywordUsing
	KeywordWhen
	KeywordWhere
	KeywordWindow
	KeywordWith
	KeywordWithin
)

func KeywordCodeOf(keyword string) KeywordCode {
	switch strings.ToUpper(keyword) {
	default:
		assert.Params(false, `unsupported keyword: %s`, keyword)
	case "ALL":
		return KeywordAll
	case "AND":
		return KeywordAnd
	case "ANY":
		return KeywordAny
	case "ARRAY":
		return KeywordArray
	case "AS":
		return KeywordAs
	case "ASC":
		return KeywordAsc
	case "ASSERT_ROWS_MODIFIED":
		return KeywordAssertRowsModified
	case "AT":
		return KeywordAt
	case "BETWEEN":
		return KeywordBetween
	case "BY":
		return KeywordBy
	case "CASE":
		return KeywordCase
	case "CAST":
		return KeywordCast
	case "COLLATE":
		return KeywordCollate
	case "CONTAINS":
		return KeywordContains
	case "CREATE":
		return KeywordCreate
	case "CROSS":
		return KeywordCross
	case "CUBE":
		return KeywordCube
	case "CURRENT":
		return KeywordCurrent
	case "DEFAULT":
		return KeywordDefault
	case "DEFINE":
		return KeywordDefine
	case "DESC":
		return KeywordDesc
	case "DISTINCT":
		return KeywordDistinct
	case "ELSE":
		return KeywordElse
	case "END":
		return KeywordEnd
	case "ENUM":
		return KeywordEnum
	case "ESCAPE":
		return KeywordEscape
	case "EXCEPT":
		return KeywordExcept
	case "EXCLUDE":
		return KeywordExclude
	case "EXISTS":
		return KeywordExists
	case "EXTRACT":
		return KeywordExtract
	case "FALSE":
		return KeywordFalse
	case "FETCH":
		return KeywordFetch
	case "FOLLOWING":
		return KeywordFollowing
	case "FOR":
		return KeywordFor
	case "FROM":
		return KeywordFrom
	case "FULL":
		return KeywordFull
	case "GROUP":
		return KeywordGroup
	case "GROUPING":
		return KeywordGrouping
	case "GROUPS":
		return KeywordGroups
	case "HASH":
		return KeywordHash
	case "HAVING":
		return KeywordHaving
	case "IF":
		return KeywordIf
	case "IGNORE":
		return KeywordIgnore
	case "IN":
		return KeywordIn
	case "INNER":
		return KeywordInner
	case "INTERSECT":
		return KeywordIntersect
	case "INTERVAL":
		return KeywordInterval
	case "INTO":
		return KeywordInto
	case "IS":
		return KeywordIs
	case "JOIN":
		return KeywordJoin
	case "LATERAL":
		return KeywordLateral
	case "LEFT":
		return KeywordLeft
	case "LIKE":
		return KeywordLike
	case "LIMIT":
		return KeywordLimit
	case "LOOKUP":
		return KeywordLookup
	case "MERGE":
		return KeywordMerge
	case "NATURAL":
		return KeywordNatural
	case "NEW":
		return KeywordNew
	case "NO":
		return KeywordNo
	case "NOT":
		return KeywordNot
	case "NULL":
		return KeywordNull
	case "NULLS":
		return KeywordNulls
	case "OF":
		return KeywordOf
	case "ON":
		return KeywordOn
	case "OR":
		return KeywordOr
	case "ORDER":
		return KeywordOrder
	case "OUTER":
		return KeywordOuter
	case "OVER":
		return KeywordOver
	case "PARTITION":
		return KeywordPartition
	case "PRECEDING":
		return KeywordPreceding
	case "PROTO":
		return KeywordProto
	case "RANGE":
		return KeywordRange
	case "RECURSIVE":
		return KeywordRecursive
	case "RESPECT":
		return KeywordRespect
	case "RIGHT":
		return KeywordRight
	case "ROLLUP":
		return KeywordRollup
	case "ROWS":
		return KeywordRows
	case "SELECT":
		return KeywordSelect
	case "SET":
		return KeywordSet
	case "SOME":
		return KeywordSome
	case "STRUCT":
		return KeywordStruct
	case "TABLESAMPLE":
		return KeywordTablesample
	case "THEN":
		return KeywordThen
	case "TO":
		return KeywordTo
	case "TREAT":
		return KeywordTreat
	case "TRUE":
		return KeywordTrue
	case "UNBOUNDED":
		return KeywordUnbound
	case "UNION":
		return KeywordUnion
	case "UNNEST":
		return KeywordUnnest
	case "USING":
		return KeywordUsing
	case "WHEN":
		return KeywordWhen
	case "WHERE":
		return KeywordWhere
	case "WINDOW":
		return KeywordWindow
	case "WITH":
		return KeywordWith
	case "WITHIN":
		return KeywordWithin
	}

	return assert.Unexpected1[KeywordCode](`unexpected keyword value: %s`, keyword)
}
func (k KeywordCode) Keyword() string {
	return strings.ToUpper(k.String()[8:])
}
func (k KeywordCode) String() string {
	switch k {
	default:
		assert.State(false, `KeywordCode=%d`, k)
	case KeywordAll:
		return "Keyword_ALL"
	case KeywordAnd:
		return "Keyword_AND"
	case KeywordAny:
		return "Keyword_ANY"
	case KeywordArray:
		return "Keyword_ARRAY"
	case KeywordAs:
		return "Keyword_AS"
	case KeywordAsc:
		return "Keyword_ASC"
	case KeywordAssertRowsModified:
		return "Keyword_ASSERT_ROWS_MODIFIED"
	case KeywordAt:
		return "Keyword_AT"
	case KeywordBetween:
		return "Keyword_BETWEEN"
	case KeywordBy:
		return "Keyword_BY"
	case KeywordCase:
		return "Keyword_CASE"
	case KeywordCast:
		return "Keyword_CAST"
	case KeywordCollate:
		return "Keyword_COLLATE"
	case KeywordContains:
		return "Keyword_CONTAINS"
	case KeywordCreate:
		return "Keyword_CREATE"
	case KeywordCross:
		return "Keyword_CROSS"
	case KeywordCube:
		return "Keyword_CUBE"
	case KeywordCurrent:
		return "Keyword_CURRENT"
	case KeywordDefault:
		return "Keyword_DEFAULT"
	case KeywordDefine:
		return "Keyword_DEFINE"
	case KeywordDesc:
		return "Keyword_DESC"
	case KeywordDistinct:
		return "Keyword_DISTINCT"
	case KeywordElse:
		return "Keyword_ELSE"
	case KeywordEnd:
		return "Keyword_END"
	case KeywordEnum:
		return "Keyword_ENUM"
	case KeywordEscape:
		return "Keyword_ESCAPE"
	case KeywordExcept:
		return "Keyword_EXCEPT"
	case KeywordExclude:
		return "Keyword_EXCLUDE"
	case KeywordExists:
		return "Keyword_EXISTS"
	case KeywordExtract:
		return "Keyword_EXTRACT"
	case KeywordFalse:
		return "Keyword_FALSE"
	case KeywordFetch:
		return "Keyword_FETCH"
	case KeywordFollowing:
		return "Keyword_FOLLOWING"
	case KeywordFor:
		return "Keyword_FOR"
	case KeywordFrom:
		return "Keyword_FROM"
	case KeywordFull:
		return "Keyword_FULL"
	case KeywordGroup:
		return "Keyword_GROUP"
	case KeywordGrouping:
		return "Keyword_GROUPING"
	case KeywordGroups:
		return "Keyword_GROUPS"
	case KeywordHash:
		return "Keyword_HASH"
	case KeywordHaving:
		return "Keyword_HAVING"
	case KeywordIf:
		return "Keyword_IF"
	case KeywordIgnore:
		return "Keyword_IGNORE"
	case KeywordIn:
		return "Keyword_IN"
	case KeywordInner:
		return "Keyword_INNER"
	case KeywordIntersect:
		return "Keyword_INTERSECT"
	case KeywordInterval:
		return "Keyword_INTERVAL"
	case KeywordInto:
		return "Keyword_INTO"
	case KeywordIs:
		return "Keyword_IS"
	case KeywordJoin:
		return "Keyword_JOIN"
	case KeywordLateral:
		return "Keyword_LATERAL"
	case KeywordLeft:
		return "Keyword_LEFT"
	case KeywordLike:
		return "Keyword_LIKE"
	case KeywordLimit:
		return "Keyword_LIMIT"
	case KeywordLookup:
		return "Keyword_LOOKUP"
	case KeywordMerge:
		return "Keyword_MERGE"
	case KeywordNatural:
		return "Keyword_NATURAL"
	case KeywordNew:
		return "Keyword_NEW"
	case KeywordNo:
		return "Keyword_NO"
	case KeywordNot:
		return "Keyword_NOT"
	case KeywordNull:
		return "Keyword_NULL"
	case KeywordNulls:
		return "Keyword_NULLS"
	case KeywordOf:
		return "Keyword_OF"
	case KeywordOn:
		return "Keyword_ON"
	case KeywordOr:
		return "Keyword_OR"
	case KeywordOrder:
		return "Keyword_ORDER"
	case KeywordOuter:
		return "Keyword_OUTER"
	case KeywordOver:
		return "Keyword_OVER"
	case KeywordPartition:
		return "Keyword_PARTITION"
	case KeywordPreceding:
		return "Keyword_PRECEDING"
	case KeywordProto:
		return "Keyword_PROTO"
	case KeywordRange:
		return "Keyword_RANGE"
	case KeywordRecursive:
		return "Keyword_RECURSIVE"
	case KeywordRespect:
		return "Keyword_RESPECT"
	case KeywordRight:
		return "Keyword_RIGHT"
	case KeywordRollup:
		return "Keyword_ROLLUP"
	case KeywordRows:
		return "Keyword_ROWS"
	case KeywordSelect:
		return "Keyword_SELECT"
	case KeywordSet:
		return "Keyword_SET"
	case KeywordSome:
		return "Keyword_SOME"
	case KeywordStruct:
		return "Keyword_STRUCT"
	case KeywordTablesample:
		return "Keyword_TABLESAMPLE"
	case KeywordThen:
		return "Keyword_THEN"
	case KeywordTo:
		return "Keyword_TO"
	case KeywordTreat:
		return "Keyword_TREAT"
	case KeywordTrue:
		return "Keyword_TRUE"
	case KeywordUnbound:
		return "Keyword_UNBOUNDED"
	case KeywordUnion:
		return "Keyword_UNION"
	case KeywordUnnest:
		return "Keyword_UNNEST"
	case KeywordUsing:
		return "Keyword_USING"
	case KeywordWhen:
		return "Keyword_WHEN"
	case KeywordWhere:
		return "Keyword_WHERE"
	case KeywordWindow:
		return "Keyword_WINDOW"
	case KeywordWith:
		return "Keyword_WITH"
	case KeywordWithin:
		return "Keyword_WITHIN"
	}
	return assert.Unexpected1[string](`unexpected NodeKind value: %d`, k)
}

type KeywordNode interface {
	Node
	KeywordCode() KeywordCode
	AsIdentifier() IdentifierNode
}

func Keyword(value string) NewNodeFunc[KeywordNode] {
	return func(begin, end int) KeywordNode {
		return keyword{
			nodeBase: nodeBase{kind: NodeKeyword, begin: begin, end: end},
			code:     KeywordCodeOf(value),
			value:    value,
		}
	}
}

type keyword struct {
	nodeBase
	value string
	code  KeywordCode
}

var (
	_ Node        = keyword{}
	_ KeywordNode = keyword{}
)

func (n keyword) Children() []Node {
	return nil
}

func (n keyword) KeywordCode() KeywordCode {
	return n.code
}

func (n keyword) AsIdentifier() IdentifierNode {
	return Identifier(n.value)(n.Begin(), n.End())
}
