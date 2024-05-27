package types

//go:generate go run "golang.org/x/tools/cmd/stringer" -type TypeCode type_code.go
type TypeCode int

const (
	TypeCodeUnspecified TypeCode = iota
	TypeCodeArray
	TypeCodeBool
	TypeCodeBytes
	TypeCodeDate
	TypeCodeJSON
	TypeCodeInt64
	TypeCodeNumeric
	TypeCodeFloat64
	TypeCodeString
	TypeCodeStruct
	TypeCodeTimestamp
)
