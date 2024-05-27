package types_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node/types"
	types2 "github.com/Jumpaku/sqanner/parse/parser/types"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseTypeSize(t *testing.T) {
	testcases := []test.Case[types.TypeSizeNode]{
		{
			Message: `type size`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSize(123),
		},
		{
			Message: `type size`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0xFF")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSize(255),
		},
		{
			Message: `type size`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("0X0A")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSize(10),
		},
		{
			Message: `float`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenLiteralFloat, Content: []rune("123.5")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `max`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSizeMax(),
		},
		{
			Message: `max`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("mAx")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSizeMax(),
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
			Message: `identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune("_GRouP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			ShouldErr: true,
		},
		{
			Message: `quoted identifier`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
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
				{Kind: tokenize.TokenLiteralInteger, Content: []rune("123")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewTypeSize(123),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, types2.ParseTypeSize)
		})
	}
}
