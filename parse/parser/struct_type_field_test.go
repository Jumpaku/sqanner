package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParseStructTypeField(t *testing.T) {
	testcases := []testcase[node.StructTypeFieldNode]{
		{
			message: `unnamed field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`INT64`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeFieldUnnamed(nodeOf(node.Int64Type()))),
		},
		{
			message: `unnamed field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeFieldUnnamed(nodeOf(node.StringType()))),
		},
		{
			message: `unnamed field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeFieldUnnamed(nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))),
		},
		{
			message: `unnamed field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenKeyword, Content: []rune(`ARRAY`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`<`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`>`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeFieldUnnamed(nodeOf(node.ArrayType(nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))))),
		},
		{
			message: `named field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`INT64`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeField(nodeOf(node.Identifier(`abc`)), nodeOf(node.Int64Type()))),
		},
		{
			message: `named field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeField(nodeOf(node.Identifier(`abc`)), nodeOf(node.StringType()))),
		},
		{
			message: `named field`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenIdentifier, Content: []rune(`abc`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`STRING`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`(`)},
				{Kind: tokenize.TokenIdentifier, Content: []rune(`MAX`)},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(`)`)},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.StructTypeField(nodeOf(node.Identifier(`abc`)), nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))),
		},
		{
			message: `named field`,
			input: []tokenize.Token{
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
			wantNode: nodeOf(node.StructTypeField(nodeOf(node.Identifier(`abc`)), nodeOf(node.ArrayType(nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))))),
		},
		{
			message: `with spaces`,
			input: []tokenize.Token{
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
			wantNode: nodeOf(node.StructTypeField(nodeOf(node.Identifier(`abc`)), nodeOf(node.ArrayType(nodeOf(node.StringTypeSized(nodeOf(node.TypeSizeMax()))))))),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParseStructField)
		})
	}
}
