package main

import "fmt"

func main() {
	/*
		myStruct is defined as a struct type of object,
		which can input seceral types of elements.
		struct{field: a type of field}
	*/
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	/*
	 Using (The name of struct).(the name of struct's field) allows
	 to access fields in the struct you decided.
	*/
	myStruct.number = 3.14
	/*
		"%#v" verb prints the value in myStruct as a struct literal.
	*/
	fmt.Printf("%#v\n", myStruct)
}
