package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	Port             string `envconfig:"port"`
	Env              string `envconfig:"env"`
	PostgresHost     string `envconfig:"postgres_host"`
	PostgresPort     string `envconfig:"postgres_port"`
	PostgresUser     string `envconfig:"postgres_user"`
	PostgresPassword string `envconfig:"postgres_password"`
	PostgresDB       string `envconfig:"postgres_db"`
	PostgresTimezone string `envconfig:"postgres_timezone"`
	GithubBaseURL    string `envconfig:"github_base_url"`
	GithubToken      string `envconfig:"github_token"`
	GithubOwner      string `envconfig:"github_owner"`
	GithubRepo       string `envconfig:"github_repo"`
	CronInterval     string `envconfig:"cron_interval"`
}

var Config = &Configuration{}

func Init(envFile string) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	if envFile == "" {
		envFile = ".env"
	}
	log.Printf("sourcing %v", envFile)
	if err := godotenv.Load(fmt.Sprintf("%s/../../%s", basepath, envFile)); err != nil {
		log.Fatalf("couldn't load env vars: %v", err)
	}
	err := envconfig.Process("houdini", Config)
	if err != nil {
		log.Fatalf("could not process env config: %v", err)
	}
}
