package main

import "fmt"

/*
type part defines the type of struct with
{description(= the type of string), count(= the type of int)}
*/
type part struct {
	description string
	count       int
}

/*
func minimumOrder returns part (= the type of struct),
and its argument is the string type, whose name is "description".
*/
func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)
}
