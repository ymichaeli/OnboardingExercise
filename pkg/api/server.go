package api

import (
	"OnboardingExercise/pkg/api/lifecycle"
	"OnboardingExercise/pkg/api/middlewares"
	"OnboardingExercise/pkg/api/profile"
	"OnboardingExercise/pkg/repository/profile"
	"OnboardingExercise/pkg/service/lifecycle"
	"OnboardingExercise/pkg/service/profile"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() Server {
	engine := gin.New()
	server := Server{engine}

	server.initMiddlewares()
	server.initRoutes()

	return server
}

func (server Server) Start() error {
	return server.engine.Run("localhost:8080")
}

func (server Server) initRoutes() {
	lifecycleHandler := lifecycle.NewHandler(lifecycle_service.NewService())
	profileHandler := profile.NewHandler(profile_service.NewService(profile_repository.NewDAL(profile_repository.Profiles)))

	routes := []Router{
		lifecycle.NewRouter(lifecycleHandler),
		profile.NewRouter(profileHandler),
	}

	for _, route := range routes {
		route.InitRoutes(server.engine)
	}
}

func (server Server) initMiddlewares() {
	server.engine.Use(gin.Logger())
	server.engine.Use(gin.Recovery())
	server.engine.Use(middlewares.ErrorHandler)
}
