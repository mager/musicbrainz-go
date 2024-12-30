package musicbrainz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type GetWorkRequest struct {
	Includes Includes `json:"includes"`
	ID       string   `json:"id"`
}

type GetWorkResponse struct {
	Work
}

func (c *MusicbrainzClient) GetWork(req GetWorkRequest) (GetWorkResponse, error) {
	var resp GetWorkResponse
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/work/%s", c.baseURL, req.ID))
	q := u.Query()
	q.Add("fmt", "json")
	incs := make([]string, 0)
	if IncludesContains(req.Includes, "artist-rels") {
		incs = append(incs, "artist-rels")
	}
	if IncludesContains(req.Includes, "url-rels") {
		incs = append(incs, "url-rels")
	}
	if len(incs) > 0 {
		q.Add("inc", strings.Join(incs, "+"))
	}
	u.RawQuery = q.Encode()

	// Make the request
	httpResp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error getting work: %s", err)
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
