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
		c := NewCPU()
		ItActsLikeADD(c, c.ADD)

		Convey("It does not use previously stored carry flag", func() {
			c.carryFlag = true

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.ADD(16, 16), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 2)
			So(c.carryFlag, ShouldBeTrue)
		})
	})

	Convey("ADC", t, func() {
		c := NewCPU()
		ItActsLikeADD(c, c.ADC)

		Convey("It wraps around for results >255", func() {
			c := NewCPU()

			So(c.LDI(16, 255), ShouldBeNil)
			So(c.LDI(17, 1), ShouldBeNil)
			So(c.ADC(16, 17), ShouldBeNil) // 255 + 1
			So(c.registers[16], ShouldEqual, 0)

			Convey("It uses stored carry next ADC op", func() {
				// R0 is empty, so result will set only if carry is there
				So(c.ADC(0, 0), ShouldBeNil) // 0 + 0 + carry (1)
				So(c.registers[0], ShouldEqual, 1)

				Convey("It clears carry after being used", func() {
					// R1 is empty, so result will set only if carry is there
					So(c.ADC(1, 1), ShouldBeNil) // 0 + 0 + carry (0)
					So(c.registers[2], ShouldEqual, 0)
				})
			})
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
			So(c.INC(31), ShouldBeNil) // 1 + 1
			So(c.registers[31], ShouldEqual, 2)
		})

		Convey("It starts again at 0 if OP causes overflows", func() {
			c := NewCPU()

			So(c.LDI(31, 255), ShouldBeNil)
			So(c.INC(31), ShouldBeNil) // 255 + 1
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

			So(c.DEC(31), ShouldBeNil) // 0 - 1
			So(c.registers[31], ShouldEqual, 255)
		})

		Convey("It decrements existing value at register by 1", func() {
			c := NewCPU()

			So(c.LDI(31, 1), ShouldBeNil)
			So(c.DEC(31), ShouldBeNil) // 1 - 0
			So(c.registers[31], ShouldEqual, 0)
		})

		Convey("It starts again at 255 if OP causes overflows", func() {
			c := NewCPU()

			So(c.LDI(31, 0), ShouldBeNil)
			So(c.DEC(31), ShouldBeNil) // 0 - 1
			So(c.registers[31], ShouldEqual, 255)
		})
	})

	Convey("CLR", t, func() {
		Convey("It returns error if register is <0 or >31", func() {
			c := NewCPU()

			So(c.CLR(32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It clears values at given register", func() {
			c := NewCPU()

			So(c.LDI(31, 1), ShouldBeNil)
			So(c.CLR(31), ShouldBeNil)

			So(c.registers[31], ShouldEqual, 0)
		})
	})

	Convey("MOV", t, func() {
		Convey("It returns error if given register indexes are not valid", func() {
			c := NewCPU()

			So(c.MOV(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
			So(c.MOV(1, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		})

		Convey("It stores value from origin register to destination register", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.MOV(17, 16), ShouldBeNil)
			So(c.registers[17], ShouldEqual, 1)
		})
	})

	Convey("SUB", t, func() {
		c := NewCPU()
		ItActsLikeSUB(c, c.SUB)

		Convey("It does not use previously stored carry flag", func() {
			c.carryFlag = true

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.SUB(16, 16), ShouldBeNil)
			So(c.registers[16], ShouldEqual, 0)
			So(c.carryFlag, ShouldBeTrue)
		})
	})

	Convey("SBC", t, func() {
		c := NewCPU()
		ItActsLikeSUB(c, c.SBC)

		Convey("It wraps around for results <0", func() {
			c := NewCPU()

			So(c.LDI(16, 1), ShouldBeNil)
			So(c.LDI(17, 2), ShouldBeNil)
			So(c.SUB(16, 17), ShouldBeNil) // 1 - 2
			So(c.registers[16], ShouldEqual, 255)

			Convey("It uses stored carry next SBC op", func() {
				So(c.LDI(18, 2), ShouldBeNil)
				So(c.SBC(18, 19), ShouldBeNil) // 2 - 0 - carry (1)
				So(c.registers[18], ShouldEqual, 1)

				Convey("It clears carry after being used", func() {
					So(c.SBC(2, 2), ShouldBeNil)
					So(c.registers[2], ShouldEqual, 0)
				})
			})

			Convey("It uses stored carry next ADC op", func() {
				// R0 is empty, so result will set only if carry is there
				So(c.ADC(1, 1), ShouldBeNil) // 0 + 0 + carry (1)
				So(c.registers[1], ShouldEqual, 1)

				Convey("It clears carry after being used", func() {
					So(c.ADC(2, 2), ShouldBeNil) // 0 + 0 + carry (0)
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

	Convey("CLR", t, func() {
		Convey("It clears carry flag for next OP", func() {
			c := NewCPU()
			So(c.LDI(16, 255), ShouldBeNil)
			So(c.ADD(16, 16), ShouldBeNil)

			c.CLC()

			So(c.ADC(1, 1), ShouldBeNil) // 0 + 0 + carry (0)
			So(c.registers[1], ShouldEqual, 0)
		})
	})
}

func ItActsLikeADD(c *CPU, fx func(uint8, uint8) error) {
	Convey("It returns error if given register indexes are not valid", func() {
		So(fx(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		So(fx(1, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
	})

	Convey("It adds two register with empty value", func() {
		So(fx(1, 2), ShouldBeNil)
		So(c.registers[1], ShouldEqual, 0)
	})

	Convey("It adds register with value to itself", func() {
		So(c.LDI(16, 1), ShouldBeNil)
		So(fx(16, 16), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 2)
	})

	Convey("It sets result of addition of 2 registers to 1st register", func() {
		So(c.LDI(16, 1), ShouldBeNil)
		So(c.LDI(17, 2), ShouldBeNil)
		So(fx(16, 17), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 3)
	})

	Convey("It sets carry flag if result overflows", func() {
		So(c.LDI(16, 255), ShouldBeNil)
		So(fx(16, 16), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 254)
		So(c.carryFlag, ShouldBeTrue)
	})
}

func ItActsLikeSUB(c *CPU, fx func(uint8, uint8) error) {
	Convey("It returns error if given register indexes are not valid", func() {
		So(fx(32, 1), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
		So(fx(1, 32), ShouldResemble, fmt.Errorf("R32 is not a valid register"))
	})

	Convey("It subtracts two registers with empty value", func() {
		So(fx(1, 2), ShouldBeNil)
		So(c.registers[1], ShouldEqual, 0)
	})

	Convey("It substracts register with value from itself", func() {
		So(c.LDI(16, 1), ShouldBeNil)
		So(fx(16, 16), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 0)
	})

	Convey("It sets result of subtracting of 2 registers to 1st register", func() {
		So(c.LDI(16, 1), ShouldBeNil)
		So(c.LDI(17, 2), ShouldBeNil)
		So(fx(16, 17), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 255)
	})

	Convey("It sets carry flag if result overflows", func() {
		So(c.LDI(16, 0), ShouldBeNil)
		So(c.LDI(17, 1), ShouldBeNil)
		So(fx(16, 17), ShouldBeNil)
		So(c.registers[16], ShouldEqual, 255)
		So(c.carryFlag, ShouldBeTrue)
	})
}
