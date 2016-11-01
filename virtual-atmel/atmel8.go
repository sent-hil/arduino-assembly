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
	registers   []uint8
	registerMin int
	registerMax int
}

func NewCPU() *CPU {
	return &CPU{
		registers:   make([]uint8, 16, 16),
		registerMin: 16,
		registerMax: 31,
	}
}

// LDI is 'Load Immediate'; it loads given value into given register.
// registerIndex must be: 16 <= d <= 31.
func (c *CPU) LDI(registerIndex int, value uint8) error {
	if registerIndex < 16 {
		return ErrLDILowRegister
	}

	if registerIndex > 31 {
		return fmt.Errorf(ErrLDIInvalidRegisterStr, registerIndex)
	}

	c.registers[registerIndex-c.registerMin] = value

	return nil
}
