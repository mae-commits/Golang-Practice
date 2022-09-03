/*
This code inputs the data of subscribers' name, rate, and status.
*/
package main

import "fmt"

func main() {
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly:", subscriber.rate)
	fmt.Println("Active?", subscriber.active)
}
