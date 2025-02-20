package v1

import (
	"github.com/Dffarhn/bakulenapi/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userHandler *UserHandler) {
	// Register user routes
	router.GET("/users", middleware.AuthMiddleware() ,userHandler.GetUser)
	router.PUT("/users", middleware.AuthMiddleware() ,userHandler.UpdateUser)
}
