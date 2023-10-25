package db

import (
	adminClient "backend/clients/admin"
	clienteClient "backend/clients/cliente"
	hotelClient "backend/clients/hotel"
	reservaClient "backend/clients/reserva"
	telefonoClient "backend/clients/telefono"

	"backend/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsoft"
	DBUser := "root"
	DBPass := "2001Mama"
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("La conexión no se pudo abrir")
		log.Fatal(err)
	} else {
		log.Info("Conexión establecida")
	}

	adminClient.Db = db
	clienteClient.Db = db
	hotelClient.Db = db
	telefonoClient.Db = db
	reservaClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Cliente{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Reserva{})
	db.AutoMigrate(&model.Telefono{})

	log.Info("Finalización de las tablas de la base de datos de migración")
}
