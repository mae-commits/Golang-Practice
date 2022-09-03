package main

import (
	"fmt"
	"math"
)

func floatParts(number float64) (integerPart int, fractinalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	cans, remainder := floatParts(1.26) // floatParts returns two values
	fmt.Println(cans, remainder)
}
