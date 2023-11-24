package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	testcases := []test.Case[node.IdentifierNode]{
		{
			Message: `unquoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("_5abc")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewIdentifier("_5abc", false),
		},
		{
			Message: `unquoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewIdentifier("abc5", false),
		},
		{
			Message: `quoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewIdentifier("GROUP", true),
		},
		{
			Message: `quoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`gRouP`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewIdentifier("gRouP", true),
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
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewIdentifier("abc", false),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser.ParseIdentifier)
		})
	}
}
