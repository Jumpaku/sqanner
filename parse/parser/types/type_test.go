package types_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/node/types"
	perser_types "github.com/Jumpaku/sqanner/parse/parser/types"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseType_Scalar(t *testing.T) {
	testcases := []test.Case[types.TypeNode]{
		{
			Message: `bool`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("bool")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBool(),
		},
		{
			Message: `bool`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("BOOL")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBool(),
		},
		{
			Message: `int64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("int64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewInt64(),
		},
		{
			Message: `int64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewInt64(),
		},
		{
			Message: `float64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("float64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewFloat64(),
		},
		{
			Message: `float64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("FLOAT64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewFloat64(),
		},
		{
			Message: `numeric`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("numeric")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewNumeric(),
		},
		{
			Message: `numeric`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewNumeric(),
		},
		{
			Message: `json`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("json")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewJSON(),
		},
		{
			Message: `json`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewJSON(),
		},
		{
			Message: `date`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("date")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewDate(),
		},
		{
			Message: `date`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewDate(),
		},
		{
			Message: `timestamp`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("timestamp")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTimestamp(),
		},
		{
			Message: `timestamp`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTimestamp(),
		},
		{
			Message: `string`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("string")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStringSized(types.NewTypeSize(123)),
		},
		{
			Message: `string`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStringSized(types.NewTypeSizeMax()),
		},
		{
			Message: `not sized string`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewString(),
		},
		{
			Message: `bytes`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("bytes")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBytesSized(types.NewTypeSize(123)),
		},
		{
			Message: `bytes`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBytesSized(types.NewTypeSizeMax()),
		},
		{
			Message: `not sized bytes`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBytes(),
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("_GRouP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `quoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`INT64`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `number`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `special char`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpecialChar, Content: []rune("@")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewBytesSized(types.NewTypeSize(123)),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, perser_types.ParseType)
		})
	}
}

func TestParseType_Array(t *testing.T) {
	testcases := []test.Case[types.TypeNode]{
		{
			Message: `array of bool`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BOOL")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewBool()),
		},
		{
			Message: `array of int64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewInt64()),
		},
		{
			Message: `array of float64`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("FLOAT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewFloat64()),
		},
		{
			Message: `array of numeric`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewNumeric()),
		},
		{
			Message: `array of json`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewJSON()),
		},
		{
			Message: `array of date`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewDate()),
		},
		{
			Message: `array of timestamp`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewTimestamp()),
		},
		{
			Message: `array of string`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewStringSized(types.NewTypeSizeMax())),
		},
		{
			Message: `array of bytes`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewBytesSized(types.NewTypeSize(123))),
		},
		{
			Message: `including spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(types.NewBytesSized(types.NewTypeSize(123))),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, perser_types.ParseType)
		})
	}
}

func TestParseType_Struct(t *testing.T) {
	testcases := []test.Case[types.TypeNode]{
		{
			Message: `struct with unnamed fields`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("int64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(",")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("string")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(",")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("bytes")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStruct([]types.StructFieldNode{
				types.NewStructFieldUnnamed(types.NewInt64()),
				types.NewStructFieldUnnamed(types.NewStringSized(types.NewTypeSize(123))),
				types.NewStructFieldUnnamed(types.NewBytesSized(types.NewTypeSizeMax())),
			}),
		},
		{
			Message: `empty struct`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStruct([]types.StructFieldNode{}),
		},
		{
			Message: `struct with named fields`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("a")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("int64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(",")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("b")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("string")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStruct([]types.StructFieldNode{
				types.NewStructField(node.NewIdentifier("a", false), types.NewInt64()),
				types.NewStructField(node.NewIdentifier("b", false), types.NewString()),
			}),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, perser_types.ParseType)
		})
	}
}

func TestParseType_Complex(t *testing.T) {
	testcases := []test.Case[types.TypeNode]{
		{
			Message: `struct in struct`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("x")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("y")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(",")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("z")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStruct([]types.StructFieldNode{
				types.NewStructField(
					node.NewIdentifier("x", false),
					types.NewStruct([]types.StructFieldNode{
						types.NewStructField(node.NewIdentifier("y", false), types.NewInt64()),
						types.NewStructField(node.NewIdentifier("z", false), types.NewInt64()),
					})),
			}),
		},
		{
			Message: `array in struct`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("inner_array")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStruct([]types.StructFieldNode{
				types.NewStructField(
					node.NewIdentifier("inner_array", false),
					types.NewArray(types.NewInt64())),
			}),
		},
		{
			Message: `struct in array`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(",")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(
				types.NewStruct([]types.StructFieldNode{
					types.NewStructFieldUnnamed(types.NewInt64()),
					types.NewStructFieldUnnamed(types.NewInt64()),
				})),
		},
		{
			Message: ``,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenKeyword, Content: []rune("STRUCT")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewArray(
				types.NewStruct([]types.StructFieldNode{
					types.NewStructFieldUnnamed(
						types.NewArray(types.NewInt64())),
				})),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, perser_types.ParseType)
		})
	}
}
