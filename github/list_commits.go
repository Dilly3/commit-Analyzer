package github

import (
	"github.com/dilly3/houdini/model"
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) ListCommits(owner, repo string) ([]model.CommitInfo, error) {
	var commits []interface{}
	err := gh.listCommits(owner, repo, &commits)
	if err != nil {
		return nil, err
	}
	var commitsSlice []model.CommitInfo
	for i := 0; i < len(commits); i++ {
		commit := model.CommitResponse{}
		err = mapstructure.Decode(commits[i], &commit)
		if err != nil {
			return nil, err
		}
		commitInfo := model.MapCommitResponse(&commit)
		commitInfo.ID = model.SplitID(commit.URL)
		commitsSlice = append(commitsSlice, commitInfo)
	}
	return commitsSlice, nil
}
