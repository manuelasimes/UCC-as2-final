package model

type Hotel struct {
	Id        int `gorm:"primaryKey"`   
	HotelName      string `gorm:"type:varchar(50);not null"` 
	IdMongo   string `gorm:"type:varchar(250);not null;unique"` 
	IdAmadeus string `gorm:"type:varchar(250);not null;unique"` 
}


type Hotels []Hotel