package handlers

import (
	"depoguna-api/controllers"
	"depoguna-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler interface {
	Handler(rg *gin.RouterGroup)
}

type orderHandler struct {
	DB *gorm.DB
}

func NewOrderHandler(db *gorm.DB) OrderHandler {
	return &orderHandler{
		DB: db,
	}
}

func (h *orderHandler) Handler(rg *gin.RouterGroup) {
	orderController := controllers.NewOrderController(h.DB)

	order := rg.Group("/api")
	order.Use(middleware.JWTAuth)
	order.GET("/order", orderController.FindAll)
	order.GET("/order/:id", orderController.GetDetail)
	order.POST("/order", orderController.Insert)
	order.PUT("/order/:id", orderController.Update)
	order.DELETE("/order/:id", orderController.Delete)
	order.GET("/order/search", orderController.Search)
}
