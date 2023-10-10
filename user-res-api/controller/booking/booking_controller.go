package booking

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"user-res-api/dto"
	service "user-res-api/service"
	"net/http"
	"strconv"	
	"fmt" 
	"bytes"
    // "io/ioutil"
	"encoding/json"

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

func GetAvailabilityByIdAndDate(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	
	log.Debug("Booking startDate to load: " + c.Param("start_date"))

	startDate, _ := strconv.Atoi(c.Param("start_date"))
	
	log.Debug("Booking endDate to load: " + c.Param("end_date"))

	endDate, _ := strconv.Atoi(c.Param("end_date"))
	
	var request dto.CheckRoomDto
	request.StartDate = startDate
	request.EndDate = endDate
	IsAvailable, err := service.BookingService.GetBookingByHotelIdAndDate(request,id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, IsAvailable)
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
var bookingDto dto.BookingDto
err := c.BindJSON(&bookingDto)

// Error Parsing json param
 	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	    // Serializa el objeto BookingDto a formato JSON
	requestData, err := json.Marshal(bookingDto)
	if err != nil {
		fmt.Println("Error al serializar BookingDto a JSON:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	// antes de llamar la funcion insert del service deberiamos hacer el llamado a amadeus 
	// URL de la API externa
	
	 apiUrl := "https://test.api.amadeus.com/v1/booking/hotel-bookings"
	 // Crear una solicitud HTTP POST
	 solicitud, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestData))
	 if err != nil {
		 fmt.Println("Error al crear la solicitud:", err)
		 c.JSON(http.StatusInternalServerError, err.Error())

		 return
	 }
 
	 // Agregar el encabezado de autorización Bearer con tu token
	 token := "kGQbUjrndsKlKrpBc6Rqpg5lFZ2G" // Reemplaza con tu token real
	 solicitud.Header.Set("Authorization", "Bearer " + token)
	 // solicitud.Header.Set("Content-Type", "application/json") // Especifica el tipo de contenido si es necesario
 
	 // Realiza la solicitud HTTP
	 cliente := &http.Client{}
	 respuesta, err := cliente.Do(solicitud)
	 if err != nil {
		 fmt.Println("Error al realizar la solicitud:", err)
		 c.JSON(http.StatusInternalServerError, err.Error())
		 return
	 } else  if err == nil {
		bookingDto, err := service.BookingService.InsertBooking(bookingDto)
	 	// Error del Insert
		if err != nil {
		c.JSON(err.Status(), err)
		return
	} 
	c.JSON(http.StatusCreated, bookingDto)
	 }

	 defer respuesta.Body.Close()
	 
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
