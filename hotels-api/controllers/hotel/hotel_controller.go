package hotel

import (
	"hotels-api/dtos"
	service "hotels-api/services"
	"hotels-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	rateLimiter = make(chan bool, 3)
)

func Get(c *gin.Context) {

	id := c.Param("id")

	if len(rateLimiter) == cap(rateLimiter) {
		apiErr := errors.NewTooManyRequestsError("too many requests")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	rateLimiter <- true
	hotelDto, er := service.HotelService.GetHotel(id)
	<-rateLimiter

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

func Insert(c *gin.Context) {
	var hotelDto dtos.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func Update(c *gin.Context) {
    // Obtener el ID del hotel a actualizar desde los parámetros de la URL
    id := c.Param("id")

    // Parsear el objeto JSON del cuerpo de la solicitud
    var hotelDto dtos.HotelDto
    if err := c.BindJSON(&hotelDto); err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // Llamar al servicio para actualizar el hotel
    updatedHotelDto, err := service.HotelService.UpdateHotel(id, hotelDto)

    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

    c.JSON(http.StatusOK, updatedHotelDto)
}