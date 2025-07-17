package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	for x := 5; x >= 1; x-- {
		print(x)
	}

	fX := 0
	for fX < 5 {
		print(fX)
		fX++
	}
}
