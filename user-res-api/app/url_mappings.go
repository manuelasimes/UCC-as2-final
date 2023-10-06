package app

import (

	log "github.com/sirupsen/logrus"
	
	userController "UCC-as2-final/controller/user"
	bookingController "UCC-as2-final/controller/booking"
	
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert) // Sign In

	
	// Bookings Mapping
	router.GET("/booking/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	// router.POST("/booking", bookingController.InsertBooking)
	router.GET("/booking/user/:user_id", bookingController.GetBookingsByUserId)
	router.GET("/hotel/availability/:id/:start_date/:end_date", bookingController.GetAvailabilityByIdAndDate)

	
	// Login
	// router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
