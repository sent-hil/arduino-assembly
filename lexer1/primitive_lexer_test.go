package lexer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimitiveLexers(t *testing.T) {
	Convey("IntegerLex", t, func() {
		Convey("It implments Lexer interface", func() {
			var _ Lexer = NewIntegerLex()
		})

		Convey("It matches only if first char is an integer", func() {
			So(NewIntegerLex().Match('1'), ShouldBeTrue)
		})

		Convey("It parses every digit till break", func() {
			f := newFileLex("1000")
			NewIntegerLex().Run(f)

			So(len(f.Tokens), ShouldEqual, 1)

			token := f.Tokens[0]
			So(token.ID, ShouldEqual, TokenInteger)
			So(token.Value, ShouldEqual, "1000")
		})
	})
}
