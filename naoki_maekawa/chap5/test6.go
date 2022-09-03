package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	// the loop below sums the all elememts in the array.
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
