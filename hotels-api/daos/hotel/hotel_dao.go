package hotel

import (
	model "hotels-api/models"
	"hotels-api/utils/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetById(id string) model.Hotel {
	var hotel model.Hotel
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return hotel
	}
	err = db.Collection("hotels").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		fmt.Println(err)
		return hotel
	}
	return hotel

}

func Insert(hotel model.Hotel) model.Hotel {
	db := db.MongoDb
	insertHotel := hotel
	insertHotel.Id = primitive.NewObjectID()
	_, err := db.Collection("hotels").InsertOne(context.TODO(), &insertHotel)

	if err != nil {
		fmt.Println(err)
		return hotel
	}
	hotel.Id = insertHotel.Id
	return hotel
}
