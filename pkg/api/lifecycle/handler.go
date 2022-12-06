package lifecycle

import (
	"OnboardingExercise/pkg/service/lifecycle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service lifecycle_service.Service
}

func NewHandler(service lifecycle_service.Service) Handler {
	return Handler{service: service}
}

func (h Handler) health(c *gin.Context) {
	if h.service.IsAlive() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
}

func (h Handler) readiness(c *gin.Context) {
	if h.service.IsReady() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
}
