package controller

import (
	"hotels-api/dtos"
	service "hotels-api/services"
	//"hotels-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Get(c *gin.Context) {

	id := c.Param("id")

	
	hotelDto, err := service.HotelService.GetHotel(id)
	

	// Error del Insert
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	origin := c.Request.Header.Get("Origin")

	if origin == "http://localhost:3000" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	}
	c.JSON(http.StatusOK, hotelDto)
}

func Insert(c *gin.Context) {
	var hotelDto dto.HotelDto
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
    var hotelDto dto.HotelDto

    if err := c.BindJSON(&hotelDto); err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // Llamar al servicio para actualizar el hotel
    updatedHotelDto, err := service.HotelService.UpdateHotel(id, hotelDto) // error

    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

	
    c.JSON(http.StatusOK, updatedHotelDto)
}
