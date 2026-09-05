package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http/dto"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type TripConfigHandler struct {
	createTripConfigUseCase input.CreateTripConfigUseCase
	deleteTripConfigUseCase input.DeleteTripConfigUseCase
	updateTripConfigUseCase input.UpdateTripConfigUseCase
	listTripConfigUseCase   input.ListTripConfigUseCase
	getTripConfigUseCase    input.GetTripConfigUseCase
}

func NewTripConfigHandler(
	createTripConfigUC input.CreateTripConfigUseCase,
	deleteTripConfigUC input.DeleteTripConfigUseCase,
	updateTripConfigUC input.UpdateTripConfigUseCase,
	listTripConfigUC input.ListTripConfigUseCase,
	getTripConfigUC input.GetTripConfigUseCase,
) *TripConfigHandler {
	return &TripConfigHandler{
		createTripConfigUseCase: createTripConfigUC,
		deleteTripConfigUseCase: deleteTripConfigUC,
		updateTripConfigUseCase: updateTripConfigUC,
		listTripConfigUseCase:   listTripConfigUC,
		getTripConfigUseCase:    getTripConfigUC,
	}
}

// CreateTrip godoc
// @Summary      Create a new Trip Config
// @Description  Registers a new Trip Config with the provided details.
// @Tags         trips-config
// @Accept       json
// @Produce      json
// @Param        organization   body      dto.CreateTripConfigRequest  true  "Trip Config Data to Create"
// @Success      201  {object}  dto.TripConfigResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips-config [post]
func (h *TripConfigHandler) CreateTripConfig(c *gin.Context) {
	var req dto.CreateTripConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	exp, err := time.Parse(time.RFC3339, req.ExpirationDate)
	if err != nil {
		logger.Error("Invalid expiration date format", zap.Error(err))
		c.JSON(400, dto.ErrorResponse{Error: "Invalid expiration date format"})
		return
	}
	start, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		logger.Error("Invalid start date format", zap.Error(err))
		c.JSON(400, dto.ErrorResponse{Error: "Invalid start date format"})
		return
	}

	tripInput := input.CreateTripConfigInput{
		ShipID:         req.ShipID,
		RouteID:        req.RouteID,
		Recurrence:     req.Recurrence,
		ExpirationDate: exp,
		DurationDays:   req.DurationDays,
		StartDate:      start,
		DepartureTime:  req.DepartureTime,
		ArrivalTime:    req.ArrivalTime,
	}

	errUC := h.createTripConfigUseCase.Execute(c.Request.Context(), &tripInput)
	if errUC != nil {
		logger.Error("Failed to create config trip", zap.Error(errUC))
		c.JSON(500, dto.ErrorResponse{Error: fmt.Sprintf("Failed to create config trip error: %v", errUC)})
		return
	}
	logger.Info("Trip created successfully", zap.String("trip_date", req.DepartureTime))
	c.JSON(201, dto.SuccessMessageResponse{Message: "config trip created successfully"})

}

// UpdateTrip godoc
// @Summary      Update an existing Trip Config
// @Description  Updates the details of an existing Trip Config.
// @Tags         trips-config
// @Accept       json
// @Produce      json
// @Param        id             path      string                 true  "Trip Config ID"
// @Param        trip           body      dto.UpdateTripConfigRequest  true  "Trip Config Data to Update"
// @Success      200  {object}   dto.TripConfigResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips-config/{id} [put]

func (h *TripConfigHandler) UpdateTripConfig(c *gin.Context) {
	tripID := c.Param("id")
	if tripID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "TRip ID is required"})
		return
	}

	var req dto.UpdateTripConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	tripInput := input.UpdateTripConfigInput{
		TripID:         tripID,
		ShipID:         req.ShipID,
		RouteID:        req.RouteID,
		Recurrence:     req.Recurrence,
		ExpirationDate: req.ExpirationDate,
		DurationDays:   req.DurationDays,
		StartDate:      req.StartDate,
		DepartureTime:  req.DepartureTime,
	}

	err := h.updateTripConfigUseCase.Execute(c.Request.Context(), tripID, &tripInput)
	if err != nil {
		logger.Error("Failed to update trip", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update trip"})
		return
	}

	logger.Info("Trip updated successfully", zap.String("trip_id", tripID))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Trip updated successfully"})
}

// FindTripByID godoc
// @Summary      Get Trip Config Config by ID
// @Description  Retrieves a Trip Config by its ID.
// @Tags         trips-config
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Trip Config ID"
// @Success      200  {object}   dto.TripConfigResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid Trip Config ID"
// @Failure      404	{object}  dto.ErrorResponse "Trip Config not found"
// @Failure 		 500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips-config/{id} [get]

func (h *TripConfigHandler) FindByIDTripConfig(c *gin.Context) {
	tripID := c.Param("id")
	trip, err := h.getTripConfigUseCase.Execute(c.Request.Context(), tripID)
	if err != nil {
		logger.Error("Error retrieving trip", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	if trip == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Trip not found"})
		return
	}
	response := dto.TripConfigResponse{
		ID:             trip.ID,
		ShipID:         trip.ShipID,
		RouteID:        trip.RouteID,
		ShipName:       trip.ShipName,
		RouteName:      trip.RouteName,
		Recurrence:     trip.Recurrence,
		ExpirationDate: trip.ExpirationDate,
		DepartureTime:  trip.DepartureTime,
		DurationDays:   trip.DurationDays,
		StartDate:      trip.StartDate,
		CreatedAt:      trip.CreatedAt,
		UpdatedAt:      trip.UpdatedAt,
	}
	logger.Info("Trip retrieved successfully", zap.String("trip_id", trip.ID.String()))
	c.JSON(http.StatusOK, response)
}

// ListTrips godoc
// @Summary      List Trip Config
// @Description  Retrieves a paginated list of trips.
// @Tags         trips-config
// @Accept       json
// @Produce      json
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        per_page query     int  false  "Items per page"  default(10)
// @Success      200  {object}   dto.TripConfigResponse
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips-config [get]
func (h *TripConfigHandler) ListTripConfig(c *gin.Context) {
	page := 1
	perPage := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if pp := c.Query("per_page"); pp != "" {
		fmt.Sscanf(pp, "%d", &perPage)
	}

	inputDTO := input.ListTripConfigInput{
		Page:    page,
		PerPage: perPage,
	}

	outputDTO, err := h.listTripConfigUseCase.Execute(c.Request.Context(), inputDTO)
	if err != nil {
		logger.Error("Failed to list trips", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list trips"})
		return
	}
	trips := make([]dto.TripConfigListItemResponse, 0, len(outputDTO.Trips))
	for _, t := range outputDTO.Trips {
		trips = append(trips, dto.TripConfigListItemResponse{
			ID:             t.ID,
			ShipID:         t.ShipID,
			RouteID:        t.RouteID,
			ShipName:       t.ShipName,
			RouteName:      t.RouteName,
			Recurrence:     string(t.Recurrence),
			ExpirationDate: t.ExpirationDate,
			DepartureTime:  t.DepartureTime,
			ArrivalTime:    t.ArrivalTime,
			DurationDays:   t.DurationDays,
			StartDate:      t.StartDate,
		})
	}

	response := dto.ListTripConfigResponse{
		Trips:       trips,
		TotalCount:  outputDTO.TotalCount,
		TotalPages:  outputDTO.TotalPages,
		CurrentPage: outputDTO.CurrentPage,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTrip godoc
// @Summary      Delete a Trip Config
// @Description  Deletes a Trip Config by its ID.
// @Tags         trips-config
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Trip Config ID"
// @Success      200  {object}  dto.SuccessMessageResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid Trip Config ID"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips-config/{id} [delete]
func (h *TripConfigHandler) DeleteTripConfig(c *gin.Context) {
	tripID := c.Param("id")
	if tripID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Trip ID is required"})
		return
	}
	err := h.deleteTripConfigUseCase.Execute(c.Request.Context(), tripID)
	if err != nil {
		logger.Error("Failed to delete trip", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete trip"})
		return
	}
	logger.Info("Trip deleted successfully", zap.String("trip_id", tripID))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Trip deleted successfully"})
}
