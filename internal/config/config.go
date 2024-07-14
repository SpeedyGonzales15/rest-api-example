package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port        string
	LogFilePath string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}
type Config struct {
	Server ServerConfig
	DB     DBConfig
}

func LoadEnvVariables() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		Server: ServerConfig{
			Port:        os.Getenv("PORT"),
			LogFilePath: os.Getenv("LOG_FILE_PATH"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
	}
	return config, nil
}
