package main

import "fmt"

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myBool bool = true
	printPointer(&myBool)
}
