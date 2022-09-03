package main

import "fmt"

func calmDown() {
	// Returns an interface {} value.
	p := recover()
	// Assert that the type of the panic value is "error".
	err, ok := p.(error)
	if ok {
		// Now that we have an "error" value, we can call the Error method.
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	// Instead of a string, pass an error value to "panic".
	panic(err)
}
