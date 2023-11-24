package tokenize_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/tokenize"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcaseTokenize struct {
	message    string
	input      string
	wantTokens []tokenize.TokenKind
	shouldErr  bool
}

func testTokenize(t *testing.T, testcase testcaseTokenize) {
	t.Helper()

	gotTokens, gotErr := tokenize.Tokenize([]rune(testcase.input))
	if (gotErr != nil) != testcase.shouldErr {
		if testcase.shouldErr {
			t.Errorf("%s: input=%q:\n	err is expected but got nil", testcase.message, testcase.input)
		} else {
			t.Errorf("%s: input=%q:\n	err is not expected but got %v", testcase.message, testcase.input, gotErr)
		}
	}

	size := len(gotTokens)
	if size < len(testcase.wantTokens) {
		size = len(testcase.wantTokens)
	}

	ok := len(testcase.wantTokens) == len(gotTokens)
	diff := ``
	for i := 0; i < size; i++ {
		var want, got tokenize.TokenKind
		s := fmt.Sprintf(`	token[%d]:	`, i)
		if i < len(testcase.wantTokens) {
			want = testcase.wantTokens[i]
			s += fmt.Sprintf(`want=%-17s`, want.String()[5:])
		} else {
			s += fmt.Sprintf(`want=%-17s`, "(nothing)")
		}

		if i < len(gotTokens) {
			got = gotTokens[i].Kind
			s += fmt.Sprintf(`got=%s(%q)`, got.String()[5:], string(gotTokens[i].Content))
		} else {
			s += fmt.Sprintf(`got=(nothing)`)
		}

		diff += s + "\n"

		ok = ok && (want == got)
	}
	if !ok {
		t.Errorf("%s: input=%q\n%s", testcase.message, testcase.input, diff)
	}
}

func TestTokenize_Identifiers(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#identifiers
	testcases := []testcaseTokenize{
		{
			message: `valid identifiers`,
			input:   "-- Valid. _5abc and dataField are valid identifiers.\n_5abc.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid quoted identifiers`,
			input:   "-- Valid. `5abc` and dataField are valid identifiers.\n`5abc`.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifierQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message:    `invalid identifiers`,
			input:      "-- Invalid. 5abc is an invalid identifier because it is unquoted and starts\n-- with a number rather than a letter or underscore.\n5abc.dataField",
			wantTokens: nil,
			shouldErr:  true,
		},
		{
			message: `valid identifiers`,
			input:   "-- Valid. abc5 and dataField are valid identifiers.\nabc5.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `identifiers`,
			input:   "-- Invalid. abc5! is an invalid identifier because it is unquoted and contains\n-- a character that is not a letter, number, or underscore.\nabc5!.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid quoted identifiers`,
			input:   "-- Valid. `GROUP` and dataField are valid identifiers.\n`GROUP`.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifierQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `invalid identifiers`,
			input:   "-- Invalid. GROUP is an invalid identifier because it is unquoted and is a\n-- stand-alone reserved keyword.\nGROUP.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenComment,
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `identifiers`,
			input:   "-- Valid. abc5 and GROUP are valid identifiers.\nabc5.GROUP",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenKeyword,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `function call`,
			input:   "-- Valid. dataField is a valid identifier in a function called foo().\nfoo().dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `function call`,
			input:   "-- Valid. dataField is a valid identifier in an array called items.\nitems[OFFSET(3)].dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `parameter`,
			input:   "-- Valid. param and dataField are valid identifiers.\n@param.dataField",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_PathExpression(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#path_expressions
	testcases := []testcaseTokenize{
		{
			message: `valid path expression`,
			input:   "foo.bar",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid path expression`,
			input:   "foo.bar/25",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid path expression`,
			input:   "foo/bar:25",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid path expression`,
			input:   "foo/bar/25-31",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid path expression`,
			input:   "/foo/bar",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `valid path expression`,
			input:   "/25/foo/bar",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_StringAndByteLiteral(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#string_and_bytes_literals
	testcases := []testcaseTokenize{
		{
			message: `quoted string`,
			input:   "\"abc\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `quoted string`,
			input:   "\"it's\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `quoted string`,
			input:   "'it\\'s'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `quoted string`,
			input:   "'Title: \"Boy\"'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `triple-quoted string`,
			input:   "\"\"\"abc\"\"\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `triple-quoted string`,
			input:   "'''it's'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `triple-quoted string`,
			input:   "'''Title:\"Boy\"'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `triple-quoted string`,
			input:   "'''two\nlines'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `triple-quoted string`,
			input:   "'''why\\?'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw string`,
			input:   "r\"abc+\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw string`,
			input:   "r'''abc+'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw string`,
			input:   "r\"\"\"abc+\"\"\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw string`,
			input:   "r'f\\(abc,(.*),def\\)'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `bytes`,
			input:   "B\"abc\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `bytes`,
			input:   "B'''abc'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `bytes`,
			input:   "b\"\"\"abc\"\"\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw bytes`,
			input:   "br'abc+'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw bytes`,
			input:   "RB\"abc+\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `raw bytes`,
			input:   "RB'''abc'''",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_IntegerLiteral(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#integer_literals
	testcases := []testcaseTokenize{
		{
			message: `integer`,
			input:   "123",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `integer`,
			input:   "0xABC",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `integer`,
			input:   "-123",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_NumericLiteral(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#numeric_literals
	testcases := []testcaseTokenize{
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '0';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '123456';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '-3.14';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '-0.54321';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '1.23456e05';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `numeric`,
			input:   "SELECT NUMERIC '-9.876e-3';",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_FloatLiteral(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#floating_point_literals
	testcases := []testcaseTokenize{
		{
			message: `float`,
			input:   "123.456e-67",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralFloat,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `float`,
			input:   ".1E4",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralFloat,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `float`,
			input:   "58.",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralFloat,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `float`,
			input:   "4e2",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenLiteralFloat,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_ArrayLiteral(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#array_literals
	testcases := []testcaseTokenize{
		{
			message: `array of integers`,
			input:   "[1, 2, 3]",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `array of strings`,
			input:   "['x', 'y', 'xy']",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `array of integers`,
			input:   "ARRAY[1, 2, 3]",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `array of strings`,
			input:   "ARRAY<string>['x', 'y', 'xy']",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `array of integers`,
			input:   "ARRAY<int64>[]",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_Struct(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#struct_literals
	testcases := []testcaseTokenize{
		{
			message: `struct value`,
			input:   "(1, 2, 3)",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct value`,
			input:   "(1, 'abc')",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct value`,
			input:   "STRUCT(1 AS foo, 'abc' AS bar)",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct value`,
			input:   "STRUCT<INT64, STRING>(1, 'abc')",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct value`,
			input:   "STRUCT(1)",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct value`,
			input:   "STRUCT<INT64>(1)",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct type`,
			input:   "STRUCT<INT64, INT64, INT64>",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `struct type`,
			input:   "STRUCT<foo INT64, bar STRING>",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_Date_Timestamp_JSON(t *testing.T) {
	testcases := []testcaseTokenize{
		// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#date_literals
		{
			message: `date`,
			input:   "DATE '2014-09-27'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `statement using date`,
			input:   "SELECT * FROM foo WHERE date_col = \"2014-09-27\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#timestamp_literals
		{
			message: `timestamp`,
			input:   "TIMESTAMP '2014-09-27 12:30:00.45-08'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `statement using timestamp`,
			input:   "SELECT * FROM foo\nWHERE timestamp_col = \"2014-09-27 12:30:00.45 America/Los_Angeles\"",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#json_literals
		{
			message: `JSON`,
			input:   "JSON '\n{\n  \"id\": 10,\n  \"type\": \"fruit\",\n  \"name\": \"apple\",\n  \"on_menu\": true,\n  \"recipes\":\n    {\n      \"salads\":\n      [\n        { \"id\": 2001, \"type\": \"Walnut Apple Salad\" },\n        { \"id\": 2002, \"type\": \"Apple Spinach Salad\" }\n      ],\n      \"desserts\":\n      [\n        { \"id\": 3001, \"type\": \"Apple Pie\" },\n        { \"id\": 3002, \"type\": \"Apple Scones\" },\n        { \"id\": 3003, \"type\": \"Apple Crumble\" }\n      ]\n    }\n}\n'",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_Parameter(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#named_query_parameters
	testcases := []testcaseTokenize{
		{
			message: `parameter`,
			input:   "SELECT * FROM Roster WHERE LastName = @myparam",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_Hint(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#hints
	testcases := []testcaseTokenize{
		{
			message: `hint`,
			input:   "@{ database_engine_a.file_count=23, database_engine_b.file_count=10 }",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenSpecialChar,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenLiteralInteger,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestTokenize_Comment(t *testing.T) {
	// https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#comments
	testcases := []testcaseTokenize{
		{
			message: `single line comment with #`,
			input:   "# this is a single-line comment\nSELECT book FROM library;",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `single line comment with --`,
			input:   "-- this is a single-line comment\nSELECT book FROM library;",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `single line comment with /* and */`,
			input:   "/* this is a single-line comment */\nSELECT book FROM library;",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenComment,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `single line comment with /* and */`,
			input:   "SELECT book FROM library\n/* this is a single-line comment */\nWHERE book = \"Ulysses\";",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `inline comment with #`,
			input:   "SELECT book FROM library; # this is an inline comment",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `inline comment with --`,
			input:   "SELECT book FROM library; -- this is an inline comment",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `inline comment with /* and */`,
			input:   "SELECT book FROM library; /* this is an inline comment */",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `inline comment with /* and */`,
			input:   "SELECT book FROM library /* this is an inline comment */ WHERE book = \"Ulysses\";",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `multiline comment with /* and */`,
			input:   "SELECT book FROM library\n/*\n  This is a multiline comment\n  on multiple lines\n*/\nWHERE book = \"Ulysses\";",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `multiline comment with /* and */`,
			input:   "SELECT book FROM library\n/* this is a multiline comment\non two lines */\nWHERE book = \"Ulysses\";",
			wantTokens: []tokenize.TokenKind{
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenComment,
				tokenize.TokenSpace,
				tokenize.TokenKeyword,
				tokenize.TokenSpace,
				tokenize.TokenIdentifier,
				tokenize.TokenSpace,
				tokenize.TokenSpecialChar,
				tokenize.TokenSpace,
				tokenize.TokenLiteralQuoted,
				tokenize.TokenSpecialChar,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testTokenize(t, testcase)
		})
	}
}

func TestScanNext(t *testing.T) {
	testcases := []struct {
		message   string
		kind      tokenize.TokenKind
		input     string
		want      string
		shouldErr bool
	}{
		{
			message: `EOF`,
			kind:    tokenize.TokenEOF,
			input:   ``,
			want:    ``,
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   "-- comment\n not comment",
			want:    "-- comment\n",
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   "-- comment",
			want:    `-- comment`,
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   "# comment\n not comment",
			want:    "# comment\n",
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   "# comment",
			want:    `# comment`,
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   `/* comment */`,
			want:    `/* comment */`,
		},
		{
			message: `comment`,
			kind:    tokenize.TokenComment,
			input:   `/* comment */ not comment`,
			want:    `/* comment */`,
		},
		{
			message: `float literal`,
			kind:    tokenize.TokenLiteralFloat,
			input:   `.0e123 not literal`,
			want:    `.0e123`,
		},
		{
			message: `integer literal`,
			kind:    tokenize.TokenLiteralInteger,
			input:   `123 not literal`,
			want:    `123`,
		},
		{
			message: `keyword`,
			kind:    tokenize.TokenKeyword,
			input:   `GROUP not keyword`,
			want:    `GROUP`,
		},
		{
			message: `quoted literal`,
			kind:    tokenize.TokenLiteralQuoted,
			input:   `rb"ab'c" not literal`,
			want:    `rb"ab'c"`,
		},
		{
			message: `quoted literal`,
			kind:    tokenize.TokenLiteralQuoted,
			input:   `BR'''ab'c''' not literal`,
			want:    `BR'''ab'c'''`,
		},
		{
			message: `quoted identifier`,
			kind:    tokenize.TokenIdentifierQuoted,
			input:   "`GROUP` not identifier",
			want:    "`GROUP`",
		},
		{
			message: `identifier`,
			kind:    tokenize.TokenIdentifier,
			input:   "ABC not identifier",
			want:    "ABC",
		},
		{
			message: `empty spaces`,
			kind:    tokenize.TokenSpace,
			input:   "ABC",
			want:    "",
		},
		{
			message: `spaces`,
			kind:    tokenize.TokenSpace,
			input:   " \n\t ABC",
			want:    " \n\t ",
		},
		{
			message: `dot`,
			kind:    tokenize.TokenSpecialChar,
			input:   ".`123`",
			want:    ".",
		},
		{
			message: `minus`,
			kind:    tokenize.TokenSpecialChar,
			input:   "-123",
			want:    "-",
		},
		{
			message: `slash`,
			kind:    tokenize.TokenSpecialChar,
			input:   "/abc",
			want:    "/",
		},
		{
			message: `plus`,
			kind:    tokenize.TokenSpecialChar,
			input:   "+123",
			want:    "+",
		},
	}
	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			got, err := tokenize.ScanNext([]rune(testcase.input), testcase.kind)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testcase.want, string(got))
			}
		})
	}
}
