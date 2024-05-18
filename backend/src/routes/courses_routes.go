package routes

import (
	controllers "backend/src/controllers/courses"

	"github.com/gin-gonic/gin"
)

func CoursesRoutes(engine *gin.Engine, controller *controllers.CoursesControllers) {
	engine.POST("/courses/create", controller.CreateCourse)
}
