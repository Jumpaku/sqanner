package parse

//go:generate go run github.com/dmarkham/enumer -type=Keyword
type Keyword int

const (
	KeywordUnspecified Keyword = iota
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

func IsAKeyword(s string) bool {
	_, err := KeywordString(s)
	return err == nil
}
