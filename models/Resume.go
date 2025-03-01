package models

type Resume struct {
	VkStats         string `json:"vk_stats"`
	GitHubStats     string `json:"github_stats"`
	CodeforcesStats string `json:"codeforces_stats"`
	UpdatedAt       string `json:"updated_at"`
}
