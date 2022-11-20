package handlers

import (
	"depoguna-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler interface {
	Handler(rg *gin.RouterGroup)
}

type authHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) AuthHandler {
	return &authHandler{
		DB: db,
	}
}

func (h *authHandler) Handler(rg *gin.RouterGroup) {
	authController := controllers.NewAuthController(h.DB)

	auth := rg.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
}
