package config

import "os"

type Env struct {
	AppPort                 string
	GinMode                 string
	DbHost                  string
	DbPort                  string
	DbUser                  string
	DbPass                  string
	DbName                  string
	KafkaBroker             string
	KafkaGithubCommitsTopic string
	GithubApiKey            string
	GithubCommitsApi        string
}

var AppEnv Env

func InitAppEnv() {
	AppEnv.AppPort = getOrDefault("APP_PORT", "9191")
	AppEnv.GinMode = getOrDefault("GIN_MODE", "debug")

	AppEnv.DbHost = getOrDefault("DB_HOST", "localhost")
	AppEnv.DbPort = getOrDefault("DB_PORT", "5432")
	AppEnv.DbUser = getOrDefault("DB_USER", "postgres")
	AppEnv.DbPass = getOrDefault("DB_PASS", "postgres")
	AppEnv.DbName = getOrDefault("DB_NAME", "github")

	AppEnv.KafkaBroker = getOrDefault("KAFKA_BROKER", "localhost:9092")
	AppEnv.KafkaGithubCommitsTopic = getOrDefault("KAFKA_GITHUB_COMMITS_TOPIC", "github_commits")

	AppEnv.GithubApiKey = getOrDefault("GITHUB_API_KEY", "")
	AppEnv.GithubCommitsApi = getOrDefault("GITHUB_COMMITS_API", "https://api.github.com/repos/matplotlib/matplotlib/commits")
}

func getOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
