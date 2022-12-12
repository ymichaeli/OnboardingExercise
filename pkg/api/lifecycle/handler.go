package lifecycle

import (
	"OnboardingExercise/pkg/service/lifecycle"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler implements service lifecycle verifications
type Handler struct {
	service lifecycle_service.Service
}

func NewHandler(service lifecycle_service.Service) Handler {
	return Handler{service: service}
}

// Health checks if the service is up and running
func (h Handler) Health(c *gin.Context) {
	if h.service.IsAlive() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.AbortWithStatus(http.StatusServiceUnavailable)
}

// Readiness checks if the service is ready to accept new requests
func (h Handler) Readiness(c *gin.Context) {
	if h.service.IsReady() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.AbortWithStatus(http.StatusServiceUnavailable)
}
