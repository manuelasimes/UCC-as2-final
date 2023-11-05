package dto

type BookingPostDto struct {
	Id int `json:"booking_id"`

	UserId int `json:"user_booked_id"`

	HotelId string `json:"booked_hotel_id"`

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`

}

type BookingsPostDto []BookingDto