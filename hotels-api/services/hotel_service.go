package services

import (
	hotelDao "hotels-api/daos/hotel"
	"hotels-api/dtos"
	queue "hotels-api/utils/queue"
	model "hotels-api/models"
	e "hotels-api/utils/errors"
	// "strconv"
	"fmt"
	"net/http"
	"hotels-api/config"
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotel(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)

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

	hotelDto.Images = existingHotel.Images
	hotelDto.Amenities = existingHotel.Amenities

	/* // Mapea las imágenes y amenidades
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
	} */

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.Country = hotelDto.Country
	hotel.City = hotelDto.City
	hotel.Adress = hotelDto.Adress

	hotel.Images = hotelDto.Images
	hotel.Amenities = hotelDto.Amenities

	/* hotel.Images = make([]model.Image, len(hotelDto.Images))
	hotel.Amenities = make([]model.Amenitie, len(hotelDto.Amenities))

	for i, imgDto := range hotelDto.Images {
		hotel.Images[i] = model.Image{
			Url: imgDto.Url,
		}
	}

	for i, amenityDto := range hotelDto.Amenities {
		hotel.Amenities[i] = model.Amenitie{
			Description: amenityDto.Description,
			Image:      amenityDto.Image,
		} 
	} */


	hotel = hotelDao.Insert(hotel)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("Error in insert")
	}
	hotelDto.Id = hotel.Id.Hex()

	// Assuming hotel.Id is of type primitive.ObjectID
	idHexString := hotel.Id.Hex()
    // Now you can pass idInt to the SendMessage function
	
	var postIdDto dto.PostID

	postIdDto.IdMongo = idHexString
	postIdDto.HotelName = hotelDto.Name

	// jsonData, err := json.Marshal(postIdDto)

	// if err != nil {
		// log.Debug(err)
		// return hotelDto, e.NewBadRequestApiError("Marshal failed")
	// }

	var buf bytes.Buffer
    err := json.NewEncoder(&buf).Encode(postIdDto)
    if err != nil {
        log.Fatal(err)
    }


	url := fmt.Sprintf("http://%s:%d/user-res-api/hotel", config.USERAPIHOST, config.USERAPIPORT)

	resp, err := http.Post(url, "application/json", &buf)

	if resp != nil {
		if err != nil {
			log.Debug(err)
			return hotelDto, e.NewBadRequestApiError("user-res-api failed")
		}
	}

	if err != nil {
		log.Debug(err)
		return hotelDto, e.NewBadRequestApiError("user-res-api failed2")
	}

	queue.SendMessage(idHexString, "INSERT")

	return hotelDto, nil
}


func (s *hotelService) UpdateHotel(id string, updatedHotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	// Obtén el hotel existente por ID
	existingHotel := hotelDao.GetById(id)

	// Verifica si el hotel existente fue encontrado
	if existingHotel.Id.Hex() == "000000000000000000000000" {
		return dto.HotelDto{}, e.NewBadRequestApiError("Hotel not found")
	}

	// Actualiza los campos del hotel existente con los valores proporcionados en updatedHotelDto
	existingHotel.Name = updatedHotelDto.Name
	existingHotel.Description = updatedHotelDto.Description
	existingHotel.Country = updatedHotelDto.Country
	existingHotel.City = updatedHotelDto.City
	existingHotel.Adress = updatedHotelDto.Adress

	existingHotel.Images = updatedHotelDto.Images
	existingHotel.Amenities = updatedHotelDto.Amenities
	
	/* existingHotel.Images = make([]model.Image, len(updatedHotelDto.Images))
	existingHotel.Amenities = make([]model.Amenitie, len(updatedHotelDto.Amenities))

	for i, imgDto := range updatedHotelDto.Images {
		existingHotel.Images[i] = model.Image{
			Url: imgDto.Url,
		}
	}

	for i, amenityDto := range updatedHotelDto.Amenities {
		existingHotel.Amenities[i] = model.Amenitie{
			Description: amenityDto.Description,
			Image:      amenityDto.Image,
		}
	} */
	
	// Realiza la actualización en la base de datos
	err := hotelDao.Update(id, existingHotel)

	if err != nil {
    return dto.HotelDto{}, e.NewBadRequestApiError("Error in update")
	}

	// Construye un HotelDto actualizado para la respuesta
	updatedHotelDto.Name = existingHotel.Name 
	updatedHotelDto.Description = existingHotel.Description
	updatedHotelDto.Country = existingHotel.Country
	updatedHotelDto.City = existingHotel.City
	updatedHotelDto.Adress = existingHotel.Adress

	// Assuming hotel.Id is of type primitive.ObjectID
	idHexString := existingHotel.Id.Hex()

    // Now you can pass idInt to the SendMessage function
    queue.SendMessage(idHexString, "UPDATE")
	
	return updatedHotelDto, nil

}


