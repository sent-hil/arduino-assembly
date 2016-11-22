package main

import (
	"bufio"
	"os"
	"unicode"
)

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

func (c *CommentLexer) Match(char byte) bool {
	return char == byte('/')
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
