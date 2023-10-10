package dto

type HotelDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Country string `json:"country"`
	City string `json:"city"`
	Adress string `json:"address"`
	Images []*ImageDto `json:"images"`
	Amenities []*AmenitieDto `json:"amenities"`
}

type HotelsDto []HotelDto
