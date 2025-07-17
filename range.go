package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	aNums := []int{1,2,3}
	for _, value := range aNums {
		print(value)
	}
}
