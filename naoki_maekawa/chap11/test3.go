package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Powering legs")
}

func play(n NoiseMaker) {
	// Call only methods that are part of the interface.
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
	play(Robot("Botco Ambler"))
}
