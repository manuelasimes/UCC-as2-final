package services

import (
	hotelDao "hotels-api/daos/hotel"
	"hotels-api/dtos"
	model "hotels-api/models"
	e "hotels-api/utils/errors"
	"time"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotel(id string) (dtos.HotelDto, e.ApiError)

	InsertHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotel(id string) (dtos.HotelDto, e.ApiError) {

	time.Sleep(15 * time.Second)

	

	var hotel model.Hotel = hotelDao.GetById(id)
	var hotelDto dtos.HotelDto

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}
	hotelDto.Name = hotel.Name
	hotelDto.Id = hotel.Id.Hex()


	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Name = hotelDto.Name

	hotel = hotelDao.Insert(hotel)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("error in insert")
	}
	hotelDto.Id = hotel.Id.Hex()

	return hotelDto, nil
}
