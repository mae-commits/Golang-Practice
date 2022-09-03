package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") // replace "#" into "o"
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}

/*

string.NewReplacer returns the arguments which replaces the string A to B,when you write as follows:
strings.NewReplacer(A,B)
strings(value) > NewReplacer(method name of strings/ value) > Replace(method name of NewReplacer)

*/
