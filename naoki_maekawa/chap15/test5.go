package main

import "fmt"

// The doMath function accept another fucnction as a parameter.
// The passed-in function must accept two integers and return a float64.
func doMath(passedFunction func(int, int) float64) {
	// Call the passed-in function.
	result := passedFunction(10, 2)
	// Print the passed-in function's return value.
	fmt.Println(result)
}

// A function that can be passed into doMath.
func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

// Another function that can be passed into doMath.
func multiply(a int, b int) float64 {
	return float64(a) * float64(b)
}

func main() {
	// Pass the "divide" function to doMath.
	doMath(divide)
	// Pass the "multiply" function to doMath.
	doMath(multiply)
}
