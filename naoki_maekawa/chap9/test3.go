package main

import "fmt"

// Define a new type Mytype by string
type MyType string

/*
func (receiver parameter the method will be defined on MyType)
*/
func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	// Create a MyType value.
	value := MyType("a MyType value")
	// Call sayHi on that value.
	value.sayHi()
	// Create another MyType value.
	anotherValue := MyType("another value")
	// Call sayHi on the new value.
	anotherValue.sayHi()
}
