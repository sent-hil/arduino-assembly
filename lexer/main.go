package main

import (
	"bufio"
	"io"
	"os"
)

type TokenID string

const (
	TokenERR     TokenID = "ERR"
	TokenEOF             = "EOF"
	TokenComment         = "COMMENT"
	TokenString          = "STRING"
	TokenInteger         = "INT"
)

type Peeker interface {
	Peek() (byte, error)
	PeekAt(i int) (byte, error)
}

type Lexer interface {
	Match(byte) bool
	Lex(Peeker) (*Token, Lexer)
}

type Token struct {
	ID    TokenID
	Value string
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

func (s *StartLexer) Lex(p Peeker) (*Token, Lexer) {
	c, err := p.Peek()
	if err != nil && err == io.EOF {
		return nil, nil
	}

	if err != nil {
		panic(err)
	}

	switch {
	case NewCommentLexer().Match(c):
		return nil, NewCommentLexer()
	case NewWordLexer().Match(c):
		return nil, NewWordLexer()
	}

	return nil, nil
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
	chars, err := f.reader.Peek(1)
	if err != nil {
		return byte('"'), err
	}

	return chars[0], nil
}

func (f *FileLexer) PeekAt(i int) (byte, error) {
	chars, err := f.reader.Peek(i)
	if err != nil {
		return byte('"'), err
	}

	if len(chars) == 0 {
		return byte('"'), io.EOF
	}

	return chars[i-1:][0], nil
}

func (f *FileLexer) Lex() {
	//var currentLexer Lexer = NewStartLexer()

	//for _, err := f.Peek(); err == nil && currentLexer != nil; {
	//  nextLexer := currentLexer.Lex(f)

	//  if nextLexer != nil {
	//    currentLexer = nextLexer
	//  }
	//}
}
