package main

import (
	"backend/src/config"
	"backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	db := config.NewDBConnection()

	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type,X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		c.Next()
	})

	routes.AppRouter(engine, db)
	engine.Run(":8080")
}
