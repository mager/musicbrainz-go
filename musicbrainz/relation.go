package musicbrainz

type Relation struct {
	Type       string          `json:"type"`
	Attributes []string        `json:"attributes"`
	TargetType string          `json:"target-type"`
	Artist     *RelationArtist `json:"artist,omitempty"`
	Work       *RelationWork   `json:"work,omitempty"`
	URL        URL             `json:"url,omitempty"`
}

type RelationArtist struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	Genres         *[]Genre `json:"genres,omitempty"`
	JoinPhrase     string   `json:"joinphrase"`
}

type RelationWork struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Disambiguation string `json:"disambiguation"`
}
