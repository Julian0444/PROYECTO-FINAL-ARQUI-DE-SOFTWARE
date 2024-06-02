package routes

import (
	"backend/src/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.Engine, controller *auth.AuthController) {
	engine.POST("/auth/register", controller.Register)
	engine.POST("/auth/login", controller.Login)
	engine.GET("/auth/refresh", controller.Refresh)
}
