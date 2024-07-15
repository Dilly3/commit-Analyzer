package model

import "github.com/dilly3/houdini/internal/config"

var repoName = config.Config.GithubRepo
var ownerName = config.Config.GithubOwner

func SetRepoName(name string) {
	repoName = name
}
func SetOwnerName(name string) {
	ownerName = name
}
func GetRepoName() string {
	return repoName
}
func GetOwnerName() string {
	return ownerName
}

type CommitInfo struct {
	ID          string `gorm:"primary_key"`
	RepoName    string `json:"repo_name"`
	Message     string `gorm:"index"`
	AuthorName  string
	AuthorEmail string
	Date        string
	URL         string
}

// TableName returns the table name for the CommitInfo struct
func (CommitInfo) TableName() string {
	return "commits"
}
func MapCommitResponse(commit *CommitResponse, repoName string) CommitInfo {
	id := SplitID(commit.URL)
	return CommitInfo{
		ID:          id,
		Message:     commit.Message,
		AuthorName:  commit.Author.Name,
		AuthorEmail: commit.Author.Email,
		Date:        commit.Author.Date,
		URL:         commit.URL,
		RepoName:    repoName,
	}
}
