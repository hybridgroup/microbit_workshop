package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

var display microbitmatrix.Device

func main() {
	display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	display.ClearDisplay()

	c := color.RGBA{255, 255, 255, 0}
	happyFace := true
	then := time.Now()
	for {
		if time.Since(then).Nanoseconds() > 800000000 {
			then = time.Now()
			display.ClearDisplay()
			happyFace = !happyFace
			if happyFace {
				// HAPPY FACE
				// eyes
				display.SetPixel(1, 1, c)
				display.SetPixel(3, 1, c)
				// smile
				display.SetPixel(0, 3, c)
				display.SetPixel(1, 4, c)
				display.SetPixel(2, 4, c)
				display.SetPixel(3, 4, c)
				display.SetPixel(4, 3, c)
			} else {
				// SAD FACE
				// eyes
				display.SetPixel(1, 1, c)
				display.SetPixel(3, 1, c)
				// smile
				display.SetPixel(0, 4, c)
				display.SetPixel(1, 3, c)
				display.SetPixel(2, 3, c)
				display.SetPixel(3, 3, c)
				display.SetPixel(4, 4, c)
			}
		}

		display.Display()
	}
}
