package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CodeforcesResponse struct {
	Status string `json:"status"`
	Result []struct {
		Handle    string `json:"handle"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Rating    int    `json:"rating"`
		MaxRating int    `json:"maxRating"`
		Rank      string `json:"rank"`
		MaxRank   string `json:"maxRank"`
	} `json:"result"`
}

func GetCodeforcesUser(handle string) (string, error) {
	url := "https://codeforces.com/api/user.info?handles=" + handle
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка запроса к Codeforces: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	var result CodeforcesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	if result.Status != "OK" || len(result.Result) == 0 {
		return "Пользователь не найден", nil
	}

	user := result.Result[0]
	return fmt.Sprintf("%s %s (@%s), рейтинг: %d (макс: %d), звание: %s (макс: %s)",
		user.FirstName, user.LastName, user.Handle, user.Rating, user.MaxRating, user.Rank, user.MaxRank), nil
}
