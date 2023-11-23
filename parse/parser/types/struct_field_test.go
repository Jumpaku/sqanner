package types_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/node/types"
	perser_types "github.com/Jumpaku/sqanner/parse/parser/types"
	"github.com/Jumpaku/sqanner/parse/test"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseStructField(t *testing.T) {
	testcases := []test.Case[types.StructFieldNode]{
		{
			Message: `unnamed field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`INT64`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructFieldUnnamed(types.NewInt64()),
		},
		{
			Message: `unnamed field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructFieldUnnamed(types.NewString()),
		},
		{
			Message: `unnamed field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructFieldUnnamed(types.NewStringSized(types.NewTypeSizeMax())),
		},
		{
			Message: `unnamed field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune(`ARRAY`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`<`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`>`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructFieldUnnamed(types.NewArray(types.NewStringSized(types.NewTypeSizeMax()))),
		},
		{
			Message: `named field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`INT64`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructField(node.NewIdentifier(`abc`), types.NewInt64()),
		},
		{
			Message: `named field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructField(node.NewIdentifier(`abc`), types.NewString()),
		},
		{
			Message: `named field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructField(node.NewIdentifier(`abc`), types.NewStringSized(types.NewTypeSizeMax())),
		},
		{
			Message: `named field`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenKeyword, Content: []rune(`ARRAY`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`<`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`>`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructField(node.NewIdentifier(`abc`), types.NewArray(types.NewStringSized(types.NewTypeSizeMax()))),
		},
		{
			Message: `with spaces`,
			Input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenKeyword, Content: []rune(`ARRAY`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`<`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`>`)},
				{Kind: tokenize.TokenSpace, Content: []rune(` `)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			WantNode: types.NewStructField(node.NewIdentifier(`abc`), types.NewArray(types.NewStringSized(types.NewTypeSizeMax()))),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.Message), func(t *testing.T) {
			test.TestParse(t, testcase, perser_types.ParseStructField)
		})
	}
}
