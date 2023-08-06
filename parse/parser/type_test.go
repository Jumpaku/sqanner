package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseType_Scalar(t *testing.T) {
	testcases := []testcase[node.TypeNode]{
		{
			message: `bool`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("bool")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.BoolType()),
		},
		{
			message: `bool`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("BOOL")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.BoolType()),
		},
		{
			message: `int64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("int64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Int64Type()),
		},
		{
			message: `int64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Int64Type()),
		},
		{
			message: `float64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("float64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Float64Type()),
		},
		{
			message: `float64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("FLOAT64")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Float64Type()),
		},
		{
			message: `numeric`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("numeric")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.NumericType()),
		},
		{
			message: `numeric`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.NumericType()),
		},
		{
			message: `json`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("json")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.JSONType()),
		},
		{
			message: `json`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.JSONType()),
		},
		{
			message: `date`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("date")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.DateType()),
		},
		{
			message: `date`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.DateType()),
		},
		{
			message: `timestamp`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("timestamp")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TimestampType()),
		},
		{
			message: `timestamp`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TimestampType()),
		},
		{
			message: `string`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("string")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StringTypeSized(nodeOf(node.TypeSize(123)))),
		},
		{
			message: `string`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))),
		},
		{
			message: `bytes`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("bytes")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.BytesTypeSized(nodeOf(node.TypeSize(123)))),
		},
		{
			message: `bytes`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.BytesTypeSized(nodeOf(node.TypeSizeMax()))),
		},
		{
			message: `identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("_GRouP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `quoted identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`INT64`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `number`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `dot`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `special char`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenSpecialChar, Content: []rune("@")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `starts with spaces`,
			input: []tokenize.Token{
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
			wantNode: nodeOf(node.BytesTypeSized(nodeOf(node.TypeSize(123)))),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseType)
		})
	}
}

func TestParseType_Array(t *testing.T) {
	testcases := []testcase[node.TypeNode]{
		{
			message: `array of bool`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BOOL")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.BoolType()))),
		},
		{
			message: `array of int64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("INT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.Int64Type()))),
		},
		{
			message: `array of float64`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("FLOAT64")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.Float64Type()))),
		},
		{
			message: `array of numeric`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.NumericType()))),
		},
		{
			message: `array of json`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.JSONType()))),
		},
		{
			message: `array of date`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.DateType()))),
		},
		{
			message: `array of timestamp`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.TimestampType()))),
		},
		{
			message: `array of string`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))),
		},
		{
			message: `array of bytes`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ARRAY")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("BYTES")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.ArrayType(nodeOf(node.BytesTypeSized(nodeOf(node.TypeSize(123)))))),
		},
		{
			message: `including spaces`,
			input: []tokenize.Token{
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
			wantNode: nodeOf(node.ArrayType(nodeOf(node.BytesTypeSized(nodeOf(node.TypeSize(123)))))),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseType)
		})
	}
}
