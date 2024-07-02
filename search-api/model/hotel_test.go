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
		Id:          "test",
		Name:        "test",
		Description: "test",
		Adress:      "test",
	}

	// Verificar que Id de Hotel
	assert.Equal("test", hotel.Id, "El ID del hotel no coincide")
	// Verificar que nombre de Hotel
	assert.Equal("test", hotel.Name, "El nombre del hotel no coincide")
	assert.Equal("test", hotel.Description, "La descripcion del hotel no coincide")
	assert.Equal("test", hotel.Adress, "El address del hotel no coincide")
}
