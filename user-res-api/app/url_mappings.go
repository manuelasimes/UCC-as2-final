package app

import (

	log "github.com/sirupsen/logrus"
	
	userController "user-res-api/controller/user"
	bookingController "user-res-api/controller/booking"
	hotelController "user-res-api/controller/hotel"
	
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert) // Sign In

	
	// Bookings Mapping
	router.GET("/booking/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	router.POST("/booking", bookingController.InsertBooking)
	router.GET("/booking/user/:user_id", bookingController.GetBookingsByUserId)
	router.GET("/hotel/availability/:id/:start_date/:end_date", bookingController.GetAvailabilityByIdAndDate)

	// Hotels Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.InsertHotel)
	//router.PUT("/hotel/update", hotelController.UpdateHotel)
	//router.DELETE("/hotel/delete/:hotel_id/:user_id", hotelController.DeleteHotel)

	
	// Login
	// router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
