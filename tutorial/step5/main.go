package main

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

func main() {
	bzrPin := machine.BUZZER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttonA := machine.BUTTONA
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInput})

	bzr := buzzer.New(bzrPin)

	for {
		if !buttonA.Get() {
			bzr.Tone(buzzer.A5, 0.2)
		}
	}
}
