package user

import (
	"user-res-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := Db.Where("user_name = ?", username).First(&user)

	log.Debug("User: ", user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) bool {
	var user model.User
	result := Db.Where("email = ?", email).First(&user)

	log.Debug("User: ", user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false // No se encontró el usuario, el email no está registrado
		}
		// Manejo de otros errores, podría ser útil añadir un log aquí
		log.Error("Error buscando usuario por email: ", result.Error)
		return false // Asumimos que el email no está registrado si hay un error distinto
	}

	return true // El usuario existe, el email está registrado
}

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

//Checkear si existe un usuario en el sistema

func CheckUserById(id int) bool {
	var user model.User

	// realza consulta a la base de datos: (con el id proporcionado como parametro)
	result := Db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return false
	}

	return true
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
