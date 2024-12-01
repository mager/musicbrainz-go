package main

import (
	"fmt"
	"log"

	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new MusicBrainz client. Make sure logging is properly configured.
	client := mb.NewMusicbrainzClient()

	// Search for recordings
	artist := "The Beatles"
	track := "Hey Jude"
	req := mb.SearchRecordingsByArtistAndTrackRequest{
		Artist: artist,
		Track:  track,
	}

	resp, err := client.SearchRecordingsByArtistAndTrack(req)
	if err != nil {
		log.Printf("Error searching recordings: %v", err)
		return
	}

	fmt.Printf("Found %d recordings matching '%s' by '%s':\n", len(resp.Recordings), track, artist)
	for _, recording := range resp.Recordings {
		fmt.Printf("  - ID: %s, Title: %s\n", recording.ID, recording.Title)
	}
}
