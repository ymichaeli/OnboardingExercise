package api

import (
	"OnboardingExercise/pkg/api/lifecycle"
	"OnboardingExercise/pkg/api/middlewares"
	"OnboardingExercise/pkg/api/profile"
	"OnboardingExercise/pkg/repository/profile"
	"OnboardingExercise/pkg/service/lifecycle"
	"OnboardingExercise/pkg/service/profile"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Server hold gin framework engine and allow to create and start a new server at selected port and domain
type Server struct {
	engine *gin.Engine
}

// NewServer initialize the server using gin engine and declares middleware and routes
func NewServer() Server {
	engine := gin.New()
	server := Server{engine}

	server.initMiddlewares()
	server.initRoutes()

	return server
}

// Start activate the server on the specified domain and port
func (server Server) Start(domain string, port int) error {
	return server.engine.Run(fmt.Sprintf("%s:%v", domain, port))
}

func (server Server) initRoutes() {
	lifecycleHandler := lifecycle.NewHandler(lifecycle_service.NewService())
	profileHandler := profile_api.NewHandler(profile_service.NewService(profile_repository.NewRepository(profile_repository.Profiles)))

	routes := []Router{
		lifecycle.NewRouter(lifecycleHandler),
		profile_api.NewRouter(profileHandler),
	}

	for _, route := range routes {
		route.InitRoutes(server.engine)
	}
}

func (server Server) initMiddlewares() {
	server.engine.Use(gin.Logger())
	server.engine.Use(gin.Recovery())
	server.engine.Use(middlewares.ErrorHandlerMiddleware)
}
