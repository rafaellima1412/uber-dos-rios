package dto

type TripsSelected struct {
	TripConfigurationID string `json:"trip_configuration_id"`
	ShipID              string   `json:"ship_id"`
	RouteID             string   `json:"route_id"`
	DepartureAt         string   `json:"departure_date"`
	ArrivalAt           string   `json:"arrival_date"`
	OccupiedUnits       []string `json:"occupied_units"`
}

type ReservationTripsSelectedRequest struct {
	ReservationID       string   `json:"reservation_id"`
	SelectedTrips []TripsSelected `json:"selected_trips"`
}

type TripsSelectedResponse struct {
	TripConfigurationID string `json:"trip_configuration_id"`
	ShipID              string `json:"ship_id"`
	RouteID             string `json:"route_id"`
	DepartureAt         string `json:"departure_date"`
	ArrivalAt           string `json:"arrival_date"`
}


