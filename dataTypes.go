package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, "", ""

	print(reflect.TypeOf(25)) // int
	print(reflect.TypeOf(3.14)) // float64
	print(reflect.TypeOf(true)) // bool
	print(reflect.TypeOf("Hello")) // string
	print(reflect.TypeOf('s')) // rune, like ascii but better with code to each symbol
}
