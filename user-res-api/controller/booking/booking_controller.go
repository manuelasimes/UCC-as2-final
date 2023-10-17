package booking

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"user-res-api/dto"
	service "user-res-api/service"
	"net/http"
	"fmt" 
	"io/ioutil"
	"encoding/json"
	"strconv"
	"net/url"
	"strings"
)

// funcion para generar un token de amadeus cada vez que voy a hacer la consulta 
func GetAmadeustoken () (string) {

	fmt.Printf("entro al f d token")
	 // Define los datos que deseas enviar en el cuerpo de la solicitud.
	 data := url.Values{}
	 data.Set("grant_type", "client_credentials")
	 data.Set("client_id", "sCkSnG1piA4ApGUWTfWsYhj1MDGQZ8Ob")
	 data.Set("client_secret", "2Jrxf1ZBL46bfj6c")
 
	 // Realiza la solicitud POST a la API externa.
	 resp, err := http.Post("https://test.api.amadeus.com/v1/security/oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	 if err != nil {
		 fmt.Println("Error al hacer la solicitud:", err)
		 return ""
	 }
	 defer resp.Body.Close()
	  // Lee la respuesta de la API.
	  body, err := ioutil.ReadAll(resp.Body)
	  if err != nil {
		  fmt.Println("Error al leer la respuesta:", err)
		  return ""
	  }
	  // Parsea la respuesta JSON para obtener el token (asumiendo que la respuesta es JSON).
    // Si la respuesta es en otro formato, ajusta esto en consecuencia.
    var response map[string]interface{}
    if err := json.Unmarshal(body, &response); err != nil {
        return ""
    }
	token, ok := response["access_token"].(string)
    if !ok {
        return ""
    }
	fmt.Println("token:", token)
    return token

}


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
	// ahora con los datos del booking dto rellenamos una nueva estructura para la requesta  amadeus 
	// Serializa el objeto BookingDto a formato JSON
	id := bookingDto.HotelId
	fmt.Println("El id mysql del hotel es:", id)
	// necesito llmara a una funcion que me traiga el id amadeus del hotel con el id que ya tengo (tengo el id mysql) 
	// GetHotelById(id int) (dto.HotelDto, e.ApiError)
	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		// c.JSON(err.Status(), err)
		fmt.Println("No se encontro un hotel con ese id")
		return
	}

	idAm := hotelDto.IdAmadeus 
	
	startdatebooking := strconv.Itoa(bookingDto.StartDate)
	fechaConGuiones := startdatebooking
	startdateconguiones := fmt.Sprintf(
        "%s-%s-%s",
        fechaConGuiones[:4],
        fechaConGuiones[4:6],
        fechaConGuiones[6:8],
    )
	enddatebooking := strconv.Itoa(bookingDto.EndDate)
	fechaConGuiones2 := enddatebooking
	enddateconguiones := fmt.Sprintf(
        "%s-%s-%s",
        fechaConGuiones2[:4],
        fechaConGuiones2[4:6],
        fechaConGuiones2[6:8],
    )

	fmt.Println("fecha de ida", startdateconguiones)
	fmt.Println("fecha de vuelta", enddateconguiones)

	// antes de llamar la funcion insert del service deberiamos hacer el llamado a amadeus 
	// URL de la API externa
	
	// apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"

	//  // Agrega los parámetros a la URL
	//  queryParams := make(url.Values)
	//  queryParams.Add("hotelIds", idAm ) // Reemplaza con el valor deseado
	//  queryParams.Add("checkInDate", startdateconguiones) // Reemplaza con la fecha deseada
	//  queryParams.Add("checkOutDate", enddateconguiones) // Reemplaza con la fecha deseada
 
	//  apiUrl += "?" + queryParams.Encode()

	// Construye la URL manualmente
apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
apiUrl += "?hotelIds=" + idAm
apiUrl += "&checkInDate=" + startdateconguiones
apiUrl += "&checkOutDate=" + enddateconguiones

	fmt.Println(apiUrl)


	 // Crear una solicitud HTTP POST
	 solicitud, err := http.NewRequest("GET", apiUrl, nil)
	 if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	 }
	 
	// Agregar el encabezado de autorización Bearer con tu token
	token := GetAmadeustoken() // Reemplaza con tu token real

	solicitud.Header.Set("Authorization", "Bearer " + token)
	// solicitud.Header.Set("Content-Type", "application/json") // Especifica el tipo de contenido si es necesario
 
	fmt.Println(solicitud)
	// Realiza la solicitud HTTP
	cliente := &http.Client{}
	respuesta, err := cliente.Do(solicitud)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else  if err == nil {
	// Verifica el código de estado de la respuesta
	if respuesta.StatusCode != http.StatusOK {
    fmt.Printf("La solicitud a la API de Amadeus no fue exitosa. Código de estado: %d\n", respuesta.StatusCode)
    c.JSON(http.StatusInternalServerError, "La solicitud a la API de Amadeus no fue exitosa.")
    return
	}
		// Lee el cuerpo de la respuesta
		responseBody, err := ioutil.ReadAll(respuesta.Body)
		if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	   // Crear una estructura para deserializar el JSON de la respuesta
	   var responseStruct struct {
		Data []struct {
			Type                 string `json:"type"`
			ID                   string `json:"id"`
			ProviderConfirmationID string `json:"providerConfirmationId"`
		} `json:"data"`
    }

    // Decodificar el JSON y extraer el campo "id"
    if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
        fmt.Println("Error al decodificar el JSON de la respuesta:", err)
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }
		// Obtén el ID del hotel del primer elemento en "data"
		if len(responseStruct.Data) > 0 {
			// si el largo de la respuesta es mayor q cero es pq hay disponibilidad --> llamo al service 
    	
		bookingDto, err := service.BookingService.InsertBooking(bookingDto)
	 	// Error del Insert
		if err != nil {
			// c.JSON(err.Status(), err)
		return
		} 
	c.JSON(http.StatusCreated, bookingDto)
	} else {
		fmt.Println("No hay disponibilidad en esas fechas")
	}

	defer respuesta.Body.Close()
	 
}
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
