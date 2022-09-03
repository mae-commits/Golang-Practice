package main

import "fmt"

// Define variable "Car" as a string type.
type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Speeding Up")
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Brake()
	Steer(string)
	Accelerate()
}

func main() {
	var vehicle Vehicle = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}
