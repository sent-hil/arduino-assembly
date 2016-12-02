package main

import "unicode"

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

func (c *CommentLexer) Match(char byte) bool {
	return char == byte('/')
}

func (c *CommentLexer) Lex(p Peeker) (*Token, Lexer) {
	return nil, nil
}

type WordLexer struct{}

func NewWordLexer() *WordLexer {
	return &WordLexer{}
}

func (w *WordLexer) Match(char byte) bool {
	return unicode.IsLetter(rune(char))
}

func (w *WordLexer) Lex(p Peeker) (*Token, Lexer) {
	return nil, nil
}

type IntegerLexer struct{}

func NewIntegerLexer() *IntegerLexer {
	return &IntegerLexer{}
}

func (i *IntegerLexer) Match(char byte) bool {
	return unicode.IsDigit(rune(char))
}

func (i *IntegerLexer) Lex(p Peeker) (*Token, Lexer) {
	var accum string = ""

	for i := 1; ; i++ {
		char, err := p.PeekAt(i)
		if err != nil || !unicode.IsDigit(rune(char)) {
			break
		}

		accum += string(char)
	}

	return &Token{ID: TokenInteger, Value: accum}, nil
}
