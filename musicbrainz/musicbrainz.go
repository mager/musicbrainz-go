package musicbrainz

import (
	"fmt"
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

	// User agent configuration
	appName string
	version string
	github  string
}

// NewMusicbrainzClient creates a new Musicbrainz client with configuration.
func NewMusicbrainzClient() *MusicbrainzClient {
	logger, _ := zap.NewProduction()

	return &MusicbrainzClient{
		Log: logger.Sugar(),

		baseURL: "https://musicbrainz.org/ws/2",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		requestDelay: time.Millisecond * 250,

		// Default values
		appName: "musicbrainz-go",
		version: "0.0.14",
		github:  "https://github.com/mager/musicbrainz-go",
	}
}

// WithUserAgent configures the user agent for the client.
func (c *MusicbrainzClient) WithUserAgent(appName, version, github string) *MusicbrainzClient {
	c.appName = appName
	c.version = version
	c.github = github
	return c
}

// NewRequest creates a new request and adds authentication headers.
func (c *MusicbrainzClient) GetRequest(u *url.URL) *http.Request {
	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Set("Content-Type", "application/json")
	userAgent := fmt.Sprintf("%s/%s ( %s )", c.appName, c.version, c.github)
	req.Header.Set("User-Agent", userAgent)

	return req
}

// Get does a GET request.
func (c *MusicbrainzClient) Get(u *url.URL) (*http.Response, error) {
	req := c.GetRequest(u)
	return c.client.Do(req)
}
