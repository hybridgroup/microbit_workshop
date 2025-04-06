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

	buttonTouch := machine.CAP_TOUCH
	buttonTouch.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !buttonTouch.Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 200)
	}
}
