package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(300);not null"`
	LastName string `gorm:"type:varchar(300);not null"`
	UserName string `gorm:"type:varchar(200);not null;unique"`
	Phone 	 int    `gorm:"not null"`
	Address  string `gorm:"type:varchar(200);not null"`
	Password string `gorm:"type:varchar(500);not null"`
	Email    string `gorm:"type:varchar(320);not null;unique"`
	Type     bool   `gorm:"not null;default:false"`
}

type Users []User
