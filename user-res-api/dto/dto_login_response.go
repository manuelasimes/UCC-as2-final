package dto

type LoginResponseDto struct {
	UserId   int    `json:"user_id"`
	Token    string `json:"token"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Type     bool   `json:"type"`
}