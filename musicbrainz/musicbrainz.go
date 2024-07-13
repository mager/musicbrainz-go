package musicbrainz

import (
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// MusicbrainzClient represents the client for the Musicbrainz API.
type MusicbrainzClient struct {
	Log *zap.SugaredLogger

	client       *http.Client
	baseURL      string
	requestDelay time.Duration
}

// NewMusicbrainzClient creates a new Musicbrainz client with configuration.
func NewMusicbrainzClient() *MusicbrainzClient {
	logger, _ := zap.NewProduction()

	return &MusicbrainzClient{
		Log: logger.Sugar(),

		baseURL: "https://musicbrainz.com/ws/2/",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		requestDelay: time.Millisecond * 250,
	}
}

// NewRequest creates a new request and adds authentication headers.
func (c *MusicbrainzClient) GetRequest(u *url.URL) *http.Request {
	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Set("Content-Type", "application/json")

	return req
}

// Get does a GET request.
func (c *MusicbrainzClient) Get(u *url.URL) (*http.Response, error) {
	req := c.GetRequest(u)
	return c.client.Do(req)
}
