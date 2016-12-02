package main

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexers(t *testing.T) {
	Convey("StartLexer", t, func() {
		Convey("It matches '/'", func() {
			So(NewStartLexer().Match(byte('/')), ShouldBeTrue)
		})

		Convey("It maches if first char is unicode", func() {
			So(NewStartLexer().Match(byte('a')), ShouldBeTrue)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewStartLexer()
		})

		Convey("It returns comment lexer if first char is a '/'", func() {
			f := newFileLexer("//")
			_, s := NewStartLexer().Lex(f)

			_, ok := s.(*CommentLexer)
			So(ok, ShouldBeTrue)
		})

		Convey("It returns word lexer if first char is unicode", func() {
			f := newFileLexer("world")
			_, s := NewStartLexer().Lex(f)

			_, ok := s.(*WordLexer)
			So(ok, ShouldBeTrue)
		})

		Convey("It returns nil lexer if first char is not a comment or word", func() {
			f := newFileLexer(" ")
			_, s := NewStartLexer().Lex(f)

			So(s, ShouldBeNil)
		})
	})

	Convey("FileLexer", t, func() {
		tmpFile, err := ioutil.TempFile("", "")
		So(err, ShouldBeNil)

		_, err = tmpFile.Write([]byte("hello world"))
		So(err, ShouldBeNil)

		Convey("It should initialize with given filename", func() {
			_, err = NewFileLexer(tmpFile.Name())
			So(err, ShouldBeNil)
		})

		Convey("It peeks first char in file", func() {
			f, err := NewFileLexer(tmpFile.Name())
			So(err, ShouldBeNil)

			char, err := f.Peek()
			So(err, ShouldBeNil)
			So(char, ShouldEqual, byte('h'))
		})
	})
}

func newFileLexer(s string) *FileLexer {
	tmpFile, err := ioutil.TempFile("", "")
	So(err, ShouldBeNil)

	_, err = tmpFile.Write([]byte(s))
	So(err, ShouldBeNil)

	f, err := NewFileLexer(tmpFile.Name())
	So(err, ShouldBeNil)

	return f
}
