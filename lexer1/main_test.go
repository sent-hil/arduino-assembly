package lexer

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimitiveLexers(t *testing.T) {
	Convey("IntegerLexer", t, func() {
		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewIntegerLexer()
		})

		Convey("It matches only if first char is an integer", func() {
			So(NewIntegerLexer().Match('1'), ShouldBeTrue)
		})

		Convey("It parses every digit till break", func() {
			f := newFileLexer("1000\n")
			NewIntegerLexer().Run(f)

			for t := range f.Tokens {
				So(t.Id, ShouldEqual, TokenInteger)
				So(t.Value, ShouldEqual, "1000")
			}

			//So(token.ID, ShouldEqual, TokenInteger)
			//So(token.Value, ShouldEqual, "1000")
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
