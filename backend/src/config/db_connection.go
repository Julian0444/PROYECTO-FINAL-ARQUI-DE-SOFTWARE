package config

import (
	model "backend/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dsn := "root:Congo2020@tcp(127.0.0.1:3306)/arq-soft?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful.")

	db.AutoMigrate(&model.User{}, &model.Course{}, &model.Category{}, &model.Inscriptos{})

	return db
}
