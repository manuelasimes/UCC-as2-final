package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateHotel(t *testing.T) {
	// Configurar el assert
	assert := assert.New(t)

	// Crear un ObjectID para el hotel
	id := primitive.NewObjectID()

	// Crear valores de prueba para Hotel
	hotel := Hotel{
		Id:          id,
		Name:        "test",
		Description: "test",
		Adress:      "test",
	}

	// Verificar que Id de Hotel
	assert.Equal(id, hotel.Id, "El ID del hotel no coincide")
	// Verificar que nombre de Hotel
	assert.Equal("test", hotel.Name, "El nombre del hotel no coincide")
	assert.Equal("test", hotel.Description, "La descripcion del hotel no coincide")
	assert.Equal("test", hotel.Adress, "El address del hotel no coincide")
}
