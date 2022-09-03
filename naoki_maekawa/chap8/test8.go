package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	/*
		To get pointer, parentheses are needed for wrapping variables.
		But it's tedious to write for each time you get pointers in field.
	*/
	fmt.Println((*pointer).myField)
	fmt.Println(pointer.myField)
}
