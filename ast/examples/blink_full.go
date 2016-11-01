package main

var (
	OUTPUT = 0
	LOW    = 0
	HIGH   = 1
)

func main() {
	pinOutput(13, OUTPUT)

	for {
		digitalWrite(13, HIGH)
		sleep(1000)
		digitalWrite(13, LOW)
		sleep(1000)
	}
}

func pinOutput(pin, mode int) {
}

func digitalWrite(pin, mode int) {
}

func sleep(time int) {
}
