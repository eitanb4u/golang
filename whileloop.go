package main

import  (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var print = fmt.Println

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Print("Guess a number between 0 and 50: ")
		print("Random Number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			print("Pick a Lower Value")
		} else if iGuess < randNum {
			print("Pick a Higher Value")
		} else {
			print("You guessed it!")
			break
		}
			

	}
}
