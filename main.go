package main

import (
	"github.com/dilly3/houdini/api/server"
	"github.com/dilly3/houdini/internal/config"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/pkg/cron"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/dilly3/houdini/pkg/postgres"
	"github.com/rs/zerolog"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"net/http"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config.Init(".env")
	pgs := postgres.New(config.Config, &logger)
	repository.NewStore(pgs)
	model.SetOwnerName(config.Config.GithubOwner)
	model.SetRepoName(config.Config.GithubRepo)
	model.SetSince(config.Config.GithubSince)
	github.DefaultGHClient = github.NewGHClient(config.Config)
	handler := server.NewHandler(&logger)
	cron.InitCron()
	cron.SetCronJob(github.DefaultGHClient.GetCommitsCron, cron.GetTimeDuration())
	cron.SetCronJob(github.DefaultGHClient.GetRepoCron, cron.GetTimeDuration())
	go cron.StartCronJob()
	httpHandler := server.NewChiRouter(handler, time.Minute)
	httpServer := &http.Server{
		Addr:    config.Config.Port,
		Handler: httpHandler,
	}
	go server.GetLimiter().CleanUp()
	logger.Info().Msgf("Server started on port %s", config.Config.Port)
	if err := httpServer.ListenAndServe(); err != nil {
	}
}
