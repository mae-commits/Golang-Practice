package main

import (
	"fmt"
)

func main() {
	counters := make(map[string]int)
	// Increase the value for each key.
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
}
