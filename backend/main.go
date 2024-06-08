package main

import (
	"backend/src/config"
	"backend/src/routes"
	"backend/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	db := config.NewDBConnection()

	engine.Use(utils.CorsMiddleware)

	routes.AppRouter(engine, db)
	engine.Run(":8080")
}
