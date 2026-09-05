package dto

import (
	"time"

	"github.com/google/uuid"
)

type JSONTime time.Time

type TripConfigListItemResponse struct {
	ID             uuid.UUID `json:"id"`
	ShipID         uuid.UUID `json:"ship_id"`
	RouteID        uuid.UUID `json:"route_id"`
	ShipName       string    `json:"ship_name"`
	RouteName      string    `json:"route_name"`
	Recurrence     string    `json:"recurrence"`
	ExpirationDate time.Time `json:"expiration_date" biding:"required"`
	DepartureTime  string    `json:"departure_time"`
	ArrivalTime    string    `json:"arrival_time"`
	StartDate      time.Time `json:"start_date" biding:"required"`
	DurationDays   int       `json:"duration_days"`
}

type ListTripConfigResponse struct {
	Trips       []TripConfigListItemResponse `json:"trips_config"`
	TotalCount  int                          `json:"total_count"`
	TotalPages  int                          `json:"total_pages"`
	CurrentPage int                          `json:"current_page"`
}

type TripConfigResponse struct {
	ID             uuid.UUID `json:"id"`
	ShipID         uuid.UUID `json:"ship_id"`
	RouteID        uuid.UUID `json:"route_id"`
	ShipName       string    `json:"ship_name"`
	RouteName      string    `json:"route_name"`
	Recurrence     string    `json:"recurrence" biding:"required"`
	ExpirationDate time.Time `json:"expiration_date" biding:"required"`
	DepartureTime  string    `json:"departure_time" biding:"required"`
	DurationDays   int       `json:"duration_days" biding:"required"`
	StartDate      string    `json:"start_date" biding:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateTripConfigRequest struct {
	ShipID         string `json:"ship_id" binding:"required"`
	RouteID        string `json:"route_id" binding:"required"`
	Recurrence     string `json:"recurrence" biding:"required"`
	ExpirationDate string `json:"expiration_date" biding:"required"`
	DepartureTime  string `json:"departure_time" biding:"required"`
	ArrivalTime    string `json:"arrival_time" biding:"required"`
	DurationDays   int    `json:"duration_days" biding:"required"`
	StartDate      string `json:"start_date" biding:"required"`
}

type UpdateTripConfigRequest struct {
	ShipID         string    `json:"ship_id" biding:"required"`
	RouteID        string    `json:"route_id" biding:"required"`
	Recurrence     string    `json:"recurrence" biding:"required"`
	ExpirationDate time.Time `json:"expiration_date" biding:"required"`
	DepartureTime  time.Time `json:"departure_time" biding:"required"`
	ArrivalTime    time.Time `json:"arrival_time" biding:"required"`
	DurationDays   int       `json:"duration_days" biding:"required"`
	StartDate      time.Time `json:"start_date" biding:"required"`
}

type Itinerary struct {
	City      string `json:"city"`
	UF        string `json:"uf"`
	StopOrder int    `json:"stop_order"`
}

type TripConfigResult struct {
	RouteID          string        `json:"route_id"`
	RouteName        string        `json:"route_name"`
	DepartureTime    JSONTime      `json:"departure_time"`
	ArrivalTime      JSONTime      `json:"arrival_time"`
	OrganizationID   string        `json:"organization_id"`
	OrganizationName string        `json:"organization_name"`
	HoraExtra        time.Duration `json:"hora_extra" swaggertype:"primitive,integer"`
	IsComplete       bool          `json:"is_complete"`
	Itinerary        []Itinerary   `json:"itinerary"`
}

type TripConfigSearchResponse struct {
	TotalCost int        `json:"total_cost"`
	Path      []PathStep `json:"path"`
}

type PathStep struct {
	OrganizationID      uuid.UUID `json:"organization_id"`
	CityID              int       `json:"city_id"`
	Cost                int       `json:"cost"`
	RouteID             uuid.UUID `json:"route_id"`
	RouteName           string    `json:"route_name"`
	DateDepartureTime   time.Time `json:"departure_date"`
	DateArrivalTime     time.Time `json:"arrival_date"`
	ShipID              uuid.UUID `json:"ship_id"`
	ShipName            string    `json:"ship_name"`
	ShipURL             string    `json:"ship_url"`
	TripConfigurationID uuid.UUID `json:"trip_configuration_id"`
}

type TripConfigSearchRequest struct {
	DateOrigin           string `json:"date_origin" biding:"required"`
	DateDest             string `json:"date_dest" biding:"required"`
	CityOriginID         int    `json:"city_origin_id" biding:"required"`
	CityDestID           int    `json:"city_dest_id" biding:"required"`
	MaxStops             int    `json:"max_stops" biding:"required"`
	MaxCost              int    `json:"max_cost" biding:"required"`
	ShipID               string `json:"ship_id"`
	PassengerCountSeats  int    `json:"passenger_count_seats"`
	PassengerCountCabins int    `json:"passenger_count_cabins"`
	MaxResults           int    `json:"max_results"`
}
