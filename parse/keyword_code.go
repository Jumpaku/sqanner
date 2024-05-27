package parse

import (
	"fmt"
	"strings"
)

//go:generate go run "golang.org/x/tools/cmd/stringer" -type KeywordCode keyword_code.go
type KeywordCode int

const (
	KeywordCodeUnspecified KeywordCode = iota
	KeywordCodeAll
	KeywordCodeAnd
	KeywordCodeAny
	KeywordCodeArray
	KeywordCodeAs
	KeywordCodeAsc
	KeywordCodeAssertRowsModified
	KeywordCodeAt
	KeywordCodeBetween
	KeywordCodeBy
	KeywordCodeCase
	KeywordCodeCast
	KeywordCodeCollate
	KeywordCodeContains
	KeywordCodeCreate
	KeywordCodeCross
	KeywordCodeCube
	KeywordCodeCurrent
	KeywordCodeDefault
	KeywordCodeDefine
	KeywordCodeDesc
	KeywordCodeDistinct
	KeywordCodeElse
	KeywordCodeEnd
	KeywordCodeEnum
	KeywordCodeEscape
	KeywordCodeExcept
	KeywordCodeExclude
	KeywordCodeExists
	KeywordCodeExtract
	KeywordCodeFalse
	KeywordCodeFetch
	KeywordCodeFollowing
	KeywordCodeFor
	KeywordCodeFrom
	KeywordCodeFull
	KeywordCodeGroup
	KeywordCodeGrouping
	KeywordCodeGroups
	KeywordCodeHash
	KeywordCodeHaving
	KeywordCodeIf
	KeywordCodeIgnore
	KeywordCodeIn
	KeywordCodeInner
	KeywordCodeIntersect
	KeywordCodeInterval
	KeywordCodeInto
	KeywordCodeIs
	KeywordCodeJoin
	KeywordCodeLateral
	KeywordCodeLeft
	KeywordCodeLike
	KeywordCodeLimit
	KeywordCodeLookup
	KeywordCodeMerge
	KeywordCodeNatural
	KeywordCodeNew
	KeywordCodeNo
	KeywordCodeNot
	KeywordCodeNull
	KeywordCodeNulls
	KeywordCodeOf
	KeywordCodeOn
	KeywordCodeOr
	KeywordCodeOrder
	KeywordCodeOuter
	KeywordCodeOver
	KeywordCodePartition
	KeywordCodePreceding
	KeywordCodeProto
	KeywordCodeRange
	KeywordCodeRecursive
	KeywordCodeRespect
	KeywordCodeRight
	KeywordCodeRollup
	KeywordCodeRows
	KeywordCodeSelect
	KeywordCodeSet
	KeywordCodeSome
	KeywordCodeStruct
	KeywordCodeTablesample
	KeywordCodeThen
	KeywordCodeTo
	KeywordCodeTreat
	KeywordCodeTrue
	KeywordCodeUnbound
	KeywordCodeUnion
	KeywordCodeUnnest
	KeywordCodeUsing
	KeywordCodeWhen
	KeywordCodeWhere
	KeywordCodeWindow
	KeywordCodeWith
	KeywordCodeWithin
)

func (k KeywordCode) Keyword() string {
	return strings.ToUpper(k.String()[8:])
}

func OfKeyword(keyword string) (KeywordCode, error) {
	switch strings.ToUpper(keyword) {
	default:
		return KeywordCodeUnspecified, fmt.Errorf("not a reserved keyword: %q", keyword)
	case "ALL":
		return KeywordCodeAll, nil
	case "AND":
		return KeywordCodeAnd, nil
	case "ANY":
		return KeywordCodeAny, nil
	case "ARRAY":
		return KeywordCodeArray, nil
	case "AS":
		return KeywordCodeAs, nil
	case "ASC":
		return KeywordCodeAsc, nil
	case "ASSERT_ROWS_MODIFIED":
		return KeywordCodeAssertRowsModified, nil
	case "AT":
		return KeywordCodeAt, nil
	case "BETWEEN":
		return KeywordCodeBetween, nil
	case "BY":
		return KeywordCodeBy, nil
	case "CASE":
		return KeywordCodeCase, nil
	case "CAST":
		return KeywordCodeCast, nil
	case "COLLATE":
		return KeywordCodeCollate, nil
	case "CONTAINS":
		return KeywordCodeContains, nil
	case "CREATE":
		return KeywordCodeCreate, nil
	case "CROSS":
		return KeywordCodeCross, nil
	case "CUBE":
		return KeywordCodeCube, nil
	case "CURRENT":
		return KeywordCodeCurrent, nil
	case "DEFAULT":
		return KeywordCodeDefault, nil
	case "DEFINE":
		return KeywordCodeDefine, nil
	case "DESC":
		return KeywordCodeDesc, nil
	case "DISTINCT":
		return KeywordCodeDistinct, nil
	case "ELSE":
		return KeywordCodeElse, nil
	case "END":
		return KeywordCodeEnd, nil
	case "ENUM":
		return KeywordCodeEnum, nil
	case "ESCAPE":
		return KeywordCodeEscape, nil
	case "EXCEPT":
		return KeywordCodeExcept, nil
	case "EXCLUDE":
		return KeywordCodeExclude, nil
	case "EXISTS":
		return KeywordCodeExists, nil
	case "EXTRACT":
		return KeywordCodeExtract, nil
	case "FALSE":
		return KeywordCodeFalse, nil
	case "FETCH":
		return KeywordCodeFetch, nil
	case "FOLLOWING":
		return KeywordCodeFollowing, nil
	case "FOR":
		return KeywordCodeFor, nil
	case "FROM":
		return KeywordCodeFrom, nil
	case "FULL":
		return KeywordCodeFull, nil
	case "GROUP":
		return KeywordCodeGroup, nil
	case "GROUPING":
		return KeywordCodeGrouping, nil
	case "GROUPS":
		return KeywordCodeGroups, nil
	case "HASH":
		return KeywordCodeHash, nil
	case "HAVING":
		return KeywordCodeHaving, nil
	case "IF":
		return KeywordCodeIf, nil
	case "IGNORE":
		return KeywordCodeIgnore, nil
	case "IN":
		return KeywordCodeIn, nil
	case "INNER":
		return KeywordCodeInner, nil
	case "INTERSECT":
		return KeywordCodeIntersect, nil
	case "INTERVAL":
		return KeywordCodeInterval, nil
	case "INTO":
		return KeywordCodeInto, nil
	case "IS":
		return KeywordCodeIs, nil
	case "JOIN":
		return KeywordCodeJoin, nil
	case "LATERAL":
		return KeywordCodeLateral, nil
	case "LEFT":
		return KeywordCodeLeft, nil
	case "LIKE":
		return KeywordCodeLike, nil
	case "LIMIT":
		return KeywordCodeLimit, nil
	case "LOOKUP":
		return KeywordCodeLookup, nil
	case "MERGE":
		return KeywordCodeMerge, nil
	case "NATURAL":
		return KeywordCodeNatural, nil
	case "NEW":
		return KeywordCodeNew, nil
	case "NO":
		return KeywordCodeNo, nil
	case "NOT":
		return KeywordCodeNot, nil
	case "NULL":
		return KeywordCodeNull, nil
	case "NULLS":
		return KeywordCodeNulls, nil
	case "OF":
		return KeywordCodeOf, nil
	case "ON":
		return KeywordCodeOn, nil
	case "OR":
		return KeywordCodeOr, nil
	case "ORDER":
		return KeywordCodeOrder, nil
	case "OUTER":
		return KeywordCodeOuter, nil
	case "OVER":
		return KeywordCodeOver, nil
	case "PARTITION":
		return KeywordCodePartition, nil
	case "PRECEDING":
		return KeywordCodePreceding, nil
	case "PROTO":
		return KeywordCodeProto, nil
	case "RANGE":
		return KeywordCodeRange, nil
	case "RECURSIVE":
		return KeywordCodeRecursive, nil
	case "RESPECT":
		return KeywordCodeRespect, nil
	case "RIGHT":
		return KeywordCodeRight, nil
	case "ROLLUP":
		return KeywordCodeRollup, nil
	case "ROWS":
		return KeywordCodeRows, nil
	case "SELECT":
		return KeywordCodeSelect, nil
	case "SET":
		return KeywordCodeSet, nil
	case "SOME":
		return KeywordCodeSome, nil
	case "STRUCT":
		return KeywordCodeStruct, nil
	case "TABLESAMPLE":
		return KeywordCodeTablesample, nil
	case "THEN":
		return KeywordCodeThen, nil
	case "TO":
		return KeywordCodeTo, nil
	case "TREAT":
		return KeywordCodeTreat, nil
	case "TRUE":
		return KeywordCodeTrue, nil
	case "UNBOUNDED":
		return KeywordCodeUnbound, nil
	case "UNION":
		return KeywordCodeUnion, nil
	case "UNNEST":
		return KeywordCodeUnnest, nil
	case "USING":
		return KeywordCodeUsing, nil
	case "WHEN":
		return KeywordCodeWhen, nil
	case "WHERE":
		return KeywordCodeWhere, nil
	case "WINDOW":
		return KeywordCodeWindow, nil
	case "WITH":
		return KeywordCodeWith, nil
	case "WITHIN":
		return KeywordCodeWithin, nil
	}
}
