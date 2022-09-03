package main

import (
	"fmt"
)

func main() {
	// var mySlice []string
	// var myArray []int
	var notes []string
	notes = make([]string, 7) // create slice with seven strings.
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	// create a slice with five integers, and set up a variable to hold it.
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(len(notes)) // output the length of the array.
	fmt.Println(len(primes))
}
