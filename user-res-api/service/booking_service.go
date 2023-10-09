package service

import (
	
	"fmt"
	json "github.com/json-iterator/go"
	bookingClient "user-res-api/client/booking"
	// userClient "user-res-api/client/user"
	"user-res-api/dto"
	"user-res-api/model"
	cache "user-res-api/cache"
	// "time"
	e "user-res-api/utils/errors"
	"strconv"
)

type bookingService struct{}

type bookingServiceInterface interface {
	GetBookingById(id int) (dto.BookingDetailDto, e.ApiError)
	GetBookings() (dto.BookingsDetailDto, e.ApiError)
	// InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
	GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, e.ApiError)
	GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError)
	GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError)
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

		bookingDto, _ = s.GetBookingById(id)

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}


// recibe un request del tipo CheckRoomDto que trae start date y end date y el id del hotel
// devuelve un responseDto el cual me dice si esta okey o no (disponible o no)
func (s *bookingService) GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, e.ApiError) {
	
	// var cacheDTO dto.Availability
	

	var IsAvailable bool 
	

	startDate := request.StartDate //del dto de parametro saca el start y end date 
	endDate := request.EndDate

	
	var responseDto dto.Availability // la respuesta q vamos a devolver 

	// debo encontrar la forma de chequear en MONGO com ver si este id existe --> por ahora sacamos 
	// if hotel.Id == 0 {
	// 	return responseDto, e.NewBadRequestApiError("El hotel no se encuentra en el sistema")
	// } 

	for i := startDate; i < endDate; i = i + 1 { // i va a ser cada dia de los q vamos a chequear 
		key := strconv.Itoa(idHotel) + strconv.Itoa(startDate)
		cacheDTO, err := cache.Get(key)

		if err == nil { // does a hit 
			fmt.Println("hit de cache!")
			// creo q si lo encuentra ya quiere decir q no esta disponible 
			
			return cacheDTO, nil 
		
		}

		// es un miss --> mem principal 
		fmt.Println("miss de cache!")
		IsAvailable = bookingClient.GetAvailabilityByIdAndDate(idHotel, i) // me devuelve si existe reserva en ese hotel en ese dia 
		if IsAvailable == true {
			responseDto.OkToBook = false 
		} else if IsAvailable == false {
			responseDto.OkToBook = true 
		}


		// save in cache 
		availability, _ := json.Marshal(responseDto)
		cache.Set(key, availability)
		fmt.Println("Saved in cache!")
		return responseDto, nil
		// mucho x ver --> como x ej si se cancela reserva! 
		
		

	}

	responseDto.OkToBook = true

	return responseDto, nil

}

func (s *bookingService) GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError) {
	var booking model.Booking = bookingClient.GetBookingByUserId(id)
	var bookingDto dto.BookingDetailDto

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("Booking not found")
	}

	/*	bookingDto.StartDay = booking.StartDay
		bookingDto.StartMonth = booking.StartMonth
		bookingDto.StartYear = booking.StartYear
		bookingDto.EndDay = booking.EndDay
	*/
	bookingDto.Id = booking.Id
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.UserName
	bookingDto.HotelId = booking.HotelId


	return bookingDto, nil
}

func (s *bookingService) GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError) {

	//var bookings model.Bookings = bookingClient.GetBookingsByUserId(id)
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
// func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {

	
// }