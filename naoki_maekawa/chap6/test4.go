package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b"}     // define the array "slice"
	fmt.Println(slice, len(slice))  // output the array "slice" itself, and the length of it.
	slice = append(slice, "c")      // append the new element "c" into the array "slice" with using "append" method.
	slice = append(slice, "d", "e") // append the new elements "d" and "e" sumaltaneously.
	fmt.Println(slice, len(slice))  // output the array "slice" itself, and the length of it.
}
