package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http/dto"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"

	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type ShipConfigHandler struct {
	getSeatUseCase           input.GetSeatUseCase
	createSeatUseCase        input.CreateSeatUseCase
	updateSeatUseCase        input.UpdateSeatUseCase
	listSeatUseCase          input.ListSeatUseCase
	searchSeatUseCase        input.SearchSeatUseCase
	listSeatAvailableUseCase input.ListAvailableUseCase
	createCabinUsecase       input.CreateCabinUseCase
	listCabinUseCase         input.ListShipCabinUseCase
	deleteSeatUseCase        input.DeleteSeatUseCase
	updateCabinUseCase       input.UpdateCabinUseCase
	deleteCabinUseCase       input.DeleteCabinUseCase
}

func NewSeatHandler(
	getSeatUC input.GetSeatUseCase,
	createSeatUC input.CreateSeatUseCase,
	updateSeatUC input.UpdateSeatUseCase,
	listSeatUC input.ListSeatUseCase,
	searchSeatUC input.SearchSeatUseCase,
	createCabinUC input.CreateCabinUseCase,
	listCabinUC input.ListShipCabinUseCase,
	deleteSeatUC input.DeleteSeatUseCase,
	updateCabinUC input.UpdateCabinUseCase,
	deleteCabinUC input.DeleteCabinUseCase,
	listSeatAvailableUC input.ListAvailableUseCase,
) *ShipConfigHandler {
	return &ShipConfigHandler{
		getSeatUseCase:           getSeatUC,
		createSeatUseCase:        createSeatUC,
		updateSeatUseCase:        updateSeatUC,
		listSeatUseCase:          listSeatUC,
		searchSeatUseCase:        searchSeatUC,
		createCabinUsecase:       createCabinUC,
		listCabinUseCase:         listCabinUC,
		deleteSeatUseCase:        deleteSeatUC,
		updateCabinUseCase:       updateCabinUC,
		deleteCabinUseCase:       deleteCabinUC,
		listSeatAvailableUseCase: listSeatAvailableUC,
	}
}

// CreateSeat godoc
// @Summary      Create a new seat
// @Description  Registers a new seat with the provided details.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param   organization   body dto.CreateSeatRequest true "Seat Data to Create"
// @Success      201  {object}  dto.SeatDetailResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/seats [post]
func (h *ShipConfigHandler) CreateSeat(c *gin.Context) {
	var req dto.CreateSeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}
	useCaseInput := input.CreateSeatInput{

		IsActive:          req.IsActive,
		LeftColumnsCount:  req.LeftColumnsCount,
		RightColumnsCount: req.RightColumnsCount,
		Prefix:            req.Prefix,
	}
	_, err := h.createSeatUseCase.Execute(c.Request.Context(), &useCaseInput)
	if err != nil {
		logger.Error("Failed to create seat", zap.Error(err))
		c.JSON(http.StatusCreated, dto.SeatDetailResponse{ID: "Failed to create seat"})
		return
	}

	logger.Info("Seat created")
	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Seat created successfully"})
}

// UpdateSeat godoc
// @Summary      Update a seat
// @Description  Updates an existing seat with the provided details.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        seat   body      dto.UpdateSeatRequest  true  "Seat Data to Update"
// @Success      201  {object}  dto.SeatDetailResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/seats/{id} [put]
func (h *ShipConfigHandler) UpdateSeat(c *gin.Context) {
	seatIDStr := c.Param("id")
	seatUUID, err := uuid.Parse(seatIDStr)
	if err != nil {
		logger.Error("Invalid seat ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var request dto.UpdateSeatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	useCaseInput := input.UpdateSeatInput{
		IsActive:          request.IsActive,
		LeftColumnsCount:  request.LeftColumnsCount,
		RightColumnsCount: request.RightColumnsCount,
		Prefix:            request.Prefix,
	}

	err = h.updateSeatUseCase.Execute(c.Request.Context(), &useCaseInput, seatUUID)
	if err != nil {
		logger.Error("Failed to update seat", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	logger.Info("Seat updated", zap.String("seat_id", seatIDStr))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "ship updated successfully"})
}

// LIstSeat godoc
// @Summary      List seats
// @Description  Retrieves a seats.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        pageSize query     int  false  "Number of items per page"  default(10)
// @Success      200  {object}  dto.SeatDetailResponse
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/seats [get]
func (h *ShipConfigHandler) ListSeat(c *gin.Context) {
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
	inputDTO := input.ListSeatInput{
		Page:    page,
		PerPage: perPage,
	}
	output, err := h.listSeatUseCase.Execute(c.Request.Context(), inputDTO)
	if err != nil {
		logger.Error("Failed to list seats", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list seats"})
		return
	}
	var seatResponses []dto.SeatDetailResponse
	for _, seat := range output.Seats {
		seatResponses = append(seatResponses, dto.SeatDetailResponse{
			ID:                seat.ID.String(),
			IsActive:          seat.IsActive,
			LeftColumnsCount:  seat.LeftColumnsCount,
			RightColumnsCount: seat.RightColumnsCount,
			Prefix:            seat.Prefix,
		})
	}
	c.JSON(http.StatusOK, dto.SeatListResponse{
		Seats:       seatResponses,
		TotalCount:  output.TotalCount,
		TotalPages:  output.TotalPages,
		CurrentPage: output.CurrentPage,
	})

}

// SearchSeat godoc
// @Summary      Search Available Seats
// @Description  Searches for seats based on available date range.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        dateorigin query string true "Data de Início (YYYY-MM-DD)"
// @Param        datedest   query string true "Data de Fim (YYYY-MM-DD)"
// @Param        shipid    query string true "ID Barco: uuid"
// @Success      200  {array}   dto.ListUnitAvailableResponse
// @Router       /ships-config/units/search [get]
func (h *ShipConfigHandler) SearchSeat(c *gin.Context) {

	dateOrigin := c.Query("dateorigin")
	dateDest := c.Query("datedest")
	ID := c.Query("shipid")

	shipID, err := uuid.Parse(ID)
	if err != nil {
		return
	}

	output, err := h.listSeatAvailableUseCase.Execute(c.Request.Context(), shipID, dateOrigin, dateDest)
	if err != nil {
		if err.Error() == "sql: no rows in result set" || strings.Contains(err.Error(), "no rows") {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "No seats found"})
			return
		}
		logger.Error("Failed to search seats", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Internal server error"})
		return
	}
	seats := dto.Seats{
		TotalSeats:     output.Seats.TotalSeats,
		OccupiedUnits:  output.Seats.OccupiedUnits,
		AvailableSeats: output.Seats.AvailableSeats,
	}
	cabins := dto.Cabins{
		TotalCabins:     output.Cabins.TotalCabins,
		OccupiedUnits:   output.Cabins.OccupiedUnits,
		AvailableCabins: output.Cabins.AvailableCabins,
	}

	ListResponse := dto.ListUnitAvailableResponse{
		ShipID:   output.ShipID.String(),
		ShipName: output.ShipName,
		Seats:    seats,
		Cabins:   cabins,
	}

	c.JSON(http.StatusOK, ListResponse)
}

// DeleteSeat godoc
// @Summary      Delete seat
// @Description  Delete a seat by ID
// @Tags         seats
// @Param        id   path      string  true  "Seat ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse "Invalid ID supplied"
// @Failure      404  {object}  dto.ErrorResponse "Seat not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/seats/{id} [delete]
func (h *ShipConfigHandler) DeleteSeat(c *gin.Context) {
	seatIDStr := c.Param("id")
	seatUUID, err := uuid.Parse(seatIDStr)
	if err != nil {
		logger.Error("Invalid seat ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	err = h.deleteSeatUseCase.Execute(c.Request.Context(), seatUUID)
	if err != nil {
		logger.Error("Failed to delete seat", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// CreateCabin godoc
// @Summary      Create a new cabin
// @Description  Registers a new cabin with the provided details.
// @Tags         cabins
// @Accept       json
// @Produce      json
// @Param   organization   body dto.CreateCabinRequest true "Cabin Data to Create"
// @Success      201  {object}  dto.CabinDetailResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/cabins [post]
func (h *ShipConfigHandler) CreateCabin(c *gin.Context) {
	var req dto.CreateCabinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	cabins := make([]input.Cabin, len(req.Cabins))
	for i, c := range req.Cabins {
		cabins[i] = input.Cabin{
			Name:        c.Name,
			BedType:     c.BedType,
			Capacity:    c.Capacity,
			Description: c.Description,
		}
	}
	useCaseInput := input.CreateCabinInput{
		Cabins: cabins,
	}

	err := h.createCabinUsecase.Execute(c.Request.Context(), &useCaseInput)
	if err != nil {
		logger.Error("Failed to create cabin", zap.Error(err))
		c.JSON(http.StatusCreated, dto.CabinDetailResponse{ID: "Failed to create cabin"})
		return
	}

	logger.Info("Cabin created")
	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Cabin created successfully"})
}

// ListCabin godoc
// @Summary      List Cabins
// @Description  Retrieves a cabins.
// @Tags				 cabins
// @Accept       json
// @Produce      json
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        pageSize query     int  false  "Number of items per page"  default(10)
// @Success      200	{object}  dto.CabinDetailResponse
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/cabins [get]
func (h *ShipConfigHandler) ListCabin(c *gin.Context) {
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

	inputDTO := input.ListCabinInput{
		Page:    page,
		PerPage: perPage,
	}
	output, err := h.listCabinUseCase.Execute(c.Request.Context(), inputDTO)
	if err != nil {
		logger.Error("Failed to list cabins", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list cabins"})
		return
	}
	var cabinResponses []dto.CabinDetailResponse
	for _, sc := range output.Cabins {
		cabinResponses = append(cabinResponses, dto.CabinDetailResponse{
			ID:          sc.ID.String(),
			Name:        sc.Name,
			BedType:     sc.BedType,
			Capacity:    sc.Capacity,
			Description: sc.Description,
		})
	}
	c.JSON(http.StatusOK, dto.CabinListItemResponse{
		Cabins:      cabinResponses,
		TotalCount:  output.TotalCount,
		TotalPages:  output.TotalPages,
		CurrentPage: output.CurrentPage,
	})
}

// DeleteCabin godoc
// @Summary      Delete cabin
// @Description  Delete a cabin by ID
// @Tags         cabins
// @Param        id   path      string  true  "Cabin ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse "Invalid ID supplied"
// @Failure      404  {object}  dto.ErrorResponse "Seat not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/cabins/{id} [delete]
func (h *ShipConfigHandler) DeleteCabin(c *gin.Context) {
	cabinIDStr := c.Param("id")
	cabinUUID, err := uuid.Parse(cabinIDStr)
	if err != nil {
		logger.Error("Invalid cabin ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	err = h.deleteCabinUseCase.Execute(c.Request.Context(), cabinUUID)
	if err != nil {
		logger.Error("Failed to delete cabin", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateCabin godoc
// @Summary      Update a cabin
// @Description  Updates an existing cabin with the provided details.
// @Tags         cabins
// @Accept       json
// @Produce      json
// @Param        cabin  body dto.UpdateCabinRequest true "Cabin Data to Update"
// @Success      201  {object}  dto.SeatDetailResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships-config/cabins/{id} [put]
func (h *ShipConfigHandler) UpdateCabin(c *gin.Context) {
	cabinIDStr := c.Param("id")
	cabinUUID, err := uuid.Parse(cabinIDStr)
	if err != nil {
		logger.Error("Invalid cabin ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var request dto.UpdateCabinRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	useCaseInput := input.UpdateCabinInput{
		Name:        request.Name,
		BedType:     request.BedType,
		Capacity:    request.Capacity,
		Description: request.Description,
	}
	err = h.updateCabinUseCase.Execute(c.Request.Context(), &useCaseInput, cabinUUID)
	if err != nil {
		logger.Error("Failed to update cabin", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	logger.Info("Cabin updated", zap.String("cabin_id", cabinIDStr))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "Cabin updated successfully"})
}
