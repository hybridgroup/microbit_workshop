package main

import (
	"machine"
	"time"
)

func main() {
	ledrow := machine.LED_COL_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led := machine.LED_ROW_1
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
