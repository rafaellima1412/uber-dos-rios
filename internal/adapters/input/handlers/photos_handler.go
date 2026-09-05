package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/http/dto"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type PhotosHandler struct {
	uploadDirectUseCase input.UploadPhotosDirectUseCase
}

func NewPhotosHandler(
	uploadDirectUseCase input.UploadPhotosDirectUseCase,
) *PhotosHandler {
	return &PhotosHandler{
		uploadDirectUseCase: uploadDirectUseCase,
	}
}

// UploadDirectMultiple godoc
// @Summary      Faz upload de múltiplos arquivos direto para Banco
// @Description  Recebe múltiplos urls e envia para banco
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param request body dto.PhotosRequest true "Upload urls photos request"
// @Success 204 "No Content"
// @Failure 400  {object}  dto.ErrorResponse "Invalid request"
// @Failure 404  {object}  dto.ErrorResponse "Not found"
// @Failure 500  {object}  dto.ErrorResponse "Internal server error"
// @Router       /photos/upload-urls [post]
func (h *PhotosHandler) UploadDirectMultiple(c *gin.Context) {
	var req dto.PhotosRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	photosInput := input.UploadShipImagesInput{
		Urls: req.Urls,
	}
	shipID := req.ShipID
	err := h.uploadDirectUseCase.Execute(c.Request.Context(), &photosInput, shipID)
	if err != nil {
		logger.Error("failed to upload files", zap.Error(err))
		c.JSON(500, dto.ErrorResponse{Error: "failed to upload files: " + err.Error()})
		return
	}

	c.JSON(200, dto.SuccessMessageResponse{Message: "Urls updated successfully"})
}
