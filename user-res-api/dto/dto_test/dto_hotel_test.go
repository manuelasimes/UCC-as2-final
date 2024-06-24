package dto_test

import (
	"testing"
	"user-res-api/dto"

	"github.com/stretchr/testify/assert"
)

func TestHotelDto(t *testing.T) {
	// Crear una instancia del DTO de Hotel, si modifico alguna y deja de ser igual, da la alerta
	hotelDto := dto.HotelDto{
		Id:        1,
		HotelName: "Mandarin",
		IdMongo:   "hola",
		IdAmadeus: "chau",
	}

	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, 1, hotelDto.Id, "El ID de la reserva no coincide")
	assert.Equal(t, "Mandarin", hotelDto.HotelName, "EL nombre del hotel no coincide")
	assert.Equal(t, "hola", hotelDto.IdMongo, "El id de mongo no coincide")
	assert.Equal(t, "chau", hotelDto.IdAmadeus, "El ID de Amadeus no coincide")
}
