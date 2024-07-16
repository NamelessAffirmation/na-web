package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var AppConfiguration AppConfig

type AppConfig struct {
	DBConnectionString string
	ClientUserAPIKey   string
	ClientAdminAPIKey  string
	DiscordWebhookURL  string
}

const (
	dBConnectionString = "DB_CONNECTION_STRING"
	clientUserApiKey   = "CLIENT_USER_API_KEY"
	clientAdminApiKey  = "CLIENT_ADMIN_API_KEY"
	discordWebhookUrl  = "DISCORD_WEBHOOK_URL"
)

func LoadEnvVars() (config *AppConfig) {
	envErrorMessages := []string{}

	dbConnectionString := os.Getenv(dBConnectionString)
	if dbConnectionString == "" {
		envErrorMessages = append(envErrorMessages, dBConnectionString)
	}

	clientUserAPIKey := os.Getenv(clientUserApiKey)
	if clientUserAPIKey == "" {
		envErrorMessages = append(envErrorMessages, clientUserApiKey)
	}

	clientAdminAPIKey := os.Getenv(clientAdminApiKey)
	if clientAdminAPIKey == "" {
		envErrorMessages = append(envErrorMessages, clientAdminApiKey)
	}

	discordWebhookURL := os.Getenv(discordWebhookUrl)
	if discordWebhookURL == "" {
		envErrorMessages = append(envErrorMessages, discordWebhookUrl)
	}

	if len(envErrorMessages) > 0 {
		panic(fmt.Sprintf("Missing required ENV vars: %v", envErrorMessages))
	}

	appConfig := AppConfig{}
	appConfig.DBConnectionString = dbConnectionString
	appConfig.ClientUserAPIKey = clientUserAPIKey
	appConfig.ClientAdminAPIKey = clientAdminAPIKey
	appConfig.DiscordWebhookURL = discordWebhookURL

	AppConfiguration = appConfig
	return &AppConfiguration
}
