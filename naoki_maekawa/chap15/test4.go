package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	// This variable will hold a function with no parameters and no return value.
	var greeterFunction func()
	// This variable will hold a function with two int parameters and a float64 return value.
	var mathFunction func(int, int) float64
	// Assign the "sayHi" function to the greeterFunction variable.
	greeterFunction = sayHi
	// Assign the "divide" function to the mathFunction variable.
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
	// fmt.Println(divide(5, 2))
}
