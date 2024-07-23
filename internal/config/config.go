package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	owner   = "github_owner"
	repo    = "github_repo"
	since   = "github_since"
	perPage = "github_per_page"
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
	GithubPerPage    string `envconfig:"github_per_page"`
	GithubSince      string `envconfig:"github_since"`
	GithubToken      string `envconfig:"github_token"`
	GithubOwner      string `envconfig:"github_owner"`
	GithubRepo       string `envconfig:"github_repo"`
	CronInterval     string `envconfig:"cron_interval"`
	NetworkRetry     int    `envconfig:"network_retry"`
	RedisHost        string `envconfig:"redis_host"`
	RedisADDR        string `envconfig:"redis_addr"`
	RedisPassword    string `envconfig:"redis_password"`
	RedisUser        string `envconfig:"redis_user"`
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

func GetSettings() map[string]string {
	return map[string]string{
		owner:   Config.GithubOwner,
		repo:    Config.GithubRepo,
		since:   Config.GithubSince,
		perPage: Config.GithubPerPage,
	}
}

func GetTimeDuration() time.Duration {
	aInt, err := strconv.Atoi(Config.CronInterval)
	if err != nil {
		log.Println("failed to convert cron interval to int")
		return 1
	}
	return time.Minute * time.Duration(aInt)

}
