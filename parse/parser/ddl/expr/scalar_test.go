package expr_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node/ddl/expr"
	parser_expr "github.com/Jumpaku/sqanner/parse/parser/ddl/expr"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseScalar(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		// BOOL
		{
			Message: `true`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("tRuE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("TRUE"),
		},
		{
			Message: `false`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("FaLsE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("FALSE"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("TRUE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("TRUE"),
		},

		// INT64
		{
			Message: `decimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("123"),
		},
		{
			Message: `lower hexadecimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0xdef")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("0xdef"),
		},
		{
			Message: `upper hexadecimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0XABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("0xABC"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("123"),
		},

		// FLOAT64
		{
			Message: `exponent`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.456e-67")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("123.456e-67"),
		},
		{
			Message: `starts with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune(".1E4")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64(".1e4"),
		},
		{
			Message: `ends with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("58.")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("58."),
		},
		{
			Message: `no dots`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("4e2")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("4e2"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.456e-67")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("123.456e-67"),
		},

		// DATE
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dAtE")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'2023-11-24'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewDate("2023-11-24"),
		},

		// JSON
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("JsOn")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune(`"` + `{"a":-123.456,"b":"abc","c":true,"d":{"x":-1},"e":[1,"",false,{"x":{}},[]]}` + `"`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewJSON(`{"a":-123.456,"b":"abc","c":true,"d":{"x":-1},"e":[1,"",false,{"x":{}},[]]}`),
		},

		// TIMESTAMP
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("tImEsTaMp")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'" + "2023-11-24T15:20:10Z" + "'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewTimestamp("2023-11-24T15:20:10Z"),
		},

		// NUMERIC
		{
			Message: `exponent`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'123.456e-67'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("123.456e-67"),
		},
		{
			Message: `starts with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'.1E4'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric(".1E4"),
		},
		{
			Message: `ends with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'58.'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("58."),
		},
		{
			Message: `no dots`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'4e2'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("4e2"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'4e2'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("4e2"),
		},

		// STRING
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'abc'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewString("'abc'"),
		},

		// BYTES
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("b'abc'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBytes("b'abc'"),
		},
	}
	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseScalar)
		})
	}
}

func TestParseBool(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `true`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("tRuE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("TRUE"),
		},
		{
			Message: `false`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("FaLsE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("FALSE"),
		},

		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("TRUE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBool("TRUE"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseBool)
		})
	}
}

func TestParseInt64(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `decimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("123"),
		},
		{
			Message: `lower hexadecimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0xdef")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("0xdef"),
		},
		{
			Message: `upper hexadecimals`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0XABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("0xABC"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewInt64("123"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseInt64)
		})
	}
}

func TestParseFloat64(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `exponent`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.456e-67")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("123.456e-67"),
		},
		{
			Message: `starts with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune(".1E4")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64(".1e4"),
		},
		{
			Message: `ends with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("58.")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("58."),
		},
		{
			Message: `no dots`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("4e2")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("4e2"),
		},
		{
			Message: `starts with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.456e-67")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewFloat64("123.456e-67"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseFloat64)
		})
	}
}

func TestParseDate(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("dAtE")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'2023-11-24'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewDate("2023-11-24"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseDate)
		})
	}
}

func TestParseDate_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := "2023-11-24"
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewDate(want),
			})
		}
	}
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			date := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: date,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("DATE")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(date)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseDate)
		})
	}
}

func TestParseJSON(t *testing.T) {
	want := `{"a":-123.456,"b":"abc","c":true,"d":{"x":-1},"e":[1,"",false,{"x":{}},[]]}`
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("JsOn")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune(`"` + want + `"`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewJSON(want),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseJSON)
		})
	}
}

func TestParseJSON_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := `{"a":-123.456,"b":"abc","c":true,"d":{"x":-1},"e":[1,"",false,{"x":{}}]}`
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewJSON(want),
			})
		}
	}
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("JSON")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseJSON)
		})
	}
}

func TestParseTimestamp(t *testing.T) {
	want := "2023-11-24T15:20:10Z"
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("tImEsTaMp")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'" + want + "'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewTimestamp(want),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseTimestamp)
		})
	}
}

func TestParseTimestamp_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := "2023-11-24T15:20:10Z"
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewTimestamp(want),
			})
		}
	}
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("TIMESTAMP")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseTimestamp)
		})
	}
}

func TestParseNumeric(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `exponent`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'123.456e-67'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("123.456e-67"),
		},
		{
			Message: `starts with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'.1E4'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric(".1E4"),
		},
		{
			Message: `ends with dot`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'58.'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("58."),
		},
		{
			Message: `no dots`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'4e2'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("4e2"),
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("nUmeRiC")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'4e2'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewNumeric("4e2"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseNumeric)
		})
	}
}

func TestParseNumeric_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := "123"
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewNumeric(want),
			})
		}
	}
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			quoted := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: quoted,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenIdentifier, Content: []rune("NUMERIC")},
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(quoted)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseNumeric)
		})
	}
}

func TestParseString(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("DATE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("JSON")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("'abc'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewString("'abc'"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseString)
		})
	}
}

func TestParseString_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := "abc"
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			w := prefix + quote + "abc" + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: w,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(w)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewString(w),
			})
		}
	}
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			w := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: w,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(w)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseString)
		})
	}
}

func TestParseBytes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("ABC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("TIMESTAMP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("DATE")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("NUMERIC")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("JSON")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenLiteralQuoted, Content: []rune("b'abc'")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: expr.NewBytes("b'abc'"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseBytes)
		})
	}
}

func TestParseBytes_Quotes(t *testing.T) {
	testcases := []test.Case[expr.ScalarNode]{}
	want := "abc"
	for _, prefix := range []string{"b", "br", "bR", "B", "Br", "BR", "rB", "RB"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			w := prefix + quote + "abc" + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: w,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(w)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				WantNode: expr.NewBytes(w),
			})
		}
	}
	for _, prefix := range []string{"", "r", "R"} {
		for _, quote := range []string{`"`, `'`, `'''`, `"""`} {
			w := prefix + quote + want + quote
			testcases = append(testcases, test.Case[expr.ScalarNode]{
				Message: w,
				Input: []tokenize.Token{
					{Kind: tokenize.TokenLiteralQuoted, Content: []rune(w)},
					{Kind: tokenize.TokenEOF, Content: []rune("")},
				},
				ShouldErr: true,
			})
		}
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser_expr.ParseBytes)
		})
	}
}
