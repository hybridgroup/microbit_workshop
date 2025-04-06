package main

import (
	"time"
)

func main() {

	i := 0
	for {
		println("Hello Gophers!", i)
		i++
		time.Sleep(time.Millisecond * 1000)
	}

}
