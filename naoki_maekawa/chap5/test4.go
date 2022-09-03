package main

import "fmt"

func main() {
	// the next line sumltaneously defines the string notes.
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	// the next line is similar to the last line: it also defines int strings.
	primes := [5]int{2, 3, 5, 7, 11}
	// the next line outputs all elements of the string "notes".
	fmt.Println(notes)
	// the next line outputs all elements of the string "primes"
	fmt.Println(primes)
	// the next line shows the line 7 itself.
	fmt.Printf("%#v\n", notes)
	// the next line show the line 9 itself.
	fmt.Printf("%#v\n", primes)
}
