package adapters

import (
	"backend/src/clients"
	controllers "backend/src/controllers/categories"
	"backend/src/services/categories"

	"gorm.io/gorm"
)

func CategoryAdapter(db *gorm.DB) *controllers.CategoriesController {
	client := clients.NewCategoriesClient(db)
	service := categories.NewCourseService(client)
	return controllers.NewCategoriesController(service)
}
