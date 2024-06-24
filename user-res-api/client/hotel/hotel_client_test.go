package hotel_test

import (
	// "fmt"
	"testing"
	"user-res-api/model"

	"github.com/stretchr/testify/assert"
)

// Mock para simular el comportamiento de la base de datos
type MockHotelDAO struct{}

func (m *MockHotelDAO) InsertHotel(hotel model.Hotel) model.Hotel {
	// Simular la lógica de inserción en la base de datos
	// Se establece un ID para el hotel
	hotel.Id = 1 // Si modifico a cero, genera la alerta
	return hotel
}

func TestInsertHotel(t *testing.T) {
	// Crear una instancia del mock del DAO de HOtel
	mockDAO := &MockHotelDAO{}

	// Crear una nueva reserva
	newHotel := model.Hotel{

		Id:        1,
		HotelName: "Mandarin",
		IdMongo:   "hola",
		IdAmadeus: "chau",
	}

	// Insertar la reserva utilizando el mock del DAO
	inserted := mockDAO.InsertHotel(newHotel)

	// Verificar que la reserva tenga un ID asignado
	assert.NotZero(t, inserted.Id, "La insercion dle hotel no se pudo realizar")

	// Verificar otros atributos de la reserva
	assert.Equal(t, newHotel.Id, inserted.Id)
	assert.Equal(t, newHotel.HotelName, inserted.HotelName)
	assert.Equal(t, newHotel.IdMongo, inserted.IdMongo)
	assert.Equal(t, newHotel.IdAmadeus, inserted.IdAmadeus)
}

// TEST PARA LA FUNCION GETBOOKINGBYID

func (m *MockHotelDAO) GetHotelById(id int) model.Hotel {
	// Simular la búsqueda en la base de datos
	hotel := model.Hotel{
		Id:        1,
		HotelName: "Mandarin",
		IdMongo:   "hola",
		IdAmadeus: "chau",
	}

	return hotel
}

func TestGetHotelById(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockDAO := &MockHotelDAO{}

	// ID de reserva a buscar - Si la cambio deja de funcionar
	hotelId := 1

	// Obtener la reserva utilizando el mock del DAO
	hotel := mockDAO.GetHotelById(hotelId)

	// Verificar que la reserva obtenida tenga el ID correcto
	assert.Equal(t, hotelId, hotel.Id, "El ID del hotel no existe")

	// Verificar otros atributos de la reserva
	assert.Equal(t, 1, hotel.Id)
	assert.Equal(t, "Mandarin", hotel.HotelName)
	assert.Equal(t, "hola", hotel.IdMongo)
	assert.Equal(t, "chau", hotel.IdAmadeus)
}
