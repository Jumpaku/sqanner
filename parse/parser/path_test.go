package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestParsePath(t *testing.T) {
	testcases := []testcase[node.PathNode]{
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. _5abc and dataField are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("_5abc")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("_5abc")),
				nodeOf(node.Identifier("dataField")),
			})),
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. `5abc` and dataField are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`5abc`")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("`5abc`")),
				nodeOf(node.Identifier("dataField")),
			})),
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. `5abc` and dataField are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`5abc`")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("`dataField`")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("`5abc`")),
				nodeOf(node.Identifier("`dataField`")),
			})),
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. abc5 and dataField are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("abc5")),
				nodeOf(node.Identifier("dataField")),
			})),
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Invalid. abc5! is an invalid identifier because it is unquoted and contains\n")},
				{Kind: tokenize.TokenComment, Content: []rune("-- a character that is not a letter, number, or underscore.\n")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("!")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("abc5")),
			})),
		},
		{
			message: `invalid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Invalid. abc5! is an invalid identifier because it is unquoted and contains\n")},
				{Kind: tokenize.TokenComment, Content: []rune("-- a character that is not a letter, number, or underscore.\n")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune("!")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. `GROUP` and dataField are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifierQuoted, Content: []rune("`GROUP`")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("`GROUP`")),
				nodeOf(node.Identifier("dataField")),
			})),
		},
		{
			message: `invalid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Invalid. GROUP is an invalid identifier because it is unquoted and is a\n")},
				{Kind: tokenize.TokenComment, Content: []rune("-- stand-alone reserved keyword.\n")},
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("dataField")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			shouldErr: true,
		},
		{
			message: `valid path`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenComment, Content: []rune("-- Valid. abc5 and GROUP are valid identifiers.\n")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("abc5")),
				nodeOf(node.Identifier("GROUP")),
			})),
		},
		{
			message: `valid path with spaces`,
			input: []tokenize.Token{
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenIdentifier, Content: []rune("abc5")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenSpecialChar, Content: []rune(".")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenKeyword, Content: []rune("GROUP")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenComment, Content: []rune("/* comment */")},
				{Kind: tokenize.TokenSpace, Content: []rune(" ")},
				{Kind: tokenize.TokenEOF, Content: []rune("")},
			},
			wantNode: nodeOf(node.Path([]node.IdentifierNode{
				nodeOf(node.Identifier("abc5")),
				nodeOf(node.Identifier("GROUP")),
			})),
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf(`case[%d]:%s`, i, testcase.message), func(t *testing.T) {
			testParse(t, testcase, parser.ParsePath)
		})
	}
}
