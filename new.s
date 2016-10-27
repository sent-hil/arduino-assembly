#define CLOCK_MHZ       16UL
#define DELAY_LENGTH_MS 1000UL
#define DELAY_VALUE     (uint32_t)((CLOCK_MHZ * 1000UL * DELAY_LENGTH_MS) / 5UL)

void setup() {
asm  volatile (
"sbi %0, %1 \n" // pinMode(13, OUTPUT);

:
:
	"I" (_SFR_IO_ADDR(DDRB)), "I" (DDB5)
	)
	}

	void loop() {
	asm  volatile (
	"mov r18, %D2 \n" // save for second delay iteration
	"mov r20, %C2 \n"
	"mov r21, %B2 \n"

	"sbi %0, %1   \n" // turn LED on

"1:
	\n"               // delay ~1 second
	"subi %A2, 1  \n"
	"sbci %B2, 0  \n"
	"sbci %C2, 0  \n"
	"brcc 1b      \n"

	"cbi %0, %1   \n" // turn LED off

"2:
	\n"               // delay ~1 second
	"subi r18, 1  \n"
	"sbci r19, 0  \n"
	"sbci r20, 0  \n"
	"brcc 2b      \n"

:
:
	"I" (_SFR_IO_ADDR(PORTB)), "I" (PORTB5), "r" (DELAY_VALUE) : "r18", "r19", "r20"
	)
	}
