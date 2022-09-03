package main

import "fmt"

// The function "average" defines the calculation in the array.
func average(numbers ...float64) float64 {
	// The variable "sum" is the summation of all elements in the array "numbers".
	var sum float64 = 0
	// The next line extracts indices and elements from the array "numbers".
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}
