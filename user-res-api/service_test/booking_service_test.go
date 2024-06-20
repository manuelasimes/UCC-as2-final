package service_test

import (
	"net/http"
	"testing"
	"user-res-api/dto"
	"user-res-api/model"

	// "user-res-api/service"
	e "user-res-api/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestBookings struct{}

func (t *TestBookings) InsertBooking(bookingPDto dto.BookingPostDto) (dto.BookingDto, e.ApiError) {
	if bookingPDto.UserId == 0 {
		return dto.BookingDto{}, e.NewApiError("Error al insertar la reserva", "booking_insert_error", http.StatusInternalServerError, nil)
	}
	return dto.BookingDto{}, nil
}

func (t *TestBookings) GetBookingById(id int) (dto.BookingDetailDto, e.ApiError) {
	return dto.BookingDetailDto{}, nil
}

func (t *TestBookings) GetBookings() (dto.BookingsDetailDto, e.ApiError) {
	return dto.BookingsDetailDto{}, nil
}

func (t *TestBookings) GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError) {
	return dto.BookingDetailDto{}, nil
}

func (t *TestBookings) Availability(startdateconguiones string, enddateconguiones string, idAm string) bool {
	return true
	//chequear debe estar mal
}

func (t *TestBookings) GetAmadeustoken() string {
	token := " hola "
	return token
}

func (t *TestBookings) GetAvailabilityByIdAndDate(idAm string, startDate int, endDate int) (dto.Availability, e.ApiError) {
	return dto.Availability{}, nil
}

func (t *TestBookings) GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError) {
	return dto.BookingsDetailDto{}, nil
}

type MockHotelClient struct {
	mock.Mock
}

func (m *MockHotelClient) GetHotelByIdMongo(hotelId string) model.Hotel {
	args := m.Called(hotelId)
	return args.Get(0).(model.Hotel)
}

type MockUserClient struct {
	mock.Mock
}

func (m *MockUserClient) CheckUserById(userId int) bool {
	args := m.Called(userId)
	return args.Bool(0)
}

type MockBookingService struct {
	mock.Mock
}

func (m *MockBookingService) InsertBooking(bookingPDto dto.BookingPostDto) (dto.BookingDto, e.ApiError) {
	args := m.Called(bookingPDto)
	return args.Get(0).(dto.BookingDto), args.Get(1).(e.ApiError)
}

func (m *MockBookingService) GetAvailabilityByIdAndDate(idAm string, startDate, endDate int) (dto.Availability, e.ApiError) {
	args := m.Called(idAm, startDate, endDate)
	return args.Get(0).(dto.Availability), args.Get(1).(e.ApiError)
}

func TestInsertBooking(t *testing.T) {
	// Crear instancia de MockBookingService
	bookingService := new(MockBookingService)

	// Configurar comportamiento esperado en MockBookingService
	bookingPDto := dto.BookingPostDto{
		Id:        1,
		UserId:    1,
		HotelId:   "hola",
		StartDate: 22042025,
		EndDate:   28042025,
	}

	bookingDto := dto.BookingDto{
		Id:        1,
		UserId:    1,
		HotelId:   1,
		StartDate: 22042025,
		EndDate:   28042025,
	}

	apiError := e.NewBadRequestApiError("Error")

	bookingService.On("InsertBooking", bookingPDto).Return(bookingDto, apiError)

	// Ejecutar la función que se está probando
	resultBookingDto, resultApiError := bookingService.InsertBooking(bookingPDto)

	// Verificar los resultados esperados
	assert.Equal(t, bookingDto, resultBookingDto)
	assert.Equal(t, apiError, resultApiError)

	// Verificar que las expectativas se cumplieron
	bookingService.AssertExpectations(t)
}
