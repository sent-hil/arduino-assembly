package atmel8

import "errors"

var (
	ErrLessThanMinRegister = errors.New("Given register index is less than minimum allowed value")
	ErrMoreThanMaxRegister = errors.New("Given register index is more than maximum allowed value")
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
		return ErrLessThanMinRegister
	}

	if registerIndex > 31 {
		return ErrMoreThanMaxRegister
	}

	c.registers[registerIndex-c.registerMin] = value

	return nil
}
