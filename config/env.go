package config

import "os"

type Env struct {
	AppPort                 string
	GinMode                 string
	KafkaBroker             string
	KafkaGithubCommitsTopic string
}

var AppEnv Env

func InitAppEnv() {
	AppEnv.AppPort = getOrDefault("APP_PORT", "9191")
	AppEnv.GinMode = getOrDefault("GIN_MODE", "debug")

	AppEnv.KafkaBroker = getOrDefault("KAFKA_BROKER", "localhost:9092")
	AppEnv.KafkaGithubCommitsTopic = getOrDefault("KAFKA_GITHUB_COMMITS_TOPIC", "github_commits")
}

func getOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
