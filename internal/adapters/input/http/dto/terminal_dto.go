package dto

// CreateTerminalRequest represents the request payload for creating a terminal.
type CreateTerminalRequest struct {
	Name      string  `json:"name" binding:"required"`
	UF        string  `json:"uf" binding:"required"`
	CityID    int     `json:"city_id" binding:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TerminalDetailResponse represents the detailed response for a terminal.
type TerminalDetailResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UF        string  `json:"uf"`
	CityName  string  `json:"city_name"`
	CityID    int     `json:"city_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Active    bool    `json:"active"`
}

// UpdateTerminalRequest represents the request payload for updating a terminal.
type UpdateTerminalRequest struct {
	Name      string  `json:"name" binding:"required"`
	UF        string  `json:"uf" binding:"required"`
	CityID    int     `json:"city_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TerminalListItemResponse represents a summarized item in the terminal listing.
type TerminalListItemResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UF        string  `json:"uf"`
	CityName  string  `json:"city_name"`
	CityID    int     `json:"city_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Active    bool    `json:"active"`
}

// CityListItemResponse represents a city in the city listing.
type UFCitiesResponse struct {
	UF     string                 `json:"uf"`
	Cities []CityListItemResponse `json:"cities"`
}

type CityListItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TerminalListResponse is the DTO for the terminal list response with pagination details.
type TerminalListResponse struct {
	Terminals   []TerminalListItemResponse `json:"terminals"`
	TotalCount  int                        `json:"total_count"`
	TotalPages  int                        `json:"total_pages"`
	CurrentPage int                        `json:"current_page"`
}
