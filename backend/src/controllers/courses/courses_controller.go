package courses

import (
	dtos "backend/src/domain/dtos/courses"
	"backend/src/services/courses"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CoursesControllers struct {
	service courses.ICourseService
}

type ICoursesControllers interface {
	GetCourseById(g *gin.Context)
	GetAllCourses(g *gin.Context)
	CreateCourse(g *gin.Context)
	UpdateCourse(g *gin.Context)
}

func NewCoursesControllers(service courses.ICourseService) *CoursesControllers {
	return &CoursesControllers{service: service}
}

func (c *CoursesControllers) GetCourseById(g *gin.Context) {
	id := g.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(400, gin.H{
			"Ok":      false,
			"Message": "El id debe ser un numero",
		})
		return
	}
	response := c.service.GetCourseById(idParsed)

	g.JSON(200, gin.H{
		"Ok":      true,
		"Message": "Curso obtenido con exito",
		"Data":    response,
	})

}

func (c *CoursesControllers) GetAllCourses(g *gin.Context) {
	response := c.service.GetAllCourses()

	g.JSON(200, gin.H{
		"Ok":      true,
		"Message": "Cursos obtenidos con exito",
		"Data":    response,
	})

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

func (c *CoursesControllers) UpdateCourse(g *gin.Context) {
	var updateCourseDto dtos.UpdateCoursesRequestDto

	fmt.Println("updateCourseDto", updateCourseDto)

	if err := g.ShouldBindJSON(&updateCourseDto); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.service.UpdateCourse(updateCourseDto)

	g.JSON(200, gin.H{
		"Ok":      true,
		"Message": "Curso actualizado con exito",
		"Data":    response,
	})

}
