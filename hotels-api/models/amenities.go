package model

type Amenitie struct {
	Description string `bson:"amenitie_description"`
	Image string `bson:"amenitie_image"`
}

type Amenities []Amenitie