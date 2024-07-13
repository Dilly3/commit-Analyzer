package github

import "fmt"

func (gh *GHClient) ListCommits(owner, repo string, expectedResponse interface{}) error {
	endPointURl := fmt.Sprintf("repos/%s/%s/commits", owner, repo)
	return gh.Get(endPointURl, expectedResponse)
}
