package domain
type TripRecurrence string

const (
	DAILY TripRecurrence = "DAILY" // viagens regulares
	MONTHLY TripRecurrence = "MONTHLY" // viagens mensais
	WEEKLY  TripRecurrence = "WEEKLY"  // viagens semanais
)

