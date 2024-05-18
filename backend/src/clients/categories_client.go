package clients

import (
	model "backend/src/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CategoriesClient struct {
	Db *gorm.DB
}

func NewCategoriesClient(db *gorm.DB) *CategoriesClient {
	return &CategoriesClient{Db: db}
}

func (c *CategoriesClient) Create(category model.Category) model.Category {
	result := c.Db.Create(&category)

	if result.Error != nil {
		log.Error()
	}
	log.Debug("Categoria creado con exito, su id es: ", result)
	return category
}
