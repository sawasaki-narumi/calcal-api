package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sawasaki-narumi/calcal-api/handlers"
)

func InitializeRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.POST("/register", handlers.HandleRegistration)
	auth.POST("/login", handlers.HandleLogin)
}

