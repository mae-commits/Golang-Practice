package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	/*
		The returnable value is an array.
	*/
	var result []float64
	for _, number := range numbers {
		/*
			The condition of defining the range of number between min and max.
		*/
		if number >= min && number <= max {
			// append the element "number" with the array "result".
			result = append(result, number)
		}
	}
	return result
}

func main() {
	/*
		the next line is imported as follows: min = 1, max = 100, numbers = [-12.5, 3.2, 0, 50, 103.5]
	*/
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
