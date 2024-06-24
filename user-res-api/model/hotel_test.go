package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHotel(t *testing.T) {
	// Configurar el assert
	assert := assert.New(t)

	// Crear valores de prueba para Hotel
	hotel := Hotel{
		Id:        1,
		HotelName: "Mandarin",
		IdMongo:   "hola",
		IdAmadeus: "chau",
	}

	// Verificar que Id de Hotel
	assert.Equal(1, hotel.Id, "El ID del hotel no coincide")
	// Verificar que nombre de Hotel
	assert.Equal("Mandarin", hotel.HotelName, "El nombre del hotel no coincide")
	// Verificar Id Mongo
	assert.Equal("hola", hotel.IdMongo, "El idmongo del hotel no coincide")
	// Verificar que idamadeus de Hotel
	assert.Equal("chau", hotel.IdAmadeus, "El idamadeus del hotel no coincide")

}
