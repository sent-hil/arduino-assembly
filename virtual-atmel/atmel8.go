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
	rMin      uint8
	rMax      uint8
	carryFlag bool
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
// rDestIndex must be: 16 <= d <= 31.
func (c *CPU) LDI(rDestIndex, value uint8) error {
	if rDestIndex < c.rMin {
		return ErrLDILowRegister
	}

	if err := c.checkRegisterOutofRange(rDestIndex); err != nil {
		return err
	}

	c.registers[rDestIndex] = value

	return nil
}

// ADD is 'Add without Carry'; it adds two given registers. It stores the
// results in the 1st register. It sets C flag if result overflows.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADD(rDestIndex, rIndex uint8) error {
	return c.add(rDestIndex, rIndex, false)
}

// ADD is 'Add with Carry'; it adds two given registers & carry from last op.
// It stores the results in the 1st register. It sets C flag if result
// overflows.
//
// rDestIndex and rIndex must be: 0 <= d <= 31.
func (c *CPU) ADC(rDestIndex, rIndex uint8) error {
	return c.add(rDestIndex, rIndex, true)
}

// SEC is 'Set Carry Flag'; it sets carry flag.
func (c *CPU) SEC() {
	c.carryFlag = true
}

// INC is 'Increment'; it increments value at register by 1. It OP
// causes overflow, it does NOT set carry flag.
func (c *CPU) INC(rDestIndex uint8) error {
	if err := c.checkRegisterOutofRange(rDestIndex); err != nil {
		return err
	}

	c.registers[rDestIndex] += 1

	return nil
}

///// private

func (c *CPU) add(rDestIndex, rIndex uint8, carry bool) error {
	if err := c.checkRegisterOutofRange(rDestIndex, rIndex); err != nil {
		return err
	}

	v1, v2 := uint16(c.registers[rDestIndex]), uint16(c.registers[rIndex])

	var v3 uint16 = 0
	if carry && c.carryFlag {
		v3 = 1
		c.carryFlag = false
	}

	result := v1 + v2 + v3
	if result > 255 {
		c.carryFlag = true
	}

	c.registers[rDestIndex] = uint8(v1 + v2 + v3)

	return nil
}

func (c *CPU) checkRegisterOutofRange(rIndexes ...uint8) error {
	for _, rIndex := range rIndexes {
		if rIndex > c.rMax {
			return fmt.Errorf(ErrLDIInvalidRegisterStr, rIndex)
		}
	}

	return nil
}
