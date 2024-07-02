package app

import (
	log "github.com/sirupsen/logrus"

	bookingController "user-res-api/controller/booking"
	hotelController "user-res-api/controller/hotel"
	userController "user-res-api/controller/user"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user-res-api/user/:id", userController.GetUserById)
	router.GET("/user-res-api/user", userController.GetUsers)
	router.POST("/user-res-api/user", userController.UserInsert) // Sign In

	// Bookings Mapping
	router.GET("/user-res-api/booking/:id", bookingController.GetBookingById)
	router.GET("/user-res-api/booking", bookingController.GetBookings)
	router.POST("/user-res-api/booking", bookingController.InsertBooking)
	router.GET("/user-res-api/booking/user/:user_id", bookingController.GetBookingsByUserId)
	router.GET("/user-res-api/hotel/availability/:id/:start_date/:end_date", bookingController.GetAvailabilityByIdAndDate)
	router.DELETE("/user-res-api/booking/:booking_id", bookingController.DeleteBooking)

	// Hotels Mapping
	router.GET("/user-res-api/hotel/:id", hotelController.GetHotelById)
	router.GET("/user-res-api/hotel", hotelController.GetHotels)
	router.POST("/user-res-api/hotel", hotelController.InsertHotel)
	//router.PUT("/hotel/update", hotelController.UpdateHotel)
	router.DELETE("/user-res-api/hotel/delete/:idMongo", hotelController.DeleteHotel)

	// Login
	router.POST("/user-res-api/login", userController.Login)

	// Refresh Token
	router.POST("/user-res-api/refresh", userController.Refresh)

	log.Info("Finishing mappings configurations")
}
