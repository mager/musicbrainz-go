package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Get Recording
	isrc := "USUM72401991"
	c, err := client.SearchRecordingsByISRC(isrc)
	if err != nil {
		client.Log.Errorw("Error fetching recording", "isrc", isrc)
	}

	// Loop through recordings and print the ID
	for _, recording := range c.Recordings {
		client.Log.Infow("Fetched recording", "ID", recording.ID)
	}
}
