package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" // import path and you can use rand directly without calling math package.
	"os"
	"strconv"
	"strings"
	"time" // import "time" package
)

func main() {
	seconds := time.Now().Unix() // get the current date and time, as an integer.
	rand.Seed(seconds)           // seed the random number generator.
	target := rand.Intn(100) + 1 // when you call rand.Intn, you can generate a random integer from 0 to 99.
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) // read keyboard input.
	success := false                    // the detector of correct guess.
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n') // read what the user types, up until they press enter.
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // remove the new line.
		guess, err := strconv.Atoi(input) // convert the input string to an integer.

		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break // if you guess correctly, the game will be over.
		}
	}
	if !success { // if value of success is still false, output such as saying as follows.
		fmt.Println("Sorry,you didn't guess my number. It was:", target)
	}
}
