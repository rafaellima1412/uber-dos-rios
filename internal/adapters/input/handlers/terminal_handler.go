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

type TerminalHandler struct {
	createTerminalUseCase input.CreateTerminalUseCase
	getTerminalUseCase    input.GetTerminalUseCase
	updateTerminalUseCase input.UpdateTerminalUseCase
	deleteTerminalUseCase input.DeleteTerminalUseCase
	listTerminalUseCase   input.ListTerminalUseCase
	searchTerminalUseCase input.SearchTerminalUseCase
	listCitiesUseCase     input.ListCitiesUseCase
}

func NewTerminalHandler(
	createTerminalUseCase input.CreateTerminalUseCase,
	getTerminalUseCase input.GetTerminalUseCase,
	updateTerminalUseCase input.UpdateTerminalUseCase,
	deleteTerminalUseCase input.DeleteTerminalUseCase,
	listTerminalUseCase input.ListTerminalUseCase,
	searchTerminalUseCase input.SearchTerminalUseCase,
	listCitiesUseCase input.ListCitiesUseCase,
) *TerminalHandler {
	return &TerminalHandler{
		createTerminalUseCase: createTerminalUseCase,
		getTerminalUseCase:    getTerminalUseCase,
		updateTerminalUseCase: updateTerminalUseCase,
		deleteTerminalUseCase: deleteTerminalUseCase,
		listTerminalUseCase:   listTerminalUseCase,
		searchTerminalUseCase: searchTerminalUseCase,
		listCitiesUseCase:     listCitiesUseCase,
	}
}

// CreateTerminal godoc
// @Summary Create a new terminal
// @Description Create a new terminal with the provided details
// @Tags terminals
// @Accept json
// @Produce json
// @Param terminal body dto.CreateTerminalRequest true "Terminal details"
// @Success 201 {object} dto.SuccessMessageResponse "Terminal created successfully"
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals [post]
func (h *TerminalHandler) CreateTerminal(c *gin.Context) {
	var req dto.CreateTerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	terminalInput := input.CreateTerminalInput{
		Name:      req.Name,
		UF:        req.UF,
		CityID:    req.CityID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	if err := h.createTerminalUseCase.Execute(c.Request.Context(), terminalInput); err != nil {
		logger.Error("Error executing CreateTerminal use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	logger.Info("Terminal created successfully", zap.String("terminal_name", req.Name))
	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Terminal created successfully"})
}

// GetTerminal godoc
// @Summary Get a terminal by ID
// @Description Retrieve terminal details by its ID
// @Tags terminals
// @Accept json
// @Produce json
// @Param id path string true "Terminal ID"
// @Success 200 {object} dto.TerminalDetailResponse
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 404  {object}  dto.ErrorResponse "Terminal not found"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals/{id} [get]
func (h *TerminalHandler) GetTerminal(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		logger.Error("Terminal ID is required")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Terminal ID is required"})
		return
	}

	terminal, err := h.getTerminalUseCase.Execute(c.Request.Context(), idParam)
	if err != nil {
		logger.Error("Error executing GetTerminal use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	responseTerminal := dto.TerminalDetailResponse{
		ID:        terminal.ID.String(),
		Name:      terminal.Name,
		UF:        terminal.UF,
		CityID:    terminal.CityID,
		Latitude:  terminal.Latitude,
		Longitude: terminal.Longitude,
		Active:    terminal.Active,
	}

	logger.Info("Terminal retrieved successfully", zap.String("terminal_id", idParam))
	c.JSON(http.StatusOK, responseTerminal)
}

// UpdateTerminal godoc
// @Summary Update an existing terminal
// @Description Update terminal details by its ID
// @Tags terminals
// @Accept json
// @Produce json
// @Param id path string true "Terminal ID"
// @Param terminal body dto.UpdateTerminalRequest true "Updated terminal details"
// @Success 200 {object} dto.SuccessMessageResponse "Terminal updated successfully"
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 404  {object}  dto.ErrorResponse "Terminal not found"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals/{id} [put]
func (h *TerminalHandler) UpdateTerminal(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		logger.Error("Terminal ID is required")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Terminal ID is required"})
		return
	}

	var req dto.UpdateTerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	terminalInput := input.UpdateTerminalInput{
		ID:        idParam,
		Name:      req.Name,
		UF:        req.UF,
		CityID:    req.CityID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	if err := h.updateTerminalUseCase.Execute(c.Request.Context(), idParam, terminalInput); err != nil {
		logger.Error("Error executing UpdateTerminal use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	logger.Info("Terminal updated successfully", zap.String("terminal_id", idParam))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Terminal updated successfully"})
}

// DeleteTerminal godoc
// @Summary Delete a terminal by ID
// @Description Delete a terminal using its ID
// @Tags terminals
// @Accept json
// @Produce json
// @Param id path string true "Terminal ID"
// @Success 204 "No Content"
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 404  {object}  dto.ErrorResponse "Terminal not found"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals/{id} [delete]
func (h *TerminalHandler) DeleteTerminal(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		logger.Error("Terminal ID is required")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Terminal ID is required"})
		return
	}

	if err := h.deleteTerminalUseCase.Execute(c.Request.Context(), idParam); err != nil {
		logger.Error("Error executing DeleteTerminal use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	logger.Info("Terminal deleted successfully", zap.String("terminal_id", idParam))
	c.Status(http.StatusNoContent)
}

// ListTerminals godoc
// @Summary List terminals with pagination
// @Description Retrieve a paginated list of terminals
// @Tags terminals
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} dto.TerminalListResponse
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals [get]
func (h *TerminalHandler) ListTerminals(c *gin.Context) {
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

	listInput := input.ListTerminalInput{
		Page:    page,
		PerPage: perPage,
	}

	listOutput, err := h.listTerminalUseCase.Execute(c.Request.Context(), listInput)
	if err != nil {
		logger.Error("Error executing ListTerminals use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var terminalResponses []dto.TerminalListItemResponse
	for _, t := range listOutput.Terminals {
		terminalResponses = append(terminalResponses, dto.TerminalListItemResponse{
			ID:        t.TerminalID,
			Name:      t.Name,
			UF:        t.UF,
			CityID:    t.CityID,
			CityName:  t.CityName,
			Latitude:  t.Latitude,
			Longitude: t.Longitude,
			Active:    t.Active,
		})
	}

	response := dto.TerminalListResponse{
		Terminals:   terminalResponses,
		TotalCount:  listOutput.TotalCount,
		TotalPages:  listOutput.TotalPages,
		CurrentPage: listOutput.CurrentPage,
	}

	logger.Info("Terminals listed successfully", zap.Int("count", len(terminalResponses)))
	c.JSON(http.StatusOK, response)
}

// SearchTerminals godoc
// @Summary Search terminals with pagination
// @Description Search for terminals using a query string with pagination
// @Tags terminals
// @Accept json
// @Produce json
// @Param query query int false "Search query"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} dto.TerminalListResponse
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals/search [get]
func (h *TerminalHandler) SearchTerminals(c *gin.Context) {
	query := c.Query("query")
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

	searchInput := input.SearchTerminalInput{
		Page:    page,
		PerPage: perPage,
	}

	searchOutput, err := h.searchTerminalUseCase.Execute(c.Request.Context(), query, searchInput)
	if err != nil {
		logger.Error("Error executing SearchTerminals use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var terminalResponses []dto.TerminalListItemResponse
	for _, t := range searchOutput.Terminals {
		terminalResponses = append(terminalResponses, dto.TerminalListItemResponse{
			ID:        t.TerminalID,
			Name:      t.Name,
			UF:        t.UF,
			CityID:    t.CityID,
			CityName:  t.CityName,
			Latitude:  t.Latitude,
			Longitude: t.Longitude,
			Active:    t.Active,
		})
	}

	response := dto.TerminalListResponse{
		Terminals:   terminalResponses,
		TotalCount:  searchOutput.TotalCount,
		TotalPages:  searchOutput.TotalPages,
		CurrentPage: searchOutput.CurrentPage,
	}

	logger.Info("Terminals searched successfully", zap.Int("count", len(terminalResponses)))
	c.JSON(http.StatusOK, response)
}

// ListCities godoc
// @Summary List all cities with terminals
// @Description Retrieve a list of all cities that have terminals
// @Tags terminals
// @Accept json
// @Produce json
// @Success 200 {object} dto.CityListItemResponse
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /terminals/cities [get]
func (h *TerminalHandler) ListCities(c *gin.Context) {
	cities, err := h.listCitiesUseCase.Execute(c.Request.Context())
	if err != nil {
		logger.Error("Error executing ListCities use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	groups := make(map[string][]dto.CityListItemResponse)

	for _, city := range cities {
		item := dto.CityListItemResponse{
			ID:   city.ID,
			Name: city.Name,
		}
		groups[city.State] = append(groups[city.State], item)
	}

	var response []dto.UFCitiesResponse
	for uf, cityList := range groups {
		response = append(response, dto.UFCitiesResponse{
			UF:     uf,
			Cities: cityList,
		})
	}
	logger.Info("Cities listed successfully")
	c.JSON(http.StatusOK, response)
}
