package main

import (
	"backend/src/config"
	"backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	db := config.NewDBConnection()
	routes.AppRouter(engine, db)

	engine.Run(":8000")
}
