package main

import (
	"fmt"
)

func main() {
	// err := errors.New("height can't be negative")
	// fmt.Println(err.Error())
	// log.Fatal(err)
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error()) // prints the error meassage
	fmt.Println(err)         // same as the line above
}
