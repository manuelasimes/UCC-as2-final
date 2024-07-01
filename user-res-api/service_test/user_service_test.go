package service_test

import (
	"fmt"
	"testing"
	"time"
	"user-res-api/dto"
	"user-res-api/model"
	"user-res-api/service"
	"user-res-api/utils/errors"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// GetUserById mock del método GetUserById
func (m *MockUserClient) GetUserById(id int) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

// GetUsers mock del método GetUsers
func (m *MockUserClient) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

// GetUserByEmail mock del método GetUserByEmail
func (m *MockUserClient) GetUserByEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

// InsertUser mock del método InsertUser
func (m *MockUserClient) InsertUser(user model.User) model.User {
	args := m.Called(user)
	return args.Get(0).(model.User)
}

// GetUserByUsername mock del método GetUserByUsername
func (m *MockUserClient) GetUserByUsername(username string) (model.User, error) {
	args := m.Called(username)
	return args.Get(0).(model.User), args.Error(1)
}

// MockUserService estructura del servicio mock
type MockUserService struct {
	UserClient *MockUserClient
}

// NewMockUserService crea una nueva instancia de MockUserService
func NewMockUserService() *MockUserService {
	mockUserClient := new(MockUserClient)
	return &MockUserService{
		UserClient: mockUserClient,
	}
}

// InsertUser implementa el método InsertUser del servicio
func (s *MockUserService) InsertUser(userDto dto.UserDto) (*dto.UserDto, errors.ApiError) {
	if s.UserClient.GetUserByEmail(userDto.Email) {
		return nil, errors.NewBadRequestApiError("El email ya está registrado")
	}

	hashedPassword, err := s.HashPassword(userDto.Password)
	if err != nil {
		return nil, errors.NewBadRequestApiError("No se puede utilizar esa contraseña")
	}

	user := model.User{
		Name:     userDto.Name,
		LastName: userDto.LastName,
		UserName: userDto.UserName,
		Password: hashedPassword,
		Phone:    userDto.Phone,
		Address:  userDto.Address,
		Email:    userDto.Email,
		Type:     userDto.Type,
	}

	user = s.UserClient.InsertUser(user)
	if user.Id == 0 {
		return nil, errors.NewBadRequestApiError("Nombre de usuario repetido")
	}

	userDto.Id = user.Id
	return &userDto, nil
}

// Login implementa el método Login del servicio
func (s *MockUserService) Login(loginDto *dto.LoginDto) (*dto.LoginResponseDto, errors.ApiError) {
	user, err := s.UserClient.GetUserByUsername(loginDto.Username)
	if err != nil {
		return nil, errors.NewBadRequestApiError("Usuario no encontrado")
	}

	if err := s.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, errors.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	accessToken, err := service.GenerateAccessToken(user)
	if err != nil {
		return nil, errors.NewInternalServerApiError("Error al generar el access token", err)
	}

	refreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.NewInternalServerApiError("Error al generar el refresh token", err)
	}

	loginResponseDto := &dto.LoginResponseDto{
		UserId:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Name:         user.Name,
		LastName:     user.LastName,
		UserName:     user.UserName,
		Email:        user.Email,
		Type:         user.Type,
	}
	return loginResponseDto, nil
}

// HashPassword implementa el método HashPassword del servicio
func (s *MockUserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("No se pudo hashear la password %w", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword implementa el método VerifyPassword del servicio
func (s *MockUserService) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

// GetUserById implementa el método GetUserById del servicio
func (s *MockUserService) GetUserById(id int) (*dto.UserDto, errors.ApiError) {
	user, err := s.UserClient.GetUserById(id)
	if err != nil {
		return nil, errors.NewBadRequestApiError("Usuario no encontrado")
	}

	userDto := &dto.UserDto{
		Name:     user.Name,
		LastName: user.LastName,
		UserName: user.UserName,
		Phone:    user.Phone,
		Address:  user.Address,
		Email:    user.Email,
		Id:       user.Id,
		Type:     user.Type,
	}
	return userDto, nil
}

// GetUsers implementa el método GetUsers del servicio
func (s *MockUserService) GetUsers() (dto.UsersDto, errors.ApiError) {
	users, err := s.UserClient.GetUsers()
	if err != nil {
		return nil, errors.NewInternalServerApiError("Error al obtener los usuarios", err)
	}

	var usersDto dto.UsersDto
	for _, user := range users {
		userDto := dto.UserDto{
			Name:     user.Name,
			LastName: user.LastName,
			UserName: user.UserName,
			Phone:    user.Phone,
			Address:  user.Address,
			Email:    user.Email,
			Id:       user.Id,
			Type:     user.Type,
		}
		usersDto = append(usersDto, userDto)
	}
	return usersDto, nil
}

// Refresh implementa el método Refresh del servicio
func (s *MockUserService) Refresh(refreshTokenDto *dto.RefreshTokenDto) (*dto.LoginResponseDto, errors.ApiError) {
	claims, err := service.ValidateRefreshToken(refreshTokenDto.RefreshToken)
	if err != nil {
		return nil, err.(errors.ApiError)
	}

	user, apiErr := s.UserClient.GetUserByUsername(claims.Username)
	if apiErr != nil {
		return nil, errors.NewInternalServerApiError("Error al obtener el usuario", apiErr)
	}

	accessToken, err := service.GenerateAccessToken(user)
	if err != nil {
		return nil, errors.NewInternalServerApiError("Error al generar el access token", err)
	}

	refreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.NewInternalServerApiError("Error al generar el refresh token", err)
	}

	loginResponseDto := &dto.LoginResponseDto{
		UserId:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Name:         user.Name,
		LastName:     user.LastName,
		UserName:     user.UserName,
		Email:        user.Email,
		Type:         user.Type,
	}
	return loginResponseDto, nil
}

// Tests de los métodos del servicio

func TestInsertUser(t *testing.T) {
	mockService := NewMockUserService()

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
	mockService.UserClient.On("InsertUser", mock.AnythingOfType("model.User")).Return(expectedUser)

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

	result, err := mockService.InsertUser(userDto)

	assert.NoError(t, err, "Se recibió un error inesperado al insertar el usuario")
	assert.Equal(t, expectedUser.Id, result.Id, "El ID del usuario no es el esperado")
	assert.Equal(t, expectedUser.Name, result.Name, "El nombre del usuario no es el esperado")
	assert.Equal(t, expectedUser.LastName, result.LastName, "El apellido del usuario no es el esperado")
	assert.Equal(t, expectedUser.UserName, result.UserName, "El nombre de usuario no es el esperado")
	assert.Equal(t, expectedUser.Email, result.Email, "El email del usuario no es el esperado")
	assert.Equal(t, expectedUser.Phone, result.Phone, "El teléfono del usuario no es el esperado")
	assert.Equal(t, expectedUser.Address, result.Address, "La dirección del usuario no es la esperada")
	assert.Equal(t, expectedUser.Type, result.Type, "El tipo de usuario no es el esperado")

	mockService.UserClient.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockService := NewMockUserService()

	password := "password"
	hashedPassword, _ := mockService.HashPassword(password)
	expectedUser := model.User{
		Id:       1,
		UserName: "testuser",
		Password: hashedPassword,
	}
	mockService.UserClient.On("GetUserByUsername", "testuser").Return(expectedUser, nil)

	loginDto := &dto.LoginDto{
		Username: "testuser",
		Password: password,
	}

	result, err := mockService.Login(loginDto)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, result.UserId)
	assert.Equal(t, expectedUser.UserName, result.UserName)
	assert.NotEmpty(t, result.AccessToken)

	mockService.UserClient.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	mockService := NewMockUserService()

	expectedUser := model.User{
		Id:       1,
		Name:     "Test",
		LastName: "User",
		UserName: "testuser",
		Phone:    1234567890,
		Address:  "Test Address",
		Email:    "test1@example.com",
		Type:     true,
	}
	mockService.UserClient.On("GetUserById", 1).Return(expectedUser, nil)

	result, err := mockService.GetUserById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, result.Id)
	assert.Equal(t, expectedUser.Name, result.Name)
	assert.Equal(t, expectedUser.LastName, result.LastName)
	assert.Equal(t, expectedUser.UserName, result.UserName)
	assert.Equal(t, expectedUser.Phone, result.Phone)
	assert.Equal(t, expectedUser.Address, result.Address)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.Type, result.Type)

	mockService.UserClient.AssertExpectations(t)
}

func TestGetUsers(t *testing.T) {
	mockService := NewMockUserService()

	expectedUsers := []model.User{
		{
			Id:       1,
			Name:     "Test",
			LastName: "User",
			UserName: "testuser1",
			Phone:    1234567890,
			Address:  "Test Address 1",
			Email:    "test1@example.com",
			Type:     true,
		},
		{
			Id:       2,
			Name:     "Another",
			LastName: "User",
			UserName: "testuser2",
			Phone:    987654321,
			Address:  "Test Address 2",
			Email:    "test2@example.com",
			Type:     false,
		},
	}
	mockService.UserClient.On("GetUsers").Return(expectedUsers, nil)

	result, err := mockService.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, len(expectedUsers), len(result))

	for i, user := range expectedUsers {
		assert.Equal(t, user.Id, result[i].Id)
		assert.Equal(t, user.Name, result[i].Name)
		assert.Equal(t, user.LastName, result[i].LastName)
		assert.Equal(t, user.UserName, result[i].UserName)
		assert.Equal(t, user.Phone, result[i].Phone)
		assert.Equal(t, user.Address, result[i].Address)
		assert.Equal(t, user.Email, result[i].Email)
		assert.Equal(t, user.Type, result[i].Type)
	}

	mockService.UserClient.AssertExpectations(t)
}

func TestRefresh(t *testing.T) {
	mockService := NewMockUserService()

	refreshTokenDto := &dto.RefreshTokenDto{}
	claims := &service.Claims{
		Username: "testuser",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, _ := token.SignedString([]byte("refresh_secret_key"))

	refreshTokenDto.RefreshToken = refreshToken

	expectedUser := model.User{
		Id:       1,
		UserName: "testuser",
	}
	mockService.UserClient.On("GetUserByUsername", "testuser").Return(expectedUser, nil)

	result, err := mockService.Refresh(refreshTokenDto)

	assert.NoError(t, err, "Se recibió un error inesperado al refrescar el token")
	assert.Equal(t, expectedUser.Id, result.UserId, "El ID del usuario no es el esperado")
	assert.NotEmpty(t, result.AccessToken, "El access token no debería estar vacío")
	assert.NotEmpty(t, result.RefreshToken, "El refresh token no debería estar vacío")

	mockService.UserClient.AssertExpectations(t)
}
