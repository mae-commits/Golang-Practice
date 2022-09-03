package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) (float64, []float64) {
	// the next line defines starting max from -1.
	max := math.Inf(-1)
	// extract from indices and elememts in the array numbers.
	for _, number := range numbers {
		if number > max {
			// renew the maximum number if the value "number" is bigger than the latest number.
			max = number
		}
	}
	return max, numbers
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}
