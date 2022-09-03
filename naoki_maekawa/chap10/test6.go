package main

import (
	"calender"
	"fmt"
)

func main() {
	// date := Date{Year: 2019, Month: 5, Day: 27}
	// fmt.Println(date)
	date := calender.Date{}
	date.year = 2019
	date.month = 14
	date.day = 50
	fmt.Println(date)
}
