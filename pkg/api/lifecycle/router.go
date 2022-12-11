package lifecycle

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler Handler
}

func NewRouter(handler Handler) Router {
	return Router{handler: handler}
}

func (r Router) InitRoutes(engine *gin.Engine) {
	engine.GET("/is-alive", r.handler.Health)
	engine.GET("/is-ready", r.handler.Readiness)
}
