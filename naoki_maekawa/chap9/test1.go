package main

import "fmt"

/*
Define two new types, each with an underlying type of float64.
*/
type Liters float64
type Gallons float64

func main() {
	/*
		Define a variable with a type of Gallons.
	*/
	var carFuel Gallons
	/*
		Define a variable with a type of Liters.
	*/
	var busFuel Liters
	/*
		Convert a float64 to Gallons.
	*/
	carFuel = Gallons(10.0)
	/*
		Convert a float64 to Liters.
	*/
	busFuel = Liters(240.0)
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}
