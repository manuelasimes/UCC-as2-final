package model

type Booking struct {
	Id int `gorm:"primaryKey"`

	StartDate int `gorm:"not null"`
	EndDate   int `gorm:"not null"`
	Rooms     int `gorm:"not null"`
	
	User   User `gorm:"foreignkey:UserId"`
	UserId int

	// Hotel   Hotel `gorm:"foreignkey:HotelId"`
	// HotelId int
}

type Bookings []Booking