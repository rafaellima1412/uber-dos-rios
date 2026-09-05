package dto

import (
	"github.com/google/uuid"
)

type ShipResponse struct {
	ID                  uuid.UUID  `json:"id"`
	ConfigurationID *[]uuid.UUID `json:"configuration_id,omitempty"`
	OrganizationID      *uuid.UUID `json:"organization_id,omitempty"`
	IMO                 string     `json:"imo"`
	Name                string     `json:"name"`
	Status              string     `json:"status"`
	TypeShip            string     `json:"type_ship"`
	TotalSeats          int        `json:"total_seats"`
	WeightCapacity      float64    `json:"weight_capacity"`
	PassengerCapacity   int        `json:"passenger_capacity"`
	ImageUrl            []string   `json:"image_url"`
}

type CreateShipRequest struct {
	Name              string  `json:"name" binding:"required"`
	TypeShip          string  `json:"type_ship" binding:"required"`
	IMO               string  `json:"imo" binding:"required"`
	TotalSeats        int     `json:"total_seats" binding:"required"`
	TotalCabins       int     `json:"total_cabins" binding:"required"`
	Status            string  `json:"status" binding:"required"`
	PassengerCapacity int     `json:"passenger_capacity" binding:"required"`
	WeightCapacity    float64 `json:"weight_capacity" binding:"required"`
	OrganizationID    string  `json:"organization_id"`
	ConfigurationsID  []string  `json:"configuration_id"`
}

type ShipListItemResponse struct {
	ID                uuid.UUID  `json:"id"`
	ConfigurationID   *[]uuid.UUID `json:"configuration_id"`
	OrganizationID    *uuid.UUID `json:"organization_id"`
	IMO               string     `json:"imo"`
	Name              string     `json:"name"`
	Status            string     `json:"status"`
	TypeShip          string     `json:"type_ship"`
	TotalSeats        int        `json:"total_seats"`
	WeightCapacity    float64    `json:"weight_capacity"`
	PassengerCapacity int        `json:"passenger_capacity"`
	ImageUrl          []string   `json:"image_url"`
}

type ListShipResponse struct {
	Ships       []ShipListItemResponse `json:"ships"`
	TotalCount  int                    `json:"total_count"`
	TotalPages  int                    `json:"total_pages"`
	CurrentPage int                    `json:"current_page"`
}

type UpdateShipRequest struct {
	Name              string  `json:"name"`
	TypeShip          string  `json:"type_ship"`
	IMO               string  `json:"imo"`
	TotalSeats        int     `json:"total_seats"`
	TotalCabins       int     `json:"total_cabins"`
	Status            string  `json:"status"`
	PassengerCapacity int     `json:"passenger_capacity"`
	WeightCapacity    float64 `json:"weight_capacity"`
	//ConfigurationsID  []string  `json:"configuration_id"`
	OrganizationID    string  `json:"organization_id"`
}
