package github

import (
	"github.com/dilly3/houdini/pkg/github"
)

var gitHubItr *GHubITR

type GHubITR struct {
	ghc *github.GHClient
}

func setGitHubAdp(adaptor *GHubITR) {
	gitHubItr = adaptor
}

func GetGitHubAdp() *GHubITR {
	return gitHubItr
}

// NewGHubITR sets up new github interactor
func NewGHubITR(ghc *github.GHClient) *GHubITR {
	ghi := &GHubITR{
		ghc,
	}
	setGitHubAdp(ghi)
	return ghi
}
