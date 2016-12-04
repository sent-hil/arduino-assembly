package lexer

import "unicode"

type IntegerLex struct{}

func NewIntegerLex() *IntegerLex {
	return &IntegerLex{}
}

func (i *IntegerLex) Match(char rune) bool {
	return unicode.IsDigit(char)
}

func (i *IntegerLex) Run(l Lexable) Lexer {
	l.Lexed(&Token{
		ID:    TokenInteger,
		Value: l.ReadTill(i.Match),
	})

	return nil
}
