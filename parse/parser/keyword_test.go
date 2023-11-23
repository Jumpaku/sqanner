package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseKeyword(t *testing.T) {
	testcases := []test.Case[node.KeywordNode]{
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ALL")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewKeyword(parse.KeywordCodeAll),
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("aSsErT_rOwS_mOdIfIeD")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewKeyword(parse.KeywordCodeAssertRowsModified),
		},
		{
			Message: `keyword`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewKeyword(parse.KeywordCodeGroup),
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
			Message: `quotedidentifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
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
				{Kind: tokenize.TokenKeyword, Content: []rune("SELECT")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: node.NewKeyword(parse.KeywordCodeSelect),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, parser.ParseKeyword)
		})
	}
}
