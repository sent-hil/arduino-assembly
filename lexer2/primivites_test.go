package lexer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimitiveLexer(t *testing.T) {
	Convey("IntegerLexer", t, func() {
		s := newStringLexer("100\n")
		i := NewIntegerLexer()

		Convey("Match", func() {
			So(i.Match(s), ShouldBeTrue)
		})

		Convey("Lex", func() {
			tok := i.Lex(s)
			So(tok.ID, ShouldEqual, TokenInteger)
			So(tok.Value, ShouldEqual, "100")
		})
	})
}
