package github

import (
	"github.com/dilly3/houdini/model"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) ListCommits(owner, repo string) ([]model.Commit, error) {
	var commits []interface{}
	err := gh.listCommits(owner, repo, &commits)
	if err != nil {
		return nil, err
	}
	var commitsSlice []model.Commit
	for i := 0; i < len(commits); i++ {
		commit := model.Commit{}
		err = mapstructure.Decode(commits[i], &commit)
		if err != nil {
			return nil, err
		}
		commitsSlice = append(commitsSlice, commit)
	}
	return commitsSlice, nil
}
