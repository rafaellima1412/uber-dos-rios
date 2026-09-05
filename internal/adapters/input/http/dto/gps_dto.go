package dto

type GPSRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Speed     float64 `json:"speed"`
	Course    float64 `json:"course"`
	ShipName  string  `json:"ship_name" binding:"required"`
}
