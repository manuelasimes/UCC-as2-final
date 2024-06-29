package hotel_test

import (
	"context"
	model "hotels-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"hotels-api/daos/hotel"
)

// MockHotelDAO es un mock del DAO de hoteles para pruebas
type MockHotelDAO struct {
	mock.Mock
}

// GetById implementa el método GetById del DAO de hoteles utilizando mocks
func (m *MockHotelDAO) GetById(id string) model.Hotel {
	args := m.Called(id)
	return args.Get(0).(model.Hotel)
}

// Insert implementa el método Insert del DAO de hoteles utilizando mocks
func (m *MockHotelDAO) Insert(hotel model.Hotel) model.Hotel {
	args := m.Called(hotel)
	return args.Get(0).(model.Hotel)
}

// Update implementa el método Update del DAO de hoteles utilizando mocks
func (m *MockHotelDAO) Update(id string, updatedHotel model.Hotel) error {
	args := m.Called(id, updatedHotel)
	return args.Error(0)
}

// MockCollection es un mock de la colección de MongoDB para pruebas
type MockCollection struct {
	mock.Mock
}

// Implementación de métodos de la colección para el mock

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

// Ejemplo de prueba para GetById utilizando el mock
func TestGetById(t *testing.T) {
	// Configurar el mock de Collection
	//mockCollection := new(MockCollection)

	// Crear instancia del mock DAO utilizando el mock de la colección
	mockDAO := &MockHotelDAO{}

	// Definir el ID de ejemplo
	id := primitive.NewObjectID()

	// Definir el hotel esperado
	expectedHotel := model.Hotel{
		Id:          id,
		Name:        "MockHotel",
		Description: "Mock Description",
		Country:     "Mock Country",
		City:        "Mock City",
		Adress:      "Mock Address",
		Images:      []string{"image1.jpg", "image2.jpg"},
		Amenities:   []string{"wifi", "pool"},
	}

	// Configurar el comportamiento esperado del mock para GetById
	mockDAO.On("GetById", id.Hex()).Return(expectedHotel)

	// Llamar a GetById del mock DAO
	result := mockDAO.GetById(id.Hex())

	// Verificar que el resultado sea el esperado
	assert.Equal(t, expectedHotel, result)
}
