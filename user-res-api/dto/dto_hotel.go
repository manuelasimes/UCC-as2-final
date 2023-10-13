package dto

type HotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	IdMongo          string `json:"id_mongo"`
	IdAmadeus        string `json:"id_amadeus"`
}

type HotelsDto []HotelDto