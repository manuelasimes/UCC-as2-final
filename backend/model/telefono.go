package model

type Telefono struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Codigo  int `gorm:"not null"`
	Numero  int `gorm:"not null;unique"`

	HotelID int
}

type Telefonos []Telefono