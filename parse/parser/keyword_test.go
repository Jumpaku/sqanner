package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseKeyword(t *testing.T) {
	testcases := []testcase[node.KeywordNode]{
		{
			message: `keyword`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("ALL")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Keyword(node.KeywordAll)),
		},
		{
			message: `keyword`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("aSsErT_rOwS_mOdIfIeD")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Keyword(node.KeywordAssertRowsModified)),
		},
		{
			message: `keyword`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Keyword(node.KeywordGroup)),
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
			message: `quotedidentifier`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
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
				{Kind: tokenize.TokenKeyword, Content: []rune("SELECT")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Keyword(node.KeywordSelect)),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseKeyword)
		})
	}
}
