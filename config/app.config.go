package config

import "os"

type AppConfig struct {
	DBConnectionString string
	JWTSecret          string
	HashPepper         string
	AppPort            string
	SendgridAPIKey     string
}

func NewAppConfig() *AppConfig {

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		panic("DB_CONNECTION_STRING is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET is required")
	}

	hashPepper := os.Getenv("HASH_PEPPER")
	if hashPepper == "" {
		panic("HASH_PEPPER is required")
	}

	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	if sendgridAPIKey == "" {
		panic("SENDGRID_API_KEY is required")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	return &AppConfig{
		DBConnectionString: dbConnectionString,
		JWTSecret:          jwtSecret,
		HashPepper:         hashPepper,
		AppPort:            appPort,
		SendgridAPIKey:     sendgridAPIKey,
	}
}
