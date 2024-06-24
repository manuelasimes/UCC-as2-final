package dto_test

import (
	"testing"
	"user-res-api/dto"

	"github.com/stretchr/testify/assert"
)

func TestUserDto(t *testing.T) {
	// Crear una instancia del DTO de Hotel, si modifico alguna y deja de ser igual, da la alerta
	userDto := dto.UserDto{
		Id:       1,
		Name:     "Manuela",
		LastName: "Simes",
		Phone:    351554433,
		Address:  "Pacha 5678",
		Password: "contraseña",
		Email:    "manusimes@gmail.com",
		Type:     true,
	}

	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, 1, userDto.Id, "El ID del user no coincide")
	assert.Equal(t, "Manuela", userDto.Name, "EL nombre del user no coincide")
	assert.Equal(t, "Simes", userDto.LastName, "El apellido del user no coincide")
	assert.Equal(t, 351554433, userDto.Phone, "El numero del user no coincide")
	assert.Equal(t, "Pacha 5678", userDto.Address, "La direccion no coincide")
	assert.Equal(t, "contraseña", userDto.Password, "La contraseña del user no coincide")
	assert.Equal(t, "manusimes@gmail.com", userDto.Email, "El mail del user no coincide")
	assert.Equal(t, true, userDto.Type, "El type del user no coincide")

}
