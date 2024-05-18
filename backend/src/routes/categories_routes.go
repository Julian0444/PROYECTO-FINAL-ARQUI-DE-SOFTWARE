package routes

import (
	"backend/src/controllers/categories"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(engine *gin.Engine, controller *categories.CategoriesController) {
	engine.POST("/categories/create", controller.CreateCategory)
}
