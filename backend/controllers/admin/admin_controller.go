package adminController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAdminById(c *gin.Context) {
	log.Debug("ID de administrador para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdminByUsername(c *gin.Context) {
	log.Debug("Admin a cargar: " + c.Param("username"))

	username := c.Param("username")
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdminByEmail(c *gin.Context) {
	log.Debug("Admin a cargar: " + c.Param("email"))

	email := c.Param("email")
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdmins(c *gin.Context) {
	var adminsDto dto.AdminsDto
	adminsDto, err := service.AdminService.GetAdmins()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, adminsDto)
}

func InsertAdmin(c *gin.Context) {
	var adminDto dto.AdminDto
	err := c.BindJSON(&adminDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	adminDto, er := service.AdminService.InsertAdmin(adminDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, adminDto)
}

func GetClienteById(c *gin.Context) {
	log.Debug("ID de Cliente para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var clienteDto dto.ClienteDto

	clienteDto, err := service.AdminService.GetClienteById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClientes(c *gin.Context) {
	var clientesDto dto.ClientesDto
	clientesDto, err := service.AdminService.GetClientes()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, clientesDto)
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

func GetHotelByEmail(c *gin.Context) {
	log.Debug("Email de Hotel para cargar: " + c.Param("email"))

	email := c.Param("email")
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotelByNombre(c *gin.Context) {
	log.Debug("Nombre de Hotel para cargar: " + c.Param("nombre"))

	nombre := c.Param("nombre")
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelByNombre(nombre)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHoteles(c *gin.Context) {
	var hotelesDto dto.HotelesDto
	hotelesDto, err := service.AdminService.GetHoteles()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.AdminService.InsertHotel(hotelDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func AddTelefono(c *gin.Context) {

	log.Debug("Agregar Teléfono al hotel: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var telefonoDto dto.TelefonoDto
	err := c.BindJSON(&telefonoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	telefonoDto.HotelID = id

	var hotelDto dto.HotelDto

	hotelDto, er := service.AdminService.AddTelefono(telefonoDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := service.AdminService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}