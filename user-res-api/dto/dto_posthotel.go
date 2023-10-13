package dto

type HotelPostDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	IdMongo          string `json:"id_mongo"`
	// este no me trae el id de amadeus 
}
