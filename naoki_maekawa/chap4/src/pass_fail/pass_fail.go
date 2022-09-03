package main

import (
	"fmt"
	"keyboard" // import keyboard packages from .goenv/versions/1.18.3/src/keyboard
	"log"
)

func main() {
	fmt.Print("Enter a grade:")
	grade, err := keyboard.Getfloat() // use Getfloat package from keyboard folders.
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
