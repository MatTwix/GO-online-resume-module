package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetGithubUser(username string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка запроса к GitHub: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	var user GithubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	if user.Name == "" {
		user.Name = user.Login
	}

	if user.Email == "" {
		user.Email = "Не указан"
	}

	result := fmt.Sprintf("%s (%s) Email: %s", user.Name, user.Login, user.Email)

	return result, nil
}
