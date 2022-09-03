package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	// Pause the main goroutine for 1 second.
	// The time.Sleep call delays the exit long enough for the other goroutines to run.
	time.Sleep(time.Second)
	fmt.Println("end main()")
}
