package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	Id 			primitive.ObjectID `bson:"_id"`
	Name 		string 			`bson:"name"`
	Description string 			`bson:"description"`
	Country 	string 			`bson:"country"`
	City 		string			`bson:"city"`
	Adress 		string 			`bson:"address"`
	Images 		[]Image 		`bson:"images"`
	Amenities 	[]Amenitie 		`bson:"amenities"`
}

type Hotels []Hotel

//falta poner toda la estructura