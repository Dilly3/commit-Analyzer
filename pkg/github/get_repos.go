package github

import (
	"github.com/dilly3/houdini/internal/model"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) GetRepos() ([]model.RepoInfo, error) {
	var expectedResponse []interface{}
	err := gh.getRepos(&expectedResponse)
	if err != nil {
		return nil, err
	}
	var resultFromRepos []model.RepoInfo
	for i := 0; i < len(expectedResponse); i++ {
		repo := model.RepoResponse{}
		err = mapstructure.Decode(expectedResponse[i], &repo)
		if err != nil {
			return nil, err
		}
		resultFromRepos = append(resultFromRepos, model.MapRepoResponse(&repo))

	}
	return resultFromRepos, nil

}
