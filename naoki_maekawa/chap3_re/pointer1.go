package main

import (
	"fmt"
)

func main() {
	myInt := 4
	myIntPointer := &myInt // & refers the address of value itself
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer) // * shows values itself

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}
