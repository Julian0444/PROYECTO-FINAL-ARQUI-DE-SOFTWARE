package routes

import (
	"backend/src/adapters"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRouter(engine *gin.Engine, db *gorm.DB) {
	CategoriesRoutes(engine, adapters.CategoryAdapter(db))
	CoursesRoutes(engine, adapters.CoursesAdapter(db))
}
