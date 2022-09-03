package main

import "fmt"

/*
	Delete the field("homeaddress"), leaving only the type address.
*/
type subscriber struct {
	name   string
	rate   float64
	active bool
	address
}

type employee struct {
	name   string
	salary float64
	address
}

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

func main() {
	subscriber1 := subscriber{name: "Aman Singh"}
	/*
		The variable "street" is in the struct "address",
		which is embedded in the struct "subscriber1",
		so you can omit wrting struct "address".
	*/
	subscriber1.street = "123 Oak street"
	subscriber1.city = "Omaha"
	fmt.Println("Street:", subscriber1.street)
	fmt.Println("City:", subscriber1.city)
	employee1 := employee{name: "Joy Carr"}
	employee1.state = "OR"
	employee1.postalCode = "97222"
	fmt.Println("State:", employee1.state)
	fmt.Println("State:", employee1.postalCode)
}
