package github

import (
	"github.com/dilly3/houdini/pkg/github"
)

var gitHubAdp *GHubAdaptor

type GHubAdaptor struct {
	ghc *github.GHClient
}

func setGitHubAdp(adaptor *GHubAdaptor) {
	gitHubAdp = adaptor
}

func GetGitHubAdp() *GHubAdaptor {
	return gitHubAdp
}

func NewGHubAdaptor(ghc *github.GHClient) *GHubAdaptor {
	gha := &GHubAdaptor{
		ghc,
	}
	setGitHubAdp(gha)
	return gha
}
