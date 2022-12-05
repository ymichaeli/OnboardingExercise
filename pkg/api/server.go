package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() (Server, error) {
	engine := gin.New()
	server := Server{engine}

	server.initMiddlewares()
	server.initRoutes()

	return server, nil
}

func (server Server) Start() error {
	return server.engine.Run("localhost:8080")
}

func (server Server) initRoutes() {
	routes := GetRoutes()

	for _, route := range routes {
		route.InitRoutes(server.engine)
	}
}

func (server Server) initMiddlewares() {
	server.engine.Use(gin.Logger())
	server.engine.Use(gin.Recovery())
}
