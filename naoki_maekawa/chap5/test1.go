package main

import "fmt"

func main() {
	// create an array of seven strings (all components are the same type.)
	var notes [7]string
	// the initial index in an array is 0.
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
}
