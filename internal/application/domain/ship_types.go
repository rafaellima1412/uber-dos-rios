package domain

type ShipType string

const (
	FERRYBOAT ShipType = "FERRYBOAT" // passageiro (camarotes) + carga (veículos) + encomendas (grandes) cabin
	SHIP      ShipType = "SHIP"      // navio de grande porte (passageiro (camarotes) + carga + encomendas) cabin
	BOAT      ShipType = "BOAT"      // lancha (passageiro + encomendas pequenas)  seats
)

type ShipStatus string

const (
	ACTIVE      ShipStatus = "ACTIVE"      // disponível para viagens
	INACTIVE    ShipStatus = "INACTIVE"    // não disponível para viagens
	MAINTENANCE ShipStatus = "MAINTENANCE" // em manutenção
)


