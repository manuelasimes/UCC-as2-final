package service_test

import (
	// "net/http"
	"testing"
	"user-res-api/dto"
	"user-res-api/model"

	// "user-res-api/service"
	"user-res-api/utils/errors"
	e "user-res-api/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestHotels struct{}

func (t *TestHotels) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	return dto.HotelDto{}, nil
}

func (t *TestHotels) CheckHotelByIdAmadeus(id string) (bool, e.ApiError) {

	return true, nil

}

func (t *TestHotels) GetHotels() (dto.HotelsDto, e.ApiError) {
	return dto.HotelsDto{}, nil
}

func (m *MockHotelClient) InsertHotel(hotel model.Hotel) model.Hotel {
	args := m.Called(hotel)
	return args.Get(0).(model.Hotel)
}

// Implementación del servicio HotelService con MockHotelClient
type MockHotelService struct {
	HotelClient *MockHotelClient
}

func NewMockHotelService() *MockHotelService {
	mockHotelClient := new(MockHotelClient)
	return &MockHotelService{
		HotelClient: mockHotelClient,
	}
}

func (s *MockHotelService) InsertHotel(hotelDto dto.HotelPostDto, idAmadeus string) (dto.HotelDto, errors.ApiError) {
	var hotel model.Hotel
	hotel.HotelName = hotelDto.HotelName
	hotel.IdAmadeus = idAmadeus
	hotel.IdMongo = hotelDto.IdMongo

	hotel = s.HotelClient.InsertHotel(hotel)

	var response dto.HotelDto
	if hotel.Id == 0 {
		return response, errors.NewBadRequestApiError("Error al insertar hotel")
	}

	response.Id = hotel.Id
	response.HotelName = hotel.HotelName
	response.IdAmadeus = hotel.IdAmadeus
	response.IdMongo = hotel.IdMongo

	return response, nil
}

func TestInsertHotel(t *testing.T) {
	mockService := NewMockHotelService()

	// Configurar el comportamiento esperado en el mock
	expectedHotel := model.Hotel{
		Id:        1,
		HotelName: "Test Hotel",
		IdAmadeus: "test_amadeus_id",
		IdMongo:   "test_mongo_id",
	}

	mockService.HotelClient.On("InsertHotel", mock.AnythingOfType("model.Hotel")).Return(expectedHotel)

	// Crear el DTO de entrada
	hotelDto := dto.HotelPostDto{
		HotelName: "Test Hotel",
		IdMongo:   "test_mongo_id",
	}

	// Llamar a la función que estamos probando
	result, err := mockService.InsertHotel(hotelDto, "test_amadeus_id")

	// Verificar los resultados esperados
	assert.NoError(t, err)
	assert.Equal(t, expectedHotel.Id, result.Id)
	assert.Equal(t, expectedHotel.HotelName, result.HotelName)
	assert.Equal(t, expectedHotel.IdAmadeus, result.IdAmadeus)
	assert.Equal(t, expectedHotel.IdMongo, result.IdMongo)

	mockService.HotelClient.AssertExpectations(t)
}
