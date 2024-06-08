package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type,X-Auth-Token,Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204) // No content
		return
	}
	c.Next()
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, BearerSchema))

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		_, err := VerifyToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
