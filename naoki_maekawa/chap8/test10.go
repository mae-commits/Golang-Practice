package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var s subscriber
	s.rate = 4.99
	fmt.Printf("%#v\n", s)
}
