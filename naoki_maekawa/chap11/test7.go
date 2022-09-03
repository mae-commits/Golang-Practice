package main

import (
	"fmt"
	"gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	// Assign the second return value to a variable.
	recorder, ok := player.(gadget.TapeRecoder)
	// Call the Record method only if the original value was a TapeRecorder.
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecoder{})
	TryOut(gadget.TapePlayer{})
}
