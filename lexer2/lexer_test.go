package lexer

import "strings"

func newStringLexer(s string) *Lexer {
	return NewLexer(strings.NewReader(s))
}
