package storage

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

type ICommitStore interface {
	// GetCommitsByRepoName retrieves commits by repo ID
	GetCommitsByRepoName(ctx context.Context, repoName string) ([]model.CommitInfo, error)
	GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error)
	SaveCommit(ctx context.Context, commit *model.CommitInfo) error
	SaveCommits(ctx context.Context, commits []model.CommitInfo) error
}

type CommitStore struct {
	storage *Storage
}

func NewCommitStore(storage *Storage) *CommitStore {
	return &CommitStore{storage: storage}
}

func (cs *CommitStore) GetCommitsByRepoName(ctx context.Context, repoName string) ([]model.CommitInfo, error) {
	var commits []model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("repo_name = ?", repoName).Find(&commits).Error
	return commits, err
}

func (cs *CommitStore) GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error) {
	var commit model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("id = ?", id).First(&commit).Error
	return &commit, err
}

func (cs *CommitStore) SaveCommit(ctx context.Context, commit *model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).FirstOrCreate(commit).Error
}

func (cs *CommitStore) SaveCommits(ctx context.Context, commits []model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).Save(commits).Error
}
