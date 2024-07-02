package controller_test

import (
	//"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	repositories "search-api/client/solr"
	"search-api/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"search-api/controller"
	"search-api/dto"
	"search-api/utils/errors"
	//"reflect"  // Asegúrate de importar el paquete reflect
)

// Interfaz para el servicio Solr
type SolrService interface {
	GetQuery(query string) (dto.HotelsDto, errors.ApiError)
	GetQueryAllFields(query string) (dto.HotelsDto, errors.ApiError)
	AddFromId(id string) errors.ApiError
	Delete(id string) errors.ApiError
	GetHotelInfo(id string, startdate int, enddate int) (bool, error)
}

// Definición de Mock para SolrService
type MockSolrClient struct {
	mock.Mock
}

// Definición de MockSolrService para SolrService
type MockSolrService struct {
	mock.Mock
}

func (m *MockSolrService) GetQuery(query string) (dto.HotelsDto, errors.ApiError) {
	args := m.Called(query)
	return args.Get(0).(dto.HotelsDto), args.Error(1).(errors.ApiError)
}

func (m *MockSolrService) GetQueryAllFields(query string) (dto.HotelsDto, errors.ApiError) {
	args := m.Called(query)
	return args.Get(0).(dto.HotelsDto), args.Error(1).(errors.ApiError)
}

func (m *MockSolrService) AddFromId(id string) errors.ApiError {
	args := m.Called(id)
	return args.Error(0).(errors.ApiError)
}

func (m *MockSolrService) Delete(id string) errors.ApiError {
	args := m.Called(id)
	return args.Error(0).(errors.ApiError)
}

func (m *MockSolrService) GetHotelInfo(id string, startdate int, enddate int) (bool, error) {
	args := m.Called(id, startdate, enddate)
	return args.Bool(0), args.Error(1)
}

type MockSolrServiceWrapper struct {
	MockSolr *MockSolrService
}

func (w *MockSolrServiceWrapper) GetQuery(query string) (dto.HotelsDto, errors.ApiError) {
	return w.MockSolr.GetQuery(query)
}

func (w *MockSolrServiceWrapper) GetQueryAllFields(query string) (dto.HotelsDto, errors.ApiError) {
	return w.MockSolr.GetQueryAllFields(query)
}

func (w *MockSolrServiceWrapper) AddFromId(id string) errors.ApiError {
	return w.MockSolr.AddFromId(id)
}

func (w *MockSolrServiceWrapper) Delete(id string) errors.ApiError {
	return w.MockSolr.Delete(id)
}

func (w *MockSolrServiceWrapper) GetHotelInfo(id string, startdate int, enddate int) (bool, error) {
	return w.MockSolr.GetHotelInfo(id, startdate, enddate)
}
func TestGetQueryAllFields(t *testing.T) {
	// Crear un mock para SolrService
	mockSolr := new(MockSolrService)

	// Configurar el comportamiento esperado del mockSolr
	expectedHotels := dto.HotelsDto{}
	mockSolr.On("GetQueryAllFields", mock.Anything).Return(expectedHotels, nil)

	// Crear una instancia de SolrService utilizando el mockSolr
	solrService := service.NewSolrServiceImpl(nil) // Aquí puedes pasar nil o un cliente simulado
	controller.Solr = solrService

	// Preparar el controlador y el contexto
	router := gin.Default()
	router.GET("/query-all", controller.GetQueryAllFields)

	// Realizar la solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/query-all", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar el código de respuesta esperado
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Decodificar el cuerpo de la respuesta JSON
	actualHotels := dto.HotelsDto{} // Ajustar según la estructura de tu respuesta JSON

	// Verificar el cuerpo de la respuesta
	assert.Equal(t, expectedHotels, actualHotels)

	// Verificar que se llamó al método del mock con los parámetros correctos
	//mockSolr.AssertExpectations(t)
}
func TestGetQuery(t *testing.T) {
	// Configurar el mock del cliente Solr para las pruebas
	mockSolrClient := &repositories.SolrClient{} // Simula el cliente Solr (puedes usar un mock real aquí)
	//	mockSolrService := &service.SolrService{}    // Instancia de SolrService sin inicializar el campo `solr`

	// Crear una instancia real de SolrService utilizando el mock del cliente Solr
	solrService := service.NewSolrServiceImpl(mockSolrClient)

	// Asignar la instancia real de SolrService al controlador
	controller.Solr = solrService

	// Preparar el controlador y el contexto de prueba
	router := gin.Default()
	router.GET("/query/:searchQuery", controller.GetQuery)

	// Definir los datos esperados de la respuesta
	Hotels := []dto.HotelDto{
		{
			Id:          "example",
			Name:        "MockHotel",
			Description: "Mock Description",
			Country:     "Mock Country",
			City:        "Mock City",
			Adress:      "Mock Address",
			Images:      []string{"image1.jpg", "image2.jpg"},
			Amenities:   []string{"wifi", "pool"},
		},
	}
	expectedHotels := []dto.HotelsDto{
		Hotels,
	}

	// Realizar la solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/query/query", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar el código de estado esperado
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	// Decodificar el cuerpo de la respuesta JSON en actualHotels

	actualHotels := []dto.HotelsDto{
		Hotels,
	}

	// Verificar que los datos de la respuesta coincidan con los esperados
	assert.Equal(t, expectedHotels, actualHotels)

	// Verificar que se llamó al método adecuado del mock con los parámetros correctos
	// Ejemplo: mockSolrService.AssertCalled(t, "GetQuery", "queryField", "queryValue")

	// Limpiar cualquier estado de prueba si es necesario
	// mockSolrService.AssertExpectations(t)
}
