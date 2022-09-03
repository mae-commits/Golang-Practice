package main

import "fmt"

/*
type (struct name) (struct) {
	field1 a type of field1
	field2 a type of field2
	...
}
This method defines types themselves, and its scope is only within the function.
That is why this method should be written out of the main function.
*/
type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println(porsche.name)
	fmt.Println(porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}
