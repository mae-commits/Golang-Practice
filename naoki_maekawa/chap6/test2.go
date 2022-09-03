package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"} // define the array "letters" and all emements in it by slice literals.
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i]) // output elements in the array "letters" in an every line.
	}
	for _, letter := range letters { // extract indices and elements for each index in the array "letters",
		fmt.Println(letter)
	}
	fmt.Println(letters[0:2])
	/*
		decide the slice operator. note: be careful for deciding the range of slices;
		slice [a:b] means the range of the array from index value a to index value b-1.
	*/
}
