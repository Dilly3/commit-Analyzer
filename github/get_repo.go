package github

import (
	"github.com/dilly3/houdini/model"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) GetRepo(owner, repo string) (*model.RepoResponse, error) {
	expectedResponse := map[string]interface{}{}
	err := gh.getRepo(owner, repo, &expectedResponse)
	if err != nil {
		return nil, err
	}
	resultFromRepo := model.RepoResponse{}
	err = mapstructure.Decode(expectedResponse, &resultFromRepo)
	if err != nil {
		return nil, err
	}
	return &resultFromRepo, nil
}
