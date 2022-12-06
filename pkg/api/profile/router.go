package profile

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler Handler
}

// NewRouter must return *Handler because of the Handler interface and that the functions get *Handler
func NewRouter(handler Handler) *Router {
	return &Router{handler: handler}
}

func (router *Router) InitRoutes(engine *gin.Engine) {
	group := engine.Group("/profiles")

	group.GET("", router.handler.GetAllProfiles)
	group.POST("", router.handler.CreateProfile)
	group.GET("/:userId", router.handler.GetProfileByUserID)
	group.PUT("/:userId", router.handler.UpdateProfile)
}
