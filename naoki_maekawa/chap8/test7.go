package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

/*
To take a pointer to a struct, not the struct inself.
*/
func applyDiscount(s *subscriber) {
	/*
		Update the struct field without * operator
		because dot notation itself works on struct pointers
		as well as the structs themselves.
	*/
	s.rate = 4.99
}

func main() {
	var s subscriber
	// Pass a pointer, not a struct.
	applyDiscount(&s)
	fmt.Println(s.rate)
}
