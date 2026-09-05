package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http/dto"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type ReservationHandler struct {
	createReservationUsecase     input.CreateReservationUseCase
	updateReservationUsecase     input.UpdateReservationUseCase
	getReservationUseCase        input.GetReservationUseCase
	listReservationUseCase       input.ListReservationUseCase
	deleteReservationUseCase     input.DeleteReservationUseCase
	validatedTripUseCase         input.ValidatedTripsUseCase
	addTripsToReservationUseCase input.AddTripsToReservationUseCase
}

func NewReservationHandler(
	createReservationUsecase input.CreateReservationUseCase,
	updateReservationUsecase input.UpdateReservationUseCase,
	getReservationUseCase input.GetReservationUseCase,
	listReservationUseCase input.ListReservationUseCase,
	deleteReservationUseCase input.DeleteReservationUseCase,
	validatedTripUseCase input.ValidatedTripsUseCase,
	addTripsToReservationUseCase input.AddTripsToReservationUseCase,

) *ReservationHandler {
	return &ReservationHandler{
		createReservationUsecase:     createReservationUsecase,
		updateReservationUsecase:     updateReservationUsecase,
		getReservationUseCase:        getReservationUseCase,
		listReservationUseCase:       listReservationUseCase,
		deleteReservationUseCase:     deleteReservationUseCase,
		validatedTripUseCase:         validatedTripUseCase,
		addTripsToReservationUseCase: addTripsToReservationUseCase,
	}
}

// CreateReservation godoc
// @Summary Create a new reservation
// @Description Create a new reservation with the provided details
// @Tags reservations
// @Accept json
// @Produce json
// @Param terminal body dto.CreateReservationRequest true "Reservation details"
// @Success 201 {object} dto.SuccessMessageResponse "Reservation created successfully"
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router /reservations [post]
func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var req dto.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	tReservation, err := time.Parse(time.RFC3339, req.ReservationDate)
	if err != nil {
		logger.Error("Erro ao converter data da reserva", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Formato de data inválido. Use ISO 8601 (ex: 2025-10-22T14:30:00Z)",
		})
		return
	}

	ID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	reservationInput := input.CreateReservationInput{
		UserID:          ID,
		ReservationDate: &tReservation,
	}

	resID, err := h.createReservationUsecase.Execute(c.Request.Context(), &reservationInput)
	if err != nil {
		logger.Error("Error executing CreateReservation use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               resID,
	})
}

// UpdateReservation godoc
// @Summary Update a reservation
// @Description Update an existing reservation with the provided details
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Param terminal body dto.UpdateReservationRequest true "Updated reservation details"
// @Success 200 {object} dto.SuccessMessageResponse "Reservation updated successfully"
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      404  {object}  dto.ErrorResponse "SHIP not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /reservations/{id} [put]
func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	reservationStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Reservation ID is required"})
		return
	}

	var req dto.UpdateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}
	t, err := time.Parse(time.RFC3339, req.ReservationDate)
	if err != nil {
		logger.Error("Erro ao converter data da reserva", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Formato de data inválido. Use ISO 8601 (ex: 2025-10-22T14:30:00Z)"})
		return
	}
	reservationInput := input.UpdateReservationInput{
		UserID: req.UserID,
		// TripConfigID:    req.TripConfigurationIDs,
		ReservationDate: &t,
		Status:          req.Status,
	}
	err = h.updateReservationUsecase.Execute(c.Request.Context(), &reservationInput, reservationID)
	if err != nil {
		logger.Error("Error executing UpdateReservation use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update reservation"})
		return
	}
	//atualizar acentos ocupados
	logger.Info("Reservation updated successfully", zap.String("reservation_id", reservationStr))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Reservation updated successfully"})
}

// FindReservationByID godoc
// @Summary Get Reservation by ID
// @Description Retrieves a Reservation by its ID.
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} dto.ReservationListItemResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} dto.ErrorResponse "Reservation not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router /reservations/{id} [get]
func (h *ReservationHandler) FindByID(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservation, err := h.getReservationUseCase.Execute(c.Request.Context(), reservationIDStr)
	if err != nil {
		logger.Error("Error executing GetReservation use case", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}
	if reservation == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Reservation not found"})
		return
	}
	response := dto.ReservationListItemResponse{
		ID:     reservation.ID,
		UserID: reservation.UserID,
		// TripConfigurationIDs: reservation.TripID,
		ReservationDate: reservation.ReservationDate,
		Status:          string(reservation.Status),
	}
	logger.Info("Reservation found", zap.String("reservation_id", reservationIDStr))
	c.JSON(http.StatusOK, response)

}

// ListReservations godoc
// @Summary List Reservations
// @Description Retrieves a paginated list of reservations.
// @Tags reservations
// @Accept json
// @Produce json
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        pageSize query     int  false  "Number of items per page"  default(10)
//
//	@Success      200  {array}   dto.ListReservationResponse
//
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /reservations [get]
func (h *ReservationHandler) ListReservations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	inputDTO := input.ListReservationInput{
		Page:    page,
		PerPage: perPage,
	}

	result, err := h.listReservationUseCase.Execute(c.Request.Context(), inputDTO)
	if err != nil {
		logger.Error("Failed to list reservations", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to list reservations",
		})
		return
	}

	reservations := make([]dto.ReservationListItemResponse, 0, len(result.Reservations))

	for _, r := range result.Reservations {
		var reservationDate string
		if r.ReservationDate != nil {
			reservationDate = r.ReservationDate.Format(time.RFC3339)
		}

		reservations = append(reservations, dto.ReservationListItemResponse{
			ID:              r.ID,
			UserID:          r.UserID,
			ReservationDate: reservationDate,
			Status:          string(r.Status),
			CreatedAt:       r.CreatedAt,
			UpdatedAt:       r.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, dto.ListReservationResponse{
		Reservations: reservations,
		TotalCount:   result.TotalCount,
		TotalPages:   result.TotalPages,
		CurrentPage:  result.CurrentPage,
	})
}

// DeleteReservation godoc
// @Summary Delete a reservation
// @Description Delete a reservation by its ID.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Ship ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse "Invalid ID supplied"
// @Failure      404  {object}  dto.ErrorResponse "Reservation not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /reservations/{id} [delete]
func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationUUID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		logger.Error("Invalid ID supplied", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID is required"})
		return
	}
	err = h.deleteReservationUseCase.Execute(c.Request.Context(), reservationUUID)
	if err != nil {
		logger.Error("Error executing DeleteReservation use case", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
