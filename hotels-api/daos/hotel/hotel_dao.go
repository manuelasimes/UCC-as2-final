package hotel

import (
	model "hotels-api/models"
	"hotels-api/utils/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgraph-io/ristretto"
)

var cache *ristretto.Cache

func init() {
    // Initialize Ristretto cache
    cache, _ = ristretto.NewCache(&ristretto.Config{
        NumCounters: 1e6,     // Number of counters (1 million)
        MaxCost:     1 << 30, // Maximum cost of cache (1GB)
        BufferItems: 64,      // Number of keys per Get buffer
    })
}

func GetById(id string) model.Hotel {

	fmt.Println("Inside GetById")

	// Check cache first
    if cached, found := cache.Get(id); found {
        if hotel, ok := cached.(model.Hotel); ok {
			fmt.Println("Hotel found in cache")
			fmt.Println(hotel)
            return hotel
        }
    }

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

	fmt.Println("got hotel from mongo")

	// Add to cache
    cache.Set(id, hotel, 1) // Assuming a constant cost of 1 for simplicity
	fmt.Println("Hotel added to cache")

	return hotel

}

func Insert(hotel model.Hotel) model.Hotel {

	fmt.Println("inside insert")

	db := db.MongoDb
	insertHotel := hotel
	insertHotel.Id = primitive.NewObjectID()
	_, err := db.Collection("hotels").InsertOne(context.TODO(), &insertHotel)

	if err != nil {
		fmt.Println(err)
		return hotel
	}
	hotel.Id = insertHotel.Id

	// Convert ObjectID to hex string
	idHexString := hotel.Id.Hex()
	// Update cache
    cache.Set(idHexString, hotel, 1)
	fmt.Println("Hotel added to cache")

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


