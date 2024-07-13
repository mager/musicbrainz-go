package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Metallica - One
	ID := "56d2735d-abc7-4070-9c3f-bc27593d922d"
	req := mb.GetRecordingRequest{
		ID:       ID,
		Includes: []mb.Include{"artist-rels", "genres"},
	}
	// Get Recording
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

	if r.Genres != nil && len(*r.Genres) > 0 {
		for _, genre := range *r.Genres {
			client.Log.Infow(
				"Recording genre",
				"ID", genre.ID,
				"Count", genre.Count,
				"Name", genre.Name,
			)
		}
	}
}
