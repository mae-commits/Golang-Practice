package main

import "fmt"

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	intSlice := []int{1, 2, 3}
	/*
		The next line imports the strings from intSlice as strings with "...",
		in order to input as the array.
	*/
	severalInts(intSlice...)
	stringSlice := []string{"a", "b", "c", "d"}
	mix(1, true, stringSlice...)
}
