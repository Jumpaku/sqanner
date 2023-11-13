package node

import (
	"github.com/Jumpaku/go-assert"
	"strings"
)

//go:generate go run "golang.org/x/tools/cmd/stringer" -type KeywordCode keyword.go
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

type KeywordNode interface {
	Node
	KeywordCode() KeywordCode
}

func Keyword(code KeywordCode) NewNodeFunc[KeywordNode] {
	return func(begin, end int) KeywordNode {
		return keyword{
			nodeBase: nodeBase{kind: KindKeyword, begin: begin, end: end},
			code:     code,
		}
	}
}

type keyword struct {
	nodeBase
	code KeywordCode
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
