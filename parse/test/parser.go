package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type testNodeDiffLine struct {
	wantPath  string
	gotPath   string
	wantValue string
	gotValue  string
}

func TestParse[T node.Node](t *testing.T, testcase Case[T], sut func(s *parse.ParseState) (T, error)) {
	t.Helper()

	gotNode, gotErr := sut(parse.NewParseState(testcase.Input))

	if testcase.ShouldErr {
		assert.NotNilf(t, gotErr, "%s:\n	err is expected but got nil", failMessage(testcase.Message, testcase.Input, testcase.WantNode, gotNode))
		return
	} else {
		assert.Nilf(t, gotErr, "%s:\n	nil is expected but got err", failMessage(testcase.Message, testcase.Input, testcase.WantNode, gotNode))
	}

	assertNodeMatch(t, testcase.WantNode, gotNode, failMessage(testcase.Message, testcase.Input, testcase.WantNode, gotNode))
}

func failMessage[T node.Node](msg string, input []tokenize.Token, want, got T) string {

	wantNodeArr := nodeArray(want, []int{0})
	gotNodeArr := nodeArray(got, []int{0})

	nDiffLines := max(len(gotNodeArr), len(wantNodeArr))

	diffLines := []testNodeDiffLine{}
	for line := 0; line < nDiffLines; line++ {
		diffLine := testNodeDiffLine{}
		if line < len(wantNodeArr) {
			want := wantNodeArr[line]
			diffLine.wantPath = want.treePathText()
			diffLine.wantValue = fmt.Sprintf("%T", want.node)
		}

		if line < len(gotNodeArr) {
			got := gotNodeArr[line]
			diffLine.gotPath = got.treePathText()
			diffLine.gotValue = fmt.Sprintf("%T", got.node)
		}

		diffLines = append(diffLines, diffLine)
	}

	diff := ``
	{
		msg := fmt.Sprintf("%s: input=", msg)
		for _, token := range input {
			msg += fmt.Sprintf("[%q]", string(token.Content))
		}

		diff += msg
	}
	diff += "\n"

	wantPathLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.wantPath) }))
	wantValueLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.wantValue) }))
	gotPathLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.gotPath) }))
	gotValueLen := lo.Max(lo.Map(diffLines, func(item testNodeDiffLine, i int) int { return len(item.gotValue) }))
	for line, diffLine := range diffLines {
		lineFormat := fmt.Sprintf("node[%-2d]| %%-%ds:%%-%ds | %%-%ds:%%-%ds\n", line, wantPathLen, wantValueLen, gotPathLen, gotValueLen)
		diff += fmt.Sprintf(lineFormat, diffLine.wantPath, diffLine.wantValue, diffLine.gotPath, diffLine.gotValue)
	}

	return diff
}

type Case[T node.Node] struct {
	Message   string
	Input     []tokenize.Token
	WantNode  T
	ShouldErr bool
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

func assertNodeMatch[T node.Node](t *testing.T, want T, got T, failMsg string) {
	var v any = want
	switch v.(type) {
	default:
		t.Fatal(fmt.Sprintf("unsupported node type: %T: \n%s", v, failMsg))
	case node.IdentifierNode:
		w, g := cast[node.IdentifierNode](want, got)
		assert.Equal(t, w.Value(), g.Value(), failMsg)
		assert.Equal(t, w.RequiresQuotes(), g.RequiresQuotes(), failMsg)
		assert.Equal(t, w.QuotedValue(), g.QuotedValue(), failMsg)
	case types.TypeNode:
		w, g := cast[types.TypeNode](want, got)
		assert.Equal(t, w.TypeCode(), g.TypeCode(), failMsg)
		assert.Equal(t, w.IsScalar(), g.IsScalar(), failMsg)
		assert.Equal(t, w.IsStruct(), g.IsStruct(), failMsg)
		assert.Equal(t, w.IsArray(), g.IsArray(), failMsg)
		switch {
		default:
			panic("invalid type")
		case w.IsScalar():
			assert.Equal(t, w.ScalarSized(), g.ScalarSized(), failMsg)
			if w.ScalarSized() {
				assertNodeMatch(t, w.ScalarSize(), g.ScalarSize(), failMsg)
			}
		case w.IsStruct():
			assert.Equal(t, len(w.StructFields()), len(g.StructFields()), failMsg)
			for i := 0; i < len(w.StructFields()); i++ {
				assertNodeMatch(t, w.StructFields()[i], g.StructFields()[i], failMsg)
			}
		case w.IsArray():
			assertNodeMatch(t, w.ArrayElement(), g.ArrayElement(), failMsg)
		}
	case types.TypeSizeNode:
		w, g := cast[types.TypeSizeNode](want, got)
		assert.Equal(t, w.Max(), g.Max(), failMsg)
		if !w.Max() {
			assert.Equal(t, w.Size(), g.Size(), failMsg)
		}
	case types.StructFieldNode:
		w, g := cast[types.StructFieldNode](want, got)
		assert.Equal(t, w.Named(), g.Named(), failMsg)
		if w.Named() {
			assertNodeMatch(t, w.Name(), g.Name(), failMsg)
		}
		assertNodeMatch(t, w.Type(), g.Type(), failMsg)
	case node.KeywordNode:
		w, g := cast[node.KeywordNode](want, got)
		assert.Equal(t, w.KeywordCode(), g.KeywordCode(), failMsg)
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
