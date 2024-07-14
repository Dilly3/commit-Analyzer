package github

import (
	"github.com/dilly3/houdini/model"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) GetRepo(owner, repo string) (*model.RepoInfo, error) {
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
	var repoData model.RepoInfo
	repoData = model.MapRepoResponse(&resultFromRepo)
	return &repoData, nil
}
