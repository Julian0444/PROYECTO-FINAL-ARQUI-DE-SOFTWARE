package routes

import (
	controllers "backend/src/controllers/courses"

	"github.com/gin-gonic/gin"
)

func CoursesRoutes(engine *gin.Engine, controller *controllers.CoursesControllers) {
	engine.GET("/courses/:id", controller.GetCourseById)
	engine.GET("/courses", controller.GetAllCourses)
	engine.POST("/courses/create", controller.CreateCourse)
	engine.PUT("/courses/update", controller.UpdateCourse)
}
