package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int // declare the pointer of int type
	myInt = 42
	myIntPointer = &myInt // refer variable myInt by pointer type of variable myIntPointer
	fmt.Println(*myIntPointer)
	fmt.Println(*myIntPointer)
}
