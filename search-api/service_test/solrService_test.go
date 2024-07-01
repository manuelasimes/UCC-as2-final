package service_test

import (
	"testing"

	client "search-api/client/solr"
	"search-api/dto"
	"search-api/service"
	"search-api/utils/errors"

	"github.com/stretchr/testify/mock"
)

// MockSolrClient implementa SolrClientInterface para pruebas
type MockSolrClient struct {
	mock.Mock
}

func (m *MockSolrClient) GetQuery(query string) (dto.HotelsDto, errors.ApiError) {
	args := m.Called(query)
	return args.Get(0).(dto.HotelsDto), args.Get(1).(errors.ApiError)
}

func (m *MockSolrClient) GetQueryAllFields(query string) (dto.HotelsDto, errors.ApiError) {
	args := m.Called(query)
	return args.Get(0).(dto.HotelsDto), args.Get(1).(errors.ApiError)
}

func (m *MockSolrClient) Add(hotel dto.HotelDto) errors.ApiError {
	args := m.Called(hotel)
	return args.Error(0).(errors.ApiError)
}

func (m *MockSolrClient) Delete(id string) errors.ApiError {
	args := m.Called(id)
	return args.Error(0).(errors.ApiError)
}

// SolrClientWrapper es un wrapper para adaptar MockSolrClient a client.SolrClient
type SolrClientWrapper struct {
	*client.SolrClient
	mockClient *MockSolrClient
}

func (w *SolrClientWrapper) GetQuery(query string) (dto.HotelsDto, errors.ApiError) {
	return w.mockClient.GetQuery(query)
}

func (w *SolrClientWrapper) GetQueryAllFields(query string) (dto.HotelsDto, errors.ApiError) {
	return w.mockClient.GetQueryAllFields(query)
}

func (w *SolrClientWrapper) Add(hotel dto.HotelDto) errors.ApiError {
	return w.mockClient.Add(hotel)
}

func (w *SolrClientWrapper) Delete(id string) errors.ApiError {
	return w.mockClient.Delete(id)
}

func TestSolrService_GetQuery(t *testing.T) {
	// Configurar el mock para devolver un error específico
	mockClient := new(MockSolrClient)
	mockClient.On("GetQuery", "query_query").Return(dto.HotelsDto{}, errors.NewBadRequestApiError("Solr failed;Error Code: bad_request"))

	// Crear el wrapper
	solrClientWrapper := &SolrClientWrapper{
		SolrClient: &client.SolrClient{},
		mockClient: mockClient,
	}

	// Crear instancia de SolrService con el wrapper
	solrService := service.NewSolrServiceImpl(solrClientWrapper.SolrClient)

	// Ejecutar el método bajo prueba
	result, err := solrService.GetQuery("query_query")

	// Validar que result está vacío
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	// Validar que se haya producido un error
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Validar el tipo de error
	apiErr, ok := err.(errors.ApiError)
	if !ok {
		t.Errorf("Expected ApiError, got %T", err)
	}

	// Validar el tipo específico de error
	if apiErr.Status() != 400 {
		t.Errorf("Expected status code 400, got %d", apiErr.Status())
	}

	// Validar el mensaje de error
	expectedMessage := "Solr failed"
	if apiErr.Message() != expectedMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedMessage, apiErr.Message())
	}

	// Validar que no se llamó al método del mock con los parámetros incorrectos
	mockClient.AssertNotCalled(t, "GetQuery", "query_query")
}
