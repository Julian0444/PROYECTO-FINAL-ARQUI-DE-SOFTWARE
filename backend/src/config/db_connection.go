package config

import (
	model "backend/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dsn := "myuser:mypassword@tcp(database:3306)/mydatabase"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful.")

	db.AutoMigrate(&model.User{}, &model.Course{}, &model.Category{}, &model.Inscriptos{})

	return db
}
