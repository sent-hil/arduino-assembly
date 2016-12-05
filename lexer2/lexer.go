package lexer

import (
	"bufio"
	"io"
)

type TokenID int

const (
	TokenERR TokenID = iota
	TokenEOF
	TokenInteger
	TokenComment
)

type Token struct {
	ID    TokenID
	Value string
}

type Lexable interface {
	Match(Peekable) bool
	Lex(Readable) *Token
}

type Peekable interface {
	Peek(int) ([]rune, error) // will return error if unable to peek given size
}

type Readable interface {
	Read(func(rune) bool) string
}

type Lexer struct {
	runes []rune
}

func NewLexer(reader io.Reader) *Lexer {
	runes := []rune{}
	bufReader := bufio.NewReader(reader)

	for runeChar, _, err := bufReader.ReadRune(); ; {
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}

		runes = append(runes, runeChar)
	}

	return &Lexer{runes: runes}
}

func (l *Lexer) Peek(size int) (chars []rune, err error) {
	return l.runes[:size], nil
}

func (l *Lexer) Read(fn func(rune) bool) (accum string) {
	for i, runeChar := range l.runes {
		if !fn(runeChar) {
			break
		}

		accum += string(runeChar)
		l.runes = l.runes[i:]
	}

	return accum
}
