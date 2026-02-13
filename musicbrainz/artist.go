package musicbrainz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Artist entity
type Artist struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	SortName       string      `json:"sort-name"`
	Type           string      `json:"type"`
	Gender         string      `json:"gender,omitempty"`
	Disambiguation string      `json:"disambiguation"`
	Country        string      `json:"country"`
	Area           *Area       `json:"area,omitempty"`
	BeginArea      *Area       `json:"begin-area,omitempty"`
	LifeSpan       *LifeSpan   `json:"life-span,omitempty"`
	Genres         *[]Genre    `json:"genres,omitempty"`
	Relations      *[]Relation `json:"relations,omitempty"`
}

type Area struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LifeSpan struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
	Ended bool   `json:"ended"`
}

type GetArtistRequest struct {
	ID       string   `json:"id"`
	Includes Includes `json:"includes"`
}

type GetArtistResponse struct {
	Artist
}

func (c *MusicbrainzClient) GetArtist(req GetArtistRequest) (GetArtistResponse, error) {
	var resp GetArtistResponse
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/artist/%s", c.baseURL, req.ID))
	q := u.Query()
	q.Add("fmt", "json")
	incs := make([]string, 0)
	if IncludesContains(req.Includes, "genres") {
		incs = append(incs, "genres")
	}
	if IncludesContains(req.Includes, "url-rels") {
		incs = append(incs, "url-rels")
	}
	if IncludesContains(req.Includes, "recording-rels") {
		incs = append(incs, "recording-rels")
	}
	if IncludesContains(req.Includes, "release-rels") {
		incs = append(incs, "release-rels")
	}
	if IncludesContains(req.Includes, "work-rels") {
		incs = append(incs, "work-rels")
	}
	if len(incs) > 0 {
		q.Add("inc", strings.Join(incs, "+"))
	}
	u.RawQuery = q.Encode()

	c.Log.Debugf("Making request to URL: %s", u.String())

	// Make the request
	httpResp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error getting artist: %s", err)
		return resp, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return resp, fmt.Errorf("failed to decode response: %w", err)
	}

	return resp, nil
}

type SearchArtistsRequest struct {
	Query string `json:"query"`
}

type SearchArtistsResponse struct {
	Count   int      `json:"count"`
	Artists []Artist `json:"artists"`
}

func (c *MusicbrainzClient) SearchArtists(req SearchArtistsRequest) (SearchArtistsResponse, error) {
	var resp SearchArtistsResponse
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/artist", c.baseURL))
	q := u.Query()
	q.Add("fmt", "json")
	q.Add("query", req.Query)
	u.RawQuery = q.Encode()

	c.Log.Debugf("Making request to URL: %s", u.String())

	httpResp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error searching artists: %s", err)
		return resp, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return resp, fmt.Errorf("failed to decode response: %w", err)
	}

	return resp, nil
}
