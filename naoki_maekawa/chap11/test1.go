package main

import "gadget"

func playList(device gadget.TapePlayer, songs []string) {
	// Loop over each song.
	for _, song := range songs {
		// Play the current song.
		device.Play(song)
	}
	// Stop the player once we are done.
	device.Stop()
}

func main() {
	// Create a TapePlayer.
	player := gadget.TapePlayer{}
	// Create a slice of song titles.
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	// Play the songs using the TapePlayer.
	playList(player, mixtape)
	player := gadget.TapeRecoder{}
	mixTape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
}
