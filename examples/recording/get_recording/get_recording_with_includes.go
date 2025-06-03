package main

import (
	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new client
	client := mb.NewMusicbrainzClient()

	// Kendrick Lamar - squabble up
	ID := "451690ff-5d48-4724-bd69-56951dbbf733"
	req := mb.GetRecordingRequest{
		ID: ID,
		Includes: []mb.Include{
			"artist-credits",
			"artist-rels",
			"genres",
			"work-rels",
			"releases",
		},
	}
	// Get Recording
	r, err := client.GetRecording(req)
	if err != nil {
		client.Log.Errorw("Error fetching recording", "id", ID)
	}

	client.Log.Infow("Fetched recording", "ID", r.ID, "Title", r.Title)
	if r.Relations != nil && len(*r.Relations) > 0 {
		for _, rel := range *r.Relations {
			fields := []interface{}{
				"Target Type", rel.TargetType,
				"Type", rel.Type,
				"Attributes", rel.Attributes,
			}

			if rel.Artist != nil {
				fields = append(fields, "Artist", rel.Artist)
			}
			if rel.Work != nil {
				fields = append(fields, "Work", rel.Work)
			}

			client.Log.Infow("Recording relation", fields...)
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
	if r.ArtistCredits != nil && len(*r.ArtistCredits) > 0 {
		for _, credit := range *r.ArtistCredits {
			client.Log.Infow("Recording artist credit", "Artist", credit.Artist, "Name", credit.Name, "JoinPhrase", credit.JoinPhrase)
		}
	}

	client.Log.Infow("Recording first release date", "Date", r.FirstReleaseDate)

	if r.Releases != nil && len(*r.Releases) > 0 {
		for _, release := range *r.Releases {
			client.Log.Infow("Recording release", "ID", release.ID, "Title", release.Title, "Date", release.Date, "Country", release.Country)
		}
	}
}
