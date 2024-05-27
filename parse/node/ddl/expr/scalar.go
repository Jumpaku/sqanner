package expr

import (
	"cloud.google.com/go/civil"
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"github.com/samber/lo"
	"math/big"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type ScalarNode interface {
	node.Node
	Type() types.TypeCode
	Literal() string
	UnquotedStringValue() string
}

func AcceptInt64(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeInt64,
		literalString: literalString,
	}
}

func AcceptFloat64(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeFloat64,
		literalString: literalString,
	}
}

func AcceptBool(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeBool,
		literalString: strings.ToUpper(literalString),
	}
}

func AcceptStringOrBytes(s *parse.ParseState, literalString string) *scalarNode {
	_, isBytes := unquote(literalString)
	code := types.TypeCodeString
	if isBytes {
		code = types.TypeCodeBytes
	}
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      code,
		literalString: literalString,
	}
}

func AcceptTimestamp(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeTimestamp,
		literalString: literalString,
	}
}

func AcceptDate(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeDate,
		literalString: literalString,
	}
}

func AcceptJSON(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeJSON,
		literalString: literalString,
	}
}

func AcceptNumeric(s *parse.ParseState, literalString string) *scalarNode {
	return &scalarNode{
		NodeBase:      node.NewNodeBase(s.Begin(), s.End()),
		typeCode:      types.TypeCodeNumeric,
		literalString: literalString,
	}
}

func NewInt64(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeInt64,
		literalString: literalString,
	}
}

func NewFloat64(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeFloat64,
		literalString: literalString,
	}
}

func NewBool(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeBool,
		literalString: literalString,
	}
}

func NewString(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeString,
		literalString: literalString,
	}
}

func NewBytes(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeBytes,
		literalString: literalString,
	}
}

func NewTimestamp(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeTimestamp,
		literalString: literalString,
	}
}

func NewDate(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeDate,
		literalString: literalString,
	}
}

func NewJSON(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeJSON,
		literalString: literalString,
	}
}

func NewNumeric(literalString string) *scalarNode {
	return &scalarNode{
		typeCode:      types.TypeCodeNumeric,
		literalString: literalString,
	}
}

type scalarNode struct {
	node.NodeBase
	typeCode       types.TypeCode
	literalString  string
	int64Value     int64
	float64Value   float64
	boolValue      bool
	stringValue    string
	bytesValue     []byte
	timestampValue time.Time
	dateValue      civil.Date
	jsonValue      any
	numericValue   *big.Rat
}

var _ ScalarNode = (*scalarNode)(nil)

func (n *scalarNode) Children() []node.Node {
	return nil
}

func (n *scalarNode) Type() types.TypeCode {
	return n.typeCode
}

func (n *scalarNode) Literal() string {
	return n.literalString
}

func (n *scalarNode) UnquotedStringValue() string {
	v, _ := unquote(n.literalString)
	return v
}

func unquote(quotedLiteralString string) (content string, isBytes bool) {
	content = quotedLiteralString

	prefix := regexp.MustCompile(`^[rbRB]?[rbRB]?`).FindString(content)
	isRaw := strings.ContainsAny(prefix, `rR`)
	isBytes = strings.ContainsAny(prefix, `bB`)
	content = strings.TrimPrefix(content, prefix)

	switch {
	case regexp.MustCompile(`""".*"""$`).MatchString(content):
		quote := `"""`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`'''.*'''$`).MatchString(content):
		quote := `'''`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`".*"$`).MatchString(content):
		quote := `"`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`'.*'$`).MatchString(content):
		quote := `'`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	}

	if !isRaw {
		content = regexp.MustCompile("\\\\a|\\\\b|\\\\f|\\\\n|\\\\r|\\\\t|\\\\v|\\\\\\\\|\\\\\\?|\\\\\"|\\\\'|\\\\`").ReplaceAllStringFunc(content, unescapeSpecial)
		content = regexp.MustCompile(`\\[0-7]{3}`).ReplaceAllStringFunc(content, unescapeOctal)
		content = regexp.MustCompile(`\\u[0-7a-fA-F]{4}`).ReplaceAllStringFunc(content, unescapeUnicode4)
		content = regexp.MustCompile(`\\U[0-7a-fA-F]{8}`).ReplaceAllStringFunc(content, unescapeUnicode8)
	}

	return content, isBytes
}

func quote(content string, isBytes bool) (quotedLiteralString string) {
	content = regexp.MustCompile("\\a|\\b|\\f|\\n|\\r|\\t|\\v|\\\\|\\?|\"|'|`").ReplaceAllStringFunc(content, escapeSpecial)
	content = strings.Join(lo.Map([]rune(content), func(r rune, idx int) string { return escapeUnicode(r) }), "")
	prefix := ""
	if isBytes {
		prefix = "B"
	}
	quote := `"`
	return prefix + quote + content + quote
}

func unescapeSpecial(escaped string) string {
	unescape := map[string]string{
		`\a`:  "\a",
		`\b`:  "\b",
		`\f`:  "\f",
		`\n`:  "\n",
		`\r`:  "\r",
		`\t`:  "\t",
		`\v`:  "\v",
		`\\`:  "\\",
		`\?`:  "?",
		`\"`:  "\"",
		`\'`:  "'",
		"\\`": "`",
	}
	if unescaped, ok := unescape[escaped]; ok {
		return unescaped
	}
	return escaped
}

func escapeSpecial(unescaped string) string {
	escape := map[string]string{
		"\a": `\a`,
		"\b": `\b`,
		"\f": `\f`,
		"\n": `\n`,
		"\r": `\r`,
		"\t": `\t`,
		"\v": `\v`,
		"\\": `\\`,
		"?":  `\?`,
		"\"": `\"`,
		"'":  `\'`,
		"`":  "\\`",
	}
	if escaped, ok := escape[unescaped]; ok {
		return escaped
	}
	return unescaped
}

func unescapeOctal(escaped string) string {
	runes := []rune(escaped)
	o := 0
	for i := 0; i < 3; i++ {
		o *= 8
		o += int(runes[1+i] - '0')
	}
	return string(rune('\000' + o))
}

func escapeUnicode(unescapedRune rune) string {
	if unicode.IsGraphic(unescapedRune) {
		return string(unescapedRune)
	}
	return fmt.Sprintf(`\U%08X`, unescapedRune)
}

func unescapeUnicode4(escaped string) string {
	runes := []rune(escaped)
	u := 0
	for i := 0; i < 4; i++ {
		u *= 16
		r := runes[2+i]
		switch {
		case '0' <= r && r <= '9':
			u += int(r - '0')
		case 'a' <= r && r <= 'f':
			u += int(r - 'a')
		case 'A' <= r && r <= 'F':
			u += int(r - 'A')
		}
	}
	return string(rune('\u0000' + u))
}

func unescapeUnicode8(escaped string) string {
	runes := []rune(escaped)
	u := 0
	for i := 0; i < 8; i++ {
		u *= 16
		r := runes[2+i]
		switch {
		case '0' <= r && r <= '9':
			u += int(r - '0')
		case 'a' <= r && r <= 'f':
			u += int(r - 'a')
		case 'A' <= r && r <= 'F':
			u += int(r - 'A')
		}
	}
	return string(rune('\U00000000' + u))
}
