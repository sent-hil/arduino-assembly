package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexers(t *testing.T) {
	Convey("CommentLexer", t, func() {
		Convey("It matches only if char is '/'", func() {
			So((NewCommentLexer()).Match(byte('/')), ShouldBeTrue)
			So((NewCommentLexer()).Match(byte('1')), ShouldBeFalse)
		})
	})

	Convey("WordLexer", t, func() {
		Convey("It matches only if char is a unicode character", func() {
			So((NewWordLexer()).Match(byte('a')), ShouldBeTrue)
			So((NewWordLexer()).Match(byte('/')), ShouldBeFalse)
		})
	})
}
