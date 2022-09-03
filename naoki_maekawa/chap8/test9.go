package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

/*
Update to take a pointer.
*/
func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

/*
Update to return a pointer.
*/
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	/*
		Return a pointer to a struct instead of the struct itself.
	*/
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	/*
		The variable "subscriber1" doesn't refer a struct type but does a pointer type
		because defaultSubscriber function returns the address of the argument in the function.
	*/
	subscriber1 := defaultSubscriber("Aman Singh")
	/*
		The function "applyDiscount" can access to and change the element of subscriber1.
		That is because "subscriber1" is the pointer type.
	*/
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}
