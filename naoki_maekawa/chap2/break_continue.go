package main

import "fmt"

func main() {
	for x := 1; x <= 3; x++ {
		fmt.Println("before continue")
		continue // skip lines after "continue" and return the first line of the loop.
		fmt.Println("after continue")
	}
	for x := 1; x <= 3; x++ {
		fmt.Println("before break")
		break // break the loop and proceed out of the loop.
		fmt.Println("after break")
	}
}
