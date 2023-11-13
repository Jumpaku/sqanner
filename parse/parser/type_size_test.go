package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseTypeSize(t *testing.T) {
	testcases := []testcase[node.TypeSizeNode]{
		{
			message: `type size`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSize(123)),
		},
		{
			message: `type size`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0xFF")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSize(255)),
		},
		{
			message: `type size`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0X0A")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSize(10)),
		},
		{
			message: `float`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.5")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `max`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSizeMax()),
		},
		{
			message: `max`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("mAx")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSizeMax()),
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
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
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
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.TypeSize(123)),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseTypeSize)
		})
	}
}
