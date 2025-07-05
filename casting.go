package main

import (
	"fmt"
	"strconv"
	"reflect"
)

var print = fmt.Println

func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	print(cV2, reflect.TypeOf(cV2))

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	print(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	print(cV6, reflect.TypeOf(cV6))

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {	
		print(cV8, reflect.TypeOf(cV8))
	}

	cV9 := fmt.Sprintf("%f", 3.14) // from decimal to string using the Sprintf method
	print(cV9, reflect.TypeOf(cV9))
}
