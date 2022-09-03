package main

import (
	"fmt"
	"log"
)

type ComedyError string

// Define a type with an underlying type of "string".
func (c ComedyError) Error() string {
	// The Error method needs to return a string, so do a type conversion.
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

type error interface {
	Error() string
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	// Set up a variable with a type of "error".
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
	// ComedyError satisfies the error interface, so we can assign a ComedyError to the variable.
	/*
		err = ComedyError("What's programmer's favorite beer? Logger!")
		fmt.Println(err)
		err = OverheatError(35.67)
		fmt.Println(err)
	*/

}
