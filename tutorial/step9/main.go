package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
	"tinygo.org/x/tinydraw"
)

var display microbitmatrix.Device

func main() {
	display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	display.ClearDisplay()

	c := color.RGBA{255, 255, 255, 0}
	tinydraw.Line(&display, 0, 2, 2, 0, c)
	tinydraw.Rectangle(&display, 2, 2, 3, 3, c)
	for {
		display.Display()
		time.Sleep(10 * time.Millisecond)
	}
}
