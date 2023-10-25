package dto

type TelefonoDto struct {
	ID      int    	`json:"id"`
	Codigo  int 	`json:"codigo"`
	Numero  int 	`json:"numero"`
	HotelID int    	`json:"hotel_id,omitempty"`
}

type TelefonosDto []TelefonoDto