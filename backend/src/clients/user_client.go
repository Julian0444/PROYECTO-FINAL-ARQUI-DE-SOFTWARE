package clients

import (
	model "backend/src/models"

	"gorm.io/gorm"
)

type UserClient struct {
	db *gorm.DB
}

func NewUserClient(db *gorm.DB) *UserClient {
	return &UserClient{db: db}
}

func (u *UserClient) CreateUser(user model.User) model.User {
	user = model.User{
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
		Name:     user.Name,
		Avatar:   user.Avatar,
	}
	result := u.db.Create(&user)
	if result.Error != nil {
		return model.User{}
	}
	return user
}

func (u *UserClient) GetUserByEmail(email string) model.User {
	var user model.User
	u.db.Where("email = ?", email).First(&user)
	return user
}

func (u *UserClient) GetUserById(id int) model.User {
	var user model.User
	u.db.Find(&user, id)
	return user
}
