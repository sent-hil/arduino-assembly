package atmel8

import (
	"errors"
	"fmt"
)

var (
	ErrLDILowRegister        = errors.New("ldi can only use a high register (r16 - r31)")
	ErrLDIInvalidRegisterStr = "R%v is not a valid register"
)

type CPU struct {
	registers       []uint8
	rMin            int
	rMax            int
	carryFromLastOp bool
}

func NewCPU() *CPU {
	return &CPU{
		registers: make([]uint8, 32, 32),
		rMin:      16,
		rMax:      31,
	}
}

// LDI is 'Load Immediate'; it loads given value into given register.
//
// rIndex must be: 16 <= d <= 31.
func (c *CPU) LDI(rIndex int, value uint8) error {
	if rIndex < c.rMin {
		return ErrLDILowRegister
	}

	if rIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rIndex)
	}

	c.registers[rIndex] = value

	return nil
}

// ADD is 'Add without Carry'; it adds two given registers and stores the
// results in the 1st register. It discards high bit if the result overflows.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADD(rDestIndex, rIndex int) error {
	if rDestIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rDestIndex)
	}

	if rIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rIndex)
	}

	c.registers[rDestIndex] += c.registers[rIndex]

	if c.carryFromLastOp {
		c.registers[rDestIndex] += 1
		c.carryFromLastOp = false
	}

	return nil
}

// ADD is 'Add with Carry'; it adds two given registers and stores the
// results in the 1st register. It stores the high if results overflows,
// which will be used in the next ADD/ADC OP.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADC(rDestIndex, rIndex int) error {
	if rDestIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rDestIndex)
	}

	if rIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rIndex)
	}

	v1, v2 := uint16(c.registers[rDestIndex]), uint16(c.registers[rIndex])

	c.registers[rDestIndex] += c.registers[rIndex]
	if c.carryFromLastOp {
		c.registers[rDestIndex] += 1
		c.carryFromLastOp = false
	}

	if v1+v2 > 255 {
		c.carryFromLastOp = true
	}

	return nil
}
