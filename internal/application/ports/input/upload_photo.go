package input

import (
	"context"
)
type UploadShipImagesInput struct {
    Urls  []string 
}

type UploadPhotosDirectUseCase interface {
	Execute(ctx context.Context, photoInput *UploadShipImagesInput, ShipID string) error
}

