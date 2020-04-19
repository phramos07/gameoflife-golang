package main

import (
	concrete "gol/concrete"
	"time"
)

func main() {
	gol := concrete.NewGol(180, 50, 6)
	gol.Start()
	for {
		ClearScreen()
		gol.Print()
		gol.Next()
		time.Sleep(1 * time.Second / 10)
	}
}
