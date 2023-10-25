package dto

type ReservaDto struct {
	ID             int    `json:"id"`

	HotelID        int    `json:"hotel_id"`

	ClienteID      int    `json:"cliente_id"`

	AnioInicio     int    `json:"anio_inicio"`
	AnioFinal      int    `json:"anio_final"`
	MesInicio      int    `json:"mes_inicio"`
	MesFinal       int    `json:"mes_final"`
	DiaInicio  	   int    `json:"dia_inicio"`
	DiaFinal   	   int    `json:"dia_final"`
	Dias           int    `json:"dias"`
}

type ReservasDto []ReservaDto