package input

import (
	"context"

	"github.com/google/uuid"
)
type CabinLayout struct {
	Andar         int    
	Lado          string 
	Local         string 
	Identificador string 
}
type Cabin struct {
	Name        string  
	BedType     string  
	Capacity    int     
	Description *string 
}

type CreateCabinInput struct {
	Cabins []Cabin
}

type CreateCabinOutput struct {
	ID          uuid.UUID
	Name        string
	BedType     string
	Capacity    int
	Description *string
}

type CreateCabinUseCase interface {
	Execute(ctx context.Context, input *CreateCabinInput) error
}
