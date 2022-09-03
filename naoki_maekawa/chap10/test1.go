package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

// Needs to be a pointer receiver, so original value can be updated.
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	// d.Year is automatically gets value at pointer.
	// Now updates original value, not a copy.
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 0 || month > 13 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 0 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	// date := Date{Year: 2019, Month: 5, Day: 27}
	// fmt.Println(date)
	date := Date{}
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
