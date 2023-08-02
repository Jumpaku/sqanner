package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
	"testing"
)

type testcase[T node.Node] struct {
	message   string
	input     []tokenize.Token
	wantNode  T
	shouldErr bool
}

func (tc testcase[T]) MessageWithInput() string {
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

func (e nodeArrayElement) pathText() string {
	var path []string
	for _, n := range e.path {
		path = append(path, strconv.FormatInt(int64(n), 10))
	}
	return strings.Join(path, "-")
}

func (e nodeArrayElement) tokenText(input []tokenize.Token) string {
	tokenSText := `[]string{`
	for _, token := range input[e.node.Begin():e.node.End()] {
		tokenSText += fmt.Sprintf(`%q,`, string(token.Content))
	}
	return tokenSText + `}`
}

func nodeArray(root node.Node, path []int) []nodeArrayElement {
	arr := []nodeArrayElement{{path, root}}
	for i, ch := range root.Children() {
		chPath := append(append([]int{}, path...), i)
		arr = append(arr, nodeArray(ch, chPath)...)
	}
	return arr
}

func testParse[T node.Node](t *testing.T, testcase testcase[T], sut func(s *parser.ParseState) (T, error)) {
	t.Helper()

	s := parser.NewParseState(testcase.input)
	gotNode, gotErr := sut(s)

	if (gotErr != nil) != testcase.shouldErr {
		if testcase.shouldErr {
			t.Errorf("%s:\n	err is expected but got nil", testcase.MessageWithInput())
		} else {
			t.Errorf("%s:\n	err is not expected but got %v", testcase.MessageWithInput(), gotErr)
		}
	}

	gotNodeArr := nodeArray(gotNode, []int{})
	wantNodeArr := nodeArray(testcase.wantNode, []int{})

	size := len(gotNodeArr)
	if size < len(wantNodeArr) {
		size = len(wantNodeArr)
	}

	diff := ``
	for i := 0; i < size; i++ {
		s := fmt.Sprintf(`  node[%d]:  `, i)
		if i < len(wantNodeArr) {
			want := wantNodeArr[i]
			s += fmt.Sprintf(`want=%10s:%-17s`, want.pathText(), want.node.Kind().String()[4:])
		} else {
			s += fmt.Sprintf(`want=%10s:%-17s`, ``, `(nothing)`)
		}

		if i < len(gotNodeArr) {
			got := gotNodeArr[i]
			s += fmt.Sprintf(`got=%10s:%17s(%s)`, got.pathText(), got.node.Kind().String()[4:], got.pathText())
		} else {
			s += fmt.Sprintf(`got=%10s:(nothing)`, ``)
		}

		diff += s + "\n"
	}
	if !nodeMatch(testcase.wantNode, gotNode) {
		t.Errorf("%s\n%s", testcase.MessageWithInput(), diff)
	}
}

func nodeMatch(want node.Node, got node.Node) bool {
	if want.Kind() != got.Kind() {
		return false
	}

	switch want.Kind() {
	default:
		panic(fmt.Sprintf(`unsupported node kind: %v`, want.Kind()))
	case node.NodeIdentifier:
		w := want.(node.IdentifierNode)
		g := got.(node.IdentifierNode)
		return w.Value() == g.Value() && slices.EqualFunc(w.Children(), g.Children(), nodeMatch)
	case node.NodePath:
		w := want.(node.PathNode)
		g := got.(node.PathNode)
		return slices.EqualFunc(w.Identifiers(), g.Identifiers(), func(w, g node.IdentifierNode) bool { return nodeMatch(w, g) })
	case node.NodeStructTypeField:
		w := want.(node.StructTypeFieldNode)
		g := got.(node.StructTypeFieldNode)
		return nodeMatch(w.Name(), g.Name()) && nodeMatch(w.Type(), g.Type())
	case node.NodeType:
		w := want.(node.TypeNode)
		g := got.(node.TypeNode)
		ok := w.TypeCode() == g.TypeCode() &&
			w.IsScalar() == g.IsScalar() &&
			w.IsStruct() == g.IsStruct() &&
			w.IsArray() == g.IsArray()
		switch {
		default:
			panic(`invalid type`)
		case w.IsScalar():
			ok = ok && w.ScalarName() == g.ScalarName() && w.ScalarHasSize() == g.ScalarHasSize()
			if w.ScalarHasSize() {
				ok = ok && nodeMatch(w.ScalarSize(), g.ScalarSize())
			}
			return ok
		case w.IsArray():
			return ok && nodeMatch(w.ArrayElement(), g.ArrayElement())
		case w.IsStruct():
			return ok && slices.EqualFunc(w.StructFields(), g.StructFields(), func(w, g node.StructTypeFieldNode) bool { return nodeMatch(w, g) })
		}
	case node.NodeTypeSize:
		w := want.(node.TypeSizeNode)
		g := got.(node.TypeSizeNode)
		ok := w.Max() == g.Max()
		if !w.Max() {
			ok = ok && w.Size() == g.Size()
		}
		return ok
	}
}

func TestPrintTokens(t *testing.T) {
	input := "_5abc.dataField"

	tokens, err := tokenize.Tokenize([]rune(input))
	if err != nil {
		t.Fatal(err)
	}
	out := fmt.Sprintf("{\n")
	out += "\tinput: []tokenize.Token{\n"
	for _, token := range tokens {
		out += fmt.Sprintf("\t\t{Kind: tokenize.%s, Content:[]rune(%q)},\n", token.Kind.String(), string(token.Content))
	}
	out += "\t},\n"
	out += fmt.Sprintf("},\n")
	fmt.Printf(`%s`, out)
}
