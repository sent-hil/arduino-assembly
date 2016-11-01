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

	Convey("ADD", t, func() {
		Convey("It returns error if given register indexes are not valid", func() {
			c := NewCPU()

			So(c.ADD(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
			So(c.ADD(1, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
			So(c.ADD(32, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It adds two register with empty value", func() {
			c := NewCPU()

			So(c.ADD(1, 2), ShouldBeNil)
			So(c.registers[1], ShouldEqual, 0)
		})

		Convey("It adds register with value to itself", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.ADD(16, 16), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 2)
		})

		Convey("It sets result of addition of 2 registers to 1 register", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.LDI(17, 2), ShouldBeNil)
			So(c.ADD(16, 17), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 3)
		})
	})
}
