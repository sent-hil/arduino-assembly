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
	registers []uint8
	rMin      int
	rMax      int
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

// ADD is 'Add without carry'; it adds two given registers and stores the
// results in the 1st register
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

	return nil
}
