package main

import "fmt"

type Gallons float64
type Liters float64

// The number of Gallons to Liters is just over 1/4 the number of Liters.
func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

// The number of Liters is just under four times the number of Gallons.
func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	// Convert Liters to Gallons before adding.
	carFuel += ToGallons(Liters(40.0))
	// Convert Gallons to Liters before adding.
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
