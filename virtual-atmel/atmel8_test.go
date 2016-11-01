package atmel8

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCPU(t *testing.T) {
	Convey("LDI", t, func() {
		Convey("It returns error if register is <16 or > 31", func() {
			c := NewCPU()

			So(c.LDI(15, 1), ShouldEqual, ErrLessThanMinRegister)
			So(c.LDI(32, 1), ShouldNotBeNil)
			So(c.LDI(32, 1), ShouldNotBeNil)
		})

		Convey("It loads given uint8 value into given register", func() {
			c := NewCPU()

			for i := 16; i <= 31; i++ {
				So(c.LDI(i, uint8(i)), ShouldBeNil)
				So(c.registers[i-16], ShouldEqual, i)
			}

		})

		Convey("It accepts hex as value", func() {
			c := NewCPU()

			So(c.LDI(16, 0x01), ShouldBeNil)
			So(c.registers[0], ShouldEqual, 1)
		})
	})
}
