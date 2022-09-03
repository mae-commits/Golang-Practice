package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		// the next line shows the index and the i th element in the array.
		fmt.Println(i, notes[i])
	}
	// the next line shows the length of the array "notes".
	fmt.Println(len(notes))
	/*
		first, second := range array returns
		the index of the array into "first",
		and the element of the array into "second" orderly.
		Moreover, this method requires two variables,
		so you should prepare two variables or "_" as the either variable.
	*/
	for index, note := range notes {
		fmt.Println(index, note)
	}
}
