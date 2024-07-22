package github

import (
	"context"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/rs/zerolog/log"
)

func (g *GHubITR) ListCommits(owner, repo, since string) ([]model.CommitInfo, error) {
	var commitsInfo []model.CommitInfo
	res, err := g.ghc.ListCommits(owner, repo, since)
	if err != nil {
		return nil, errs.NewAppError("ListCommits:failed to decode commits,", err)
	}
	if len(res) < 1 {
		return commitsInfo, nil
	}
	commitsInfo = mapToCommitsInfo(res, repo)
	return commitsInfo, nil
}

// GetCommitsCron runs in the background to fetch commits
func (g *GHubITR) GetCommitsCron() error {
	var since string
	cmt, err := repository.GetDefaultStore().GetLastCommit(context.Background(), model.GetRepoName())
	if err != nil {
		since = model.GetSince()
	} else {
		since = cmt.Date
	}

	res, err := g.ghc.ListCommits(model.GetOwnerName(), model.GetRepoName(), since)
	if err != nil {
		log.Error().Err(err).Msg("failed to get commits")
		return errs.NewAppError("listCommitsCron:failed to get commits,", err)
	}
	var commitsSlice []model.CommitInfo
	if len(res) < 1 {
		return nil
	}
	commitsSlice = mapToCommitsInfo(res, model.GetRepoName())
	ctx := context.Background()

	err = repository.GetDefaultStore().SaveCommits(ctx, commitsSlice)
	if err != nil {
		return errs.NewAppError("listCommitsCron:failed to save commit", err)
	}

	return nil
}
