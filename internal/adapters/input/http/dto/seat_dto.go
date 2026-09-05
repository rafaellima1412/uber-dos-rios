package dto

type CreateSeatRequest struct {
	IsActive          bool `json:"is_active"`
	LeftColumnsCount  int  `json:"left_columns_count"`
	RightColumnsCount int  `json:"right_columns_count"`
	Nivel             int  `json:"nivel"`
	Prefix            bool `json:"prefix"`
}

type UpdateSeatRequest struct {
	IsActive          bool `json:"is_active"`
	LeftColumnsCount  int  `json:"left_columns_count"`
	RightColumnsCount int  `json:"right_columns_count"`
	Prefix            bool `json:"prefix"`
}

type SeatDetailResponse struct {
	ID                string `json:"id"`
	IsActive          bool   `json:"is_active"`
	LeftColumnsCount  int    `json:"left_columns_count"`
	RightColumnsCount int    `json:"right_columns_count"`
	Prefix            bool   `json:"prefix"`
}

type SeatListResponse struct {
	Seats       []SeatDetailResponse `json:"seats"`
	TotalCount  int                  `json:"total_count"`
	TotalPages  int                  `json:"total_pages"`
	CurrentPage int                  `json:"current_page"`
}

type Seats struct {
	TotalSeats     int      `json:"total_seats"`
	OccupiedUnits  []string `json:"occupied_units"`
	AvailableSeats int      `json:"available_seats"`
}
type Cabins struct {
	TotalCabins     int      `json:"total_cabins"`
	OccupiedUnits   []string `json:"occupied_units"`
	AvailableCabins int      `json:"available_cabins"`
}

type ListUnitAvailableResponse struct {
	ShipID   string `json:"ship_id"`
	ShipName string `json:"ship_name"`
	Seats    Seats  `json:"seats"`
	Cabins   Cabins `json:"cabins"`
}
