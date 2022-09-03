package main

import "fmt"

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	// Print the current starting number.
	fmt.Println(start)
	// If we have not reached the ending number...
	if start < end {
		// the "count" function calls itself, with a starting number 1 higher than before.
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d,%d) call\n", start, end)
}

func main() {
	// Call "count" for the first time, specifying it should count from 1 to 3.
	count(1, 3)
}
