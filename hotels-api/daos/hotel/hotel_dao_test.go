package hotel_test

import (
	"context"
	"testing"

	model "hotels-api/models"
	// "hotels-api/utils/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface para el acceso a la base de datos
type DatabaseHandler interface {
	Collection(name string) CollectionHandler
}

// Interface para las colecciones
type CollectionHandler interface {
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
}

// Implementación de la interfaz para MongoDB
type MongoHandler struct{}

func (m *MongoHandler) Collection(name string) CollectionHandler {
	return &MongoCollection{}
}

// Implementación de la interfaz para las colecciones en MongoDB
type MongoCollection struct{}

func (m *MongoCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	// Implementación simulada de FindOne
	return nil // Implementa según tus necesidades de prueba
}

func (m *MongoCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	// Implementación simulada de InsertOne
	return nil, nil // Implementa según tus necesidades de prueba
}

func (m *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	// Implementación simulada de UpdateOne
	return nil, nil // Implementa según tus necesidades de prueba
}

// Mock del DAO para la prueba
type MockDAO struct {
	mock.Mock
}

func (m *MockDAO) Collection(name string) CollectionHandler {
	args := m.Called(name)
	return args.Get(0).(CollectionHandler)
}

// Mock de la colección para la prueba
type MockCollection struct {
	mock.Mock
}

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

func GetById(id string) model.Hotel {
	dbHandler := &MongoHandler{}
	collection := dbHandler.Collection("hotels")

	var hotel model.Hotel
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Handle the error
		return hotel
	}

	result := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}})
	if result.Err() != nil {
		// Handle the error
		return hotel
	}

	err = result.Decode(&hotel)
	if err != nil {
		// Handle the error
		return hotel
	}

	return hotel
}

func Insert(hotel model.Hotel) (*mongo.InsertOneResult, error) {
	dbHandler := &MongoHandler{}
	collection := dbHandler.Collection("hotels")

	insertHotel := hotel
	insertHotel.Id = primitive.NewObjectID()

	return collection.InsertOne(context.TODO(), &insertHotel)
}

// Función original del DAO para actualizar un hotel por ID
func Update(id string, updatedHotel model.Hotel) error {
	dbHandler := &MongoHandler{}
	collection := dbHandler.Collection("hotels")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Handle the error (possibly return an error or log it)
		return err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: updatedHotel.Name},
			{Key: "description", Value: updatedHotel.Description},
			{Key: "city", Value: updatedHotel.City},
			{Key: "country", Value: updatedHotel.Country},
			{Key: "address", Value: updatedHotel.Adress},
			{Key: "images", Value: updatedHotel.Images},
			{Key: "amenities", Value: updatedHotel.Amenities},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}, update)
	return err
}

func TestGetById(t *testing.T) {
	// Configurar el mock del DAO
	mockDAO := &MockDAO{}
	mockCollection := &MockCollection{}
	mockDAO.On("Collection", "hotels").Return(mockCollection)

	// Configurar el comportamiento del mock para FindOne
	mockResult := &model.Hotel{Id: primitive.NewObjectID(), Name: "MockHotel"}
	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockResult, nil)

	// Llamar a la función bajo prueba
	result := GetById("some_id")

	// Verificar el resultado esperado
	assert.Equal(t, mockResult, result)
	mockDAO.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestInsert(t *testing.T) {
	// Configurar el mock del DAO
	mockDAO := &MockDAO{}
	mockCollection := &MockCollection{}
	mockDAO.On("Collection", "hotels").Return(mockCollection)

	// Configurar el comportamiento del mock para InsertOne
	mockInsertResult := &mongo.InsertOneResult{}
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(mockInsertResult, nil)

	// Llamar a la función bajo prueba
	result, err := Insert(model.Hotel{})
	if err != nil {

	}
	// Verificar el resultado esperado
	assert.Equal(t, mockInsertResult, result)
	mockDAO.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	// Configurar el mock del DAO
	mockDAO := &MockDAO{}
	mockCollection := &MockCollection{}
	mockDAO.On("Collection", "hotels").Return(mockCollection)

	// Configurar el comportamiento del mock para UpdateOne
	mockUpdateResult := &mongo.UpdateResult{}
	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(mockUpdateResult, nil)

	// Llamar a la función bajo prueba
	err := Update("some_id", model.Hotel{})

	// Verificar el resultado esperado
	assert.NoError(t, err)
	mockDAO.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}
