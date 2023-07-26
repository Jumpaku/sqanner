package tokenize_test

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

type testcase struct {
	message   string
	input     string
	wantLen   int
	wantCode  tokenize.TokenCode
	shouldErr bool
}
type scanFunc func(*tokenize.ScanState) (int, tokenize.TokenCode, error)

func check(t *testing.T, testcase testcase, sut scanFunc) {
	t.Helper()

	gotLen, gotCode, gotErr := sut(&tokenize.ScanState{Input: []rune(testcase.input)})
	if (gotErr != nil) != testcase.shouldErr {
		if testcase.shouldErr {
			t.Errorf(`%s: input=%q:
	err is expected but got nil`, testcase.message, testcase.input)
		} else {
			t.Errorf(`%s: input=%q:
	err is not expected but got %v`, testcase.message, testcase.input, gotErr)
		}

	}

	if !(gotLen == testcase.wantLen && gotCode == testcase.wantCode) {
		t.Errorf(`%s: input=%q:
	got  : len=%5d code=%v:
	want : len=%5d code=%v`, testcase.message, testcase.input, gotLen, gotCode, testcase.wantLen, testcase.wantCode)
	}
}

func TestSpaces(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `starts with non white space`,
			input:     `a`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `spaces ends with EOF`,
			input:     " \t\n\v\f\r",
			wantLen:   6,
			wantCode:  tokenize.TokenSpace,
			shouldErr: false,
		},
		{message: `spaces ends with alphabet`,
			input:     " \t\n\v\f\r A",
			wantLen:   7,
			wantCode:  tokenize.TokenSpace,
			shouldErr: false,
		},
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.Spaces)
	}
}

func TestComment(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `non comment`,
			input:     `a`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `empty comment starts with # ends with EOF`,
			input:     "#",
			wantLen:   1,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with # ends with \n`,
			input:     "#\n",
			wantLen:   2,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with // ends with EOF`,
			input:     "//",
			wantLen:   2,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with // ends with \n`,
			input:     "//\n",
			wantLen:   3,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with -- ends with EOF`,
			input:     "--",
			wantLen:   2,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with -- ends with \n`,
			input:     "--\n",
			wantLen:   3,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `empty comment starts with /* ends with */`,
			input:     "/**/",
			wantLen:   4,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with # ends with EOF`,
			input:     "#abc",
			wantLen:   4,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with # ends with \n`,
			input:     "#abc\n",
			wantLen:   5,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with // ends with EOF`,
			input:     "//abc",
			wantLen:   5,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with // ends with \n`,
			input:     "//abc\n",
			wantLen:   6,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with -- ends with EOF`,
			input:     "--abc",
			wantLen:   5,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with -- ends with \n`,
			input:     "--abc\n",
			wantLen:   6,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment starts with /* ends with */`,
			input:     "/*abc*/",
			wantLen:   7,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `multiline comment starts with /* ends with */`,
			input:     "/*\n  This is a multiline comment\n  on multiple lines\n*/",
			wantLen:   55,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `invalid comment /*/`,
			input:     "/*/",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `incomplete comment starts with /*`,
			input:     "/* abc",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `comment ends with first */ if nested`,
			input:     "/* /* nested */ */",
			wantLen:   15,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
		{message: `comment ends with first */ if consecutive`,
			input:     "/* 1 */ /* 2 */",
			wantLen:   7,
			wantCode:  tokenize.TokenComment,
			shouldErr: false,
		},
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.Comment)
	}
}

func TestIdentifierQuoted(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `empty quoted identifier`,
			input:     "``",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `incomplete quoted identifier`,
			input:     "`a",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid escape sequence \`,
			input:     "`" + `\` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid escape sequence \c`,
			input:     "`" + `\c` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid octal escape sequence \678`,
			input:     "`" + `\678` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid octal escape sequence \67`,
			input:     "`" + `\67` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid hex escape sequence \x4`,
			input:     "`" + `\x4` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid hex escape sequence \X1y`,
			input:     "`" + `\X1y` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid unicode escape sequence \uabc`,
			input:     "`" + `\uabc` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid unicode escape sequence \uefgh`,
			input:     "`" + `\uefgh` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid unicode escape sequence \Uabcde`,
			input:     "`" + `\Uabcde` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains invalid unicode escape sequence \Uabcdefgh`,
			input:     "`" + `\Uabcdefgh` + "`",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted identifier contains escape sequence \a \b \f \n \r \t \v`,
			input:     "`" + `\a \b \f \n \r \t \v` + "`",
			wantLen:   22,
			wantCode:  tokenize.TokenIdentifierQuoted,
			shouldErr: false,
		},
		{message: `quoted identifier contains escape sequence \\ \? \" \'`,
			input:     "`" + `\\ \? \" \'` + "`",
			wantLen:   13,
			wantCode:  tokenize.TokenIdentifierQuoted,
			shouldErr: false,
		},
		{message: `quoted identifier contains escape sequence \770 \xFF \X00 \u00fF \U00ffFF00`,
			input:     "`" + `\770 \xFF \X00 \u00fF \U00ffFF00` + "`",
			wantLen:   34,
			wantCode:  tokenize.TokenIdentifierQuoted,
			shouldErr: false,
		},
		{message: `quoted identifier contains escaped back quote`,
			input:     string([]rune{'`', '\\', '`', '`'}),
			wantLen:   4,
			wantCode:  tokenize.TokenIdentifierQuoted,
			shouldErr: false,
		},
		{message: `quoted identifier contains any characters`,
			input:     "`_1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()-+={}[]|:\"'<>?,./`",
			wantLen:   93,
			wantCode:  tokenize.TokenIdentifierQuoted,
			shouldErr: false,
		},
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.IdentifierQuoted)
	}
}

func TestLiteralQuoted(t *testing.T) {
	type lqTestcase struct {
		message        string
		prefix         []string
		quote          []string
		content        string
		wantContentLen int
		wantCode       tokenize.TokenCode
		shouldErr      bool
	}
	toTestcases := func(c lqTestcase) []testcase {
		ts := []testcase{}
		for _, prefix := range c.prefix {
			for _, quote := range c.quote {
				input := prefix + quote + c.content + quote
				wantLen := 0
				if !c.shouldErr {
					wantLen = c.wantContentLen + +len(prefix) + 2*len(quote)
				}
				ts = append(ts, testcase{
					message:   c.message + `: ` + input,
					input:     input,
					wantLen:   wantLen,
					wantCode:  c.wantCode,
					shouldErr: c.shouldErr,
				})
			}
		}
		return ts
	}
	lqTestcases := []lqTestcase{
		{message: `empty string`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `r`, `b`, `R`, `B`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        ``,
			wantContentLen: 0,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\`,
			wantContentLen: 1,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\c`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\c`,
			wantContentLen: 2,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\678`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\678`,
			wantContentLen: 4,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\67`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\67`,
			wantContentLen: 3,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\x4`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\x4`,
			wantContentLen: 3,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\X1y`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\X1y`,
			wantContentLen: 4,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\uabc`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\uabc`,
			wantContentLen: 5,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\uefgh`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\uefgh`,
			wantContentLen: 6,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\Uabcde`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\Uabcde`,
			wantContentLen: 7,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having invalid escape sequence`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`},
			content:        `\Uabcdefgh`,
			wantContentLen: 0,
			wantCode:       tokenize.TokenUnspecified,
			shouldErr:      true,
		},
		{message: `quoted literal having \`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\Uabcdefgh`,
			wantContentLen: 10,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having escape sequences`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\a \b \f \n \r \t \v`,
			wantContentLen: 20,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having escape sequences`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\\ \?`,
			wantContentLen: 5,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having escape sequences`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\770 \xFF \X00 \u00fF \U00ffFF00`,
			wantContentLen: 32,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having escape sequences`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `\770 \xAF \xaf \X09 \u09fF \u28aA \U09afAF5b`,
			wantContentLen: 44,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having escaped back quote`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        "\\`",
			wantContentLen: 2,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having valid characters`,
			quote:          []string{`'`, `"`, `'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `_1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ~!@#$%^&*()-+={}[]|:<>?,./`,
			wantContentLen: 90,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having newline`,
			quote:          []string{`'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        "two\nlines",
			wantContentLen: 9,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having newline`,
			quote:          []string{`'`, `"`},
			prefix:         []string{`r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        "two\nlines",
			wantContentLen: 9,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having a unescaped quote`,
			quote:          []string{`'''`, `"""`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `it's "double" quote`,
			wantContentLen: 19,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having a unescaped quote`,
			quote:          []string{`'`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `"double quote"`,
			wantContentLen: 14,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
		{message: `quoted literal having a unescaped quote`,
			quote:          []string{`"`},
			prefix:         []string{``, `b`, `B`, `r`, `R`, `rb`, `br`, `Rb`, `Br`, `rB`, `bR`, `RB`, `BR`},
			content:        `'single quote'`,
			wantContentLen: 14,
			wantCode:       tokenize.TokenLiteralQuoted,
			shouldErr:      false,
		},
	}

	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `incomplete quoted literal starts with '`,
			input:     `'a`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `incomplete quoted literal starts with Rb"""`,
			input:     `Rb"""a`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `quoted literal followed by quotes`,
			input:     `Rb"""a""""""`,
			wantLen:   9,
			wantCode:  tokenize.TokenLiteralQuoted,
			shouldErr: false,
		},
		{message: `quoted literal followed by quotes`,
			input:     `bR'a''''`,
			wantLen:   5,
			wantCode:  tokenize.TokenLiteralQuoted,
			shouldErr: false,
		},
		{message: `quoted literal followed by tokens`,
			input:     `Rb"""a"""abc`,
			wantLen:   9,
			wantCode:  tokenize.TokenLiteralQuoted,
			shouldErr: false,
		},
		{message: `quoted literal followed by tokens`,
			input:     `bR'a'abc`,
			wantLen:   5,
			wantCode:  tokenize.TokenLiteralQuoted,
			shouldErr: false,
		},
	}

	for _, lqTestcase := range lqTestcases {
		testcases = append(testcases, toTestcases(lqTestcase)...)
	}
	for _, testcase := range testcases {
		check(t, testcase, tokenize.LiteralQuoted)
	}
}

func TestIdentifierOrKeyword(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `starts with not letter`,
			input:     `1`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `starts with not letter`,
			input:     " ",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `starts with underscore`,
			input:     "_12ab",
			wantLen:   5,
			wantCode:  tokenize.TokenIdentifier,
			shouldErr: false,
		},
		{message: `starts with alphabet`,
			input:     "ab_12",
			wantLen:   5,
			wantCode:  tokenize.TokenIdentifier,
			shouldErr: false,
		},
		{message: `identifier starts with keyword`,
			input:     "SELECT_XYZ",
			wantLen:   10,
			wantCode:  tokenize.TokenIdentifier,
			shouldErr: false,
		},
		{message: `identifier followed by tokens`,
			input:     "_XYZ123 123",
			wantLen:   7,
			wantCode:  tokenize.TokenIdentifier,
			shouldErr: false,
		},
		{message: `keyword followed by tokens`,
			input:     "SELECT 123",
			wantLen:   6,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
		{message: `case-insensitive keyword`,
			input:     "sElEcT",
			wantLen:   6,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
		{message: `keyword`,
			input:     "NUMERIC",
			wantLen:   7,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
		{message: `keyword`,
			input:     "DATE",
			wantLen:   4,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
		{message: `keyword`,
			input:     "TiMeStAmP",
			wantLen:   9,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
		{message: `keyword`,
			input:     "json",
			wantLen:   4,
			wantCode:  tokenize.TokenKeyword,
			shouldErr: false,
		},
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.IdentifierOrKeyword)
	}
}

func TestNumberOrDot(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `letter`,
			input:     `a`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `symbol`,
			input:     `@`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `operator`,
			input:     `+`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `operator`,
			input:     `-`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.`,
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.a`,
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.1`,
			wantLen:   2,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.e`,
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.e1`,
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.e1+`,
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.02e1+`,
			wantLen:   5,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.92e-1+`,
			wantLen:   6,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.10e+912`,
			wantLen:   8,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `starts with dot`,
			input:     `.1E4`,
			wantLen:   4,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `hex`,
			input:     `0x`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `starts with 0x`,
			input:     `0xyz`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `starts with 0x`,
			input:     `0X `,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `starts with 0x`,
			input:     `1x1234`,
			wantLen:   1,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `starts with 0x`,
			input:     `0xFf09aA`,
			wantLen:   8,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `starts with 0x`,
			input:     `0X1fG`,
			wantLen:   4,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `simple integer`,
			input:     `0A`,
			wantLen:   1,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `simple integer`,
			input:     `1`,
			wantLen:   1,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `simple integer`,
			input:     `9876543210+`,
			wantLen:   10,
			wantCode:  tokenize.TokenLiteralInteger,
			shouldErr: false,
		},
		{message: `float`,
			input:     `123.456e-67`,
			wantLen:   11,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `0.0E+0`,
			wantLen:   6,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `9.9E9`,
			wantLen:   5,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `58.`,
			wantLen:   3,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `58.e+3`,
			wantLen:   6,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `58.E-3`,
			wantLen:   6,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `58.E+0`,
			wantLen:   6,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `float`,
			input:     `4e2`,
			wantLen:   3,
			wantCode:  tokenize.TokenLiteralFloat,
			shouldErr: false,
		},
		{message: `invalid float`,
			input:     `4e`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `1.E`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `4e+`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `4E-`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `.1E-`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `.9e+`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `4e`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `1.Ee`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `4e+-`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `4E-@`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `.1E-+`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
		{message: `invalid float`,
			input:     `.9e+A`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: true,
		},
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.NumberOrDot)
	}
}

func TestSpecialChar(t *testing.T) {
	testcases := []testcase{
		{message: `empty`,
			input:     ``,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `number`,
			input:     `1`,
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `underscore`,
			input:     "_",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `space`,
			input:     " ",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `alphabet`,
			input:     "a",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
		{message: `double symbols`,
			input:     ">>",
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `double symbols`,
			input:     ".>",
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		},
		{message: `dot followed by number`,
			input:     ".0",
			wantLen:   0,
			wantCode:  tokenize.TokenUnspecified,
			shouldErr: false,
		},
	}

	for _, c := range []rune("@,()[]{}<>.;:/+-~*|&^=!$?") {
		testcases = append(testcases, testcase{
			message:   `special`,
			input:     string(c),
			wantLen:   1,
			wantCode:  tokenize.TokenSpecialChar,
			shouldErr: false,
		})
	}

	for _, testcase := range testcases {
		check(t, testcase, tokenize.SpecialChar)
	}
}
