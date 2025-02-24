package main

import (
	"fmt"
	"os"

	"github.com/MatTwix/Go-online-resume-module/api"
	"github.com/MatTwix/Go-online-resume-module/config"
)

func main() {
	cfg := config.LoadConfig()

	githubUser, err := api.GetGithubUser(cfg.GithubName)
	if err != nil {
		fmt.Println("Ошибка получения данных GitHub:", err)
		return
	}

	vkUser, err := api.GetVkUser(cfg.VkToken, cfg.VkUserID)
	if err != nil {
		fmt.Println("Ошибка получения данных VK:", err)
		return
	}

	codeforcesUser, err := api.GetCodeforcesUser(cfg.CodeforcesHandle)
	if err != nil {
		fmt.Println("Ошибка получения данных Codeforces:", err)
		return
	}

	data := fmt.Sprintf("Github: %s\nVK: %s\nCodeforces: %s", githubUser, vkUser, codeforcesUser)

	err = os.WriteFile("output/data.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}

	fmt.Println("Данные сохранены в output/data.txt")
}
