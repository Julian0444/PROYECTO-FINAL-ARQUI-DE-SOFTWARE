package categories

import (
	"backend/src/clients"
	"backend/src/domain/dtos/categories"
	model "backend/src/models"
)

type courseService struct {
	client clients.CategoriesClient
}

type ICourseService interface {
	CrateCategory(createCategoryDto categories.CreateCategoryRequestDto) categories.CreateCategoryResponseDto
}

func NewCourseService(client *clients.CategoriesClient) ICourseService {
	return &courseService{client: *client}
}

func (c *courseService) CrateCategory(createCategoryDto categories.CreateCategoryRequestDto) categories.CreateCategoryResponseDto {
	newCategory := model.Category{
		CategoryName: createCategoryDto.CategoryName,
	}

	category := c.client.Create(newCategory)

	return categories.CreateCategoryResponseDto{
		CategoryId:   int(category.Id),
		CategoryName: category.CategoryName,
	}
}
