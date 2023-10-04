package app

import (

	log "github.com/sirupsen/logrus"
	
	userController "UCC-as2-final/controller/user"
	
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert) // Sign In

	
	// Login
	// router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
