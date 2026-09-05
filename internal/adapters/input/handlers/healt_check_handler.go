package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {
	index string
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		index: "index",
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Check the health of the service
// @Tags health
// @Success 200 {object} map[string]string
// @Router / [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
