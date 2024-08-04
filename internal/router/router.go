package router

import (
	"github.com/gin-gonic/gin"
	"wb_tech_level_0/internal/controllers"
	"wb_tech_level_0/internal/service"
)

func SetupRoutes(r *gin.Engine, orderService *service.OrderService) {
	// Глобальная загрузка шаблонов
	r.LoadHTMLGlob("templates/*")

	orderController := controllers.NewOrderController(orderService)

	// Обработчик на домашнюю страницу
	r.GET("/", orderController.GetIndex)

	// Обработчик на получение данных о заказе
	r.POST("/orders", orderController.GetOrder)
}
