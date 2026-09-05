package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/http/dto"
	"github.com/riolivre/nautical_logistics/internal/adapters/output/service"
)

// GPSHandler lida com requisições HTTP relacionadas ao GPS.
type GPSHandler struct {
	service *service.GPSService
}

// NewGPSHandler cria uma nova instância de GPSHandler.
func NewGPSHandler(s *service.GPSService) *GPSHandler {
	return &GPSHandler{service: s}
}

// PostGPS godoc
// @Summary Recebe dados de localização GPS
// @Description Recebe latitude, longitude e timestamp enviados por um dispositivo (ex: ESP32 com módulo GPS).
// @Tags gps
// @Accept json
// @Produce json
// @Param gps body dto.GPSRequest true "Dados do GPS"
// @Success 200 {object} map[string]string "GPS data received successfully"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Failure 500 {object} map[string]string "Failed to process GPS data"
// @Router /gps [post]
func (h *GPSHandler) PostGPS(c *gin.Context) {
	var req dto.GPSRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.service.ProcessGPSData(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process GPS data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GPS data received successfully"})
}
