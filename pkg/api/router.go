package api

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRoutes(engine *gin.Engine)
}
