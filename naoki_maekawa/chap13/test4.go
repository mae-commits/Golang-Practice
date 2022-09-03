package main

import "fmt"

// Take a channel as a parameter.
func greeting(myChannel chan string) {
	// Send a value over the channel.
	myChannel <- "hi"
}

func main() {
	// Create a new channel.
	myChannel := make(chan string)
	// Pass the channel to function running in a new goroutine.
	go greeting(myChannel)
	// Receive a value from the channel.
	fmt.Println(<-myChannel)
}
