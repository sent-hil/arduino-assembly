package lexer

import "unicode"

type IntegerLexer struct{}

func NewIntegerLexer() *IntegerLexer {
	return &IntegerLexer{}
}

func (i *IntegerLexer) Match(p Peekable) bool {
	chars, err := p.Peek(1)
	if err != nil {
		return false
	}

	return unicode.IsDigit(chars[0])
}

func (i *IntegerLexer) Lex(r Readable) *Token {
	accum := r.Read(func(c rune) bool {
		return unicode.IsDigit(c)
	})

	return &Token{ID: TokenInteger, Value: accum}
}

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

func (i *CommentLexer) Match(p Peekable) bool {
	chars, err := p.Peek(2)
	if err != nil {
		return false
	}

	return chars[0] == '/' && chars[1] == '/'
}

func (i *CommentLexer) Lex(r Readable) *Token {
	accum := r.Read(func(c rune) bool {
		return c != '\n'
	})

	return &Token{ID: TokenComment, Value: accum}
}
