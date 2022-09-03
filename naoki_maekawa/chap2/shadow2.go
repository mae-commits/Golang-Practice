package main

import (
	"fmt"
)

func main() {
	var count int = 12
	var suffix string = "minutes of bonus minutes"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol")
	fmt.Println(count, suffix, "on", format, languages)
}

/*
Naming different variables allow you to refer correct variables: otherwise, could not.
*/
