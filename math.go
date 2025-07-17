package main

import (
	"fmt"
	"math/rand"
	"math"
)

var print = fmt.Println

func main() {
	print("5 + 4 =", 5+4)
	print("5 - 4 =", 5-4)
	print("5 * 4 =", 5*4)
	print("5 / 4 =", 5/4)
	print("5 % 4 =", 5%4)

	mInt := 1
	mInt += 1
	mInt++
	print("mInt", mInt)

	randNum := rand.Intn(50) + 1
	print("Random:", randNum)
}
