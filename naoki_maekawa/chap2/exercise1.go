package main

import (
	"fmt"
)

func main() {
	a := "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		// fmt.Println(d)
	}
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
}

/*

Commented out lines try to output the undifined variables
because they are out of range which firstly defined

*/
