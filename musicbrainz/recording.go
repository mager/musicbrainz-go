package musicbrainz

import (
	"encoding/json"
	"fmt"
	"net/url"
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
	if IncludesContains(req.Includes, "artist-rels") {
		q.Add("inc", "artist-rels")
	}
	if IncludesContains(req.Includes, "genres") {
		q.Add("inc", "genres")
	}
	u.RawQuery = q.Encode()

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
