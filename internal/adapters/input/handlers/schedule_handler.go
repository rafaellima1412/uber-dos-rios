package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/http/dto"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type ScheduleHandler struct {
	createScheduleUseCase  input.CreateScheduleUseCase
	getScheduleUseCase     input.GetScheduleUseCase
	updateScheduleUseCase  input.UpdateScheduleUseCase
	deleteScheduleUseCase  input.DeleteScheduleUseCase
	listSchedulesUseCase   input.ListScheduleUseCase
	searchSchedulesUseCase input.SearchScheduleUseCase
}

func NewScheduleHandler(
	createScheduleUseCase input.CreateScheduleUseCase,
	getScheduleUseCase input.GetScheduleUseCase,
	updateScheduleUseCase input.UpdateScheduleUseCase,
	deleteScheduleUseCase input.DeleteScheduleUseCase,
	listSchedulesUseCase input.ListScheduleUseCase,
	searchSchedulesUseCase input.SearchScheduleUseCase,
) *ScheduleHandler {
	return &ScheduleHandler{
		createScheduleUseCase:  createScheduleUseCase,
		getScheduleUseCase:     getScheduleUseCase,
		updateScheduleUseCase:  updateScheduleUseCase,
		deleteScheduleUseCase:  deleteScheduleUseCase,
		listSchedulesUseCase:   listSchedulesUseCase,
		searchSchedulesUseCase: searchSchedulesUseCase,
	}
}

// CreateSchedule godoc
// @Summary Create a new schedule
// @Description Create a new schedule with the provided information
// @Tags schedules
// @Accept json
// @Produce json
// @Param schedule body dto.CreateScheduleRequest true "Schedule to create"
// @Success 201 {object} dto.SuccessMessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules [post]
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var req dto.CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	inputSchedule := input.CreateScheduleInput{
		RouteID:    req.RouteID,
		TerminalID: req.TerminalID,
		Value:      req.Value,
		StopOrder:  req.StopOrder,
	}

	// Call the use case to create the schedule
	if err := h.createScheduleUseCase.Execute(c.Request.Context(), inputSchedule); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Schedule created successfully"})
}

// GetSchedule godoc
// @Summary Get a schedule by ID
// @Description Retrieve a schedule by its ID
// @Tags schedules
// @Accept json
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} dto.GetScheduleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules/{id} [get]
func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
	id := c.Param("id")

	schedule, err := h.getScheduleUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetScheduleResponse{
		ID: schedule.ID,
		Route: dto.GetRouteDetailResponse{
			ID:     schedule.Route.ID,
			Name:   schedule.Route.Name,
			Active: schedule.Route.Active,
		},
		Terminal: dto.TerminalDetailResponse{
			ID:        schedule.Terminal.ID.String(),
			Name:      schedule.Terminal.Name,
			UF:        schedule.Terminal.UF,
			CityID:      schedule.Terminal.CityID,
			Latitude:  schedule.Terminal.Latitude,
			Longitude: schedule.Terminal.Longitude,
			Active:    schedule.Terminal.Active,
		},
		Value:  schedule.Value,
		Active: schedule.Active,
	})
}

// UpdateSchedule godoc
// @Summary Update an existing schedule
// @Description Update an existing schedule with the provided information
// @Tags schedules
// @Accept json
// @Produce json
// @Param id path string true "Schedule ID"
// @Param schedule body dto.UpdateScheduleRequest true "Schedule to update"
// @Success 200 {object} dto.SuccessMessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules/{id} [put]
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	inputSchedule := input.UpdateScheduleInput{
		ID:         id,
		RouteID:    req.RouteID,
		TerminalID: req.TerminalID,
		Active:     req.Active,
		Value:      req.Value,
	}

	if err := h.updateScheduleUseCase.Execute(c.Request.Context(), inputSchedule); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Schedule updated successfully"})
}

// DeleteSchedule godoc
// @Summary Delete a schedule by ID
// @Description Delete a schedule by its ID
// @Tags schedules
// @Accept json
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} dto.SuccessMessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules/{id} [delete]
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	if err := h.deleteScheduleUseCase.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Schedule deleted successfully"})
}

// ListSchedules godoc
// @Summary List schedules with pagination
// @Description Retrieve a paginated list of schedules
// @Tags schedules
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} dto.ListScheduleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules [get]
func (h *ScheduleHandler) ListSchedules(c *gin.Context) {
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

	inputList := input.ListScheduleInput{
		Page:    page,
		PerPage: perPage,
	}

	schedules, err := h.listSchedulesUseCase.Execute(c.Request.Context(), inputList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	response := dto.ListScheduleResponse{
		Schedules:   make([]dto.ListItemScheduleResponse, len(schedules.Schedules)),
		TotalCount:  schedules.TotalCount,
		TotalPages:  schedules.TotalPages,
		CurrentPage: page,
	}

	for i, sched := range schedules.Schedules {
		response.Schedules[i] = dto.ListItemScheduleResponse{
			ID:           sched.ID,
			RouteID:      sched.RouteID,
			TerminalID:   sched.TerminalID,
			RouteName:    sched.RouteName,
			TerminalName: sched.TenimalName,
			Value:        sched.Value,
			Active:       sched.Active,
		}
	}

	c.JSON(http.StatusOK, response)
}

// SearchSchedules godoc
// @Summary Search schedules with a query string
// @Description Search schedules by a query string with pagination
// @Tags schedules
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} dto.ListScheduleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedules/search [get]
func (h *ScheduleHandler) SearchSchedules(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Query parameter is required"})
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

	inputSearch := input.SearchScheduleInput{
		Page:    page,
		PerPage: perPage,
	}

	schedules, err := h.searchSchedulesUseCase.Execute(c.Request.Context(), query, inputSearch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	response := dto.ListScheduleResponse{
		Schedules:   make([]dto.ListItemScheduleResponse, len(schedules.Schedules)),
		TotalCount:  schedules.TotalCount,
		TotalPages:  schedules.TotalPages,
		CurrentPage: page,
	}

	for i, sched := range schedules.Schedules {
		response.Schedules[i] = dto.ListItemScheduleResponse{
			ID:         sched.ID,
			RouteID:    sched.RouteID,
			TerminalID: sched.TerminalID,
			Value:      sched.Value,
			Active:     sched.Active,
		}
	}

	c.JSON(http.StatusOK, response)
}
