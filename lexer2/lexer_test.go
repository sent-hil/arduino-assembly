package lexer

import "bytes"

func newStringLexer(s string) *Lexer {
	return NewLexer(bytes.NewBufferString(s))
}
