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
	err = db.Collection("hotels").FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&hotel)
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

func Update(id string, updatedHotel model.Hotel) error {
    db := db.MongoDb
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    // Define las actualizaciones que deseas realizar en el documento
    update := bson.D{
        {Key: "$set", Value: bson.D{
            {Key: "name", Value: updatedHotel.Name},
			{Key: "description", Value: updatedHotel.Description},
			{Key: "city", Value: updatedHotel.City},
			{Key: "country", Value: updatedHotel.Country},
			{Key: "address", Value: updatedHotel.Adress},
			{Key: "images", Value: updatedHotel.Images},
			{Key:"amenities", Value: updatedHotel.Amenities},
            // Puedes agregar más campos aquí según tus necesidades
        }},
    }

    _, err = db.Collection("hotels").UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}, update)
    return err
}


