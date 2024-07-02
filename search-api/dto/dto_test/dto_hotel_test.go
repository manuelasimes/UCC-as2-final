package dto_test

import (
	dto "search-api/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelDto(t *testing.T) {
	hotelDto := dto.HotelDto{
		Id:          "test",
		Name:        "test",
		Description: "test",
		Country:     "test",
		City:        "test",
		Adress:      "test",
	}
	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, "test", hotelDto.Id, "El ID de la reserva no coincide")
	assert.Equal(t, "test", hotelDto.Name, "EL nombre del hotel no coincide")
	assert.Equal(t, "test", hotelDto.Description, "La descripcion no coincide")
	assert.Equal(t, "test", hotelDto.Country, "El pais no coincide")
	assert.Equal(t, "test", hotelDto.City, "La ciudad no coincide")
	assert.Equal(t, "test", hotelDto.Adress, "El pais no coincide")
}
