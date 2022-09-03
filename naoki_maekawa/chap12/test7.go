package main

import "fmt"

func calmDown() {
	// Call "recover" in this other function.
	recover()
}

func freakOut() {
	// Defer a call to the function that recovers.
	defer calmDown()
	// If the program panics after that, the deferred function call will recover!
	panic("oh no")
	// The lines following the line above will not be read.
	fmt.Println("I won't be run@")
}

func main() {
	freakOut()
	// But the line below runs after freakOut returns.
	fmt.Println("Exiting normally")
}
