package main

import (
	"fmt"
)

var print = fmt.Println

func main() {

	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		print("Important Bday")
	} else if (iAge == 21) || (iAge == 50) {
		print("Important Bday")
	} else if iAge >= 65 {
		print("Important Bday")
	} else {
		print("Not an Important Bday")
	}
	print()
}
