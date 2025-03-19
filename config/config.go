package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GithubName       string
	VkToken          string
	VkUserID         string
	CodeforcesHandle string
	Port             string
	AppUrl           string
	ReactPort        string
	ENV              string
	GithubToken      string
}

func LoadConfig() Config {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Ошибка загрузки .env файла")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return Config{
		Port:             port,
		GithubName:       os.Getenv("GITHUB_NAME"),
		GithubToken:      os.Getenv("GITHUB_TOKEN"),
		VkToken:          os.Getenv("VK_KEY"),
		VkUserID:         os.Getenv("VK_USER_ID"),
		CodeforcesHandle: os.Getenv("CODEFORCES_HANDLE"),
		AppUrl:           os.Getenv("APP_URL"),
		ReactPort:        os.Getenv("REACT_PORT"),
		ENV:              os.Getenv("ENV"),
	}
}
