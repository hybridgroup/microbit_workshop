package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
	"tinygo.org/x/tinyfont"
)

var display microbitmatrix.Device

func main() {
	display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	display.ClearDisplay()

	c := color.RGBA{255, 255, 255, 0}
	then := time.Now()
	i := 0
	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		if time.Since(then).Nanoseconds() > 800000000 {
			then = time.Now()
			display.ClearDisplay()
			tinyfont.DrawChar(&display, &tinyfont.Tiny3x3a2pt7b, 1, 3, rune(abc[i]), c)
			i++
			if i >= len(abc) {
				i = 0
			}
		}

		display.Display()
	}
}
