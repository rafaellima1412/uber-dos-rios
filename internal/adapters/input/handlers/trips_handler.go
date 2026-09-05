package handlers

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http/dto"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type TripHandler struct {
	loadGraphTripUsecase        input.LoadGraphTripInput
	dijkstraTripUsecase         input.DijkstraTripInput
	dfsTripsUseCase             input.DFSTripsUseCase
	validatedTripUseCase        input.ValidatedTripsUseCase
	addTriptoReservationUseCase input.AddTripsToReservationUseCase
	listUnitAvailableUseCase    input.ListAvailableUseCase
	searchTripUseCase           input.TripSearchUseCase
}

func NewTripHandler(
	loadGraphTripUC input.LoadGraphTripInput,
	dijkstraTripUC input.DijkstraTripInput,
	dfsTripsUC input.DFSTripsUseCase,
	validatedTripUC input.ValidatedTripsUseCase,
	addTriptoReservationUC input.AddTripsToReservationUseCase,
	listUnitAvailableUC input.ListAvailableUseCase,
	searchTripUC input.TripSearchUseCase,

) *TripHandler {
	return &TripHandler{
		loadGraphTripUsecase:        loadGraphTripUC,
		dijkstraTripUsecase:         dijkstraTripUC,
		dfsTripsUseCase:             dfsTripsUC,
		validatedTripUseCase:        validatedTripUC,
		addTriptoReservationUseCase: addTriptoReservationUC,
		listUnitAvailableUseCase:    listUnitAvailableUC,
		searchTripUseCase:           searchTripUC,
	}
}

// ValidateTrips godoc
// @Summary      Validate trip stops sequence and find the cheapest path
// @Description  Validates if the route has all required cyties between origin and destination
// @Tags         trips
// @Accept       json
// @Produce      json
// @Param        trip body dto.TripConfigSearchRequest true "Trip search"
// @Success      200 {array} dto.TripConfigSearchResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /trips/catalogo/menor-preco [post]
func (h *TripHandler) ValidateConnectionsTrips(c *gin.Context) {
	var req dto.TripConfigSearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	if req.CityOriginID == req.CityDestID {
		logger.Error("Origin and destination must be different")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "origin and destination must be different"})
		return
	}

	graph, err := h.loadGraphTripUsecase.Execute(c.Request.Context(), req.DateOrigin, req.DateDest)
	if err != nil {
		logger.Error("Failed to validate trips", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to validate trips " + err.Error()})
		return
	}

	path, ok := h.dijkstraTripUsecase.Execute(graph, req.CityOriginID, req.CityDestID)
	if !ok {
		logger.Error("No valid path found between origin and destination")
		c.JSON(404, dto.ErrorResponse{Error: "no valid path found between origin and destination"})
		return
	}
	out := make([]dto.PathStep, 0, len(path.Path))

	for _, p := range path.Path {
		out = append(out, dto.PathStep{
			OrganizationID:    p.OrganizationID,
			CityID:            p.CityID,
			Cost:              p.Cost,
			RouteID:           p.RouteID,
			RouteName:         p.RouteName,
			DateDepartureTime: p.DateDepartureTime,
			DateArrivalTime:   p.DateArrivalTime,
		})
	}
	resp := dto.TripConfigSearchResponse{
		TotalCost: path.TotalCost,
		Path:      out,
	}

	logger.Info("Calculated cheapest path successfully")
	c.JSON(http.StatusOK, resp)
}

// ValidateTrips godoc
// @Summary      Validate trip stops sequence
// @Description  Validates if the route has all required cyties between origin and destination
// @Tags         trips
// @Accept       json
// @Produce      json
// @Param        trip body dto.TripConfigSearchRequest true "Trip search data"
// @Success      200 {array} dto.TripConfigResult
// @Failure      400 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /trips/catalogo [post]
func (h *TripHandler) DFSTrips(c *gin.Context) {
	var req dto.TripConfigSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	if req.CityOriginID == req.CityDestID {
		logger.Error("Origin and destination must be different")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "origin and destination must be different"})
		return
	}

	graph, err := h.loadGraphTripUsecase.Execute(c.Request.Context(), req.DateOrigin, req.DateDest)
	if err != nil {
		logger.Error("Failed to validate trips", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to validate trips " + err.Error()})
		return
	}
	origin, err := time.Parse(time.RFC3339, req.DateOrigin)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Formato de data inválido. Use YYYY-MM-DDTHH:MM:SS±HH:MM",
		})
		return
	}
	dest, err := time.Parse(time.RFC3339, req.DateDest)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Formato de data inválido. Use YYYY-MM-DDTHH:MM:SS±HH:MM",
		})
		return
	}

	input := input.TripSearchInput{
		CityOriginID:         req.CityOriginID,
		CityDestID:           req.CityDestID,
		DateDepartureTime:    origin,
		DateArrivalTime:      dest,
		MaxStops:             req.MaxStops,
		MaxCost:              req.MaxCost,
		ShipID:               req.ShipID,
		PassangerCountSeats:  req.PassengerCountSeats,
		PassangerCountCabins: req.PassengerCountSeats,
		MaxResults:           req.MaxResults,
	}
	// TODO remover as  regras do handler
	for node := range graph {
		sort.Slice(graph[node], func(i, j int) bool {
			return graph[node][i].Cost < graph[node][j].Cost
		})
	}

	paths, err := h.dfsTripsUseCase.Execute(graph, input)
	resp := make([]dto.TripConfigSearchResponse, 0, len(paths))
	for _, p := range paths {
		isRouteAvailable := true
		out := make([]dto.PathStep, 0, len(p.Path))

		for i := 1; i < len(p.Path); i++ {
			step := p.Path[i]

			availability, err := h.listUnitAvailableUseCase.Execute( // Lotação
				c.Request.Context(),
				step.ShipID,
				req.DateOrigin,
				req.DateDest,
			)

			//erro real de conexão/banco
			if err != nil {
				// Se o erro for "Registro não encontrado", tratar como barco vazio
				if err.Error() == "sql: no rows in result set" || strings.Contains(err.Error(), "no rows") {
					//logger.Info("Log de debug: Nenhuma viagem hoje, barco 100% livre")
					isRouteAvailable = true
				} else {
					// Erro real de sistema
					isRouteAvailable = false
					break
				}
			} else {
				// Se o registro existe, validamos a ocupação
				if (req.PassengerCountSeats > 0 && availability.Seats.AvailableSeats < req.PassengerCountSeats) ||
					(req.PassengerCountCabins > 0 && availability.Cabins.AvailableCabins < req.PassengerCountCabins) {
					isRouteAvailable = false
					break
				}
				isRouteAvailable = true
			}

			out = append(out, dto.PathStep{
				OrganizationID:      step.OrganizationID,
				CityID:              step.CityID,
				Cost:                step.Cost,
				RouteID:             step.RouteID,
				RouteName:           step.RouteName,
				DateDepartureTime:   step.DateDepartureTime,
				DateArrivalTime:     step.DateArrivalTime,
				ShipID:              step.ShipID,
				ShipName:            step.ShipName,
				ShipURL:             step.ShipURL,
				TripConfigurationID: step.TripConfigurationID,
			})
		}
		if isRouteAvailable {
			resp = append(resp, dto.TripConfigSearchResponse{
				TotalCost: p.TotalCost,
				Path:      out,
			})
		}
	}
	if err != nil || len(paths) == 0 {
		c.JSON(404, gin.H{"error": "no routes found"})
		return
	}

	c.JSON(200, resp)
}

// CreateInstance godoc
// @Summary Validated and Create a new trip for reservation
// @Description Validated and Create a new trip for reservation
// @Tags trips
// @Accept json
// @Produce json
// @Param terminal body dto.ReservationTripsSelectedRequest true "Trips Instance Details"
// @Success 201 {object} dto.SuccessMessageResponse
// @Failure 400  {object}  dto.ErrorResponse
// @Failure 500  {object}  dto.ErrorResponse
// @Router /trips [post]
func (h *TripHandler) CreateConfig(c *gin.Context) {
	var req dto.ReservationTripsSelectedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	reservationID, err := uuid.Parse(req.ReservationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	inputUC := input.TripToReservationInput{}
	for _, trip := range req.SelectedTrips {
		tripConfigurationID, err := uuid.Parse(trip.TripConfigurationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
		shipID, err := uuid.Parse(trip.ShipID)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
		routeID, err := uuid.Parse(trip.RouteID)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		departureAt, err := time.Parse(time.RFC3339, trip.DepartureAt)
		if err != nil {
			logger.Error("Erro ao converter data da saida", zap.Error(err))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error: "Formato de data inválido. Use ISO 8601 (ex: 2025-10-22T14:30:00Z)",
			})
			return
		}
		arrivalAt, err := time.Parse(time.RFC3339, trip.ArrivalAt)
		if err != nil {
			logger.Error("Erro ao converter data da chegada", zap.Error(err))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error: "Formato de data inválido. Use ISO 8601 (ex: 2025-10-22T14:30:00Z)",
			})
			return
		}
		inputUC.SelectedTrips = append(inputUC.SelectedTrips, input.SelectedTrips{
			TripConfigurationID: tripConfigurationID,
			ShipID:              shipID,
			RouteID:             routeID,
			DepartureAt:         departureAt,
			ArrivalAt:           arrivalAt,
			OccupiedUnits:       trip.OccupiedUnits,
		})
	}
	inputUC.ReservationID = reservationID
	if err := h.addTriptoReservationUseCase.Execute(c.Request.Context(), &inputUC); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Trips Created Successfully"})

}

// FilterTrips godoc
// @Summary      List Trips
// @Description  Retrieves a list of trips with optional filters for routes, ships, and dates.
// @Tags         trips
// @Accept       json
// @Produce      json
// @Param        route_id               query    string  false  "Route UUID"
// @Param        ship_id                query    string  false  "Ship UUID"
// @Param        trip_configurations_id query    string  false  "Trip Config UUID"
// @Param        departure_after        query    string  false  "Format: RFC3339 (e.g. 2026-10-21T01:00:00Z)"
// @Param        departure_before       query    string  false  "Format: RFC3339"
// @Param        page                   query    int     false  "Page number" default(1)
// @Param        limit                  query    int     false  "Number of items per page" default(10)
// @Success      200  {array}   dto.TripResponseDTO
// @Failure      500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /trips/search [get]
func (h *TripHandler) SearchTrips(c *gin.Context) {
	var filters dto.TripFiltersDTO

	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(400, gin.H{"error": "Parâmetros inválidos", "details": err.Error()})
		return
	}
	searchTripInput := input.SearchTripInput{
		Page:    filters.Page,
		PerPage: filters.Limit,
	}

	useCaseInput := input.TripFilterInput{
		RouteID:              filters.RouteID,
		ShipID:               filters.ShipID,
		DepartureAfter:       filters.DepartureAfter,
		DepartureBefore:      filters.DepartureBefore,
	}

	result, err := h.searchTripUseCase.Execute(c.Request.Context(), &useCaseInput, searchTripInput)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar viagens", "details": err.Error()})
		return
	}

	tripsDTO := make([]dto.TripResponseDTO, 0, len(result.Trips))

	for _, trip := range result.Trips {
		tripsDTO = append(tripsDTO, dto.TripResponseDTO{
			TripConfigurationID: trip.TripConfigurationID.String(),
			ShipID:              trip.ShipID.String(),
			RouteID:             trip.RouteID.String(),
			DepartureAt:         trip.DepartureAt.Format(time.RFC3339),
			ArrivalAt:           trip.ArrivalAt.Format(time.RFC3339),
		})
	}

	finalResponse := dto.TripFilterListResponse{
		Trips:       tripsDTO,
		TotalCount:  result.TotalCount,
		TotalPages:  result.TotalPages,
		CurrentPage: result.CurrentPage,
	}

	c.JSON(200, finalResponse)

}
