package hotel

import (
	"user-res-api/model"
	// "encoding/json"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotelByIdMongo(id string) model.Hotel {
	var hotel model.Hotel

	Db.Where("id_mongo = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

// creo funcio que me permita ver si hay algun hotel en la bd con ese id de amadeus
func GetHotelByIdAmadeus(idam string) bool {
	var hotel model.Hotel

	result := Db.Where("id_amadeus = ?", idam).First(&hotel)

	if result.Error != nil {
		return false
	}

	return true // si devuelve true quiere decir q ya existe un hotel con ese id de amadeus
}

func CheckHotelById(id int) bool {
	var hotel model.Hotel

	result := Db.Where("id = ?", id).First(&hotel)

	if result.Error != nil {
		return false
	}

	return true
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
		hotel.Id = 0
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

// implementando el delete todo con el id de mongo que desde hotels api nos mandan
func DeleteHotel(idmongo string) error {
	result := Db.Where("id_mongo = ?", idmongo).Delete(&model.Hotel{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func UpdateHotelById(hotel model.Hotel) model.Hotel {

// 	result := Db.Model(&hotel).Where("id = ? ", hotel.Id).Updates(map[string]interface{}{"hotel_name": hotel.HotelName, "hotel_description": hotel.HotelDescription, "rooms": hotel.Rooms, "address": hotel.Address})

// 	if result.Error != nil {
// 		log.Error("")
// 		hotel.Id = 0
// 	}

// 	log.Debug("Hotel Updated or Created: ", hotel.Id)

// 	return hotel
// }

// func DeleteHotel(hotel model.Hotel) (bool, error) {

// 	result := Db.Delete(&hotel, hotel.Id)

// 	if result.Error != nil {
// 		return false, result.Error
// 	}

// 	return true, nil
// }

// func UpdateHotel(hotel model.Hotel) {
// 	Db.Save(&hotel)
// 	log.Debug("Hotel Updated: ", hotel.Id)
// }
