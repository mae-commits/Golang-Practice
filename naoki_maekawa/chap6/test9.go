package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// the next line gets values from command line.
	arguments := os.Args[1:]
	var sum float64 = 0
	// get both the index and the element of the array.
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
