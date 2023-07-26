package tokenize_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

type testcaseTokenize struct {
	message    string
	input      string
	wantTokens []tokenize.TokenCode
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
		var want, got tokenize.TokenCode
		s := fmt.Sprintf(`	token[%d]:	`, i)
		if i < len(testcase.wantTokens) {
			want = testcase.wantTokens[i]
			s += fmt.Sprintf(`want=%-17s`, want.String()[5:])
		} else {
			s += fmt.Sprintf(`want=%-17s`, "(nothing)")
		}

		if i < len(gotTokens) {
			got = gotTokens[i].Code
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

func TestTokenize(t *testing.T) {
	testcases := []testcaseTokenize{
		{
			message: `valid identifiers`,
			input:   "-- Valid. _5abc and dataField are valid identifiers.\n_5abc.dataField",
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenIdentifier,
				tokenize.TokenEOF,
			},
			shouldErr: false,
		},
		{
			message: `invalid identifiers`,
			input:   "-- Invalid. GROUP is an invalid identifier because it is unquoted and is a\n-- stand-alone reserved keyword.\nGROUP.dataField",
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
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
			wantTokens: []tokenize.TokenCode{
				tokenize.TokenComment,
				tokenize.TokenIdentifier,
				tokenize.TokenSpecialChar,
				tokenize.TokenKeyword,
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
			wantTokens: []tokenize.TokenCode{
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

	for _, testcase := range testcases {
		testTokenize(t, testcase)
	}
}
