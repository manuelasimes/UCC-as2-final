package services

import (
	hotelDao "hotels-api/daos/hotel"
	"hotels-api/dtos"
	model "hotels-api/models"
	e "hotels-api/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotel(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	//UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotel(id string) (dto.HotelDto, e.ApiError) {
	var hotelDto dto.HotelDto

	// Llama a la función GetById del DAO
	existingHotel := hotelDao.GetById(id)

	// Verifica si el hotel existente fue encontrado
	if existingHotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("Hotel not found")
	}

	// Mapea los campos del modelo a la estructura DTO
	hotelDto.Id = existingHotel.Id.Hex() // Id se maneja como string
	hotelDto.Name = existingHotel.Name
	hotelDto.Description = existingHotel.Description
	hotelDto.Country = existingHotel.Country
	hotelDto.City = existingHotel.City
	hotelDto.Adress = existingHotel.Adress

	// Mapea las imágenes y amenidades
	hotelDto.Images = make([]*dto.ImageDto, len(existingHotel.Images))
	for i, img := range existingHotel.Images {
		hotelDto.Images[i] = &dto.ImageDto{
			Url: img.Url,
		}
	}

	hotelDto.Amenities = make([]*dto.AmenitieDto, len(existingHotel.Amenities))
	for i, amenity := range existingHotel.Amenities {
		hotelDto.Amenities[i] = &dto.AmenitieDto{
			Description: amenity.Description,
			Image:      amenity.Image,
		}
	}

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name

	hotel = hotelDao.Insert(hotel)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("Error in insert")
	}
	hotelDto.Id = hotel.Id.Hex()

	return hotelDto, nil
}

/* Error
func (s *hotelService) UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	// Obtén el hotel existente por ID
	existingHotel := hotelDao.GetById(id)

	// Verifica si el hotel existente fue encontrado
	if existingHotel.Id.Hex() == "000000000000000000000000" {
		return dto.HotelDto{}, e.NewBadRequestApiError("Hotel not found")
	}

	// Actualiza los campos del hotel existente con los valores proporcionados en updatedHotelDto
	existingHotel.Name = updatedHotelDto.Name
	// Agrega más actualizaciones aquí según tus necesidades

	// Realiza la actualización en la base de datos
	updatedHotel := hotelDao.Update(id, existingHotel)

	if updatedHotel.Id.Hex() == "000000000000000000000000" {
		return dto.HotelDto{}, e.NewBadRequestApiError("Error in update")
	}

	// Construye un HotelDto actualizado para la respuesta
	updatedHotelDto := dto.HotelDto{
		Id:   updatedHotel.Id.Hex(),
		Name: updatedHotel.Name,
		// Agrega más campos aquí según tus necesidades
	}

	return updatedHotelDto, nil
}
*/