package auth

import (
	dtos "backend/src/domain/dtos/auth"
	services "backend/src/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service        services.IAuthService
	failedAttempts map[string]int
}

type IAuthController interface {
	Register(g *gin.Context)
	Login(g *gin.Context)
	Refresh(g *gin.Context)
}

func NewAuthController(service services.IAuthService) *AuthController {
	return &AuthController{service: service, failedAttempts: make(map[string]int)}
}

func (a *AuthController) Register(g *gin.Context) {
	var registerDto dtos.RegisterRequest
	if err := g.ShouldBindJSON(&registerDto); err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	newUser, err := a.service.Register(registerDto)
	if err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}
	g.JSON(200, gin.H{
		"Ok":      true,
		"message": "User created successfully",
		"data":    newUser,
	})
}

func (a *AuthController) Login(g *gin.Context) {
	var loginDto dtos.LoginRequest
	if err := g.ShouldBindJSON(&loginDto); err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	if a.failedAttempts[loginDto.Email] >= 3 {
		g.JSON(403, gin.H{
			"Ok":    false,
			"error": "Your account has been locked due to too many failed login attempts",
		})
		return
	}

	user, err := a.service.Login(loginDto)
	if err != nil {
		a.failedAttempts[loginDto.Email]++
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	a.failedAttempts[loginDto.Email] = 0
	g.JSON(200, gin.H{
		"Ok":      true,
		"message": "User logged in successfully",
		"data":    user,
	})
}

func (a *AuthController) Refresh(g *gin.Context) {
	token := g.Request.Header.Get("Authorization")[7:]
	if token == "" {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": "Token not found",
		})
		return
	}
	newToken, err := a.service.Refresh(token)
	if err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	g.JSON(200, gin.H{
		"Ok":      true,
		"message": "Token refreshed",
		"token":   newToken,
	})
}
