package courses

import (
	dtos "backend/src/domain/dtos/courses"
	"backend/src/services/courses"

	"github.com/gin-gonic/gin"
)

type CoursesControllers struct {
	service courses.ICourseService
}

type ICoursesControllers interface {
	CreateCourse(g *gin.Context)
}

func NewCoursesControllers(service courses.ICourseService) *CoursesControllers {
	return &CoursesControllers{service: service}
}

func (c *CoursesControllers) CreateCourse(g *gin.Context) {
	var createCourseDto dtos.CreateCoursesRequestDto

	if err := g.ShouldBindJSON(&createCourseDto); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.service.CreateCourse(createCourseDto)

	g.JSON(201, gin.H{
		"Ok":      true,
		"Message": "Curso creado con exito",
		"Data":    response,
	})

}
