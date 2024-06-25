package service_test

import (
	"testing"

	dto "hotels-api/dtos"
	e "hotels-api/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define la interfaz del servicio del hotel
type hotelServiceInterface interface {
	GetHotel(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

// Implementación del mock para el servicio del hotel
type MockHotelService struct {
	mock.Mock
}

func (m *MockHotelService) GetHotel(id string) (dto.HotelDto, e.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.HotelDto), convertError(args.Get(1))
}

func (m *MockHotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	args := m.Called(hotelDto)
	return args.Get(0).(dto.HotelDto), convertError(args.Get(1))
}

func (m *MockHotelService) UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	args := m.Called(id, updatedHotelDto)
	return args.Get(0).(dto.HotelDto), convertError(args.Get(1))
}

// Función de ayuda para convertir errores
func convertError(err interface{}) e.ApiError {
	if err == nil {
		return nil
	}
	return err.(e.ApiError)
}

func TestGetHotel(t *testing.T) {
	// Configurar el mock del servicio del hotel
	mockService := new(MockHotelService)
	expectedHotelDto := dto.HotelDto{
		Id:          "existing_id",
		Name:        "Hotel Name",
		Description: "Hotel Description",
		Country:     "Country",
		City:        "City",
		Adress:      "Address",
		Images:      []string{"image1.jpg", "image2.jpg"},
		Amenities:   []string{"Wifi", "Pool"},
	}
	mockService.On("GetHotel", "existing_id").Return(expectedHotelDto, nil)

	// Crear una instancia del servicio HotelServiceInterface y asignar el mock del servicio
	var hotelService hotelServiceInterface = mockService

	// Llamar al método bajo prueba
	hotelDto, err := hotelService.GetHotel("existing_id")

	// Verificar los resultados
	assert.NoError(t, err)
	assert.Equal(t, expectedHotelDto.Id, hotelDto.Id)
	assert.Equal(t, expectedHotelDto.Name, hotelDto.Name)
	assert.Equal(t, expectedHotelDto.Description, hotelDto.Description)
	assert.Equal(t, expectedHotelDto.Country, hotelDto.Country)
	assert.Equal(t, expectedHotelDto.City, hotelDto.City)
	assert.Equal(t, expectedHotelDto.Adress, hotelDto.Adress)
	assert.ElementsMatch(t, expectedHotelDto.Images, hotelDto.Images)
	assert.ElementsMatch(t, expectedHotelDto.Amenities, hotelDto.Amenities)

	// Verificar llamada al método del servicio
	mockService.AssertExpectations(t)
}
