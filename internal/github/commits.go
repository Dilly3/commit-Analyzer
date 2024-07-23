package github

import (
	"context"
	"github.com/dilly3/houdini/internal/config"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/repository/cache"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (g *GHubITR) ListCommits(owner, repo, since string, page int) ([]model.CommitInfo, error) {
	var commitsInfo []model.CommitInfo
	perPage := cache.GetDefaultCache().GetPerPage()
	perP, err := strconv.Atoi(perPage)
	if err != nil {
		return nil, errs.NewAppError("ListCommits:failed to convert perPage to int", err)
	}

	res, err := g.ghc.ListCommits(owner, repo, since, perP, page)
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
	var since *string
	cac := cache.GetDefaultCache()
	cmt, err := repository.GetDefaultStore().GetLastCommit(context.Background(), cac.GetRepo())
	if err != nil {
		s := cache.GetDefaultCache().GetSince()
		since = &s
	} else {
		since = &cmt.Date
	}
	perP, err := strconv.Atoi(cac.GetPerPage())

	if err != nil {
		return errs.NewAppError("GetCommitsCron:failed to convert perPage to int", err)
	}
	log.Info().Msg("fetching commits for repo:: " + cac.GetRepo())
	completeChan := make(chan bool)
	responseChan := make(chan []github.CommitResponse)
	arrInt := []int{1}
	tm := config.GetTimeDuration()
	// fetch commits in the background
	go func(chan bool, chan []github.CommitResponse, []int, *string, time.Duration) {
		startTime := time.Now()
		for {
			log.Info().Msg("fetching commits for page:: " + strconv.Itoa(arrInt[0]))
			time.Sleep(10 * time.Second)
			res, err := g.ghc.ListCommits(cac.GetOwner(), cac.GetRepo(), *since, perP, arrInt[0])
			if err != nil {
				log.Error().Err(err).Msg("failed to get commits")
				return
			}
			responseChan <- res
			if len(res) < 1 {
				completeChan <- true
				return
			}

			if time.Since(startTime) > tm-time.Minute {
				completeChan <- true
				return
			}
			arrInt[0]++

		}

	}(completeChan, responseChan, arrInt, since, tm)

	// listen for the response from fetched commits
	for {
		select {
		case <-completeChan:
			break
		case res := <-responseChan:
			var commitsSlice []model.CommitInfo
			if len(res) < 1 {
				break
			}
			commitsSlice = mapToCommitsInfo(res, cac.GetRepo())
			ctx := context.Background()
			err = repository.GetDefaultStore().SaveCommits(ctx, commitsSlice)
			if err != nil {
				log.Error().Err(err).Msg("failed to save commit")
				return errs.NewAppError("GetCommitsCron:failed to save commit", err)
			}
		}
	}

}
