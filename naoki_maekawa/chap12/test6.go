package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	// Generate a random integer between 1 and 3.
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		// No other number should be generated, but if it is, panic.
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
