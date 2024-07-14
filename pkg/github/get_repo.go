package github

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/storage"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) GetRepoCron() error {
	expectedResponse := map[string]interface{}{}
	err := gh.getRepo(model.GetOwnerName(), model.GetRepoName(), &expectedResponse)
	if err != nil {
		return err
	}
	resultFromRepo := model.RepoResponse{}
	err = mapstructure.Decode(expectedResponse, &resultFromRepo)
	if err != nil {
		return err
	}
	var repoData model.RepoInfo
	repoData = model.MapRepoResponse(&resultFromRepo)
	err = storage.GetDefaultStore().SaveRepo(context.Background(), &repoData)
	if err != nil {
		return err
	}
	return nil
}
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
