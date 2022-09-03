package main

import "fmt"

func createPointer() *float64 { // declare the function returns a float64 pointer
	myFloat := 98.5
	return &myFloat
}

func main() {
	var myFloatPointer *float64 = createPointer() // assign the returned pointer to a variable
	fmt.Println(*myFloatPointer)
}
