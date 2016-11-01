package atmel8

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCPU(t *testing.T) {
	Convey("LDI", t, func() {
		Convey("It returns error if register is <16 or > 31", func() {
			c := NewCPU()

			So(c.LDI(15, 1), ShouldEqual, ErrLDILowRegister)
			So(c.LDI(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It loads given uint8 value into given register", func() {
			c := NewCPU()

			for i := 16; i <= 31; i++ {
				So(c.LDI(i, uint8(i)), ShouldBeNil)
				So(c.registers[i], ShouldEqual, i)
			}

		})

		Convey("It accepts hex as value", func() {
			c := NewCPU()

			So(c.LDI(16, 0x01), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 1)
		})
	})
}
