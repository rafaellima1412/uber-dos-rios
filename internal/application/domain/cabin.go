package domain

import (
	"time"

	"github.com/google/uuid"
)

type CabinLayout struct {
	Andar            int
	Lado             string
	Local            string
	Identificador    string
}

type CabinDetails struct {
	ID          uuid.UUID
	Name        string
	BedType     string
	Capacity    int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Cabin struct {
	Cabins []CabinDetails
}


