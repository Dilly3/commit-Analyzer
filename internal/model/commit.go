package model

import (
	"github.com/dilly3/houdini/internal/config"
)

var repoName = config.Config.GithubRepo
var ownerName = config.Config.GithubOwner
var since = config.Config.GithubSince

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
func SetSince(sinceDate string) {
	if len(sinceDate) < 12 {
		sinceDateZ := sinceDate + "T00:00:00Z"
		since = sinceDateZ
	} else {
		since = sinceDate
	}

}
func GetSince() string {
	return since
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
