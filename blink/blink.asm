;hello.asm
; turns on an LED which is connected to PB5 (digital out 13)

.include "./m328Pdef.inc"

ldi r16, 0b00100000
out DDRB, r16
out PortB, r16

mainloop:
	eor  r17, r16; invert output bit
	out  PORTB, r17; write to port
	LDI  R18, 255
	LDI  R19, 255
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	DEC  R18
	call wait; wait some time
	rjmp mainloop; loop forever

wait:
	push r16
	push r17
	push r18

	ldi r16, 0x40; loop 0x400000 times
	ldi r17, 0x00; ~12 million cycles
	ldi r18, 0x00; ~0.7s at 16Mhz

_w0:
	dec  r18
	brne _w0
	dec  r17
	brne _w0
	dec  r16
	brne _w0

	pop r18
	pop r17
	pop r16
	ret
