package main

import (
	"fmt"
)

func main() {
	fmt.Printf("$%12s | %2s\n", "Product", "Cost in Cents")
	fmt.Println("-------------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tags", 99)
	fmt.Printf("%7.4f\n", 12.3456) // %A.Bf: A = minimun width of entire number. B = width after decimal point.
}
