package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	Id primitive.ObjectID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Country string `json:"country"`
	City string `json:"city"`
	Adress string `json:"address"`
	Images []Image `json:"images"`
	Amenities []Amenitie `json:"amenities"`
}

type Hotels []Hotel

//falta poner toda la estructura