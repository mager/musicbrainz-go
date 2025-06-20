package musicbrainz

type Include string
type Includes []Include

type Genre struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type URL struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
}

type ArtistCredit struct {
	Artist     *RelationArtist `json:"artist,omitempty"`
	Name       string          `json:"name"`
	JoinPhrase string          `json:"joinphrase,omitempty"`
}
