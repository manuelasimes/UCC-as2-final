package dto

type CheckRoomDto struct {

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`
}

type CheckRoomsDto []CheckRoomDto