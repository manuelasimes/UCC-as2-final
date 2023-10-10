package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Phone 	 int    `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Type     bool   `json:"type"`
}

type UsersDto []UserDto