package main

import "fmt"

/*
"subscriber" is defined as the type of strcut such as
{
	field1(name) string type
	field2(rate) float64 type
	field3(active) bool type
}
*/
type subscriber struct {
	name   string
	rate   float64
	active bool
}

/*
	func "printInfo" is input subscriber types of arguments,
	and print the data in arguments.
*/
func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	subscriber2.active = false
	printInfo(subscriber2)
}
