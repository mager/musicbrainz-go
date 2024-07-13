package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Get Recording
	ID := "13f11191-f67a-4f35-8b58-62410b13d6cb"
	req := mb.GetRecordingRequest{
		ID:       ID,
		Includes: []mb.Include{"artist-rels"},
	}
	r, err := client.GetRecording(req)
	if err != nil {
		client.Log.Errorw("Error fetching recording", "id", ID)
	}

	client.Log.Infow("Fetched recording", "ID", r.ID, "Title", r.Title)
	if r.Relations != nil && len(*r.Relations) > 0 {
		for _, rel := range *r.Relations {
			if len(rel.Attributes) > 0 && rel.Artist != nil {
				client.Log.Infow(
					"Recording relation",
					"Type", rel.Type,
					"Attributes", rel.Attributes,
					"Artist", rel.Artist,
				)
			}
		}
	}
}
