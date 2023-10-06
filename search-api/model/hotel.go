package model

type Hotel struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Country string `json:"country"`
	City string `json:"city"`
	Adress string `json:"address"`
	Images []Image `json:"images"`
	Amenities []Amenitie `json:"amenities"`
}

type Hotels []Hotel