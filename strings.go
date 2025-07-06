package main

import (
	"fmt"
	"strings"
)

var print = fmt.Println

func main() {
	sV1 := "A word"
	// []byte
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	print(sV2)
	print("Length :", len(sV2))
	print("Contains Another :", strings.Contains(sV2, "Another"))
	print("o index :", strings.Index(sV2, "o"))
	print("Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n"
	print(sV3)
	sV3 = strings.TrimSpace(sV3)
	print(sV3)
	print("Split :", strings.Split("a-b-c-d", "-"))
	print("Lower :", strings.ToLower(sV2))
	print("Upper :", strings.ToUpper(sV2))
	print("Prefix :", strings.HasPrefix("tacocat", "taco")) // Checks if string starting with a specific prefix
	print("Suffix :", strings.HasSuffix("tacocat", "cat"))
}
