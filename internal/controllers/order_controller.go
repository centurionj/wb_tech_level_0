package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wb_tech_level_0/internal/service"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

func (ctrl *OrderController) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (ctrl *OrderController) GetOrder(c *gin.Context) {
	idStr := c.PostForm("order_id")
	orderID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := ctrl.OrderService.GetOrder(uint(orderID))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "failed to retrieve order",
		})
		return
	}

	if order == nil {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"error": "Order not found in cache",
		})
		return
	}

	// Сериализация заказа в JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal order"})
		return
	}

	// Рендер на index.html
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Order": string(orderJSON),
	})
}
