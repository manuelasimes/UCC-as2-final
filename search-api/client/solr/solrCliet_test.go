//

package repositories

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"search-api/config"
// 	"search-api/dto"
// 	"testing"

// 	"github.com/stevenferrer/solr-go"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // SolrClientInterface define los métodos necesarios para interactuar con Solr.
// type SolrClientInterface interface {
// 	Update(ctx context.Context, collection string, contentType string, body io.Reader) (solr.UpdateResponse, error)
// 	Commit(ctx context.Context, collection string) error
// 	Add(hotel dto.HotelDto) error

// 	// Otros métodos necesarios para tu aplicación
// }

// // SolrClient es una implementación concreta de SolrClientInterface.
// type SolrClient1 struct {
// 	Client     SolrClientInterface
// 	Collection string
// }

// // MockSolrClient es un mock para SolrClientInterface.
// type MockSolrClient struct {
// 	mock.Mock
// }

// func (m *MockSolrClient) Update(ctx context.Context, collection string, contentType string, body io.Reader) (solr.UpdateResponse, error) {
// 	args := m.Called(ctx, collection, contentType, body)
// 	return args.Get(0).(solr.UpdateResponse), args.Error(1)
// }

// func (m *MockSolrClient) Commit(ctx context.Context, collection string) error {
// 	args := m.Called(ctx, collection)
// 	return args.Error(0)
// }
// func (m *MockSolrClient) Add(hotel dto.HotelDto) error {
// 	args := m.Called(hotel)
// 	return args.Error(0)
// }

// // setupSolrClientTest inicializa un cliente Solr para pruebas utilizando el mock.
// func setupSolrClientTest() *SolrClient1 {
// 	config.SOLRHOST = "localhost"
// 	config.SOLRPORT = 8983

// 	mockSolrClient := new(MockSolrClient)
// 	return &SolrClient1{
// 		Client:     mockSolrClient,
// 		Collection: "hotelSearch",
// 	}
// }

// func TestGetQuery(t *testing.T) {
// 	// Configurar el cliente Solr simulado para las pruebas
// 	client := setupSolrClientTest()

// 	// Configurar el mock de SolrClient
// 	mockSolrClient := client.Client.(*MockSolrClient)

// 	// Simular la llamada a Commit (ejemplo)
// 	mockSolrClient.On("Commit", mock.Anything, mock.AnythingOfType("string")).Return(nil)

// 	// Definir los datos esperados de la respuesta
// 	expectedHotels := dto.HotelsDto{
// 		{
// 			Id:          "example",
// 			Name:        "MockHotel",
// 			Description: "Mock Description",
// 			Country:     "Mock Country",
// 			City:        "Mock City",
// 			Adress:      "Mock Address",
// 			Images:      []string{"image1.jpg", "image2.jpg"},
// 			Amenities:   []string{"wifi", "pool"},
// 		},
// 	}

// 	responseDto := dto.SolrResponseDto{
// 		Response: dto.ResponseDto{
// 			Docs: expectedHotels,
// 		},
// 	}

// 	query := "test-query"
// 	field := "name"

// 	// Crear un servidor de prueba
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, fmt.Sprintf("/solr/hotelSearch/select?q=%s%s%s", field, "%3A", query), r.URL.String())
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseDto)
// 	}))
// 	defer ts.Close()

// 	// Sobrescribir el host y puerto de Solr para que apunte al servidor de prueba
// 	config.SOLRHOST = ts.URL[7:] // Strip "http://"

// 	// Realizar la llamada al método que quieres probar (aquí se debería llamar a un método que haga la consulta, no Update)
// 	hotels, err := client.GetQuery(context.Background(), query, field)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedHotels, hotels)

// 	// Verificar que se haya llamado a Commit después de la operación (ejemplo)
// 	mockSolrClient.AssertCalled(t, "Commit", mock.Anything, mock.AnythingOfType("string"))

// 	// Asegurarse de que todas las expectativas del mock se cumplan
// 	mockSolrClient.AssertExpectations(t)
// }

// func TestGetQueryAllFields(t *testing.T) {
// 	client := setupSolrClientTest()

// 	// Definir los datos esperados de la respuesta
// 	expectedHotels := dto.HotelsDto{
// 		{
// 			Id:          "example",
// 			Name:        "MockHotel",
// 			Description: "Mock Description",
// 			Country:     "Mock Country",
// 			City:        "Mock City",
// 			Adress:      "Mock Address",
// 			Images:      []string{"image1.jpg", "image2.jpg"},
// 			Amenities:   []string{"wifi", "pool"},
// 		},
// 	}

// 	responseDto := dto.SolrResponseDto{
// 		Response: dto.ResponseDto{
// 			Docs: expectedHotels,
// 		},
// 	}

// 	// Create a test server
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "/solr/hotelSearch/select?q=*:*", r.URL.String())
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseDto)
// 	}))
// 	defer ts.Close()

// 	// Override the Solr host and port to point to the test server
// 	config.SOLRHOST = ts.URL[7:] // Strip "http://"

// 	// Realizar la llamada al método GetQueryAllFields del cliente Solr
// 	hotels, err := client.Client.Update(context.Background(), client.Collection, "application/json", nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedHotels, hotels)
// }

// func TestAdd(t *testing.T) {
// 	// Setup
// 	client := setupSolrClientTest()

// 	// Mock Solr client
// 	mockSolrClient := client.Client.(*MockSolrClient)

// 	// Configura el comportamiento esperado del mock para Add
// 	mockSolrClient.On("Add", mock.Anything).Return(nil) // Aquí puedes ajustar el retorno según lo necesites

// 	// Datos de ejemplo para la prueba
// 	hotel := dto.HotelDto{
// 		Id:          "1",
// 		Name:        "Test Hotel",
// 		Description: "A test hotel",
// 		Country:     "Test Country",
// 		City:        "Test City",
// 		Adress:      "Test Address",
// 		Images:      []string{"image1.jpg"},
// 		Amenities:   []string{"wifi"},
// 	}

// 	// Ejecuta la función que llama a Add en el cliente
// 	err := client.Client.Add(hotel)
// 	assert.NoError(t, err) // Verifica que no haya error al añadir el hotel

// 	// Verifica que se haya llamado a Add con los datos correctos
// 	mockSolrClient.AssertCalled(t, "Add", hotel)

// 	// Asegúrate de que todas las expectativas del mock se cumplan
// 	mockSolrClient.AssertExpectations(t)
// }

// // func TestDelete(t *testing.T) {
// // 	// Setup
// // 	client := setupSolrClientTest()

// // 	// Mock Solr client
// // 	mockSolrClient := client.Client.(*MockSolrClient)

// // 	// Configura el comportamiento esperado del mock para Update
// // 	mockSolrClient.On("Update", mock.Anything, client.Collection, "application/json", nil).
// // 		Return(solr.UpdateResponse{}, nil)

// // 	// Configura el comportamiento esperado del mock para Commit
// // 	mockSolrClient.On("Commit", mock.Anything, mock.AnythingOfType("string")).
// // 		Return(nil)

// // 	// Ejecuta la función que llama a Update en el cliente
// // 	_, err := client.Client.Update(context.Background(), client.Collection, "application/json", nil)
// // 	if err != nil {
// // 		t.Errorf("Error updating Solr: %v", err)
// // 	}

// // 	// Verifica que no haya error al actualizar
// // 	assert.NoError(t, err)

// // 	// Verifica que se haya llamado a Update con los argumentos correctos
// // 	mockSolrClient.AssertCalled(t, "Update", mock.Anything, client.Collection, "application/json", nil)

// // 	// Verifica que se haya llamado a Commit con los argumentos correctos
// // 	mockSolrClient.AssertCalled(t, "Commit", mock.Anything, client.Collection)

// // 	// Asegúrate de que todas las expectativas del mock se cumplan
// // 	mockSolrClient.AssertExpectations(t)
// // }
