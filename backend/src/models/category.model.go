package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id           uint
	CategoryName string `gorm:"category_name"`
}

type Categories []Category
