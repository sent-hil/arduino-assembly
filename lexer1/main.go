package lexer

import (
	"bufio"
	"os"
	"unicode"
)

type TokenId int

const (
	TokenERR TokenId = iota
	TokenEOF
	TokenInteger
)

type Lexer interface {
	Match(rune) bool
	Run(Lexable) Lexer
}

type Lexable interface {
	Peek() (rune, error)
	Read() (rune, error)
	Lexed(*Token)
}

type Token struct {
	Id    TokenId
	Value string
}

type IntegerLexer struct{}

func NewIntegerLexer() *IntegerLexer {
	return &IntegerLexer{}
}

func (i *IntegerLexer) Match(char rune) bool {
	return unicode.IsDigit(char)
}

func (i *IntegerLexer) Run(p Lexable) Lexer {
	p.Lexed(&Token{})
	return nil
}

type FileLexer struct {
	Filename string
	Tokens   chan *Token

	reader *bufio.Reader
}

func NewFileLexer(filename string) (*FileLexer, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &FileLexer{
		Filename: filename,
		Tokens:   make(chan *Token),
		reader:   bufio.NewReader(fileReader),
	}, nil
}

func (f *FileLexer) Peek() (rune, error) {
	chars, err := f.reader.Peek(1)
	if err != nil {
		return 'a', err
	}

	return rune(chars[0]), nil
}

func (f *FileLexer) Read() (rune, error) {
	char, err := f.reader.ReadByte()
	if err != nil {
		return 'a', err
	}

	return rune(char), nil
}

func (f *FileLexer) Lexed(t *Token) {
	f.Tokens <- t
}
