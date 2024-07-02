package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-res-api/controller/hotel"
	"user-res-api/dto"
	"user-res-api/service"
	"user-res-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHotelService struct {
	mock.Mock
}

// CheckHotelByIdAmadeus implements service.hotelServiceInterface.
func (m *MockHotelService) CheckHotelByIdAmadeus(id string) (bool, errors.ApiError) {
	panic("unimplemented")
}

// GetHotels implements service.hotelServiceInterface.
func (m *MockHotelService) GetHotels() (dto.HotelsDto, errors.ApiError) {
	args := m.Called()
	hotelsDto := args.Get(0).(dto.HotelsDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Error(1).(errors.ApiError)
	}
	return hotelsDto, apiErr
}

// InsertHotel implements service.hotelServiceInterface.
func (m *MockHotelService) InsertHotel(hotelDto dto.HotelPostDto, idAmadeus string) (dto.HotelDto, errors.ApiError) {
	panic("unimplemented")
}
func (m *MockHotelService) DeleteHotel(idmongo string) errors.ApiError {
	panic("unimplemented")
}

func (m *MockHotelService) GetHotelById(id int) (dto.HotelDto, errors.ApiError) {
	args := m.Called(id)
	hotelDto := args.Get(0).(dto.HotelDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Error(1).(errors.ApiError)
	}
	return hotelDto, apiErr
}

func TestGetHotelById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/hotel/:id", hotel.GetHotelById)

	mockService := new(MockHotelService)
	service.HotelService = mockService

	expectedHotel := dto.HotelDto{
		Id:        1,
		HotelName: "Test Hotel",
	}
	mockService.On("GetHotelById", 1).Return(expectedHotel, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hotel/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualHotel dto.HotelDto
	err := json.Unmarshal(w.Body.Bytes(), &actualHotel)
	assert.NoError(t, err)
	assert.Equal(t, expectedHotel, actualHotel)

	mockService.AssertExpectations(t)
}

func TestGetHotels(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/hotels", hotel.GetHotels)

	mockService := new(MockHotelService)
	service.HotelService = mockService

	expectedHotels := dto.HotelsDto{
		{Id: 1, HotelName: "Hotel 1"},
		{Id: 2, HotelName: "Hotel 2"},
	}
	mockService.On("GetHotels").Return(expectedHotels, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hotels", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualHotels dto.HotelsDto
	err := json.Unmarshal(w.Body.Bytes(), &actualHotels)
	assert.NoError(t, err)
	assert.Equal(t, expectedHotels, actualHotels)

	mockService.AssertExpectations(t)
}
