package clients

import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAdminById(id int) model.Admin {
	var admin model.Admin

	Db.Where("id = ?", id).First(&admin)
	log.Debug("Admin: ", admin)

	return admin
}

func GetAdminByUsername(username string) model.Admin {
	var admin model.Admin

	Db.Where("user_name = ?", username).First(&admin)
	log.Debug("Admin: ", admin)

	return admin
}

func GetAdminByEmail(email string) model.Admin {
	var admin model.Admin

	Db.Where("email = ?", email).First(&admin)
	log.Debug("Admin: ", admin)

	return admin
}

func GetAdmins() model.Admins {
	var admins model.Admins
	
	Db.Find(&admins)
	log.Debug("Administradores: ", admins)

	return admins
}

func InsertAdmin(admin model.Admin) model.Admin {
	result := Db.Create(&admin)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Administrador Creado: ", admin.ID)
	return admin
}