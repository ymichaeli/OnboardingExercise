package api

import (
	"OnboardingExercise/pkg/api/lifecycle"
	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRoutes(engine *gin.Engine)
}

func GetRoutes() []Router {
	return []Router{
		lifecycle.NewRoute(),
	}
}
