package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
