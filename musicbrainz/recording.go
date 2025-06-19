package musicbrainz

import (
	"encoding/json"
	"fmt"
	"io"
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

	c.Log.Debugf("Making request to URL: %s", u.String())

	// Make the request
	httpResp, err := c.Get(u)

	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error getting recordings: %s", err)
		return resp, err
	}

	defer httpResp.Body.Close()
	bodyBytes, readErr := io.ReadAll(httpResp.Body)
	if readErr != nil {
		c.Log.Errorf("Error reading response body: %s", readErr)
		return resp, readErr
	}
	c.Log.Debugf("Response body: %s", string(bodyBytes))
	err = json.NewDecoder(strings.NewReader(string(bodyBytes))).Decode(&resp)
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
	if IncludesContains(req.Includes, "url-rels") {
		incs = append(incs, "url-rels")
	}
	if len(incs) > 0 {
		q.Add("inc", strings.Join(incs, "+"))
	}
	u.RawQuery = q.Encode()

	// Log the full URL
	c.Log.Debugf("Making request to URL: %s", u.String())

	// Make the request
	httpResp, err := c.Get(u)

	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error getting recording: %s", err)
		return resp, err
	}

	defer httpResp.Body.Close()
	bodyBytes, readErr := io.ReadAll(httpResp.Body)
	if readErr != nil {
		c.Log.Errorf("Error reading response body: %s", readErr)
		return resp, readErr
	}
	c.Log.Debugf("Response body: %s", string(bodyBytes))
	err = json.NewDecoder(strings.NewReader(string(bodyBytes))).Decode(&resp)
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
	var err error

	// Construct the search query with phrase quoting.
	query := fmt.Sprintf("artist:\"%s\" AND recording:\"%s\"", req.Artist, req.Track)

	u, err := url.Parse(fmt.Sprintf("%s/recording", c.baseURL))
	if err != nil {
		c.Log.Errorf("Error parsing URL for recording search: %v", err)
		return resp, fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	q.Add("fmt", "json")
	q.Add("query", query)
	q.Add("limit", "25")
	u.RawQuery = q.Encode()

	c.Log.Debugf("Making request to URL: %s", u.String())

	httpResp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error making HTTP request to MusicBrainz: %v", err)
		return resp, fmt.Errorf("error getting recordings: %w", err)
	}
	defer httpResp.Body.Close()
	bodyBytes, readErr := io.ReadAll(httpResp.Body)
	if readErr != nil {
		c.Log.Errorf("Error reading response body: %s", readErr)
		return resp, readErr
	}
	c.Log.Debugf("Response body: %s", string(bodyBytes))
	err = json.NewDecoder(strings.NewReader(string(bodyBytes))).Decode(&resp)
	if err != nil {
		c.Log.Errorf("Error decoding recordings: %s", err)
		return resp, err
	}

	if httpResp.StatusCode != http.StatusOK {
		c.Log.Infof("MusicBrainz API returned status code: %d for ArtistAndTrack search", httpResp.StatusCode) // Clarified log message

		bodyBytes, readErr := io.ReadAll(httpResp.Body)
		if readErr != nil {
			c.Log.Errorf("Error reading error response body from MusicBrainz: %v", readErr)
		}
		errMsg := fmt.Sprintf("MusicBrainz API returned non-OK status: %d. Response: %s", httpResp.StatusCode, string(bodyBytes))
		c.Log.Errorf(errMsg)
		return resp, fmt.Errorf("%s", errMsg)
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		c.Log.Errorf("Error decoding MusicBrainz recording search response: %v", err)
		return resp, fmt.Errorf("error decoding recordings: %w", err)
	}

	return resp, nil
}
