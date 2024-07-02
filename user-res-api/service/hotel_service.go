package service

import (
	// log "github.com/sirupsen/logrus"
	hotelClient "user-res-api/client/hotel"
	// userClient "user-res-api/client/user"

	"fmt"
	"user-res-api/dto"
	"user-res-api/model"
	e "user-res-api/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelPostDto, idAmadeus string) (dto.HotelDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	CheckHotelByIdAmadeus(id string) (bool, e.ApiError)
	// UpdateHotel(updateHotelDto dto.HandleHotelDto) (dto.HotelDto, e.ApiError)
	DeleteHotel(idmongo string) e.ApiError
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel no encontrado")
	}

	hotelDto.Id = hotel.Id
	hotelDto.HotelName = hotel.HotelName
	hotelDto.IdMongo = hotel.IdMongo
	hotelDto.IdAmadeus = hotel.IdAmadeus
	return hotelDto, nil

}
func (s *hotelService) CheckHotelByIdAmadeus(id string) (bool, e.ApiError) {

	if hotelClient.GetHotelByIdAmadeus(id) == true {
		return false, e.NewBadRequestApiError("Hotel ya en uso")
	}

	return true, nil
}

func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		id := hotel.Id

		hotelDto, _ = s.GetHotelById(id)

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelPostDto, idAmadeus string) (dto.HotelDto, e.ApiError) {
	fmt.Println("entro al service")
	var hotel model.Hotel
	var response dto.HotelDto

	hotel.HotelName = hotelDto.HotelName
	hotel.IdAmadeus = idAmadeus
	hotel.IdMongo = hotelDto.IdMongo

	hotel = hotelClient.InsertHotel(hotel)

	if hotel.Id == 0 {
		return response, e.NewBadRequestApiError("Error al insertar hotel")
	}

	hotelDto.Id = hotel.Id

	return response, nil
}

func (s *hotelService) DeleteHotel(idmongo string) e.ApiError {
	// Llama al cliente de hotel para realizar la eliminación
	err := hotelClient.DeleteHotel(idmongo)

	if err != nil {
		return e.NewBadRequestApiError("Error al eliminar hotel")
	}

	return nil
}
