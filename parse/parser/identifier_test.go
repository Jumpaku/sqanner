package parser_test

import (
	"fmt"
	"github.com/Jumpaku/sqanner/tokenize"
	"testing"
)

func TestDebugParser(t *testing.T) {
	input := ""
	tokens, err := tokenize.Tokenize([]rune(input))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf(`%#v`, tokens)

}
