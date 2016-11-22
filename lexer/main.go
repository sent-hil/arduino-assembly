package main

import "unicode"

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
