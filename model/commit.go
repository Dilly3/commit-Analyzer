package model

type CommitInfo struct {
	ID          string `gorm:"primary_key"`
	Message     string
	AuthorName  string
	AuthorEmail string
	AuthorInfo  AuthorInfo `json:"-" gorm:"foreignKey:AuthorEmail"`
	Date        string
	URL         string
}

// TableName returns the table name for the CommitInfo struct
func (CommitInfo) TableName() string {
	return "commits"
}
func MapCommitResponse(commit *CommitResponse) CommitInfo {
	id := SplitID(commit.URL)
	return CommitInfo{
		ID:          id,
		Message:     commit.Message,
		AuthorName:  commit.Author.Name,
		AuthorEmail: commit.Author.Email,
		Date:        commit.Author.Date,
		URL:         commit.URL,
	}
}
