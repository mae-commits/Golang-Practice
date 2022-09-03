package main

import (
	"errors"
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) { // the function contains the error deals
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

/*

	"errors" is the name of package.
	"error" is the type itself.

*/

func main() {
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
