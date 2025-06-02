package musicbrainz

// area, artist, event, genre, instrument, label, place, recording, release, release-group, series, url

type Include string
type Includes []Include

type Genre struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type Relation struct {
	Type       string          `json:"type"`
	Attributes []string        `json:"attributes"`
	TargetType string          `json:"target-type"`
	Artist     *RelationArtist `json:"artist,omitempty"`
	Work       *RelationWork   `json:"work,omitempty"`
}

type RelationArtist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Disambiguation string `json:"disambiguation"`
}

type RelationWork struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Disambiguation string `json:"disambiguation"`
}

type Recording struct {
	ID               string          `json:"id"`
	Title            string          `json:"title"`
	Relations        *[]Relation     `json:"relations,omitempty"`
	Genres           *[]Genre        `json:"genres,omitempty"`
	ArtistCredits    *[]ArtistCredit `json:"artist-credit,omitempty"`
	FirstReleaseDate string          `json:"first-release-date,omitempty"`
}

type RecordingWithArtistRelations struct {
	Recording
}

type Work struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Relations *[]Relation `json:"relations,omitempty"`
}

type ArtistCredit struct {
	Artist     *RelationArtist `json:"artist,omitempty"`
	Name       string          `json:"name"`
	JoinPhrase string          `json:"joinphrase,omitempty"`
}
