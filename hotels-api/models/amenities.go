package model

type Amenitie struct {
	Description string `json:"amenitie_description"`
	Image string `string:"amenitie_image"`
}

type Amenities []Amenitie