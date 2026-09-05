package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/http/dto"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/common"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type ShipHandler struct {
	createShipUseCase input.CreateShipUseCase
	deleteShipUseCase input.DeleteShipUseCase
	updateShipUseCase input.UpdateShipUseCase
	listShipUseCase   input.ListShipUseCase
	getShipUseCase    input.GetShipUseCase
	createseatUseCase input.CreateSeatUseCase
}

func NewShipHandler(
	createShipUC input.CreateShipUseCase,
	deleteShipUC input.DeleteShipUseCase,
	updateShipUC input.UpdateShipUseCase,
	listShipUC input.ListShipUseCase,
	getShipUC input.GetShipUseCase,
	createSeatUC input.CreateSeatUseCase,
) *ShipHandler {
	return &ShipHandler{
		createShipUseCase: createShipUC,
		deleteShipUseCase: deleteShipUC,
		updateShipUseCase: updateShipUC,
		listShipUseCase:   listShipUC,
		getShipUseCase:    getShipUC,
		createseatUseCase: createSeatUC,
	}
}

// CreateShip godoc
// @Summary      Create a new ship
// @Description  Registers a new ship with the provided details.
// @Tags         ships
// @Accept       json
// @Produce      json
// @Param        organization   body      dto.CreateShipRequest  true  "SHIP Data to Create"
// @Success      201  {object}  dto.ShipResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships [post]
func (h *ShipHandler) CreateShip(c *gin.Context) {
	var req dto.CreateShipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var configs []uuid.UUID
	for _, id := range req.ConfigurationsID {
		configurationID, err := uuid.Parse(id)
		if err != nil {
			logger.Error("Invalid configuration ID", zap.Error(err))

		}
		configs = append(configs, configurationID)
	}

	useCaseInput := &input.CreateShipInput{
		TypeShip:          req.TypeShip,
		Name:              req.Name,
		IMO:               req.IMO,
		TotalSeats:        req.TotalSeats,
		TotalCabins:       req.TotalCabins,
		Status:            req.Status,
		PassengerCapacity: req.PassengerCapacity,
		WeightCapacity:    req.WeightCapacity,
		OrganizationID:    req.OrganizationID,
		ConfigurationsID:  configs,
	}

	err := h.createShipUseCase.Execute(c.Request.Context(), useCaseInput)
	if err != nil {
		logger.Error("Failed to create ship", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create ship"})
		return
	}

	logger.Info("Ship created", zap.String("ship_name", req.Name))
	c.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: "Ship created successfully"})
}

// DeleteShip godoc
// @Summary      Delete a ship
// @Description  Deletes a ship by its ID.
// @Tags         ships
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Ship ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse "Invalid ID supplied"
// @Failure      404  {object}  dto.ErrorResponse "Ship not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships/{id} [delete]
func (h *ShipHandler) DeleteShip(c *gin.Context) {
	shipIDStr := c.Param("id")
	shipUUID, err := uuid.Parse(shipIDStr)
	if err != nil {
		logger.Error(common.LogErrDeleteShip, zap.String("error", "ID is required"))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID is required"})
		return
	}
	err = h.deleteShipUseCase.Execute(c.Request.Context(), shipUUID)
	if err != nil {
		logger.Error(common.LogErrDeleteShip, zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateShip godoc
// @Summary      Update a ship
// @Description  Updates an existing ship with the provided details.
// @Tags         ships
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Ship ID"
// @Param        ship   body      dto.UpdateShipRequest  true  "Ship Data to Update"
// @Success      200  {object}    dto.ShipResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid request"
// @Failure      404  {object}  dto.ErrorResponse "SHIP not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships/{id} [put]
func (h *ShipHandler) UpdateShip(c *gin.Context) {
	shipIDStr := c.Param("id")
	shipUUID, err := uuid.Parse(shipIDStr)
	if err != nil {
		logger.Error(common.LogErrFindShip, zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var request dto.UpdateShipRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	useCaseInput := input.UpdateShipInput{
		TypeShip:          request.TypeShip,
		Name:              request.Name,
		IMO:               request.IMO,
		TotalSeats:        request.TotalSeats,
		TotalCabins:       request.TotalCabins,
		Status:            request.Status,
		PassengerCapacity: request.PassengerCapacity,
		WeightCapacity:    request.WeightCapacity,
		//ConfigurationsID:  request.ConfigurationsID,
		OrganizationID: request.OrganizationID,
	}

	err = h.updateShipUseCase.Execute(c.Request.Context(), &useCaseInput, shipUUID)
	if err != nil {
		logger.Error("Failed to update ship", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update ship"})
		return
	}

	logger.Info("Ship updated", zap.String("ship_id", shipIDStr))
	c.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: "ship updated successfully"})
}

// ListShips godoc
// @Summary      List Ships
// @Description  Retrieves a list of ships with pagination.
// @Tags         ships
// @Accept       json
// @Produce      json
// @Param        typeShip   query     string  false  "Ship type (BOAT, FERRYBOAT or SHIP)"
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        pageSize query     int  false  "Number of items per page"  default(10)
// @Success      200  {array}   dto.ListShipResponse
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships [get]
func (h *ShipHandler) ListShips(c *gin.Context) {
	page := 1
	perPage := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if pp := c.Query("per_page"); pp != "" {
		fmt.Sscanf(pp, "%d", &perPage)
	}

	inputDTO := input.ListShipInput{
		Page:    page,
		PerPage: perPage,
	}

	typeShip := c.Query("typeShip")

	result, err := h.listShipUseCase.Execute(c.Request.Context(), inputDTO, typeShip)
	if err != nil {
		logger.Error("Failed to list ships", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list ships"})
		return
	}
	ships := make([]dto.ShipListItemResponse, 0, len(result.Ships))
	for _, s := range result.Ships {
		ships = append(ships, dto.ShipListItemResponse{
			ID:                s.ID,
			ConfigurationID:   s.ConfigurationID,
			OrganizationID:    s.OrganizationID,
			IMO:               s.IMO,
			Name:              s.Name,
			Status:            string(s.Status),
			TypeShip:          string(s.TypeShip),
			TotalSeats:        s.TotalSeats,
			WeightCapacity:    s.WeightCapacity,
			PassengerCapacity: s.PassengerCapacity,
			ImageUrl:          s.ImageUrl,
		})
	}

	response := dto.ListShipResponse{
		Ships:       ships,
		TotalCount:  result.TotalCount,
		TotalPages:  result.TotalPages,
		CurrentPage: result.CurrentPage,
	}
	c.JSON(http.StatusOK, response)
}

// FindShipByID godoc
// @Summary      Get Ship by ID
// @Description  Retrieves a Ship by their ID.
// @Tags         ships
// @Accept       json
// @Produce      json
// @Param        typeShip   query     string  false  "Ship type (BOAT, FERRYBOAT or SHIP)"
// @Param        id   path      string  true  "Ship ID"
// @Success      200  {object}  dto.ShipResponse
// @Failure      400  {object}  dto.ErrorResponse "Invalid SHip ID"
// @Failure      404  {object}  dto.ErrorResponse "Ship not found"
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /ships/{id} [get]
func (h *ShipHandler) FindByID(c *gin.Context) {
	shipIDStr := c.Param("id")
	shipUUID, err := uuid.Parse(shipIDStr)
	if err != nil {
		logger.Error(common.LogErrFindShip, zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	typeShip := c.Query("typeShip")
	s, err := h.getShipUseCase.Execute(c.Request.Context(), shipUUID, typeShip)
	if err != nil {
		logger.Error(common.LogErrFindShip, zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	if s == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Ship not found"})
		return
	}
	response := dto.ShipResponse{
		ID:                s.ID,
		ConfigurationID:   s.ConfigurationID,
		OrganizationID:    s.OrganizationID,
		IMO:               s.IMO,
		Name:              s.Name,
		Status:            string(s.Status),
		TypeShip:          string(s.TypeShip),
		TotalSeats:        s.TotalSeats,
		WeightCapacity:    s.WeightCapacity,
		PassengerCapacity: s.PassengerCapacity,
		ImageUrl:          s.ImageUrl,
	}
	logger.Info("Ship found", zap.String("ship_id", s.ID.String()))
	c.JSON(http.StatusOK, response)
}
