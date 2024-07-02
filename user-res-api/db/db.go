package db

import (
	bookingClient "user-res-api/client/booking"
	hotelClient "user-res-api/client/hotel"
	userClient "user-res-api/client/user"
	"user-res-api/model"

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
	DBName := "UCC_as2_final"
	DBUser := "root"
	DBPass := ""
	//DBPass := "Manuela10Simes"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "mysql"
	//DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all Clients that we build
	userClient.Db = db
	bookingClient.Db = db
	hotelClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.Hotel{})

	log.Info("Finishing Migration Database Tables")
}
