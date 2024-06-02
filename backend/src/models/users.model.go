package model

type User struct {
	Id       uint   `gorm:"id"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
	Role     int    `gorm:"role"`
	Name     string `gorm:"user_name"`
	Avatar   string `gorm:"avatar"`
}

type Users []User
