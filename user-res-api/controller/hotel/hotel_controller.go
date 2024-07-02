package hotel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-res-api/dto"
	service "user-res-api/service"
	// "crypto/tls"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// bookingController "user-res-api/controller/booking"
)

func GetHotelById(c *gin.Context) {
	log.Println("entro al controller")
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func InsertHotel(c *gin.Context) {
	// esta funcion se llama cuando desde mongo se hace un post d eun hotel
	log.Println("entro al controller")
	var insertHotelDto dto.HotelPostDto
	err := c.BindJSON(&insertHotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	apiUrl := "https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city?cityCode=MIA&radius=5&radiusUnit=KM&hotelSource=ALL"
	// Crear una solicitud HTTP GET
	solicitud, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud dentro de insert hotel:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	token := service.BookingService.GetAmadeustoken()
	solicitud.Header.Set("Authorization", "Bearer "+token)
	// Realiza la solicitud HTTP
	cliente := &http.Client{}

	/*  // Custom HTTP client with TLS configuration to skip certificate verification.
	 customTransport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    cliente := &http.Client{
        Transport: customTransport,
    } */

	respuesta, err := cliente.Do(solicitud)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer respuesta.Body.Close()

	// Leer y manejar la respuesta de la API externa
	var response struct {
		Data []struct {
			HotelID string `json:"hotelId"`
		} `json:"data"`
	}

	// Decodificar la respuesta JSON
	decoder := json.NewDecoder(respuesta.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Error("Error al decodificar la respuesta JSON:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, hotel := range response.Data {
		fmt.Printf("Id amadeus: %s\n", hotel.HotelID)
		CanIUseID, err := service.HotelService.CheckHotelByIdAmadeus(hotel.HotelID)
		if err != nil {
			// manejo error
			fmt.Println("Ocurrio un error al verificar el uso de un id amadeus")
			fmt.Println(err)
		}
		if CanIUseID == true {
			hotelDto, er := service.HotelService.InsertHotel(insertHotelDto, hotel.HotelID)
			// Error del Insert
			if er != nil {
				c.JSON(er.Status(), er)
				return
			}
			c.JSON(http.StatusCreated, hotelDto)
			log.Println("ID del primer hotel:", hotel.HotelID)
			break // Se encontró el ID, sal del bucle
		}
	}

}
func DeleteHotel(c *gin.Context) {
	fmt.Printf("entro al controller")

	idMongo := c.Param("idMongo")

	err := service.HotelService.DeleteHotel(idMongo)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
