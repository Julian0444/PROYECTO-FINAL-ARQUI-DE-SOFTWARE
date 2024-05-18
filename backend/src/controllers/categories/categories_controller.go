package categories

import (
	dtos "backend/src/domain/dtos/categories"
	categoriesService "backend/src/services/categories"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	service categoriesService.ICourseService
}

type ICategoriesController interface {
	CreateCategory(g *gin.Context)
}

func NewCategoriesController(service categoriesService.ICourseService) *CategoriesController {
	return &CategoriesController{service: service}
}

func (c *CategoriesController) CreateCategory(g *gin.Context) {
	var createCategoryDto dtos.CreateCategoryRequestDto

	if err := g.ShouldBindJSON(&createCategoryDto); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.service.CrateCategory(createCategoryDto)

	g.JSON(201, gin.H{
		"Ok":      true,
		"Message": "Categoria creada con exito",
		"Data":    response,
	})

}
