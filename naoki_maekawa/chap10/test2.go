package main

import (
	"calender"
	"fmt"
	"log"
)

func main() {
	// date := Date{Year: 2019, Month: 5, Day: 27}
	// fmt.Println(date)
	date := calender.Date{}
	// Automatically converted to a pointer.
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)

	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
