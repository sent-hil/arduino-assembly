package main

import (
	"io/ioutil"
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

	Convey("FileLexer", t, func() {
		Convey("It should initialize with given filename", func() {
			tmpFile, err := ioutil.TempFile("", "")
			So(err, ShouldBeNil)

			_, err = NewFileLexer(tmpFile.Name())
			So(err, ShouldBeNil)
		})
	})
}
