package booking

import (
	"fmt"
	"user-res-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
    
)

var Db *gorm.DB

// GetBookingById sirve para el caso de lista de
// reservas por user
func GetBookingById(id int) model.Booking {

	var booking model.Booking

	if Db == nil {
		
		return booking
	}

	Db.Where("id = ?", id).Preload("User").First(&booking)
	log.Debug("Booking: ", booking)

	
	return booking
}

// todas las reservas con su informacion
func GetBookings() model.Bookings {
	var bookings model.Bookings
	Db.Find(&bookings)

	log.Debug("Bookings: ", bookings)

	return bookings
}


func InsertBooking(booking model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Booking Created: ", booking.Id)
	return booking
}

func GetAvailabilityByIdAndDate(id_hotel int, startDate int) bool {


	var booking model.Booking
	fmt.Println("Antes de la consulta a la base de datos")
	result := Db.Where("hotel_id = ? AND start_date <= ? AND end_date > ?", id_hotel, startDate, startDate).First(&booking) // hay reserva ahi! 
	fmt.Println("Despues  de la consulta a la base de datos")
	if result.Error != nil {
		return false 
	}
	
	return true // La reserva está presente
}

func GetBookingByUserId(id int) model.Booking {
	var booking model.Booking

	Db.Where("user_id = ?", id).Preload("User").First(&booking)
	log.Debug("Booking: ", booking)

	return booking

}

func DeleteBooking(id int) error {

	var booking model.Booking
	
	result := Db.Delete(&booking, id)

	if result.Error != nil {

		return result.Error

	}

	return nil

}
