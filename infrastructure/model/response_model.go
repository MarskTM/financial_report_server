package model

type LoginResponse struct {
	ID           uint     `json:"id"`
	Role         []string `json:"role"`
	Username     string   `json:"username"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}
