package dto

type HotelDto struct {
	ID       	int    	`json:"id"`
	Nombre   	string 	`json:"nombre"`
	Descripcion	string 	`json:"descripcion"`
	Email    	string 	`json:"email"`
	Image    	string 	`json:"image"`
	Cant_Hab 	int    	`json:"cant_hab"`

	TelefonosDto TelefonosDto `json:"telefonos,omitempty"`
}


type HotelesDto []HotelDto