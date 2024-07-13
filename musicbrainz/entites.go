package musicbrainz

// area, artist, event, genre, instrument, label, place, recording, release, release-group, series, work, url

type Include string
type Includes []Include

type Genre struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type ArtistRelation struct {
	Type       string                `json:"type"`
	Attributes []string              `json:"attributes"`
	Artist     *ArtistRelationArtist `json:"artist"`
}

type ArtistRelationArtist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Disambiguation string `json:"disambiguation"`
}

type Recording struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Relations *[]ArtistRelation `json:"relations,omitempty"`
	Genres    *[]Genre          `json:"genres,omitempty"`
}

type RecordingWithArtistRelations struct {
	Recording
}
