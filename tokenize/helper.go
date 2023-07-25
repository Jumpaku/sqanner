package tokenize

import "strings"

func isOctalDigit(r rune) bool {
	return '0' <= r && r <= '7'
}
func isHexDigit(r rune) bool {
	return ('0' <= r && r <= '7') || ('A' <= r && r <= 'F') || ('a' <= r && r <= 'f')
}

// isLetter returns true if r is an alphabet or underscore; otherwise false
func isLetter(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || (r == '_')
}
func isDecimalDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

var keywords = map[string]bool{
	"ALL":                  true,
	"AND":                  true,
	"ANY":                  true,
	"ARRAY":                true,
	"AS":                   true,
	"ASC":                  true,
	"ASSERT_ROWS_MODIFIED": true,
	"AT":                   true,
	"BETWEEN":              true,
	"BY":                   true,
	"CASE":                 true,
	"CAST":                 true,
	"COLLATE":              true,
	"CONTAINS":             true,
	"CREATE":               true,
	"CROSS":                true,
	"CUBE":                 true,
	"CURRENT":              true,
	"DEFAULT":              true,
	"DEFINE":               true,
	"DESC":                 true,
	"DISTINCT":             true,
	"ELSE":                 true,
	"END":                  true,
	"ENUM":                 true,
	"ESCAPE":               true,
	"EXCEPT":               true,
	"EXCLUDE":              true,
	"EXISTS":               true,
	"EXTRACT":              true,
	"FALSE":                true,
	"FETCH":                true,
	"FOLLOWING":            true,
	"FOR":                  true,
	"FROM":                 true,
	"FULL":                 true,
	"GROUP":                true,
	"GROUPING":             true,
	"GROUPS":               true,
	"HASH":                 true,
	"HAVING":               true,
	"IF":                   true,
	"IGNORE":               true,
	"IN":                   true,
	"INNER":                true,
	"INTERSECT":            true,
	"INTERVAL":             true,
	"INTO":                 true,
	"IS":                   true,
	"JOIN":                 true,
	"LATERAL":              true,
	"LEFT":                 true,
	"LIKE":                 true,
	"LIMIT":                true,
	"LOOKUP":               true,
	"MERGE":                true,
	"NATURAL":              true,
	"NEW":                  true,
	"NO":                   true,
	"NOT":                  true,
	"NULL":                 true,
	"NULLS":                true,
	"OF":                   true,
	"ON":                   true,
	"OR":                   true,
	"ORDER":                true,
	"OUTER":                true,
	"OVER":                 true,
	"PARTITION":            true,
	"PRECEDING":            true,
	"PROTO":                true,
	"RANGE":                true,
	"RECURSIVE":            true,
	"RESPECT":              true,
	"RIGHT":                true,
	"ROLLUP":               true,
	"ROWS":                 true,
	"SELECT":               true,
	"SET":                  true,
	"SOME":                 true,
	"STRUCT":               true,
	"TABLESAMPLE":          true,
	"THEN":                 true,
	"TO":                   true,
	"TREAT":                true,
	"TRUE":                 true,
	"UNBOUNDED":            true,
	"UNION":                true,
	"UNNEST":               true,
	"USING":                true,
	"WHEN":                 true,
	"WHERE":                true,
	"WINDOW":               true,
	"WITH":                 true,
	"WITHIN":               true,

	"NUMERIC":   true,
	"DATE":      true,
	"TIMESTAMP": true,
	"JSON":      true,
}

func isKeyword(r []rune) bool {
	return keywords[strings.ToUpper(string(r))]
}

var specialCharTable = func() []bool {
	specialChars := "@,()[]{}<>.;:/+-~*|&^=!" + "$?"

	isSpecial := make([]bool, 127)
	for _, c := range []rune(specialChars) {
		isSpecial[c] = true
	}

	return isSpecial
}()

func isSpecialChar(r rune) bool {
	return specialCharTable[r]
}
