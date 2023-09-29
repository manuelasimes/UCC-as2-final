package db

import (
	
	userClient "UCC-as2-final/client/user"
	bookingClient "UCC-as2-final/client/user"
	"UCC-as2-final/model"

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
	DBPass := "Manuela10Simes"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
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
	

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Booking{})
	
	log.Info("Finishing Migration Database Tables")
}
