package lexer

import (
	"bufio"
	"os"
)

type TokenID int

const (
	TokenERR TokenID = iota
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
	ReadTill(func(rune) bool) string
	Lexed(*Token)
}

type Token struct {
	ID    TokenID
	Value string
}

type FileLex struct {
	Filename string
	Tokens   []*Token

	reader *bufio.Reader
}

func NewFileLex(filename string) (*FileLex, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &FileLex{
		Filename: filename,
		Tokens:   make([]*Token, 0),
		reader:   bufio.NewReader(fileReader),
	}, nil
}

func (f *FileLex) Peek() (rune, error) {
	chars, err := f.reader.Peek(1)
	if err != nil {
		return 'a', err
	}

	return rune(chars[0]), nil
}

func (f *FileLex) Read() (rune, error) {
	char, err := f.reader.ReadByte()
	if err != nil {
		return 'a', err
	}

	return rune(char), nil
}

func (f *FileLex) Lexed(t *Token) {
	f.Tokens = append(f.Tokens, t)
}

func (f *FileLex) ReadTill(matcher func(rune) bool) (accum string) {
	for {
		char, err := f.Read()
		if err != nil || !matcher(char) {
			break
		}

		accum += string(char)
	}

	return accum
}
