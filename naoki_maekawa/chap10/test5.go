package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
