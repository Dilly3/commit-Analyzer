package main

import (
	"github.com/dilly3/houdini/config"
	"github.com/dilly3/houdini/github"
	"github.com/dilly3/houdini/storage"
	"github.com/rs/zerolog"
	"log"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config.Init(".env")
	storage.New(config.Config, &logger)
	ghclient := github.NewGHClient(config.Config)
	res, err := ghclient.GetRepo("dilly3", "houdini")
	if err != nil {
		panic(err)
	}
	commits, err := ghclient.ListCommits("dilly3", "houdini")
	if err != nil {
		panic(err)
	}

	log.Printf("Repo Name: %v", res.Name)
	log.Printf("Description: %v", res.Description)
	log.Printf("URL: %v", res.URL)
	log.Printf("Language: %v", res.Language)
	log.Printf("Forks: %v", res.Forks)
	log.Printf("Stars: %v", res.Stars)
	log.Printf("Open Issues: %v", res.OpenIssues)
	log.Printf("Created At: %v", res.CreatedAt)
	log.Printf("Commits: %v", commits[0].Commit.Author.Name)
	log.Printf("Commits: %v", commits[0].Commit.Message)
	log.Printf("Commits: %v", commits[1].Commit.Message)

}
