package lifecycle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r Router) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (r Router) readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
