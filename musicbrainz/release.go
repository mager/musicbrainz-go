package musicbrainz

type Release struct {
	ID             string        `json:"id"`
	Title          string        `json:"title"`
	Date           string        `json:"date"`
	Country        string        `json:"country"`
	Disambiguation string        `json:"disambiguation"`
	Status         string        `json:"status"`
	Packaging      string        `json:"packaging"`
	ArtistCredit   *ArtistCredit `json:"artist-credit"`
}
