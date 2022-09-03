package main

import (
	"fmt"
)

func main() {
	s1 := []string{"s1", "s2"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4) // print slices
	s4[0] = "XX"                // change the element of index 1 in the array "s4"
	fmt.Println(s1, s2, s3, s4)
}
