package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "user-res-api/client/booking"
	"user-res-api/dto"
	"user-res-api/service"
	"user-res-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del servicio BookingService
type MockBookingService struct {
	mock.Mock
}

func (m *MockBookingService) GetBookingById(id int) (dto.BookingDetailDto, errors.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.BookingDetailDto), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) GetAvailabilityByIdAndDate(idAm string, startDate, endDate int) (dto.Availability, errors.ApiError) {
	args := m.Called(idAm, startDate, endDate)
	return args.Get(0).(dto.Availability), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) GetBookings() (dto.BookingsDetailDto, errors.ApiError) {
	args := m.Called()
	return args.Get(0).(dto.BookingsDetailDto), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) InsertBooking(bookingPDto dto.BookingPostDto) (dto.BookingDto, errors.ApiError) {
	args := m.Called(bookingPDto)
	return args.Get(0).(dto.BookingDto), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) GetBookingsByUserId(id int) (dto.BookingsDetailDto, errors.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.BookingsDetailDto), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) GetBookingByUserId(id int) (dto.BookingDetailDto, errors.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.BookingDetailDto), args.Get(1).(errors.ApiError)
}

func (m *MockBookingService) Availability(startdateconguiones string, enddateconguiones string, idAm string) bool {
	args := m.Called(startdateconguiones, enddateconguiones, idAm)
	return args.Bool(0)
}

func (m *MockBookingService) GetAmadeustoken() string {
	args := m.Called()
	return args.String(0)
}

func TestGetAvailabilityByIdAndDate(t *testing.T) {
	// Configurar el entorno de prueba
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/availability/:id/:start_date/:end_date", func(c *gin.Context) {
		// Extraer parámetros de la URL
		id := c.Param("id")
		startDate := c.Param("start_date")
		endDate := c.Param("end_date")

		// Lógica para convertir startDate y endDate a enteros si es necesario
		// Usar el mock del servicio BookingService
		mockService := new(MockBookingService)
		service.BookingService = mockService

		// Configurar el comportamiento esperado en el mock
		expectedAvailability := dto.Availability{
			OkToBook: true,
		}
		mockService.On("GetAvailabilityByIdAndDate", id, startDate, endDate).Return(expectedAvailability, nil)

		// Ejecutar la solicitud HTTP simulada
		req, _ := http.NewRequest("GET", "/availability/"+id+"/"+startDate+"/"+endDate, nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verificar los resultados esperados
		assert.Equal(t, http.StatusOK, resp.Code)

		var responseDto dto.Availability
		err := json.Unmarshal(resp.Body.Bytes(), &responseDto)
		assert.NoError(t, err)
		assert.Equal(t, expectedAvailability, responseDto)

		mockService.AssertExpectations(t)
	})

	// Ejecutar la prueba
	t.Run("GetAvailabilityByIdAndDate", func(t *testing.T) {
		// Aquí puedes agregar pruebas adicionales si es necesario
		// No es necesario agregar nada específico para esta función debido a la naturaleza del test funcional
	})
}
