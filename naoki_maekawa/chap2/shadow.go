package main

import (
	"fmt"
)

func main() {
	var int int = 12
	var append string = "minutes of bonus footage"
	// var fmt string = "DVD"
	fmt.Println(int)
	fmt.Println(append)
	var count int
	languages := append([]string{}, "Espanol")
	fmt.Println(int, append, "on", fmt, languages)
}

/*

If you name the variable names which are already defined, original methods/functions are shadowed.
That is why we cannot use the fmt.Println if we do not exclude the line 9.

*/
