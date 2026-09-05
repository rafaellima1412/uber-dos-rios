package dto

import "time"

// Create Reservation Request
type CreateReservationRequest struct {
	UserID          string `json:"user_id"`
	ReservationDate string `json:"reservation_date"`
	// Status          string `json:"status"`
}

// Update Reservation Request
type UpdateReservationRequest struct {
	UserID          string `json:"user_id"`
	ReservationDate string `json:"reservation_date"`
	Status          string `json:"status"`
}

//List Reservation
type ReservationTripConfigurationResponse struct {
	ReservationID       string   `json:"reservation_id"`
	TripConfigurationID string   `json:"trip_configuration_id"`
	OccupiedSeats       []string `json:"occupied_units"`
}

type ReservationListItemResponse struct {
	ID              string     `json:"id"`
	UserID          string     `json:"user_id"`
	ReservationDate string     `json:"reservation_date"`
	Status          string     `json:"status"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

// Reservation Response
type ListReservationResponse struct {
	Reservations []ReservationListItemResponse `json:"reservations"`
	TotalCount   int                           `json:"total_count"`
	TotalPages   int                           `json:"total_pages"`
	CurrentPage  int                           `json:"current_page"`
}
