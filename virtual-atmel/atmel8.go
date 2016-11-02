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

// NewCPU initializes CPU with 32 registers (R0-R31). R16-R31 are exposed to
// user.
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

// ADD is 'Add without Carry'; it adds two given registers. It stores the
// results in the 1st register. It sets C flag if result overflows.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADD(rDestIndex, rIndex int) error {
	return c.add(rDestIndex, rIndex, false)
}

// ADD is 'Add with Carry'; it adds two given registers & carry from last op.
// It stores the results in the 1st register. It sets C flag if result
// overflows.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADC(rDestIndex, rIndex int) error {
	return c.add(rDestIndex, rIndex, true)
}

func (c *CPU) add(rDestIndex, rIndex int, carry bool) error {
	if rDestIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rDestIndex)
	}

	if rIndex > c.rMax {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, rIndex)
	}

	v1, v2 := uint16(c.registers[rDestIndex]), uint16(c.registers[rIndex])

	var v3 uint16 = 0
	if carry && c.carryFromLastOp {
		v3 = 1
		c.carryFromLastOp = false
	}

	result := v1 + v2 + v3
	if result > 255 {
		c.carryFromLastOp = true
	}

	c.registers[rDestIndex] = uint8(v1 + v2 + v3)

	return nil
}
