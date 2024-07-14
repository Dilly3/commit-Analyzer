package github

import (
	"github.com/dilly3/houdini/internal/config"
	"net/http"
	"time"
)

type GHClient struct {
	BaseURL    string
	token      string
	HTTPClient *http.Client
}

var (
	DefaultGHClient *GHClient
)

func NewGHClient(config *config.Configuration) *GHClient {
	return &GHClient{
		BaseURL: config.GithubBaseURL,
		token:   config.GithubToken,
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}
