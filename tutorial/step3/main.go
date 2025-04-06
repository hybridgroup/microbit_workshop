package main

import (
	"machine"
	"time"
)

func main() {
	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ledA := machine.LED_ROW_1
	ledA.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ledB := machine.LED_ROW_5
	ledB.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttonA := machine.BUTTONA
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInput})

	buttonB := machine.BUTTONB
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if buttonA.Get() {
			ledA.Low()
		} else {
			ledA.High()
		}

		if buttonB.Get() {
			ledB.Low()
		} else {
			ledB.High()
		}

		time.Sleep(time.Millisecond * 20)
	}
}
