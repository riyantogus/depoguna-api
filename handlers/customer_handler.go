package handlers

import (
	"depoguna-api/controllers"
	"depoguna-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerHandler interface {
	Handler(rg *gin.RouterGroup)
}

type customerHandler struct {
	DB *gorm.DB
}

func NewCustomerHandler(db *gorm.DB) CustomerHandler {
	return &customerHandler{
		DB: db,
	}
}

func (h *customerHandler) Handler(rg *gin.RouterGroup) {
	customerController := controllers.NewCustomerController(h.DB)

	customer := rg.Group("/api")
	customer.Use(middleware.JWTAuth)
	customer.GET("/customer", customerController.FindAll)
	customer.GET("/customer/:id", customerController.GetDetail)
	customer.POST("/customer", customerController.Insert)
	customer.PUT("/customer/:id", customerController.Update)
	customer.DELETE("/customer/:id", customerController.Delete)
	customer.GET("/customer/search", customerController.Search)
}
