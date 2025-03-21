package handlers

import (
	"time"

	"github.com/MatTwix/Go-online-resume-module/api"
	"github.com/MatTwix/Go-online-resume-module/config"
	"github.com/MatTwix/Go-online-resume-module/models"
	"github.com/gofiber/fiber/v3"
)

var lastUpdatedAt string = time.Now().String()

func GetResume(c fiber.Ctx) error {
	cfg := config.LoadConfig()

	githubUser, err := api.GetGithubUser(cfg.GithubName, cfg.GithubToken)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка получения данных GitHub:" + err.Error()})
	}

	vkUser, err := api.GetVkUser(cfg.VkToken, cfg.VkUserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка получения данных VK:" + err.Error()})
	}

	codeforcesUser, err := api.GetCodeforcesUser(cfg.CodeforcesHandle)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка получения данных Codeforces:" + err.Error()})
	}

	var resume models.Resume

	resume.GitHubStats = githubUser
	resume.VkStats = vkUser
	resume.CodeforcesStats = codeforcesUser
	resume.UpdatedAt = lastUpdatedAt

	return c.JSON(resume)
}

func UpdateResume(c fiber.Ctx) error {
	lastUpdatedAt = time.Now().String()

	return c.JSON(fiber.Map{"message": "Resume updated!", "updatedAt": lastUpdatedAt})
}
