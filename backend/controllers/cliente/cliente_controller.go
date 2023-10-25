package clienteController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetClienteById(c *gin.Context) {
	log.Debug("ID de cliente para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClienteByUsername(c *gin.Context) {
	log.Debug("Cliente a cargar: " + c.Param("username"))

	username := c.Param("username")
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClienteByEmail(c *gin.Context) {
	log.Debug("Cliente a cargar: " + c.Param("email"))

	email := c.Param("email")
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func InsertCliente(c *gin.Context) {
	var clienteDto dto.ClienteDto
	err := c.BindJSON(&clienteDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	clienteDto, er := service.ClienteService.InsertCliente(clienteDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, clienteDto)
}

func GetReservaById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.ReservaDto

	reservaDto, err := service.ClienteService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservasById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservasDto dto.ReservasDto

	reservasDto, err := service.ClienteService.GetReservasById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func InsertReserva(c *gin.Context) {
	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := service.ClienteService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetHoteles(c *gin.Context) {
	var hotelesDto dto.HotelesDto
	hotelesDto, err := service.ClienteService.GetHoteles()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}

func GetHotelById(c *gin.Context) {
	log.Debug("ID de Hotel para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetDisponibilidad(c *gin.Context) {
	log.Debug("Disponibilidad de reservas para cargar: " + c.Param("id") + c.Param("AnioInicio") + c.Param("AnioFinal") + c.Param("MesInicio") + c.Param("MesFinal") + c.Param("DiaInicio") + c.Param("DiaFinal"))

	id, _ := strconv.Atoi(c.Param("id"))
	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))

	disponibilidad := service.ClienteService.GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	c.JSON(http.StatusOK, disponibilidad)
}