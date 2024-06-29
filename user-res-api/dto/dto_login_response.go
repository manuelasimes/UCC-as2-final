package dto

type LoginResponseDto struct {
	UserId       int    `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Type         bool   `json:"type"`
}
