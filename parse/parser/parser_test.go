package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type testNodeDiffLine struct {
	wantPath  string
	gotPath   string
	wantValue string
	gotValue  string
}

func testParse[T node.Node](t *testing.T, testcase testcase[T], sut func(s *parse.ParseState) (T, error)) {
	t.Helper()

	gotNode, gotErr := sut(parse.NewParseState(testcase.input))

	if testcase.shouldErr {
		assert.NotNilf(t, gotErr, "%s:\n	err is expected but got nil", testcase.messageWithInput())
		return
	} else {
		assert.Nilf(t, gotErr, "%s:\n	nil is expected but got err", testcase.messageWithInput())
	}

	wantNodeArr := nodeArray(testcase.wantNode, []int{0})
	gotNodeArr := nodeArray(gotNode, []int{0})

	nDiffLines := max(len(gotNodeArr), len(wantNodeArr))

	diffLines := []testNodeDiffLine{}
	for line := 0; line < nDiffLines; line++ {
		diffLine := testNodeDiffLine{}
		if line < len(wantNodeArr) {
			want := wantNodeArr[line]
			diffLine.wantPath = want.treePathText()
			diffLine.wantValue = reflect.TypeOf(want.node).Name()
		}

		if line < len(gotNodeArr) {
			got := gotNodeArr[line]
			diffLine.gotPath = got.treePathText()
			diffLine.gotValue = reflect.TypeOf(got.node).Name()
		}
	}

	diff := ``
	wantPathLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.wantPath) }))
	wantValueLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.wantValue) }))
	gotPathLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.gotPath) }))
	gotValueLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.gotValue) }))
	for line, diffLine := range diffLines {
		lineFormat := fmt.Sprintf("node[%-2d]: %%%ds %%%ds %%-%ds %%-%ds\n", line, wantPathLen, wantValueLen, gotPathLen, gotValueLen)
		diff += fmt.Sprintf(lineFormat, diffLine.wantPath, diffLine.wantValue, diffLine.gotPath, diffLine.gotValue)
	}

	assert.Truef(t, matchNode(testcase.wantNode, gotNode), "%s\n%s", testcase.messageWithInput(), diff)
}

type testcase[T node.Node] struct {
	message   string
	input     []tokenize.Token
	wantNode  T
	shouldErr bool
}

func (tc testcase[T]) messageWithInput() string {
	msg := fmt.Sprintf(`%s: input=`, tc.message)
	for _, token := range tc.input {
		msg += fmt.Sprintf("[%q]", string(token.Content))
	}

	return msg
}

type nodeArrayElement struct {
	path []int
	node node.Node
}

func (e nodeArrayElement) treePathText() string {
	path := lo.Map(e.path, func(n int, index int) string {
		return strconv.FormatInt(int64(n), 10)
	})

	return strings.Join(path, "-")
}

func (e nodeArrayElement) tokenText(input []tokenize.Token) string {
	tokensText := ``
	for _, token := range input[e.node.Begin():e.node.End()] {
		tokensText += fmt.Sprintf(`[%q]`, string(token.Content))
	}
	return tokensText + ``
}

func nodeArray(root node.Node, path []int) []nodeArrayElement {
	if root == nil {
		return nil
	}
	arr := []nodeArrayElement{{path, root}}
	for i, ch := range root.Children() {
		chPath := append(append([]int{}, path...), i)
		arr = append(arr, nodeArray(ch, chPath)...)
	}
	return arr
}

func cast[T node.Node](want, got any) (T, T) {
	return want.(T), got.(T)
}
func matchNode[T node.Node](want T, got T) bool {
	var v any = want
	switch v.(type) {
	default:
		panic(fmt.Sprintf("unsupported node type: %T", v))
	case *node.IdentifierNode:
		w, g := cast[*node.IdentifierNode](want, got)
		return w.Value() == g.Value() && w.IsQuoted() == g.IsQuoted() && w.UnquotedValue() == g.UnquotedValue()
	}
}

func TestPrintTokens(t *testing.T) {
	input := "ARRAY<STRUCT<ARRAY<INT64>>>"

	tokens, err := tokenize.Tokenize([]rune(input))
	if err != nil {
		t.Fatal(err)
	}
	out := fmt.Sprintf("{\n")
	out += "\tmessage: ``,\n"
	out += "\tinput: []tokenize.Token{\n"
	for _, token := range tokens {
		out += fmt.Sprintf("\t\t{Kind: tokenize.%s, Content:[]rune(%q)},\n", token.Kind.String(), string(token.Content))
	}
	out += "\t},\n"
	out += "\twantNode: nil,\n"
	out += "\tshouldErr: false,\n"
	out += fmt.Sprintf("},\n")
	fmt.Printf(`%s`, out)
}

func TestDebugParse(t *testing.T) {
	input := []tokenize.Token{
		{Kind: tokenize.TokenIdentifier, Content: []rune("ARRAY")},
		{Kind: tokenize.TokenSpecialChar, Content: []rune("<")},
		{Kind: tokenize.TokenIdentifier, Content: []rune("STRING")},
		{Kind: tokenize.TokenSpecialChar, Content: []rune("(")},
		{Kind: tokenize.TokenIdentifier, Content: []rune("MAX")},
		{Kind: tokenize.TokenSpecialChar, Content: []rune(")")},
		{Kind: tokenize.TokenSpecialChar, Content: []rune(">")},
		{Kind: tokenize.TokenEOF, Content: []rune("")},
	}
	n, err := parser.ParseIdentifier(parse.NewParseState(input))
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(n)
}
