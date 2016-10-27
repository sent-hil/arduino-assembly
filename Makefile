blink:
	avr-as -g -mmcu=atmega328p -o blink.o blink.s
	avr-ld -o blink.elf blink.o
	avr-objcopy -O ihex -R .eeprom blink.elf blink.hex
	avrdude  -p atmega328p -c arduino -P /dev/tty.usbmodem14221 -b 115200 -D -U flash:w:blink.hex:i

light:
	avra hello.asm
	avrdude -p atmega328p -c arduino -P /dev/tty.usbmodem14221 -b 115200 -D -U flash:w:hello.hex:i
