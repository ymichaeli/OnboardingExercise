package lifecycle

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRoute() Router {
	return Router{}
}

func (r Router) InitRoutes(engine *gin.Engine) {
	engine.GET("/is-alive", r.health)
	engine.GET("/is-ready", r.readiness)
}
