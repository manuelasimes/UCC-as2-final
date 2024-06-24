package user_test

import (
	// "fmt"
	"testing"
	"user-res-api/model"

	"github.com/stretchr/testify/assert"
)

// Mock para simular el comportamiento de la base de datos
type MockUserDao struct{}

func (m *MockUserDao) InsertUser(user model.User) model.User {
	// Simular la lógica de inserción en la base de datos
	// Se establece un ID para el hotel
	user.Id = 1 // Si modifico a cero, genera la alerta
	return user
}

func TestInsertUser(t *testing.T) {
	// Crear una instancia del mock del DAO de HOtel
	mockDAO := &MockUserDao{}

	// Crear una nueva reserva
	newUser := model.User{

		Id:       1,
		Name:     "test",
		LastName: "test",
		UserName: "test_test",
		Password: "testing",
		Phone:    1234,
		Email:    "test",
		Address:  "test",
		Type:     true,
	}

	// Insertar la reserva utilizando el mock del DAO
	inserted := mockDAO.InsertUser(newUser)

	// Verificar que la reserva tenga un ID asignado
	assert.NotZero(t, inserted.Id, "La insercion dle user no se pudo realizar")

	// Verificar otros atributos de la reserva
	assert.Equal(t, newUser.Id, inserted.Id)
	assert.Equal(t, newUser.Name, inserted.Name)
	assert.Equal(t, newUser.LastName, inserted.LastName)
	assert.Equal(t, newUser.UserName, inserted.UserName)
	assert.Equal(t, newUser.Password, inserted.Password)
	assert.Equal(t, newUser.Phone, inserted.Phone)
	assert.Equal(t, newUser.Address, inserted.Address)
	assert.Equal(t, newUser.Email, inserted.Email)
	assert.Equal(t, newUser.Type, inserted.Type)
}

// TEST PARA LA FUNCION GETBOOKINGBYID

func (m *MockUserDao) GetUserById(id int) model.User {
	// Simular la búsqueda en la base de datos
	user := model.User{
		Id:       1,
		Name:     "test",
		LastName: "test",
		UserName: "test_test",
		Password: "testing",
		Phone:    1234,
		Email:    "test",
		Address:  "test",
		Type:     true,
	}

	return user
}

func TestGetUserById(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockDAO := &MockUserDao{}

	// ID de reserva a buscar - Si la cambio deja de funcionar
	userId := 1

	// Obtener la reserva utilizando el mock del DAO
	user := mockDAO.GetUserById(userId)

	// Verificar que la reserva obtenida tenga el ID correcto
	assert.Equal(t, userId, user.Id, "El ID del user no existe")

	// Verificar otros atributos de la reserva
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test", user.LastName)
	assert.Equal(t, "test_test", user.UserName)
	assert.Equal(t, "test", user.Email)
	assert.Equal(t, "test", user.Address)
	assert.Equal(t, 1234, user.Phone)
	assert.Equal(t, "testing", user.Password)
	assert.Equal(t, true, user.Type)
}
