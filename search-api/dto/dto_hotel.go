package dto

type HotelDto struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Country string `json:"country"`
	City string `json:"city"`
	Adress string `json:"address"`
	Images []string `json:"images"`
	Amenities []string `json:"amenities"`
}

type HotelsDto []HotelDto