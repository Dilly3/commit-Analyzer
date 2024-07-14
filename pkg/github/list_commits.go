package github

import (
	"context"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/storage"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
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
			return nil, errs.NewAppError("ListCommits:failed to decode commits,", err)
		}
		commitInfo := model.MapCommitResponse(&commit, repo)
		commitInfo.ID = model.SplitID(commit.URL)
		commitsSlice = append(commitsSlice, commitInfo)
	}
	return commitsSlice, nil
}

func (gh *GHClient) GetCommitsCron() error {
	var commits []interface{}
	err := gh.listCommits(model.GetOwnerName(), model.GetRepoName(), &commits)
	if err != nil {
		log.Error().Err(err).Msg("failed to get commits")
		return errs.NewAppError("listCommitsCron:failed to get commits,", err)
	}
	var commitsSlice []model.CommitInfo
	for i := 0; i < len(commits); i++ {
		commit := model.CommitResponse{}
		err = mapstructure.Decode(commits[i], &commit)
		if err != nil {
			log.Error().Err(err).Msg("listCommitsCron:failed to decode commits")
			return errs.NewAppError("listCommitsCron:failed to decode commits,", err)
		}
		commitInfo := model.MapCommitResponse(&commit, model.GetRepoName())
		commitInfo.ID = model.SplitID(commit.URL)
		commitsSlice = append(commitsSlice, commitInfo)
	}
	ctx := context.Background()
	for i := 0; i < len(commitsSlice); i++ {
		err = storage.GetDefaultStore().SaveCommit(ctx, &commitsSlice[i])
		if err != nil {
			return errs.NewAppError("listCommitsCron:failed to save commit", err)
		}
	}
	return nil
}
