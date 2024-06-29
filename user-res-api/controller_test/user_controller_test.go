package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-res-api/controller/user"
	"user-res-api/dto"
	"user-res-api/service"
	"user-res-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserById(id int) (*dto.UserDto, errors.ApiError) {
	args := m.Called(id)
	userDto := args.Get(0).(*dto.UserDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return userDto, apiErr
}

func (m *MockUserService) GetUsers() (dto.UsersDto, errors.ApiError) {
	args := m.Called()
	usersDto := args.Get(0).(dto.UsersDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return usersDto, apiErr
}

func (m *MockUserService) InsertUser(userDto *dto.UserDto) (*dto.UserDto, errors.ApiError) {
	args := m.Called(userDto)
	newUserDto := args.Get(0).(*dto.UserDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return newUserDto, apiErr
}

func (m *MockUserService) Login(loginDto *dto.LoginDto) (*dto.LoginResponseDto, errors.ApiError) {
	args := m.Called(loginDto)
	loginResponseDto := args.Get(0).(*dto.LoginResponseDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return loginResponseDto, apiErr
}

func (m *MockUserService) Refresh(refreshTokenDto *dto.RefreshTokenDto) (*dto.LoginResponseDto, errors.ApiError) {
	args := m.Called(refreshTokenDto)
	refreshResponseDto := args.Get(0).(*dto.LoginResponseDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return refreshResponseDto, apiErr
}

func TestGetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/user/:id", user.GetUserById)

	mockService := new(MockUserService)
	service.UserService = mockService

	expectedUser := &dto.UserDto{
		Id:       1,
		Name:     "test",
		LastName: "test",
		UserName: "testuser",
	}
	mockService.On("GetUserById", 1).Return(expectedUser, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualUser dto.UserDto
	err := json.Unmarshal(w.Body.Bytes(), &actualUser)
	assert.NoError(t, err)
	assert.Equal(t, *expectedUser, actualUser)

	mockService.AssertExpectations(t)
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", user.GetUsers)

	mockService := new(MockUserService)
	service.UserService = mockService

	expectedUsers := dto.UsersDto{
		{Id: 1, UserName: "user1"},
		{Id: 2, UserName: "user2"},
	}
	mockService.On("GetUsers").Return(expectedUsers, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualUsers dto.UsersDto
	err := json.Unmarshal(w.Body.Bytes(), &actualUsers)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actualUsers)

	mockService.AssertExpectations(t)
}

func TestUserInsert(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/user", user.UserInsert)

	mockService := new(MockUserService)
	service.UserService = mockService

	newUser := &dto.UserDto{
		Id:       1,
		UserName: "newuser",
	}
	mockService.On("InsertUser", newUser).Return(newUser, nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/user", nil)
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var actualUser dto.UserDto
	err := json.Unmarshal(w.Body.Bytes(), &actualUser)
	assert.NoError(t, err)
	assert.Equal(t, *newUser, actualUser)

	mockService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", user.Login)

	mockService := new(MockUserService)
	service.UserService = mockService

	loginDto := &dto.LoginDto{
		Username: "testuser",
		Password: "password",
	}
	loginResponse := &dto.LoginResponseDto{
		UserId:       1,
		AccessToken:  "testaccesstoken",
		RefreshToken: "testrefreshtoken",
		Name:         "Test",
		LastName:     "User",
		UserName:     "testuser",
		Email:        "testuser@example.com",
		Type:         true,
	}
	mockService.On("Login", loginDto).Return(loginResponse, nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(loginDto)
	req, _ := http.NewRequest("POST", "/login", nil)
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualLoginResponse dto.LoginResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &actualLoginResponse)
	assert.NoError(t, err)
	assert.Equal(t, *loginResponse, actualLoginResponse)

	mockService.AssertExpectations(t)
}

func TestRefresh(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/refresh", user.Refresh)

	mockService := new(MockUserService)
	service.UserService = mockService

	refreshTokenDto := &dto.RefreshTokenDto{
		RefreshToken: "testrefresh",
	}
	refreshResponse := &dto.LoginResponseDto{
		UserId:       1,
		AccessToken:  "newaccesstoken",
		RefreshToken: "newrefreshtoken",
		Name:         "Test",
		LastName:     "User",
		UserName:     "testuser",
		Email:        "testuser@example.com",
		Type:         true,
	}
	mockService.On("Refresh", refreshTokenDto).Return(refreshResponse, nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(refreshTokenDto)
	req, _ := http.NewRequest("POST", "/refresh", nil)
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualRefreshResponse dto.LoginResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &actualRefreshResponse)
	assert.NoError(t, err)
	assert.Equal(t, *refreshResponse, actualRefreshResponse)

	mockService.AssertExpectations(t)
}
