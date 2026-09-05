package domain

type ReservationStatus string

const (
	CONFIRMED ReservationStatus = "CONFIRMED" // reserva confirmado pagamento
	CANCELLED ReservationStatus = "CANCELLED" // reserva cancelada
	PENDING   ReservationStatus = "PENDING"   // reserva pendente de associação as viagens e pagamento
	PLANNED   ReservationStatus = "PLANNED"   //criada, mas ainda não pode iniciar por falta confirmar pagamento
	COMPLETED ReservationStatus = "COMPLETED" //viagem finalizada
	ACTIVED   ReservationStatus = "ACTIVED"   //viagem iniciada
)
