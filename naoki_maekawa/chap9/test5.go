package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value")
	pointer := &value
	// Value automatically converted to pointer.
	value.pointerMethod()
	pointer.method()
	// Value at pointer automatically retrieved.
	pointer.pointerMethod()
}
