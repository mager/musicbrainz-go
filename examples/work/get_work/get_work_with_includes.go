package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Metallica - One
	ID := "bad512ce-61b6-448e-8baf-47cd0b351bf4"
	req := mb.GetWorkRequest{
		ID:       ID,
		Includes: []mb.Include{"artist-rels", "url-rels"},
	}
	// Get Recording
	r, err := client.GetWork(req)
	if err != nil {
		client.Log.Errorw("Error fetching recording", "id", ID)
	}

	client.Log.Infow("Fetched work", "ID", r.ID, "Title", r.Title)
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