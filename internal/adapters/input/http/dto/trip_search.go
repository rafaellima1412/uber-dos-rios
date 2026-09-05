package dto

import "time"

type TripFiltersDTO struct {
	RouteID *string `form:"route_id"`
	ShipID  *string `form:"ship_id"`
	//	TripConfigurationsID *string    `form:"trip_configurations_id"`
	DepartureAfter  *time.Time `form:"departure_after" time_format:"2006-01-02T15:04:05Z07:00"`
	DepartureBefore *time.Time `form:"departure_before" time_format:"2006-01-02T15:04:05Z07:00"`
	Page            int        `form:"page,default=1"`
	Limit           int        `form:"limit,default=10"`
}

type TripResponseDTO struct {
	TripConfigurationID string `json:"trip_configuration_id"`
	ShipID              string `json:"ship_id"`
	RouteID             string `json:"route_id"`
	DepartureAt         string `json:"departure_date"`
	ArrivalAt           string `json:"arrival_date"`
}

type TripFilterListResponse struct {
	Trips       []TripResponseDTO `json:"trips"`
	TotalCount  int               `json:"total_count"`
	TotalPages  int               `json:"total_pages"`
	CurrentPage int               `json:"current_page"`
}
