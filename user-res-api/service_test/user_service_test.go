package service_test

import (
	"fmt"
	// "net/http"
	"testing"
	"user-res-api/dto"
	"user-res-api/model"

	// "user-res-api/service"
	"user-res-api/utils/errors"
	e "user-res-api/utils/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type TestUsers struct{}

func (t *TestUsers) GetUserById(id int) (dto.UserDto, e.ApiError) {
	return dto.UserDto{}, nil
}

func (t *TestUsers) GetUsers() (dto.UsersDto, e.ApiError) {
	return dto.UsersDto{}, nil
}

func (m *MockUserClient) GetUserByEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

func (m *MockUserClient) InsertUser(user model.User) model.User {
	args := m.Called(user)
	return args.Get(0).(model.User)
}

func (m *MockUserClient) GetUserByUsername(username string) (model.User, error) {
	args := m.Called(username)
	return args.Get(0).(model.User), args.Error(1)
}

// Implementación del servicio UserService con MockUserClient
type MockUserService struct {
	UserClient *MockUserClient
}

func NewMockUserService() *MockUserService {
	mockUserClient := new(MockUserClient)
	return &MockUserService{
		UserClient: mockUserClient,
	}
}

func (s *MockUserService) InsertUser(userDto dto.UserDto) (dto.UserDto, errors.ApiError) {
	var user model.User

	if s.UserClient.GetUserByEmail(userDto.Email) {
		return userDto, errors.NewBadRequestApiError("El email ya esta registrado")
	}

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName

	var hashedPassword, err = s.HashPassword(userDto.Password)
	if err != nil {
		return userDto, errors.NewBadRequestApiError("No se puede utilizar esa contraseña")
	}

	user.Password = hashedPassword
	user.Phone = userDto.Phone
	user.Address = userDto.Address
	user.Email = userDto.Email
	user.Type = userDto.Type

	user = s.UserClient.InsertUser(user)
	if user.Id == 0 {
		return userDto, errors.NewBadRequestApiError("Nombre de usuario repetido")
	}

	userDto.Id = user.Id
	return userDto, nil
}

func (s *MockUserService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, errors.ApiError) {
	var loginResponseDto dto.LoginResponseDto
	loginResponseDto.UserId = -1

	// Obtener el usuario por nombre de usuario
	user, err := s.UserClient.GetUserByUsername(loginDto.Username)
	if err != nil {
		return loginResponseDto, errors.NewBadRequestApiError("Usuario no encontrado")
	}

	// Verificar la contraseña
	if err := s.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return loginResponseDto, errors.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	// Generar el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
	})
	var jwtKey = []byte("secret_key")
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return loginResponseDto, errors.NewInternalServerApiError("Error generando el token", err)
	}

	// Asignar los valores al DTO de respuesta
	loginResponseDto.UserId = user.Id
	loginResponseDto.Token = tokenString
	loginResponseDto.Name = user.Name
	loginResponseDto.LastName = user.LastName
	loginResponseDto.UserName = user.UserName
	loginResponseDto.Email = user.Email
	loginResponseDto.Type = user.Type

	return loginResponseDto, nil
}

func (s *MockUserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("No se pudo hashear la password %w", err)
	}
	return string(hashedPassword), nil
}

func (s *MockUserService) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func TestInsertUser(t *testing.T) {
	// Crear el mock del servicio
	mockService := NewMockUserService()

	// Configurar el comportamiento esperado en el mock
	mockService.UserClient.On("GetUserByEmail", "test1@example.com").Return(false)
	expectedUser := model.User{
		Id:       1,
		Name:     "Test",
		LastName: "User",
		UserName: "testuser",
		Email:    "test1@example.com",
		Password: "hashed_password",
		Phone:    1234567890,
		Address:  "Test Address",
		Type:     true,
	}
	mockService.UserClient.On("InsertUser", mock.AnythingOfType("model.User")).Return(expectedUser, nil) // Asegúrate de devolver nil como error

	// Crear el DTO de entrada
	userDto := dto.UserDto{
		Name:     "Test",
		LastName: "User",
		UserName: "testuser",
		Email:    "test1@example.com",
		Password: "password",
		Phone:    1234567890,
		Address:  "Test Address",
		Type:     true,
	}

	// Llamar a la función que estamos probando
	result, err := mockService.InsertUser(userDto)

	// Verificar los resultados esperados
	assert.NoError(t, err, "Se recibió un error inesperado al insertar el usuario")
	assert.Equal(t, expectedUser.Id, result.Id, "El ID del usuario no es el esperado")
	assert.Equal(t, expectedUser.Name, result.Name, "El nombre del usuario no es el esperado")
	assert.Equal(t, expectedUser.LastName, result.LastName, "El apellido del usuario no es el esperado")
	assert.Equal(t, expectedUser.UserName, result.UserName, "El nombre de usuario no es el esperado")
	assert.Equal(t, expectedUser.Email, result.Email, "El email del usuario no es el esperado")
	assert.Equal(t, expectedUser.Phone, result.Phone, "El teléfono del usuario no es el esperado")
	assert.Equal(t, expectedUser.Address, result.Address, "La dirección del usuario no es la esperada")
	assert.Equal(t, expectedUser.Type, result.Type, "El tipo de usuario no es el esperado")

	// Verificar que todas las expectativas del mock se cumplieron
	mockService.UserClient.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockService := NewMockUserService()

	// Configurar el comportamiento esperado en el mock
	password := "password"                                  // La contraseña sin hash
	hashedPassword, _ := mockService.HashPassword(password) // Hash de la contraseña
	expectedUser := model.User{
		Id:       1,
		UserName: "testuser",
		Password: hashedPassword,
	}
	mockService.UserClient.On("GetUserByUsername", "testuser").Return(expectedUser, nil)

	// Crear el DTO de entrada
	loginDto := dto.LoginDto{
		Username: "testuser",
		Password: password,
	}

	// Llamar a la función que estamos probando
	result, err := mockService.Login(loginDto)

	// Verificar los resultados esperados
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, result.UserId)
	assert.Equal(t, expectedUser.UserName, result.UserName)
	assert.NotEmpty(t, result.Token)

	mockService.UserClient.AssertExpectations(t)
}
