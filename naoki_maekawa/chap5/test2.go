package main

import (
	"fmt"
)

func main() {
	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	// default value of an int type array is 0.
	fmt.Println(primes[2]) // 0
	fmt.Println(primes[4]) // 0

	var notes [7]string
	notes[0] = "do"
	// default value of a string type's array is " "(blank).
	fmt.Println(notes[3])
	fmt.Println(notes[6]) // " "
	fmt.Println(notes[0]) // " "
}
