package main

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	projectDocs "github.com/rafaellima1412/uber-dos-rios/docs"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/handlers"
	httpadapter "github.com/rafaellima1412/uber-dos-rios/internal/adapters/input/http"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/output/repository/postgres"
	"github.com/rafaellima1412/uber-dos-rios/internal/adapters/output/service"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/usecase"

	"github.com/rafaellima1412/uber-dos-rios/internal/config"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

// @title           Uber dos Rios API
// @version         v1.0.21
// @description     Service for managing user identities, roles, and organizations.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8002
// @BasePath        /api/v1
// @schemes         http https
func main() {
	ctx := context.Background()

	logger.InitLogger(true)
	defer logger.Log.Sync()

	logger.Info("Initializing application...")

	cfg := loadConfig()
	setupSwaggerHost(cfg)

	dbPool := setupDatabase(ctx, cfg)
	defer dbPool.Close()

	handlers := setupHandlers(dbPool)

	router := httpadapter.SetupRouter(handlers)

	logger.Info("Starting server", zap.String("port", cfg.Server.Port))
	if err := router.Run(cfg.Server.Port); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
	}
}

func loadConfig() *config.Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.local.yaml"
	}

	return config.LoadConfig(configPath)
}

func setupSwaggerHost(cfg *config.Config) {
	projectDocs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	if projectDocs.SwaggerInfo.Host == "" {
		projectDocs.SwaggerInfo.Host = cfg.Server.Host
	}
}

func setupDatabase(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	return postgres.NewConnection(ctx, cfg)
}

func setupHandlers(dbPool *pgxpool.Pool) *handlers.Handlers {
	// healt
	healthCheckHandler := handlers.NewHealthHandler()
	// gps
	gpsRepository := postgres.NewGPSRepository(dbPool)
	gpsService := service.NewGPSService(gpsRepository)
	gpsHandler := handlers.NewGPSHandler(gpsService)
	// config do barco
	configRepository := postgres.NewSeatRepository(dbPool)
	getSeatUseCase := usecase.NewGetSeatUseCase(configRepository)
	createSeatUseCase := usecase.NewCreateSeatUseCase(configRepository)
	updateSeatUseCase := usecase.NewUpdateSeatUseCase(configRepository)
	listSeatUseCase := usecase.NewListSeatUseCase(configRepository)
	createCabinUseCase := usecase.NewCreateCabinUsecase(configRepository)
	listCabinUseCase := usecase.NewListCabinUseCase(configRepository)
	deleteSeatUseCase := usecase.NewDeleteSeatUseCase(configRepository)
	updateCabinUseCase := usecase.NewUpdateCabinUseCase(configRepository)
	deleteCabinUseCase := usecase.NewDeleteCabinUsecase(configRepository)
	listAvailableUseCase := usecase.NewListAvailableUseCase(configRepository)
	shipConfigHandler := handlers.NewSeatHandler(
		getSeatUseCase,
		createSeatUseCase,
		updateSeatUseCase,
		listSeatUseCase,
		nil,
		createCabinUseCase,
		listCabinUseCase,
		deleteSeatUseCase,
		updateCabinUseCase,
		deleteCabinUseCase,
		listAvailableUseCase,
	)
	// Ship
	shipRepository := postgres.NewShipRepository(dbPool)
	createShipUsecase := usecase.NewCreateShipUseCase(shipRepository)
	deleShipUseCase := usecase.NewDeleteShipUseCase(shipRepository)
	updateShipUseCase := usecase.NewUpdateShipUseCase(shipRepository)
	listShipUsecase := usecase.NewListShipUseCase(shipRepository)
	getShipUsecase := usecase.NewGetShipUseCase(shipRepository)
	shipHandler := handlers.NewShipHandler(
		createShipUsecase,
		deleShipUseCase,
		updateShipUseCase,
		listShipUsecase,
		getShipUsecase,
		createSeatUseCase,
	)
	// terminal
	terminalRepository := postgres.NewTerminalRepositoryImpl(dbPool)
	createTerminalUseCase := usecase.NewCreateTerminal(terminalRepository)
	getTerminalUseCase := usecase.NewGetTerminal(terminalRepository)
	updateTerminalUseCase := usecase.NewUpdateTerminal(terminalRepository)
	deleteTerminalUseCase := usecase.NewDeleteTerminalUseCase(terminalRepository)
	listTerminalUseCase := usecase.NewListTerminalUseCase(terminalRepository)
	searchTerminalUseCase := usecase.NewSearchTerminalUseCase(terminalRepository)
	listCitiesUseCase := usecase.NewListCitiesUseCase(terminalRepository)
	terminalHandler := handlers.NewTerminalHandler(
		createTerminalUseCase,
		getTerminalUseCase,
		updateTerminalUseCase,
		deleteTerminalUseCase,
		listTerminalUseCase,
		searchTerminalUseCase,
		listCitiesUseCase,
	)
	// route
	routeRepository := postgres.NewRouteRepositoryImpl(dbPool)
	createRouteUseCase := usecase.NewCreateRouteUseCase(routeRepository)
	tokenManager := &service.TokenManager{
		ClientID:     os.Getenv("ORG_SERVICE_CLIENT_ID"),
		ClientSecret: os.Getenv("ORG_SERVICE_CLIENT_SECRET"),
		AuthURL:      os.Getenv("ORG_SERVICE_AUTH_URL"), // Ex: "http://localhost:8000/api/v1"
		ExpiresAt:    time.Now(),                        // Inicializa para forçar obtenção do token na primeira chamada
	}
	serviceOrganization := service.NewOrganizationService(tokenManager)
	getRouteUseCase := usecase.NewGetRouteUseCase(routeRepository, serviceOrganization)
	updateRouteUseCase := usecase.NewUpdateRouteUseCase(routeRepository)
	deleteRouteUseCase := usecase.NewDeleteRouteUseCase(routeRepository)
	listRouteUseCase := usecase.NewListRouteUseCase(routeRepository, serviceOrganization)
	searchRouteUseCase := usecase.NewSearchRouteUseCase(routeRepository, serviceOrganization)
	routeHandler := handlers.NewRouteHandler(
		createRouteUseCase,
		getRouteUseCase,
		updateRouteUseCase,
		deleteRouteUseCase,
		listRouteUseCase,
		searchRouteUseCase,
	)
	// schedule
	scheduleRepository := postgres.NewScheduleRepositoryImpl(dbPool)
	createScheduleUseCase := usecase.NewCreateScheduleUseCase(
		scheduleRepository,
		terminalRepository,
		routeRepository,
	)
	getScheduleUseCase := usecase.NewGetScheduleUseCase(scheduleRepository, terminalRepository, routeRepository)
	updateScheduleUseCase := usecase.NewUpdateScheduleUseCase(
		scheduleRepository,
		terminalRepository,
		routeRepository,
	)
	deleteScheduleUseCase := usecase.NewDeleteScheduleUseCase(scheduleRepository)
	listScheduleUseCase := usecase.NewListScheduleUseCase(
		scheduleRepository,
		routeRepository,
		terminalRepository,
	)
	scheduleHandler := handlers.NewScheduleHandler(
		createScheduleUseCase,
		getScheduleUseCase,
		updateScheduleUseCase,
		deleteScheduleUseCase,
		listScheduleUseCase,
		nil,
	)
	// trip-config
	tripConfigRepository := postgres.NewTripConfigRepositoryImpl(dbPool)
	CreateTripConfigUseCase := usecase.NewCreateTripConfigUseCase(tripConfigRepository, shipRepository, routeRepository)
	UpdateTripUseCase := usecase.NewUpdateTripUseCase(tripConfigRepository)
	GetTripConfigUseCase := usecase.NewGetTripUseCase(tripConfigRepository)
	ListTripUsecase := usecase.NewListTripConfigUseCase(tripConfigRepository)
	DeleteTripUsecase := usecase.NewDeleteTripUseCase(tripConfigRepository)
	tripConfigHandler := handlers.NewTripConfigHandler(
		CreateTripConfigUseCase,
		DeleteTripUsecase,
		UpdateTripUseCase,
		ListTripUsecase,
		GetTripConfigUseCase,
	)
	// trip
	tripRepository := postgres.NewTripRepositoryImpl(dbPool)
	filterTripsUseCase := usecase.NewFilterTripUseCase(tripRepository)

	// Reservation
	db := postgres.NewReservationRepository(dbPool)
	LoadGraphTripUseCase := usecase.NewLoadGraphTripUseCase(tripConfigRepository)
	DijkstraTripUseCase := usecase.NewDijkstraTripUseCase()
	DFSTripsUseCase := usecase.NewDFSTripsUseCase()
	validatedUsecase := usecase.NewValidatedTripsUseCase(shipRepository, configRepository)
	addTriptoReservationUseCase := usecase.NewAddTripsToReservationUseCase(db, validatedUsecase)

	tripHandler := handlers.NewTripHandler(
		LoadGraphTripUseCase,
		DijkstraTripUseCase,
		DFSTripsUseCase,
		validatedUsecase,
		addTriptoReservationUseCase,
		listAvailableUseCase,
		filterTripsUseCase,
	)

	// reservation
	reservationRepository := postgres.NewReservationRepository(dbPool)
	createReservationUseCase := usecase.NewCreateReservationUseCase(reservationRepository)
	updateReservationUseCase := usecase.NewUpdateReservationUseCase(reservationRepository)
	getReservationUseCase := usecase.NewGetReservationUseCase(reservationRepository)
	listReservationUseCase := usecase.NewListReservationUseCase(reservationRepository)
	deleteReservationUseCase := usecase.NewDeleteReservationUseCase(reservationRepository)
	reservationHandler := handlers.NewReservationHandler(
		createReservationUseCase,
		updateReservationUseCase,
		getReservationUseCase,
		listReservationUseCase,
		deleteReservationUseCase,
		validatedUsecase,
		addTriptoReservationUseCase,
	)
	// upload photos
	uploadPhotosDirectUseCase := usecase.NewUploadDirectMultipleUseCase(shipRepository)

	photosHandler := handlers.NewPhotosHandler(
		uploadPhotosDirectUseCase,
	)

	return &handlers.Handlers{
		HealthHandler:      healthCheckHandler,
		ShipConfigHandler:  shipConfigHandler,
		ShipHandler:        shipHandler,
		TerminalHandler:    terminalHandler,
		RouteHandler:       routeHandler,
		ScheduleHandler:    scheduleHandler,
		TripConfigHandler:  tripConfigHandler,
		TripHandler:        tripHandler,
		ReservationHandler: reservationHandler,
		PhotosHandler:      photosHandler,
		GpsHandler:         gpsHandler,
	}
}
