package users

import model "backend/src/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Ok      bool       `json:"Ok"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
	Token   string     `json:"token"`
}
