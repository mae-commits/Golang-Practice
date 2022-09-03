package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	// "defer" method makes the function call until after the currect function calls.
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
