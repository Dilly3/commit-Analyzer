package main

import (
	"github.com/dilly3/houdini/config"
	"github.com/dilly3/houdini/github"
	"github.com/mitchellh/mapstructure"
	"log"
)

// repository name,
// description,
// URL,
// language,
// forks count,
// stars count,
// open issues count,
// watchers count,
// created/updated dates.
// log.Printf("Repo: %v", body)
type response struct {
	Name        string `json:"name" mapstructure:"name"`
	CreatedAt   string `json:"created_at" mapstructure:"created_at"`
	URL         string `json:"html_url" mapstructure:"html_url"`
	Description string `json:"description" mapstructure:"description"`
	Language    string `json:"language" mapstructure:"language"`
	Forks       int    `json:"forks" mapstructure:"forks"`
	Stars       int    `json:"stargazers_count" mapstructure:"stargazers_count"`
	OpenIssues  int    `json:"open_issues" mapstructure:"open_issues"`
}
type commit struct {
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

//commit message, author, date, and URL

func main() {
	config.Init(".env")
	ghclient := github.NewGHClient(config.Config)
	body := map[string]interface{}{}
	var body2 []interface{}
	err := ghclient.GetRepo("dilly3", "houdini", &body)
	if err != nil {
		panic(err)
	}
	err = ghclient.ListCommits("dilly3", "houdini", &body2)
	if err != nil {
		panic(err)

	}
	result := response{}
	result2 := commit{}
	err = mapstructure.Decode(body, &result)
	if err != nil {
		panic(err)
	}
	err = mapstructure.Decode(body2[0], &result2)
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	log.Printf("Repo Name: %v", result.Name)
	log.Printf("Description: %v", result.Description)
	log.Printf("URL: %v", result.URL)
	log.Printf("Language: %v", result.Language)
	log.Printf("Forks: %v", result.Forks)
	log.Printf("Stars: %v", result.Stars)
	log.Printf("Open Issues: %v", result.OpenIssues)
	log.Printf("Created At: %v", result.CreatedAt)
	//log.Printf("Commits: %v", result2.Message)
	log.Printf("Commits: %v", result2.Commit.Author.Email)

}
