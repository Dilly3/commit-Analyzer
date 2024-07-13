package github

import "fmt"

func (gh *GHClient) GetRepo(owner, repo string, expectedResponse interface{}) error {
	endPointURl := fmt.Sprintf("repos/%s/%s", owner, repo)
	return gh.Get(endPointURl, expectedResponse)
}
