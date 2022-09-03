package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(4.2, -3.0) // add the second return value "err" to detect the error value.
	if err != nil {
		fmt.Println(err) // if the return valus is error, print only error message.
		log.Fatal(err)   // if the code is long, you should finish the program in this line.
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}
}
