package main

import (
	"github.com/dilly3/houdini/api/server"
	"github.com/dilly3/houdini/internal/config"
	"github.com/dilly3/houdini/internal/storage"
	"github.com/dilly3/houdini/pkg/cron"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config.Init(".env")
	storage.New(config.Config, &logger)
	github.DefaultGHClient = github.NewGHClient(config.Config)
	handler := server.NewHandler(&logger)
	cron.InitCron()
	cron.SetCronJob(github.DefaultGHClient.GetCommitsCron, cron.GetTimeDuration())
	cron.SetCronJob(github.DefaultGHClient.GetRepoCron, cron.GetTimeDuration())
	cron.StartCronJob()
	httpHandler := server.NewMuxRouter(handler, time.Minute)
	httpServer := &http.Server{
		Addr:    config.Config.Port,
		Handler: httpHandler,
	}
	go server.GetLimiter().CleanUp()
	logger.Info().Msgf("Server started on port %s", config.Config.Port)
	if err := httpServer.ListenAndServe(); err != nil {
	}
}
