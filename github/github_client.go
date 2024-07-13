package github

import (
	"github.com/dilly3/houdini/config"
	"net/http"
	"time"
)

//-H "Accept: application/vnd.github+json" \
//-H "Authorization: Bearer <YOUR-TOKEN>" \
//-H "X-GitHub-Api-Version: 2022-11-28" \

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
