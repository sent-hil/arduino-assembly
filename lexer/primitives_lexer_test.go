package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimitiveLexers(t *testing.T) {
	Convey("CommentLexer", t, func() {
		Convey("It matches only if first char is '/'", func() {
			So(NewCommentLexer().Match(byte('/')), ShouldBeTrue)
			So(NewCommentLexer().Match(byte('1')), ShouldBeFalse)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewCommentLexer()
		})
	})

	Convey("WordLexer", t, func() {
		Convey("It matches only if char is a unicode char", func() {
			So(NewWordLexer().Match(byte('a')), ShouldBeTrue)
			So(NewWordLexer().Match(byte('/')), ShouldBeFalse)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewWordLexer()
		})
	})

	Convey("IntegerLexer", t, func() {
		Convey("It matches only if first char is an integer", func() {
			So(NewIntegerLexer().Match(byte('1')), ShouldBeTrue)
		})

		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewIntegerLexer()
		})

		Convey("It parses every digit till break", func() {
			f := newFileLexer("1000\n")
			token, _ := NewIntegerLexer().Lex(f)

			So(token, ShouldNotBeNil)
			So(token.ID, ShouldEqual, TokenInteger)
			So(token.Value, ShouldEqual, "1000")
		})
	})
}
