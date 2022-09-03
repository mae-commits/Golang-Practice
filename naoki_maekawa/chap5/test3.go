package main

import "fmt"

func main() {
	var counters [3]int
	// increment the first element from 0 to 1.
	counters[0]++
	// increment the first element from 1 to 2.
	counters[0]++
	// increment the third element form 0 to 1.
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
}
