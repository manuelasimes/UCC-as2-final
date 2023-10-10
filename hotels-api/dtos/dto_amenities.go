package dto

type AmenitieDto struct {
	Description string `json:"amenitie_description"`
	Image string `string:"amenitie_image"`
}

type Amenities []AmenitieDto