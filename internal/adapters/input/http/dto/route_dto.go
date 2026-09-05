package dto

// CreateRouteRequest represents the request payload for creating a new route.
type CreateRouteRequest struct {
	Name           string `json:"name" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required,uuid"`
}

// OrganizationDetailResponse represents the organization details in the GetRoute output.
type OrganizationDetailResponse struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	CNPJ        *string `json:"cnpj,omitempty"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// GetRouteDetailResponse represents the response payload for retrieving route details.
type GetRouteDetailResponse struct {
	ID                string                     `json:"id"`
	Name              string                     `json:"name"`
	Active            bool                       `json:"active"`
	OrganizationOutpt OrganizationDetailResponse `json:"organization"`
}

// UpdateRouteRequest represents the request payload for updating an existing route.
type UpdateRouteRequest struct {
	Name           string `json:"name" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required,uuid"`
}

// ListItemRouteResponse represents a summarized item in the route listing.
type ListItemRouteResponse struct {
	ID                string                     `json:"id"`
	Name              string                     `json:"name"`
	Active            bool                       `json:"active"`
	OrganizationOutpt OrganizationDetailResponse `json:"organization"`
}

// ListRouteResponse is the DTO for the route list response with pagination details.
type ListRouteResponse struct {
	Routes      []ListItemRouteResponse `json:"routes"`
	TotalCount  int                     `json:"total_count"`
	TotalPages  int                     `json:"total_pages"`
	CurrentPage int                     `json:"current_page"`
}
