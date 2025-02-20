package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *AuthHandler) {
	// Register user routes
	router.POST("/auth/register", authHandler.RegisterUser)
	router.POST("/auth/login", authHandler.LoginUser)
	router.POST("/auth/google", authHandler.GoogleLogin)
}
