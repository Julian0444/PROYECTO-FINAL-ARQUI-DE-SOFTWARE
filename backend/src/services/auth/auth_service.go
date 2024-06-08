package auth

import (
	"backend/src/clients"
	dtos "backend/src/domain/dtos/auth"
	model "backend/src/models"
	utils "backend/src/utils"
	"errors"
)

type AuthService struct {
	client *clients.UserClient
}

type IAuthService interface {
	Register(registerDto dtos.RegisterRequest) (dtos.AuthResponse, error)
	Login(loginDto dtos.LoginRequest) (dtos.AuthResponse, error)
	Refresh(token string) (dtos.AuthResponse, error)
}

func NewAuthService(client *clients.UserClient) *AuthService {
	return &AuthService{client: client}
}

func (a *AuthService) Register(registerDto dtos.RegisterRequest) (dtos.AuthResponse, error) {
	checkUser := a.client.GetUserByEmail(registerDto.Email)
	if checkUser.Id != 0 {
		return dtos.AuthResponse{}, errors.New("User already exists")
	}

	crypto := utils.NewCrypto()
	hashedPassword, err := crypto.HashPassword(registerDto.Password)
	if err != nil {
		return dtos.AuthResponse{}, err
	}
	newUser := model.User{
		Name:     registerDto.Name,
		Email:    registerDto.Email,
		Password: hashedPassword,
	}

	user := a.client.CreateUser(newUser)
	if user.Id == 0 {
		return dtos.AuthResponse{}, errors.New("User not created")
	}

	token := utils.SignDocument(user.Id, user.Role)
	return dtos.AuthResponse{
		Ok:      true,
		Message: "User created",
		Data:    user,
		Token:   token,
	}, nil
}

func (a *AuthService) Login(loginDto dtos.LoginRequest) (dtos.AuthResponse, error) {
	user := a.client.GetUserByEmail(loginDto.Email)
	if user.Id == 0 {
		return dtos.AuthResponse{}, errors.New("User not found")
	}

	crypto := utils.NewCrypto()
	if !crypto.ComparePassword(user.Password, loginDto.Password) {
		return dtos.AuthResponse{}, errors.New("Invalid password")
	}

	token := utils.SignDocument(user.Id, user.Role)
	userData := model.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return dtos.AuthResponse{
		Ok:      true,
		Message: "User logged in",
		Data:    userData,
		Token:   token,
	}, nil
}

func (a *AuthService) Refresh(token string) (dtos.AuthResponse, error) {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return dtos.AuthResponse{}, errors.New("Invalid token")
	}

	userID := int(claims["id"].(float64))
	user := a.client.GetUserById(userID)
	if user.Id == 0 {
		return dtos.AuthResponse{}, errors.New("User not found")
	}

	if user.Role != int(claims["role"].(float64)) {
		return dtos.AuthResponse{}, errors.New("Invalid role")
	}

	newToken := utils.SignDocument(user.Id, user.Role)
	userData := model.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return dtos.AuthResponse{
		Ok:      true,
		Message: "Token refreshed",
		Data:    userData,
		Token:   newToken,
	}, nil
}
