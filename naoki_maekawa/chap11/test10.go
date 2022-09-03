package main

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

type Anything interface {
}

type Whistle string

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("Toyco Canary"))
}
