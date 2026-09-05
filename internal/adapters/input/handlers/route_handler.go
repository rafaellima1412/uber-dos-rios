package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http/dto"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type RouteHandler struct {
	createRouteUseCase input.CreateRouteUseCase
	getRouteUseCase    input.GetRouteUseCase
	updateRouteUseCase input.UpdateRouteUseCase
	deleteRouteUseCase input.DeleteRouteUseCase
	listRouteUseCase   input.ListRouteUseCase
	searchRouteUseCase input.SearchRouteUseCase
}

func NewRouteHandler(createRouteUseCase input.CreateRouteUseCase, getRouteUseCase input.GetRouteUseCase, updateRouteUseCase input.UpdateRouteUseCase, deleteRouteUseCase input.DeleteRouteUseCase, listRouteUseCase input.ListRouteUseCase, searchRouteUseCase input.SearchRouteUseCase) *RouteHandler {
	return &RouteHandler{
		createRouteUseCase: createRouteUseCase,
		getRouteUseCase:    getRouteUseCase,
		updateRouteUseCase: updateRouteUseCase,
		deleteRouteUseCase: deleteRouteUseCase,
		listRouteUseCase:   listRouteUseCase,
		searchRouteUseCase: searchRouteUseCase,
	}
}

// CreateRoute godoc
// @Summary Create a new route
// @Description Create a new route with the provided details
// @Tags routes
// @Accept json
// @Produce json
// @Param route body dto.CreateRouteRequest true "Route to create"
// @Success 201 {object} dto.SuccessMessageResponse "Route created successfully"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes [post]
func (h *RouteHandler) CreateRoute(c *gin.Context) {
	var req dto.CreateRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	input := input.CreateRouteInput{
		Name:           req.Name,
		OrganizationID: req.OrganizationID,
	}

	if err := h.createRouteUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create route"})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Route created successfully"})
}

// GetRoute godoc
// @Summary Get route details
// @Description Retrieve details of a specific route by its ID
// @Tags routes
// @Accept json
// @Produce json
// @Param id path string true "Route ID"
// @Success 200 {object} dto.GetRouteDetailResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes/{id} [get]
func (h *RouteHandler) GetRoute(c *gin.Context) {
	routeID := c.Param("id")
	if routeID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Route ID is required"})
		return
	}

	output, err := h.getRouteUseCase.Execute(c.Request.Context(), routeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to retrieve route"})
		return
	}

	response := dto.GetRouteDetailResponse{
		ID:     output.ID,
		Name:   output.Name,
		Active: output.Active,
		OrganizationOutpt: dto.OrganizationDetailResponse{
			ID:            output.OrganizationOutput.ID,
			Type:          output.OrganizationOutput.Type,
			Name:          output.OrganizationOutput.Name,
			CNPJ:          output.OrganizationOutput.CNPJ,
			Email:         output.OrganizationOutput.Email,
			PhoneNumber:   output.OrganizationOutput.PhoneNumber,
		},
	}

	c.JSON(http.StatusOK, response)
}

// UpdateRoute godoc
// @Summary Update an existing route
// @Description Update the details of an existing route by its ID
// @Tags routes
// @Accept json
// @Produce json
// @Param id path string true "Route ID"
// @Param route body dto.UpdateRouteRequest true "Route details to update"
// @Success 200 {object} dto.SuccessMessageResponse "Route updated successfully"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes/{id} [put]
func (h *RouteHandler) UpdateRoute(c *gin.Context) {
	var req dto.UpdateRouteRequest
	routeID := c.Param("id")
	if routeID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Route ID is required"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	input := input.UpdateRouteInput{
		ID:             routeID,
		Name:           req.Name,
		OrganizationID: req.OrganizationID,
	}

	if err := h.updateRouteUseCase.Execute(c.Request.Context(), routeID, input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update route"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Route updated successfully"})
}

// DeleteRoute godoc
// @Summary Delete a route
// @Description Delete a specific route by its ID
// @Tags routes
// @Accept json
// @Produce json
// @Param id path string true "Route ID"
// @Success 200 {object} dto.SuccessMessageResponse "Route deleted successfully"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes/{id} [delete]
func (h *RouteHandler) DeleteRoute(c *gin.Context) {
	routeID := c.Param("id")
	if routeID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Route ID is required"})
		return
	}

	if err := h.deleteRouteUseCase.Execute(c.Request.Context(), routeID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete route"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Route deleted successfully"})
}

// ListRoutes godoc
// @Summary List routes
// @Description Retrieve a paginated list of routes
// @Tags routes
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Number of items per page" default(10)
// @Success 200 {object} dto.ListRouteResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes [get]
func (h *RouteHandler) ListRoutes(c *gin.Context) {
	page := 1
	perPage := 10

	if p := c.Query("page"); p != "" {
		if _, err := fmt.Sscanf(p, "%d", &page); err != nil || page < 1 {
			logger.Error("Invalid page parameter", zap.String("page", p))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid page parameter"})
			return
		}
	}

	if pp := c.Query("per_page"); pp != "" {
		if _, err := fmt.Sscanf(pp, "%d", &perPage); err != nil || perPage < 1 {
			logger.Error("Invalid per_page parameter", zap.String("per_page", pp))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid per_page parameter"})
			return
		}
	}

	input := input.ListRouteInput{
		Page:    page,
		PerPage: perPage,
	}

	output, err := h.listRouteUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list routes"})
		return
	}

	routes := make([]dto.ListItemRouteResponse, len(output.Routes))
	for i, route := range output.Routes {
		routes[i] = dto.ListItemRouteResponse{
			ID:     route.ID,
			Name:   route.Name,
			Active: route.Active,
			OrganizationOutpt: dto.OrganizationDetailResponse{
				ID:            route.OrganizationOutput.ID,
				Type:          route.OrganizationOutput.Type,
				Name:          route.OrganizationOutput.Name,
				CNPJ:          route.OrganizationOutput.CNPJ,
				Email:         route.OrganizationOutput.Email,
				PhoneNumber:   route.OrganizationOutput.PhoneNumber,
			},
		}
	}

	response := dto.ListRouteResponse{
		Routes:      routes,
		TotalCount:  output.TotalCount,
		TotalPages:  output.TotalPages,
		CurrentPage: output.CurrentPage,
	}

	c.JSON(http.StatusOK, response)
}

// SearchRoutes godoc
// @Summary Search routes
// @Description Search for routes based on a query string with pagination
// @Tags routes
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Number of items per page" default(10)
// @Success 200 {object} dto.ListRouteResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /routes/search [get]
func (h *RouteHandler) SearchRoutes(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Search query is required"})
		return
	}

	page := 1
	perPage := 10

	if p := c.Query("page"); p != "" {
		if _, err := fmt.Sscanf(p, "%d", &page); err != nil || page < 1 {
			logger.Error("Invalid page parameter", zap.String("page", p))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid page parameter"})
			return
		}
	}

	if pp := c.Query("per_page"); pp != "" {
		if _, err := fmt.Sscanf(pp, "%d", &perPage); err != nil || perPage < 1 {
			logger.Error("Invalid per_page parameter", zap.String("per_page", pp))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid per_page parameter"})
			return
		}
	}

	input := input.SearchRouteInput{
		Page:    page,
		PerPage: perPage,
	}

	output, err := h.searchRouteUseCase.Execute(c.Request.Context(), query, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search routes"})
		return
	}

	routes := make([]dto.ListItemRouteResponse, len(output.Routes))
	for i, route := range output.Routes {
		routes[i] = dto.ListItemRouteResponse{
			ID:     route.ID,
			Name:   route.Name,
			Active: route.Active,
			OrganizationOutpt: dto.OrganizationDetailResponse{
				ID:            route.OrganizationOutput.ID,
				Type:          route.OrganizationOutput.Type,
				Name:          route.OrganizationOutput.Name,
				CNPJ:          route.OrganizationOutput.CNPJ,
				Email:         route.OrganizationOutput.Email,
				PhoneNumber:   route.OrganizationOutput.PhoneNumber,
			},
		}
	}

	response := dto.ListRouteResponse{
		Routes:      routes,
		TotalCount:  output.TotalCount,
		TotalPages:  output.TotalPages,
		CurrentPage: output.CurrentPage,
	}

	c.JSON(http.StatusOK, response)
}
