package domain

import (
	"time"

	"github.com/google/uuid"
)

type TripConfig struct {
	ID        uuid.UUID
	ShipID    uuid.UUID
	RouteID   uuid.UUID
	ShipName  string
	RouteName string
	StartDate      time.Time      //data inicio de validade
	ExpirationDate time.Time      //data final de validade
	DepartureTime  time.Time      //hora saida
	ArrivalTime    time.Time      //hora chegada
	Recurrence     TripRecurrence //diario, semanal, mensal
	DurationDays   int            //já é o ciclo completo
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

type TripResult struct {
	OrganizationID   uuid.UUID
	OrganizationName string
	RouteID          uuid.UUID
	RouteName        string
	StopOrderOrigem  int
	StopOrderDestino int
	DepartureTime    time.Time
	ArrivalTime      time.Time

	Itinerary []Itinerary
}

type Itinerary struct {
	City      string
	UF        string
	StopOrder int
}
