package main

import (
	"fmt"
)

func main() {
	floatSlice := make([]float64, 10) // create slices without assigning values to their elements.
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])
	var intSlice []int
	var stringSlice []string
	/*
		"%#v" formats a value as it would appear in Go code.
	*/
	fmt.Printf("intSlice: %#v, stringSlices: %#v\n", intSlice, stringSlice)
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)
}
