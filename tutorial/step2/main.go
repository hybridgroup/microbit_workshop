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

	button := machine.BUTTONA
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() {
			led.Low()
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 20)
	}
}
