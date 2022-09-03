package main

import "fmt"

type car struct {
	name     string
	topspeed int
}

func main() {
	myCar := car{name: "Corvette", topspeed: 337}
	fmt.Println("Name:", myCar.name)
	fmt.Println("Rate:", myCar.topspeed)
}
