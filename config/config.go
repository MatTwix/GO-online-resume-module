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
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	return Config{
		GithubName:       os.Getenv("GITHUB_NAME"),
		VkToken:          os.Getenv("VK_KEY"),
		VkUserID:         os.Getenv("VK_USER_ID"),
		CodeforcesHandle: os.Getenv("CODEFORCES_HANDLE"),
	}
}
