package dto

type PhotosRequest struct {
	ShipID string   `json:"ship_id"`
	Urls   []string `json:"image_url" binding:"required"`
}
