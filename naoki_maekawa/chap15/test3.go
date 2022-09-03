package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

// The "twice" function accepts another function as a parameter.
func twice(theFunction func()) {
	// Call the passed-in function.
	theFunction()
	// Call the passed-in function (again).
	theFunction()
}

func main() {
	// Pass the "sayHi" function to the "twice" function.
	twice(sayHi)
	// Pass the "sayBye" function to the "twice" function.
	twice(sayBye)
}
