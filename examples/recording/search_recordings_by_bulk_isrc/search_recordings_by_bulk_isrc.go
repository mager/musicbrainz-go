package main

import (
	"fmt"
	"log"

	mb "github.com/mager/musicbrainz-go/musicbrainz"
)

func main() {
	// Create a new MusicBrainz client
	client := mb.NewMusicbrainzClient().
		WithUserAgent("bulk-isrc-example", "1.0.0", "https://github.com/mager/musicbrainz-go")

	// Example ISRCs to search for (these are real ISRCs from MusicBrainz)
	isrcs := []string{
		"GBUM71402401", // The Beatles - Let It Be
		"GBUM71800601", // The Beatles - Hey Jude
		"GBUM71800602", // The Beatles - Revolution
		"GBUM71800603", // The Beatles - Paperback Writer
		"USRC12345678", // Example ISRC (will likely not be found)
	}

	fmt.Printf("Searching for recordings with ISRCs: %v\n", isrcs)

	// Create the bulk search request
	req := mb.SearchRecordingsByBulkISRCRequest{
		ISRCs: isrcs,
	}

	// Perform the bulk search
	resp, err := client.SearchRecordingsByBulkISRC(req)
	if err != nil {
		log.Fatalf("Error performing bulk ISRC search: %v", err)
	}

	fmt.Printf("\nSearch completed!\n")
	fmt.Printf("Total recordings found: %d\n", resp.Count)
	fmt.Printf("ISRCs searched: %d\n", len(isrcs))

	// Display results
	if resp.Count > 0 {
		fmt.Printf("\nFound recordings:\n")
		fmt.Printf("================\n")

		for i, recording := range resp.Recordings {
			fmt.Printf("%d. ID: %s\n", i+1, recording.ID)
			fmt.Printf("   Title: %s\n", recording.Title)

			if recording.ArtistCredits != nil && len(*recording.ArtistCredits) > 0 {
				fmt.Printf("   Artist: %s\n", (*recording.ArtistCredits)[0].Name)
			}

			if recording.FirstReleaseDate != "" {
				fmt.Printf("   Release Date: %s\n", recording.FirstReleaseDate)
			}

			if recording.Genres != nil && len(*recording.Genres) > 0 {
				fmt.Printf("   Genres: ")
				for j, genre := range *recording.Genres {
					if j > 0 {
						fmt.Printf(", ")
					}
					fmt.Printf("%s", genre.Name)
				}
				fmt.Printf("\n")
			}

			fmt.Printf("\n")
		}
	} else {
		fmt.Printf("\nNo recordings found for the provided ISRCs.\n")
	}

	// Show ISRC mapping
	if len(resp.ISRCMap) > 0 {
		fmt.Printf("\nISRC to Recording Mapping:\n")
		fmt.Printf("==========================\n")
		for isrc, recordings := range resp.ISRCMap {
			fmt.Printf("ISRC: %s\n", isrc)
			for i, recording := range recordings {
				fmt.Printf("  %d. %s - %s\n", i+1, recording.Title, getArtistName(recording))
			}
			fmt.Printf("\n")
		}
	} else {
		fmt.Printf("\nNo ISRC mapping available (ISRCs not included in response).\n")
	}

	fmt.Printf("\nNote: This bulk search efficiently finds multiple recordings in a single request.\n")
	fmt.Printf("The ISRC mapping shows which recordings correspond to which ISRCs.\n")
}

// Helper function to get artist name from recording
func getArtistName(recording mb.Recording) string {
	if recording.ArtistCredits != nil && len(*recording.ArtistCredits) > 0 {
		return (*recording.ArtistCredits)[0].Name
	}
	return "Unknown Artist"
}
