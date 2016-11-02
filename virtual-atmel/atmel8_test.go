package atmel8

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCPU(t *testing.T) {
	Convey("LDI", t, func() {
		Convey("It returns error if register is <16 or >31", func() {
			c := NewCPU()

			So(c.LDI(15, 1), ShouldEqual, ErrLDILowRegister)
			So(c.LDI(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It loads given uint8 value into given register", func() {
			c := NewCPU()

			var i uint8
			for i = 16; i <= 31; i++ {
				So(c.LDI(i, i), ShouldBeNil)
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

		Convey("It sets result of addition of 2 registers to 1st register", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.LDI(17, 2), ShouldBeNil)
			So(c.ADD(16, 17), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 3)
		})

		Convey("It sets carry flag if result overflows", func() {
			c := NewCPU()

			So(c.LDI(16, 255), ShouldBeNil)
			So(c.ADD(16, 16), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 254)
			So(c.carryFlag, ShouldBeTrue)
		})

		Convey("It does not use previously stored carry flag", func() {
			c := NewCPU()
			c.carryFlag = true

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.ADD(16, 16), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 2)
			So(c.carryFlag, ShouldBeTrue)
		})
	})

	Convey("ADC", t, func() {
		Convey("It returns error if given register indexes are not valid", func() {
			c := NewCPU()

			So(c.ADC(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
			So(c.ADC(1, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
			So(c.ADC(32, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It does same ops as ADD for results <255", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.LDI(17, 2), ShouldBeNil)
			So(c.ADC(16, 17), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 3)
		})

		Convey("It wraps around for results >255", func() {
			c := NewCPU()

			So(c.LDI(16, 255), ShouldBeNil)
			So(c.LDI(17, 1), ShouldBeNil)
			So(c.ADC(16, 17), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 0)

			Convey("It uses stored carry next ADC op", func() {
				// R0 is empty, so result will set only if carry is there
				So(c.ADC(0, 0), ShouldBeNil)
				So(c.registers[0], ShouldEqual, 1)

				Convey("It clears carry after being used", func() {
					So(c.ADD(1, 1), ShouldBeNil)
					So(c.registers[2], ShouldEqual, 0)
				})
			})
		})
	})

	Convey("SEC", t, func() {
		Convey("It sets carry flag for next OP", func() {
			c := NewCPU()
			c.SEC()

			So(c.ADC(0, 0), ShouldBeNil)
			So(c.registers[0], ShouldEqual, 1)
		})
	})

	Convey("INC", t, func() {
		Convey("It returns error if register is <0 or >31", func() {
			c := NewCPU()

			So(c.INC(32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It increments empty value at register by 1", func() {
			c := NewCPU()

			So(c.INC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 1)
		})

		Convey("It increments existing value at register by 1", func() {
			c := NewCPU()

			So(c.LDI(31, 1), ShouldBeNil)
			So(c.INC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 2)
		})

		Convey("It starts again at 0 if OP causes overflows", func() {
			c := NewCPU()

			So(c.LDI(31, 255), ShouldBeNil)
			So(c.INC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 0)
		})
	})

	Convey("DEC", t, func() {
		Convey("It returns error if register is <0 or >31", func() {
			c := NewCPU()

			So(c.DEC(32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It decrements empty value at register by 1", func() {
			c := NewCPU()

			So(c.DEC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 255)
		})

		Convey("It decrements existing value at register by 1", func() {
			c := NewCPU()

			So(c.LDI(31, 1), ShouldBeNil)
			So(c.DEC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 0)
		})

		Convey("It starts again at 255 if OP causes overflows", func() {
			c := NewCPU()

			So(c.LDI(31, 0), ShouldBeNil)
			So(c.DEC(31), ShouldBeNil)
			So(c.registers[31], ShouldEqual, 255)
		})
	})
}
