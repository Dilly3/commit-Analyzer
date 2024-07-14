package storage

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

type IRepoStore interface {
	SaveRepo(ctx context.Context, repo *model.RepoInfo) error
	GetRepoByID(ctx context.Context, id string) (*model.RepoInfo, error)
	GetRepoByName(ctx context.Context, name string) (*model.RepoInfo, error)
	GetReposByLanguage(ctx context.Context, language string) ([]model.RepoInfo, error)
}

type RepoStore struct {
	storage *Storage
}

func NewRepoStore(storage *Storage) *RepoStore {
	return &RepoStore{storage: storage}
}

func (rs *RepoStore) SaveRepo(ctx context.Context, repo *model.RepoInfo) error {
	return rs.storage.DB.WithContext(ctx).FirstOrCreate(repo).Error
}

func (rs *RepoStore) GetRepoByID(ctx context.Context, id string) (*model.RepoInfo, error) {
	var repo model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("id = ?", id).First(&repo).Error
	return &repo, err
}

func (rs *RepoStore) GetRepoByName(ctx context.Context, name string) (*model.RepoInfo, error) {
	var repo model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("name = ?", name).First(&repo).Error
	return &repo, err
}

func (rs *RepoStore) GetReposByLanguage(ctx context.Context, language string) ([]model.RepoInfo, error) {
	var repos []model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("language = ?", language).Find(&repos).Error
	return repos, err
}
