package main

import (
	mb "github.com/mager/musicbrainz-go"
)

func main() {
	// Create a new client
	client := mb.NewMusicBrainzClient()

	// Get Recording
	isrc := "USUM72401991"
	c, err := client.SearchRecordingsByISRC(isrc)
	if err != nil {
		// TODO: Handle error
	}

	// Loop through recordings and print the ID
	for _, recording := range c.recordings {
		client.Log.Infow("Fetched recording", "name", recording.ID)
	}
}
