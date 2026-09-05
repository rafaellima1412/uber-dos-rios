package dto

// CreateScheduleRequest represents the request payload for creating a new schedule.
type CreateScheduleRequest struct {
	RouteID    string  `json:"route_id" binding:"required,uuid"`
	TerminalID string  `json:"terminal_id" binding:"required,uuid"`
	Value      float64 `json:"value" binding:"required,gt=0"`
	StopOrder  int     `json:"stop_order" binding:"required"`
}

// GetScheduleResponse represents the response payload for retrieving schedule details.
type GetScheduleResponse struct {
	ID       string                 `json:"id"`
	Route    GetRouteDetailResponse `json:"route"`
	Terminal TerminalDetailResponse `json:"terminal"`
	Active   bool                   `json:"active"`
	Value    float64                `json:"value"`
}

// UpdateScheduleRequest represents the request payload for updating an existing schedule.
type UpdateScheduleRequest struct {
	RouteID    string  `json:"route_id" binding:"required,uuid"`
	TerminalID string  `json:"terminal_id" binding:"required,uuid"`
	Active     bool    `json:"active"`
	Value      float64 `json:"value" binding:"required,gt=0"`
}

// ListItemScheduleResponse represents a summarized item in the schedule listing.
type ListItemScheduleResponse struct {
	ID           string  `json:"id"`
	RouteID      string  `json:"route_id"`
	TerminalID   string  `json:"terminal_id"`
	RouteName    string  `json:"route_name"`
	TerminalName string  `json:"terminal_name"`
	Value        float64 `json:"value"`
	Active       bool    `json:"active"`
}

// ListScheduleResponse is the DTO for the schedule list response with pagination details.
type ListScheduleResponse struct {
	Schedules   []ListItemScheduleResponse `json:"schedules"`
	TotalCount  int                        `json:"total_count"`
	TotalPages  int                        `json:"total_pages"`
	CurrentPage int                        `json:"current_page"`
}
