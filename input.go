package main

import (
	"fmt"
	"bufio"
	"os"
)

// comments

var print = fmt.Println

func main() {	
	reader := bufio.NewReader(os.Stdin)
	
	print("what is your mame")
	name, err := reader.ReadString('\n')
	if err != nil {
		print("Error reading input", err)
		return
	}
	print("Hello", name)
}

