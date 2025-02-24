package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VkResponse struct {
	Response []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

func GetVkUser(token string, userID string) (string, error) {
	url := fmt.Sprintf("https://api.vk.com/method/users.get?user_ids=%s&access_token=%s&v=5.131", userID, token)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка запроса к VK: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	var result VkResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	if len(result.Response) > 0 {
		return result.Response[0].FirstName + " " + result.Response[0].LastName, nil
	}

	return "Неизвестный пользователь", nil
}
