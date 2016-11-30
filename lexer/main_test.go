package main

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexers(t *testing.T) {
	Convey("CommentLexer", t, func() {
		Convey("It matches only if char is '/'", func() {
			So(NewCommentLexer().Match(byte('/')), ShouldBeTrue)
			So(NewCommentLexer().Match(byte('1')), ShouldBeFalse)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewCommentLexer()
		})
	})

	Convey("WordLexer", t, func() {
		Convey("It matches only if char is a unicode character", func() {
			So(NewWordLexer().Match(byte('a')), ShouldBeTrue)
			So(NewWordLexer().Match(byte('/')), ShouldBeFalse)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewWordLexer()
		})
	})

	Convey("StartLexer", t, func() {
		Convey("It matches '/'", func() {
			So(NewStartLexer().Match(byte('/')), ShouldBeTrue)
		})

		Convey("It maches if char is a unicode character", func() {
			So(NewStartLexer().Match(byte('a')), ShouldBeTrue)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewStartLexer()
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

		Convey("It peeks 1st character in file", func() {
			f, err := NewFileLexer(tmpFile.Name())
			So(err, ShouldBeNil)

			char, err := f.Peek()
			So(err, ShouldBeNil)
			So(char, ShouldEqual, byte('h'))
		})
	})
}
