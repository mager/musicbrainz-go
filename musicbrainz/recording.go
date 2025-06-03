package musicbrainz

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SearchRecordingsByISRCRequest struct {
	ISRC string `json:"isrc"`
}

type SearchRecordingsByISRCResponse struct {
	Count      int         `json:"count"`
	Recordings []Recording `json:"recordings"`
}

func (c *MusicbrainzClient) SearchRecordingsByISRC(req SearchRecordingsByISRCRequest) (SearchRecordingsByISRCResponse, error) {
	var resp SearchRecordingsByISRCResponse
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/recording", c.baseURL))
	q := u.Query()
	q.Add("fmt", "json")
	q.Add("query", fmt.Sprintf("isrc:%s", req.ISRC))
	u.RawQuery = q.Encode()

	// Make the request
	httpResp, err := c.Get(u)

	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error getting recordings: %s", err)
		return resp, err
	}

	defer httpResp.Body.Close()
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error decoding recordings: %s", err)
		return resp, err
	}

	// Log response body
	return resp, err
}

type GetRecordingRequest struct {
	Includes Includes `json:"includes"`
	ID       string   `json:"id"`
}

type GetRecordingResponse struct {
	Recording
}

func (c *MusicbrainzClient) GetRecording(req GetRecordingRequest) (GetRecordingResponse, error) {
	var resp GetRecordingResponse
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/recording/%s", c.baseURL, req.ID))
	q := u.Query()
	q.Add("fmt", "json")
	incs := make([]string, 0)
	if IncludesContains(req.Includes, "artist-rels") {
		incs = append(incs, "artist-rels")
	}
	if IncludesContains(req.Includes, "artist-credits") {
		incs = append(incs, "artist-credits")
	}
	if IncludesContains(req.Includes, "genres") {
		incs = append(incs, "genres")
	}
	if IncludesContains(req.Includes, "work-rels") {
		incs = append(incs, "work-rels")
	}
	if IncludesContains(req.Includes, "releases") {
		incs = append(incs, "releases")
	}
	if len(incs) > 0 {
		q.Add("inc", strings.Join(incs, "+"))
	}
	u.RawQuery = q.Encode()

	// Log the full URL
	c.Log.Infof("Making request to URL: %s", u.String())

	// Make the request
	httpResp, err := c.Get(u)

	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error getting recording: %s", err)
		return resp, err
	}

	defer httpResp.Body.Close()
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error decoding recording: %s", err)
		return resp, err
	}

	// Log response body
	return resp, err
}

// SearchRecordingsByArtistAndTrack searches for recordings by artist and track title.
type SearchRecordingsByArtistAndTrackRequest struct {
	Artist string `json:"artist"`
	Track  string `json:"track"`
}

type SearchRecordingsByArtistAndTrackResponse struct {
	Count      int         `json:"count"`
	Recordings []Recording `json:"recordings"`
}

func (c *MusicbrainzClient) SearchRecordingsByArtistAndTrack(req SearchRecordingsByArtistAndTrackRequest) (SearchRecordingsByArtistAndTrackResponse, error) {
	var resp SearchRecordingsByArtistAndTrackResponse

	// Construct the search query:
	query := fmt.Sprintf("artist:%s AND recording:%s", url.QueryEscape(req.Artist), url.QueryEscape(req.Track))

	u, err := url.Parse(fmt.Sprintf("%s/recording", c.baseURL))
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return resp, fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	q.Add("fmt", "json")
	q.Add("query", query)
	u.RawQuery = q.Encode()

	// Make the request
	httpResp, err := c.Get(u)
	if err != nil {
		log.Printf("Error getting recordings: %v", err)
		return resp, fmt.Errorf("error getting recordings: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		log.Printf("MusicBrainz API returned status code: %d", httpResp.StatusCode)
		return resp, fmt.Errorf("MusicBrainz API returned status code %d", httpResp.StatusCode)
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		log.Printf("Error decoding recordings: %v", err)
		return resp, fmt.Errorf("error decoding recordings: %w", err)
	}

	return resp, nil
}
