package storage

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

// ICommitStore interface
type ICommitStore interface {
	GetCommitsByRepoName(ctx context.Context, repoName string, limit int) ([]model.CommitInfo, error)
	GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error)
	SaveCommit(ctx context.Context, commit *model.CommitInfo) error
	SaveCommits(ctx context.Context, commits []model.CommitInfo) error
}

// CommitStore struct
type CommitStore struct {
	storage *Storage
}

// NewCommitStore creates a new CommitStore
func NewCommitStore(storage *Storage) *CommitStore {
	return &CommitStore{storage: storage}
}

// GetCommitsByRepoName gets commits by repo name
func (cs *CommitStore) GetCommitsByRepoName(ctx context.Context, repoName string, limit int) ([]model.CommitInfo, error) {
	var commits []model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("LOWER(repo_name) = LOWER(?)", repoName).Order("date desc").Limit(limit).Find(&commits).Error
	return commits, err
}

// GetCommitByID gets commit by id
func (cs *CommitStore) GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error) {
	var commit model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("id = ?", id).First(&commit).Error
	return &commit, err
}

// SaveCommit saves commit
func (cs *CommitStore) SaveCommit(ctx context.Context, commit *model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).FirstOrCreate(commit).Error
}

// SaveCommits saves commits
func (cs *CommitStore) SaveCommits(ctx context.Context, commits []model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).Save(commits).Error
}
