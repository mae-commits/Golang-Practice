package main

import "fmt"

func main() {
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "A" // redefine the value of the index in array1.
	fmt.Println(array1)
	fmt.Println(slice1)
}
