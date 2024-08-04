package server

import (
	"github.com/gin-gonic/gin"
	"wb_tech_level_0/config"
	"wb_tech_level_0/internal/router"
	"wb_tech_level_0/internal/service"
)

type HTTPServer struct {
	cfg          *config.Config
	orderService *service.OrderService
	router       *gin.Engine
}

func NewHTTPServer(cfg *config.Config, orderService *service.OrderService) *HTTPServer {
	r := gin.Default()

	server := &HTTPServer{
		cfg:          cfg,
		orderService: orderService,
		router:       r,
	}

	router.SetupRoutes(r, orderService)

	return server
}

func (s *HTTPServer) Run() error {
	return s.router.Run(":" + s.cfg.HTTPPort)
}
