package clients

import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetReservaById(id int) model.Reserva {
	var reserva model.Reserva

	Db.Where("id = ?", id).Preload("Hotel").Preload("Cliente").First(&reserva)
	log.Debug("Reserva: ", reserva)

	return reserva
}

func GetReservasById(id int) model.Reservas {
	var reservas model.Reservas

	Db.Where("cliente_id = ?", id).Preload("Hotel").Preload("Cliente").First(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}

func GetReservas() model.Reservas {
	var reservas model.Reservas

	Db.Preload("Hotel").Preload("Cliente").Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}
  
func InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Reserva Creada: ", reserva.ID)
	return reserva
}

func GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) model.Reservas {
	var reservas model.Reservas

	Db.Preload("Hotel").Preload("Cliente").Where("hotel_id = ? AND anio_inicio >= ? OR anio_final <= ? AND mes_inicio >= ? OR mes_final <= ? AND dia_inicio >= ? OR dia_final <= ?", id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal).Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}