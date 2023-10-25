package clients

import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).Preload("Telefonos").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotelByEmail(email string) model.Hotel {
	var hotel model.Hotel

	Db.Where("email = ?", email).Preload("Telefonos").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotelByNombre(nombre string) model.Hotel {
	var hotel model.Hotel

	Db.Where("nombre = ?", nombre).Preload("Telefonos").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHoteles() model.Hoteles {
	var hoteles model.Hoteles

	Db.Preload("Telefonos").Find(&hoteles)
	log.Debug("Hoteles: ", hoteles)

	return hoteles
}
  

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("")
	}
	
	log.Debug("hotel Creado: ", hotel.ID)
	return hotel
}