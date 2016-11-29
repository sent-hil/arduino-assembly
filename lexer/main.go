package main

import (
	"bufio"
	"os"
	"unicode"
)

type TokenId int

const (
	TokenERR TokenId = iota
	TokenEOF
	TokenComment
	TokenString
)

type Peeker interface {
	Peek() (byte, error)
}

type Lexer interface {
	Match(byte) bool
	Lex(Peeker) Lexer
}

type Token struct {
	Id    TokenId
	Value string
}

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

func (c *CommentLexer) Match(char byte) bool {
	return char == byte('/')
}

func (c *CommentLexer) Lex(p Peeker) {
}

type WordLexer struct{}

func NewWordLexer() *WordLexer {
	return &WordLexer{}
}

func (w *WordLexer) Match(char byte) bool {
	return unicode.IsLetter(rune(char))
}

type StartLexer struct {
	Comment *CommentLexer
	Word    *WordLexer
}

func NewStartLexer() *StartLexer {
	return &StartLexer{
		Comment: NewCommentLexer(),
		Word:    NewWordLexer(),
	}
}

func (s *StartLexer) Match(char byte) bool {
	return s.Comment.Match(char) || s.Word.Match(char)
}

func (s *StartLexer) Lex(p Peeker) Lexer {
	return nil
}

type FileLexer struct {
	Filename string
	reader   *bufio.Reader
}

func NewFileLexer(filename string) (*FileLexer, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &FileLexer{
		Filename: filename,
		reader:   bufio.NewReader(fileReader),
	}, nil
}

func (f *FileLexer) Peek() (byte, error) {
	return byte('a'), nil
}

func (f *FileLexer) Lex() {
	var currentLexer Lexer = NewStartLexer()

	for _, err := f.Peek(); err == nil && currentLexer != nil; {
		nextLexer := currentLexer.Lex(f)

		if nextLexer != nil {
			currentLexer = nextLexer
		}
	}
}
