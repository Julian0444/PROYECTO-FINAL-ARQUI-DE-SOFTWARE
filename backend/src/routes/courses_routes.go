package routes

import (
	controllers "backend/src/controllers/courses"
	"backend/src/utils"

	"github.com/gin-gonic/gin"
)

func CoursesRoutes(engine *gin.Engine, controller *controllers.CoursesControllers) {
	engine.GET("/courses/:id", utils.AuthRequired(), controller.GetCourseById)
	engine.GET("/courses", utils.AuthRequired(), controller.GetAllCourses)
	engine.POST("/courses/create", utils.AuthRequired(), controller.CreateCourse)
	engine.PUT("/courses/update", utils.AuthRequired(), controller.UpdateCourse)
}
