package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// Create two channels.
	channel1 := make(chan string)
	channel2 := make(chan string)
	// Pass each channel to a function running in a new goroutine.
	go abc(channel1)
	go def(channel2)
	// Receive and print values from the channels, in order.
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
