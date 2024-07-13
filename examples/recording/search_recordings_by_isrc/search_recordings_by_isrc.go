package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Search Recordings by ISRC
	isrc := "USUM72401991"
	req := mb.SearchRecordingsByISRCRequest{
		ISRC: isrc,
	}
	c, err := client.SearchRecordingsByISRC(req)
	if err != nil {
		client.Log.Errorw("Error searching recordings", "isrc", isrc)
	}

	// Loop through recordings and print the ID
	for _, recording := range c.Recordings {
		client.Log.Infow("Found recording", "ID", recording.ID)
	}
}
