package api

import (
	"OnboardingExercise/config"
	"OnboardingExercise/pkg/api/lifecycle"
	"OnboardingExercise/pkg/api/middlewares"
	"OnboardingExercise/pkg/api/profile"
	lifecycle_repository "OnboardingExercise/pkg/repository/lifecycle"
	"OnboardingExercise/pkg/repository/profile"
	"OnboardingExercise/pkg/service/lifecycle"
	"OnboardingExercise/pkg/service/profile"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Server hold gin framework engine and allow to create and start a new server at selected port and domain
type Server struct {
	engine *gin.Engine
}

// NewServer initialize the server using gin engine and declares middleware and routes
func NewServer(db *sql.DB) Server {
	engine := gin.New()
	server := Server{engine}

	server.initMiddlewares()
	server.initRoutes(db)

	return server
}

// Start activate the server on the specified domain and port
func (server Server) Start(serverInfo config.ServerInfo) error {
	return server.engine.Run(fmt.Sprintf("%s:%v", serverInfo.Host, serverInfo.Port))
}

func (server Server) initRoutes(db *sql.DB) {
	lifecycleHandler := lifecycle.NewHandler(lifecycle_service.NewService(lifecycle_repository.NewRepository(db)))
	profileHandler := profile_api.NewHandler(profile_service.NewService(profile_repository.NewRepository(db)))

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
