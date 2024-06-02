package adapters

import (
	"backend/src/clients"
	controllers "backend/src/controllers/auth"
	"backend/src/services/auth"

	"gorm.io/gorm"
)

func AuthAdapter(db *gorm.DB) *controllers.AuthController {
	client := clients.NewUserClient(db)
	service := auth.NewAuthService(client)
	return controllers.NewAuthController(service)
}
