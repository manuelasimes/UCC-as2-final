package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	cache "user-res-api/cache"
	bookingClient "user-res-api/client/booking"
	hotelClient "user-res-api/client/hotel"
	userClient "user-res-api/client/user"
	"user-res-api/dto"
	"user-res-api/model"
	e "user-res-api/utils/errors"
)

type bookingService struct{}

type bookingServiceInterface interface {
	GetBookingById(id int) (dto.BookingDetailDto, e.ApiError)
	GetBookings() (dto.BookingsDetailDto, e.ApiError)
	InsertBooking(bookingPDto dto.BookingPostDto) (dto.BookingDto, e.ApiError)
	GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError)
	GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError)
	GetAmadeustoken() string
	GetAvailabilityByIdAndDate(idAm string, startDate int, endDate int) (dto.Availability, e.ApiError)
	Availability(startdateconguiones string, enddateconguiones string, idAm string) bool
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) GetBookingById(id int) (dto.BookingDetailDto, e.ApiError) {
	var booking model.Booking = bookingClient.GetBookingById(id)
	var bookingDto dto.BookingDetailDto

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("Booking not found")
	}
	bookingDto.Id = booking.Id
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.UserName
	bookingDto.HotelId = booking.HotelId

	return bookingDto, nil
}

func (s *bookingService) GetBookings() (dto.BookingsDetailDto, e.ApiError) {

	var bookings model.Bookings = bookingClient.GetBookings()
	var bookingsDto dto.BookingsDetailDto

	for _, booking := range bookings {
		var bookingDto dto.BookingDetailDto
		id := booking.Id
		idHotel := booking.HotelId

		var hotel model.Hotel = hotelClient.GetHotelById(idHotel)

		bookingDto, _ = s.GetBookingById(id)
		bookingDto.HotelName = hotel.HotelName

		bookingsDto = append(bookingsDto, bookingDto)

	}

	return bookingsDto, nil
}

func (s *bookingService) GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError) {
	var booking model.Booking = bookingClient.GetBookingByUserId(id)
	var bookingDto dto.BookingDetailDto

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("Booking not found")
	}

	bookingDto.Id = booking.Id
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.UserName
	bookingDto.HotelId = booking.HotelId

	return bookingDto, nil
}

func (s *bookingService) GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError) {

	var bookings model.Bookings = bookingClient.GetBookings()
	var bookingsDto dto.BookingsDetailDto

	for _, booking := range bookings {
		var bookingDto dto.BookingDetailDto

		if booking.UserId == id {

			bookingDto, _ = s.GetBookingById(booking.Id)
			bookingsDto = append(bookingsDto, bookingDto)

		}
	}

	return bookingsDto, nil
}

func (s *bookingService) InsertBooking(bookingPDto dto.BookingPostDto) (dto.BookingDto, e.ApiError) {
	var booking model.Booking
	var bookingDto dto.BookingDto

	var hotel model.Hotel = hotelClient.GetHotelByIdMongo(bookingPDto.HotelId)
	idAm := hotel.IdAmadeus
	idMySQL := hotel.Id

	bookingDto.Id = bookingPDto.Id
	bookingDto.UserId = bookingPDto.UserId
	bookingDto.HotelId = idMySQL
	bookingDto.StartDate = bookingPDto.StartDate
	bookingDto.EndDate = bookingPDto.EndDate
	if userClient.CheckUserById(bookingPDto.UserId) == false {
		return bookingDto, e.NewBadRequestApiError("El usuario no esta registrado en el sistema")
	}

	var responseDto dto.Availability

	startDate := bookingPDto.StartDate
	endDate := bookingPDto.EndDate

	responseDto, err := s.GetAvailabilityByIdAndDate(idAm, startDate, endDate)

	if err != nil {
		return bookingDto, e.NewBadRequestApiError("ERROR VIENDO DISPONIBILIDAD")
	}

	if responseDto.OkToBook == true {

		booking.StartDate = bookingDto.StartDate
		booking.EndDate = bookingDto.EndDate
		booking.UserId = bookingDto.UserId
		booking.HotelId = bookingDto.HotelId

		booking = bookingClient.InsertBooking(booking)

		bookingDto.Id = booking.Id

	} else if responseDto.OkToBook == false {
		fmt.Println("no hay disponibilidad, insert del service")
		return bookingDto, e.NewBadRequestApiError("No hay disponibilidad")
	}
	return bookingDto, nil
}

// funcion para generar un token de amadeus cada vez que voy a hacer la consulta
func (s *bookingService) GetAmadeustoken() string {

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

// funcion por separado que me calcula disponibilidad
func (s *bookingService) Availability(startdateconguiones string, enddateconguiones string, idAm string) bool {
	fmt.Println("entro a availability")
	apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
	apiUrl += "?hotelIds=" + idAm
	apiUrl += "&checkInDate=" + startdateconguiones
	apiUrl += "&checkOutDate=" + enddateconguiones

	fmt.Println(apiUrl)

	// Crear una solicitud HTTP
	solicitud, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("ERROR CREANDO SOLICITUD")
		return false
	}

	// Agregar el encabezado de autorización Bearer con tu token
	token := s.GetAmadeustoken() // Reemplaza con tu token real

	solicitud.Header.Set("Authorization", "Bearer "+token)

	fmt.Println(solicitud)
	// Realiza la solicitud HTTP
	cliente := &http.Client{}
	respuesta, err := cliente.Do(solicitud)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return false

	}

	// Verifica el código de estado de la respuesta
	if respuesta.StatusCode != http.StatusOK {
		fmt.Printf("La solicitud a la API de Amadeus no fue exitosa. Código de estado: %d\n", respuesta.StatusCode)
		return false
	}
	// Lee el cuerpo de la respuesta
	responseBody, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return false
	}
	// Crear una estructura para deserializar el JSON de la respuesta
	var responseStruct struct {
		Data []struct {
			Type                   string `json:"type"`
			ID                     string `json:"id"`
			ProviderConfirmationID string `json:"providerConfirmationId"`
		} `json:"data"`
	}

	// Decodificar el JSON y extraer el campo "id"
	if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
		fmt.Println("Error al decodificar el JSON de la respuesta:", err)
		return false
	}
	// Obtén el ID del hotel del primer elemento en "data"
	if len(responseStruct.Data) > 0 {
		// si el largo de la respuesta es mayor q cero es pq hay disponibilidad --> llamo al service
		fmt.Println("Amadeus nos dice que hay disponibilidad")
		return true

	}
	fmt.Println("No hay disponibilidad en esas fechas")
	defer respuesta.Body.Close()
	return false

}

func (s *bookingService) GetAvailabilityByIdAndDate(idAm string, startDate int, endDate int) (dto.Availability, e.ApiError) {
	var responseDto dto.Availability // la respuesta q vamos a devolver

	fmt.Println("Start date before booking service: %s", startDate)
	fmt.Println("Start date before booking service: %s", endDate)

	startdate := strconv.Itoa(startDate)

	fechaConGuiones := startdate
	startdateconguiones := fmt.Sprintf(
		"%s-%s-%s",
		fechaConGuiones[:4],
		fechaConGuiones[4:6],
		fechaConGuiones[6:8],
	)

	enddate := strconv.Itoa(endDate)
	fechaConGuiones2 := enddate
	enddateconguiones := fmt.Sprintf(
		"%s-%s-%s",
		fechaConGuiones2[:4],
		fechaConGuiones2[4:6],
		fechaConGuiones2[6:8],
	)

	fmt.Println("Start date inside booking service: %s", startdateconguiones)
	fmt.Println("Start date inside booking service: %s", enddateconguiones)

	// antes de hacer eso deberiamos ver si ya esta en la cache
	key := idAm + strconv.Itoa(startDate) + strconv.Itoa(endDate) // la key sera e id del hotel junto con las fechas que se quiere saber disponibilidad
	cacheDTO, err := cache.Get(key)

	if err == nil { // does a hit
		fmt.Println("hit de cache!")

		return cacheDTO, nil

	}
	// es un miss
	IsAvailable := s.Availability(startdateconguiones, enddateconguiones, idAm)
	fmt.Println("miss de cache!")

	if IsAvailable == true {
		responseDto.OkToBook = true
	} else if IsAvailable == false {
		responseDto.OkToBook = false
	}

	// save in cache
	availability, errorMarshal := json.Marshal(responseDto)

	if errorMarshal != nil {
		fmt.Printf("Error marshalling response to JSON: %v\n", errorMarshal)
		return dto.Availability{}, e.NewBadRequestApiError("Error marshalling")
	}

	cache.Set(key, availability, 10)
	fmt.Println("Saved in cache!")

	return responseDto, nil
}
