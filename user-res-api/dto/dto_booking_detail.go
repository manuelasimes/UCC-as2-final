package dto

type BookingDetailDto struct {

	Id int `json:"booking_id"`

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`

	UserId   int    `json:"user_booked_id"`
	Username string `json:"user_name"`

	HotelId   int    `json:"booked_hotel_id"`
	HotelName string `json:"hotel_name"`
	Address   string `json:"hotel_address"`
}

type BookingsDetailDto []BookingDetailDto