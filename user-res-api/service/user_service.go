package service

import (
	"fmt"
	"time"
	userClient "user-res-api/client/user"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"user-res-api/dto"
	"user-res-api/model"
	e "user-res-api/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto *dto.UserDto) (*dto.UserDto, e.ApiError)
	GetUserById(id int) (*dto.UserDto, e.ApiError)
	Login(loginDto *dto.LoginDto) (*dto.LoginResponseDto, e.ApiError)
	Refresh(refreshTokenDto *dto.RefreshTokenDto) (*dto.LoginResponseDto, e.ApiError)
}

var (
	UserService userServiceInterface
	jwtKey      = []byte("secret_key")
	refreshKey  = []byte("refresh_secret_key")
)

func init() {
	UserService = &userService{}
}

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Type     bool   `json:"type"` // Suponiendo que Type es un booleano en el modelo
	jwt.StandardClaims
}

func GenerateAccessToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		UserId:   user.Id,
		Username: user.UserName,
		Type:     user.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GenerateRefreshToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserId:   user.Id,
		Username: user.UserName,
		Type:     user.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshKey)
}

func ValidateRefreshToken(refreshToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})
	if err != nil {
		return nil, e.NewUnauthorizedApiError("Token inválido")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, e.NewUnauthorizedApiError("Token inválido")
}

func (s *userService) GetUserById(id int) (*dto.UserDto, e.ApiError) {
	user := userClient.GetUserById(id)
	if user.Id == 0 {
		return nil, e.NewBadRequestApiError("Usuario no encontrado")
	}

	userDto := &dto.UserDto{
		Name:     user.Name,
		LastName: user.LastName,
		UserName: user.UserName,
		Phone:    user.Phone,
		Address:  user.Address,
		Email:    user.Email,
		Id:       user.Id,
		Type:     user.Type, // Suponiendo que user.Type es string
	}
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {
	users := userClient.GetUsers()
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
			Type:     user.Type, // Suponiendo que user.Type es string
		}
		usersDto = append(usersDto, userDto)
	}
	return usersDto, nil
}

func (s *userService) InsertUser(userDto *dto.UserDto) (*dto.UserDto, e.ApiError) {
	if userClient.GetUserByEmail(userDto.Email) {
		return nil, e.NewBadRequestApiError("El email ya está registrado")
	}

	hashedPassword, err := s.HashPassword(userDto.Password)
	if err != nil {
		return nil, e.NewBadRequestApiError("No se puede utilizar esa contraseña")
	}

	user := model.User{
		Name:     userDto.Name,
		LastName: userDto.LastName,
		UserName: userDto.UserName,
		Password: hashedPassword,
		Phone:    userDto.Phone,
		Address:  userDto.Address,
		Email:    userDto.Email,
		Type:     userDto.Type, // Suponiendo que userDto.Type es string
	}

	user = userClient.InsertUser(user)
	if user.Id == 0 {
		return nil, e.NewBadRequestApiError("Nombre de usuario repetido")
	}

	userDto.Id = user.Id
	return userDto, nil
}

func (s *userService) Login(loginDto *dto.LoginDto) (*dto.LoginResponseDto, e.ApiError) {
	user, err := userClient.GetUserByUsername(loginDto.Username)
	if err != nil {
		return nil, e.NewBadRequestApiError("Usuario no encontrado")
	}

	if err := s.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	accessToken, err := GenerateAccessToken(user)
	if err != nil {
		return nil, e.NewInternalServerApiError("Error al generar el access token", err)
	}

	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		return nil, e.NewInternalServerApiError("Error al generar el refresh token", err)
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
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

func (s *userService) Refresh(refreshTokenDto *dto.RefreshTokenDto) (*dto.LoginResponseDto, e.ApiError) {
	claims, err := ValidateRefreshToken(refreshTokenDto.RefreshToken)
	if err != nil {
		return nil, err.(e.ApiError) // Convertir el error a ApiError
	}

	// Obtener el usuario usando el user_id del claim
	user := userClient.GetUserById(claims.UserId)
	if user.Id == 0 {
		return nil, e.NewBadRequestApiError("Usuario no encontrado")
	}

	accessToken, err := GenerateAccessToken(user)
	if err != nil {
		return nil, e.NewInternalServerApiError("Error al generar el access token", err)
	}

	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		return nil, e.NewInternalServerApiError("Error al generar el refresh token", err)
	}

	loginResponseDto := &dto.LoginResponseDto{
		UserId:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserName:     user.UserName,
		Type:         user.Type,
	}

	return loginResponseDto, nil
}

func (s *userService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("No se pudo hashear la contraseña: %w", err)
	}
	return string(hashedPassword), nil
}

func (s *userService) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
