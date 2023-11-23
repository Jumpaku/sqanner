package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	testcases := []testcase[*node.IdentifierNode]{
		{
			message: `unquoted identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("_5abc")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: node.NewIdentifier("_5abc"),
		},
		{
			message: `unquoted identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: node.NewIdentifier("abc5"),
		},
		{
			message: `quoted identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("`GROUP`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: node.NewIdentifier("`GROUP`"),
		},
		{
			message: `quoted identifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("`gRouP`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: node.NewIdentifier("`gRouP`"),
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
			message: `keyword`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: node.NewIdentifier("abc"),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseIdentifier)
		})
	}
}
