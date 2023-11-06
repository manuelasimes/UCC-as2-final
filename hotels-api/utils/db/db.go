package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database
var client *mongo.Client

func DisconnectDB() {

	client.Disconnect(context.TODO())
}

func InitDB() error {

	clientOpts := options.Client().ApplyURI("mongodb://mongo:27017").
    SetAuth(options.Credential{
        AuthSource:   "admin", // Reemplaza "admin" con el nombre de la base de datos de autenticación que desees utilizar.
        AuthMechanism: "SCRAM-SHA-256", // Reemplaza con el mecanismo de autenticación adecuado si no es el predeterminado.
        Username:     "root", // Reemplaza con tu nombre de usuario.
        Password:     "root", // Reemplaza con tu contraseña.
    })

	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli
	if err != nil {
		return err
	}

	// Autenticación
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return err
    }

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	MongoDb = client.Database("hotels_list")

	//name db

	fmt.Println("Available databases:")
	fmt.Println(dbNames)

	return nil
}
