package main

import "fmt"

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

/*
The struct type functions can also be used in those of others,
so when you would like to use same structs in several structs,
you can use them, and they save us to write same lines.
Moreover, dividing the struct allows you to retrieve a part of the struct
that you would like to use for particular purpose.
*/

type subscriber struct {
	name        string
	rate        float64
	active      bool
	homeaddress address
}

type Employee struct {
	name        string
	salary      float64
	homeaddress address
}

func main() {
	/*
		"address" is defined by using struct "address"
		above the type definition out of main function.
	*/
	address := address{
		street:     "123 Oak St.",
		city:       "Omaha",
		state:      "NE",
		postalCode: "68111",
	}
	subscriber := subscriber{name: "Aman Singh"}
	/*
		Input the struct "address" into the field "subscriber.homeaddress."
	*/
	subscriber.homeaddress = address
	fmt.Println(subscriber.homeaddress)
	fmt.Printf("%#v\n", subscriber)
	/*
		You can create structs as follows:
		name := struct name{}
	*/
	subscriber1 := Employee{}
	fmt.Printf("%#v\n", subscriber1.homeaddress)
}
