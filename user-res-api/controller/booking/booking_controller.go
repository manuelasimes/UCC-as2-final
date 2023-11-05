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
	// e "user-res-api/utils/errors"
)

// funcion para generar un token de amadeus cada vez que voy a hacer la consulta 
// func GetAmadeustoken () (string) {

// 	fmt.Printf("entro al f d token")
// 	 // Define los datos que deseas enviar en el cuerpo de la solicitud.
// 	 data := url.Values{}
// 	 data.Set("grant_type", "client_credentials")
// 	 data.Set("client_id", "sCkSnG1piA4ApGUWTfWsYhj1MDGQZ8Ob")
// 	 data.Set("client_secret", "2Jrxf1ZBL46bfj6c")
 
// 	 // Realiza la solicitud POST a la API externa.
// 	 resp, err := http.Post("https://test.api.amadeus.com/v1/security/oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
// 	 if err != nil {
// 		 fmt.Println("Error al hacer la solicitud:", err)
// 		 return ""
// 	 }
// 	 defer resp.Body.Close()
// 	  // Lee la respuesta de la API.
// 	  body, err := ioutil.ReadAll(resp.Body)
// 	  if err != nil {
// 		  fmt.Println("Error al leer la respuesta:", err)
// 		  return ""
// 	  }
// 	  // Parsea la respuesta JSON para obtener el token (asumiendo que la respuesta es JSON).
//     // Si la respuesta es en otro formato, ajusta esto en consecuencia.
//     var response map[string]interface{}
//     if err := json.Unmarshal(body, &response); err != nil {
//         return ""
//     }
// 	token, ok := response["access_token"].(string)
//     if !ok {
//         return ""
//     }
// 	fmt.Println("token:", token)
//     return token

// }


// // funcion por separado que me calcula disponibilidad 
// func Availability (startdateconguiones string, enddateconguiones string, idAm string) (bool) {
// fmt.Println("entro a availability")
// apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
// apiUrl += "?hotelIds=" + idAm
// apiUrl += "&checkInDate=" + startdateconguiones
// apiUrl += "&checkOutDate=" + enddateconguiones

// 	fmt.Println(apiUrl)


// 	 // Crear una solicitud HTTP POST
// 	 solicitud, err := http.NewRequest("GET", apiUrl, nil)
// 	 if err != nil {
// 		fmt.Println("Error al crear la solicitud:", err)
// 		//c.JSON(http.StatusInternalServerError, err.Error())
// 		return false
// 	 }
	 
// 	// Agregar el encabezado de autorización Bearer con tu token
// 	token := GetAmadeustoken() // Reemplaza con tu token real

// 	solicitud.Header.Set("Authorization", "Bearer " + token)
// 	// solicitud.Header.Set("Content-Type", "application/json") // Especifica el tipo de contenido si es necesario
 
// 	fmt.Println(solicitud)
// 	// Realiza la solicitud HTTP
// 	cliente := &http.Client{}
// 	respuesta, err := cliente.Do(solicitud)
// 	if err != nil {
// 		fmt.Println("Error al realizar la solicitud:", err)
// 		//c.JSON(http.StatusInternalServerError, err.Error())
// 		return false
// 	} 
// 	fmt.Println("La solicitud de la api fue exitosa ")
// 	// Verifica el código de estado de la respuesta
// 	if respuesta.StatusCode != http.StatusOK {
//     	fmt.Printf("La solicitud a la API de Amadeus no fue exitosa. Código de estado: %d\n", respuesta.StatusCode)
//     	//c.JSON(http.StatusInternalServerError, "La solicitud a la API de Amadeus no fue exitosa.")
//    		return false
// 	}
// 		// Lee el cuerpo de la respuesta
// 		responseBody, err := ioutil.ReadAll(respuesta.Body)
// 		if err != nil {
// 		fmt.Println("Error al leer la respuesta:", err)
// 		//c.JSON(http.StatusInternalServerError, err.Error())
// 		return false
// 	}
// 	   // Crear una estructura para deserializar el JSON de la respuesta
// 	   var responseStruct struct {
// 		Data []struct {
// 			Type                 string `json:"type"`
// 			ID                   string `json:"id"`
// 			ProviderConfirmationID string `json:"providerConfirmationId"`
// 		} `json:"data"`
//     }

//     // Decodificar el JSON y extraer el campo "id"
//     if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
//         fmt.Println("Error al decodificar el JSON de la respuesta:", err)
//         //c.JSON(http.StatusInternalServerError, err.Error())
//         return false
//     }
// 		// Obtén el ID del hotel del primer elemento en "data"
// 		if len(responseStruct.Data) > 0 {
// 			// si el largo de la respuesta es mayor q cero es pq hay disponibilidad --> llamo al service 
//     	fmt.Println("Segun amadeus hay disp")
// 		return true
// 	 	// Error del Insert
// 		// if err != nil {
// 		// 	// c.JSON(err.Status(), err)
// 		// return false
// 		// } 
// 	//c.JSON(http.StatusCreated, bookingDto)
// 	} 
// 		fmt.Println("No hay disponibilidad en esas fechas, controller")
// 		defer respuesta.Body.Close()
// 		return false 
	

	
// }


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

	// startdate := strconv.Itoa(startDate)
	// fechaConGuiones := startdate
	// startdateconguiones := fmt.Sprintf(
    //     "%s-%s-%s",
    //     fechaConGuiones[:4],
    //     fechaConGuiones[4:6],
    //     fechaConGuiones[6:8],
    // )
	
	// enddate := strconv.Itoa(endDate)
	// fechaConGuiones2 := enddate
	// enddateconguiones := fmt.Sprintf(
    //     "%s-%s-%s",
    //     fechaConGuiones2[:4],
    //     fechaConGuiones2[4:6],
    //     fechaConGuiones2[6:8],
    // )


	// // var request dto.CheckRoomDto
	// // request.StartDate = startDate
	// // request.EndDate = endDate
	// // IsAvailable, err := service.BookingService.GetBookingByHotelIdAndDate(request,id_)
	

	// // antes de hacer eso deberiamos ver si ya esta en la cache 
	// key := id + strconv.Itoa(startDate)
	// cacheDTO, err := cache.Get(key)

	// if err == nil { // does a hit 
	// 	fmt.Println("hit de cache!")
	// 	// creo q si lo encuentra ya quiere decir q no esta disponible ANTES 
	// 	c.JSON(http.StatusOK, cacheDTO)
	// 	return 
	// 	// return cacheDTO, nil 
	
	// }
	// // es un miss --> mem principal 
	// IsAvailable := Availability (startdateconguiones, enddateconguiones, idAm)
	// fmt.Println("miss de cache!")
	// if IsAvailable == true {
	// 	responseDto.OkToBook = true 
	// } else if IsAvailable == false {
	// 	responseDto.OkToBook = false 
	// }

	// // save in cache
	// availability, _ := json.Marshal(responseDto) 
	// cache.Set(key, availability, 10)
	// fmt.Println("Saved in cache!")
	// if apiError != nil {
	// 	status := apiError.Status()
    //     c.JSON(status, gin.H{"error": apiError.Message()})
    //     return
    // }

	
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
	// necesito llmara a una funcion que me traiga el id amadeus del hotel con el id que ya tengo (tengo el id mysql) 
	// GetHotelById(id int) (dto.HotelDto, e.ApiError)
	
	// var hotel model.Hotel = hotelClient.GetHotelByIdMongo(id)
	// idAm := hotel.IdAmadeus 
	// idMySQL := hotel.Id 

	// var bookingDto dto.BookingDto

	// bookingDto.Id = bookingPDto.Id
	// bookingDto.UserId = bookingPDto.UserId
	// bookingDto.HotelId = idMySQL
	// bookingDto.StartDate = bookingPDto.StartDate
	// bookingDto.EndDate = bookingPDto.EndDate
	

	
	// startdatebooking := strconv.Itoa(bookingDto.StartDate)
	// fechaConGuiones := startdatebooking
	// startdateconguiones := fmt.Sprintf(
    //     "%s-%s-%s",
    //     fechaConGuiones[:4],
    //     fechaConGuiones[4:6],
    //     fechaConGuiones[6:8],
    // )
	// enddatebooking := strconv.Itoa(bookingDto.EndDate)
	// fechaConGuiones2 := enddatebooking
	// enddateconguiones := fmt.Sprintf(
    //     "%s-%s-%s",
    //     fechaConGuiones2[:4],
    //     fechaConGuiones2[4:6],
    //     fechaConGuiones2[6:8],
    // )

	// fmt.Println("fecha de ida", startdateconguiones)
	// fmt.Println("fecha de vuelta", enddateconguiones)
	// si mejor llamo a la funcion get availability del mismo controller
	// Available := Availability (startdateconguiones, enddateconguiones, idAm)

	// if Available == true {
	// 	bookingDto, err := service.BookingService.InsertBooking(bookingPDto)
	//  	// Error del Insert
	// 	if err != nil {
	// 		// c.JSON(err.Status(), err)
	// 	return
	// 	} 
	// 	c.JSON(http.StatusCreated, bookingDto)
	// } else if Available == false {
	// 	fmt.Println("No hay disponibilidad en esas fechas, controller")
	// }

	


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
