package model

type Cliente struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"`
	Password string `gorm:"type:varchar(150);not null"`
	Email	 string `gorm:"type:varchar(150);not null;unique"`
}

type Clientes []Cliente