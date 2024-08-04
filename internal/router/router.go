package router

import (
	"github.com/gin-gonic/gin"
	"wb_tech_level_0/internal/controllers"
	"wb_tech_level_0/internal/service"
)

func SetupRoutes(r *gin.Engine, orderService *service.OrderService) {
	// Load all HTML templates from the templates directory
	r.LoadHTMLGlob("templates/*")

	orderController := controllers.NewOrderController(orderService)

	// Handler for the root path
	r.GET("/", orderController.GetIndex)

	// Handler for creating an order
	r.POST("/orders", orderController.CreateOrder)
}
