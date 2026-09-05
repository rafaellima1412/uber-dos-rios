package dto

import "time"

type TripInstanceResponse struct {
	DepartureAt time.Time `json:"departure_at"`
	ArrivalAt   time.Time `json:"arrival_at"`
}
