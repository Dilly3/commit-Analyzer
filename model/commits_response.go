package model

type Commit struct {
	Commit CommitDetails `json:"commit" mapstructure:"commit"`
}
type CommitDetails struct {
	Message string `json:"message" mapstructure:"message"`
	Author  Author `json:"author" mapstructure:"author"`
	Date    string `json:"date" mapstructure:"date"`
	URL     string `json:"url" mapstructure:"url"`
}
type Author struct {
	Name  string `json:"name" mapstructure:"name"`
	Email string `json:"email" mapstructure:"email"`
	Date  string `json:"date" mapstructure:"date"`
}
