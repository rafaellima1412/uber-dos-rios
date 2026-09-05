package handlers

// Handlers is a struct that holds all the HTTP handlers for the application.
type Handlers struct {
	HealthHandler      *HealthHandler
	ShipConfigHandler  *ShipConfigHandler
	ShipHandler        *ShipHandler
	TerminalHandler    *TerminalHandler
	RouteHandler       *RouteHandler
	ScheduleHandler    *ScheduleHandler
	TripConfigHandler  *TripConfigHandler
	TripHandler        *TripHandler
	ReservationHandler *ReservationHandler
	PhotosHandler      *PhotosHandler
	GpsHandler         *GPSHandler   
}
