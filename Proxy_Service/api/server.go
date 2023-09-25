package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/api/routes"
	"github.com/mooha76/Kofee/Proxy_Service/config"
)

type Server struct {
	port   string
	engine *gin.Engine
}

// NewServerHTTP creates a new server with given handler functions
func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler,
	merchantHandler handler.MerchantHandler, userHandler handler.UserHandler) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.SetupUserRoutes(engine.Group("/"), userHandler, authHandler)
	routes.SetupMerchantRoutes(engine.Group("/"), merchantHandler)
	//routes.SetupAdminRoutes(engine.Group("/admin"), productHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	c.engine.Run(c.port)
}
