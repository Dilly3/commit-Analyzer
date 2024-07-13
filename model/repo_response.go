package model

type RepoResponse struct {
	Name        string `json:"name" mapstructure:"name"`
	CreatedAt   string `json:"created_at" mapstructure:"created_at"`
	URL         string `json:"html_url" mapstructure:"html_url"`
	Description string `json:"description" mapstructure:"description"`
	Language    string `json:"language" mapstructure:"language"`
	Forks       int    `json:"forks" mapstructure:"forks"`
	Stars       int    `json:"stargazers_count" mapstructure:"stargazers_count"`
	OpenIssues  int    `json:"open_issues" mapstructure:"open_issues"`
}
