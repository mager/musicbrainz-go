package musicbrainz

// area, artist, event, genre, instrument, label, place, recording, release, release-group, series, work, url

type Include string
type Includes []Include

type ArtistRelation struct {
	Type       string                    `json:"type"`
	Attributes []ArtistRelationAttribute `json:"attributes"`
}

type ArtistRelationAttribute string

type Recording struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Relations *[]ArtistRelation `json:"relations,omitempty"`
}

type RecordingWithArtistRelations struct {
	Recording
}
