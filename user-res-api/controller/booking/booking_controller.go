package booking

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"user-res-api/dto"
	service "user-res-api/service"
	"net/http"
	"fmt" 
	"user-res-api/model"
	
	"strconv"


	hotelClient "user-res-api/client/hotel"

)



func GetBookingById(c *gin.Context) {
	log.Debug("Booking id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var bookingDto dto.BookingDetailDto

	bookingDto, err := service.BookingService.GetBookingById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, bookingDto)
}

func GetAvailabilityByIdAndDate(c *gin.Context)  {
	// var apiError e.ApiError

	log.Debug("Hotel id to load: " + c.Param("id"))
	id := c.Param("id")
	
	log.Debug("Booking startDate to load: " + c.Param("start_date"))

	startDate, _ := strconv.Atoi(c.Param("start_date"))
	
	log.Debug("Booking endDate to load: " + c.Param("end_date"))

	endDate, _ := strconv.Atoi(c.Param("end_date"))
	
	var hotel model.Hotel = hotelClient.GetHotelByIdMongo(id)
	idAm := hotel.IdAmadeus



	var responseDto dto.Availability // la respuesta q vamos a devolver 
	responseDto, err := service.BookingService.GetAvailabilityByIdAndDate(idAm, startDate, endDate)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, responseDto)

	
	
}

func GetBookings(c *gin.Context) {
	var bookingsDto dto.BookingsDetailDto
	bookingsDto, err := service.BookingService.GetBookings()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}

func InsertBooking(c *gin.Context) {
	fmt.Println("Entro al controller")


	var bookingPDto dto.BookingPostDto

	err := c.BindJSON(&bookingPDto)
	
	// Error Parsing json param
	if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
	}

	id := bookingPDto.HotelId
	fmt.Println("El id mysql del hotel es:", id)

	bookingDto, er := service.BookingService.InsertBooking(bookingPDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookingDto)
	


}



func GetBookingsByUserId(c *gin.Context) {
	log.Debug("user id to load: " + c.Param("user_id"))

	id, _ := strconv.Atoi(c.Param("user_id"))

	var bookingsDto dto.BookingsDetailDto

	bookingsDto, err := service.BookingService.GetBookingsByUserId(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}

func DeleteBooking(c *gin.Context) {

	bookingIdStr := c.Param("booking_id")
    bookingId, err := strconv.Atoi(bookingIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

	err = service.BookingService.DeleteBooking(bookingId)
	
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return

	}

	c.JSON(http.StatusOK, gin.H{"delete_confirmation": true})

}
