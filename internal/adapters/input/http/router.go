package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:8080",
			"https://nautical-api.riolivre.com.br",
			"http://nautical-api.riolivre.com.br",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	router.SetTrustedProxies(nil)

	api := router.Group("/api/v1")
	{
		api.GET("/", h.HealthHandler.HealthCheck)
		api.POST("/gps", h.GpsHandler.PostGPS)
		ship := api.Group("/ships")
		{
			ship.GET("/:id", h.ShipHandler.FindByID)
			ship.POST("", h.ShipHandler.CreateShip)
			ship.GET("", h.ShipHandler.ListShips)
			ship.DELETE("/:id", h.ShipHandler.DeleteShip)
			ship.PUT("/:id", h.ShipHandler.UpdateShip)
		}
		config := api.Group("/ships-config")
		{
			// cadeiras
			config.POST("/seats", h.ShipConfigHandler.CreateSeat)
			config.PUT("/seats/:id", h.ShipConfigHandler.UpdateSeat)
			config.GET("/seats", h.ShipConfigHandler.ListSeat)
			config.DELETE("/seats/:id", h.ShipConfigHandler.DeleteSeat)
			config.GET("/units/search", h.ShipConfigHandler.SearchSeat)
			// camarote
			config.POST("/cabins", h.ShipConfigHandler.CreateCabin)
			config.GET("/cabins", h.ShipConfigHandler.ListCabin)
			config.DELETE("/cabins/:id", h.ShipConfigHandler.DeleteCabin)
			config.PUT("/cabins/:id", h.ShipConfigHandler.UpdateCabin)

		}
		terminal := api.Group("/terminals")
		{
			terminal.POST("", h.TerminalHandler.CreateTerminal)
			terminal.GET("/:id", h.TerminalHandler.GetTerminal)
			terminal.PUT("/:id", h.TerminalHandler.UpdateTerminal)
			terminal.DELETE("/:id", h.TerminalHandler.DeleteTerminal)
			terminal.GET("", h.TerminalHandler.ListTerminals)
			terminal.GET("/search", h.TerminalHandler.SearchTerminals)
			terminal.GET("cities", h.TerminalHandler.ListCities)
		}
		route := api.Group("/routes")
		{
			route.POST("", h.RouteHandler.CreateRoute)
			route.GET("/:id", h.RouteHandler.GetRoute)
			route.PUT("/:id", h.RouteHandler.UpdateRoute)
			route.DELETE("/:id", h.RouteHandler.DeleteRoute)
			route.GET("", h.RouteHandler.ListRoutes)
			route.GET("/search", h.RouteHandler.SearchRoutes)
		}
		schedule := api.Group("/schedules")
		{
			schedule.POST("", h.ScheduleHandler.CreateSchedule)
			schedule.GET("/:id", h.ScheduleHandler.GetSchedule)
			schedule.PUT("/:id", h.ScheduleHandler.UpdateSchedule)
			schedule.DELETE("/:id", h.ScheduleHandler.DeleteSchedule)
			schedule.GET("", h.ScheduleHandler.ListSchedules)
		}
		trip_config := api.Group("/trips-config")
		{
			trip_config.POST("", h.TripConfigHandler.CreateTripConfig)
			trip_config.PUT("/:id", h.TripConfigHandler.UpdateTripConfig)
			trip_config.GET("/:id", h.TripConfigHandler.FindByIDTripConfig)
			trip_config.GET("", h.TripConfigHandler.ListTripConfig)
			trip_config.DELETE("/:id", h.TripConfigHandler.DeleteTripConfig)
		}
		reservation := api.Group("/reservations")
		{
			reservation.POST("", h.ReservationHandler.CreateReservation)
			reservation.PUT("/:id", h.ReservationHandler.UpdateReservation)
			reservation.GET("/:id", h.ReservationHandler.FindByID)
			reservation.GET("", h.ReservationHandler.ListReservations)
			reservation.DELETE("/:id", h.ReservationHandler.DeleteReservation)
		}
		// trips
		trips := api.Group("/trips")
		{
			trips.GET("/search", h.TripHandler.SearchTrips)
			trips.POST("", h.TripHandler.CreateConfig)
			trips.POST("/catalogo/menor-preco", h.TripHandler.ValidateConnectionsTrips)
			trips.POST("/catalogo", h.TripHandler.DFSTrips)
		}
		photos := api.Group("/photos")
		{
			photos.POST("/upload-urls", h.PhotosHandler.UploadDirectMultiple)
		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
