package main

func main() {
	pinOutput(13, OUTPUT)

	for {
		digitalWrite(13, HIGH)
		sleep(1000)
		digitalWrite(13, LOW)
		sleep(1000)
	}
}
