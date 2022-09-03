package main

import "gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	// Loop over each song.
	for _, song := range songs {
		// Play the current song.
		device.Play(song)
	}
	// Stop the player once we are done.
	device.Stop()
}

func main() {
	// Update the variable to hold any Player.
	var player Player = gadget.TapePlayer{}
	// Create a slice of song titles.
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	// Play the songs using the TapePlayer.
	playList(player, mixtape)
	player = gadget.TapeRecoder{}
	mixtape = []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
}
