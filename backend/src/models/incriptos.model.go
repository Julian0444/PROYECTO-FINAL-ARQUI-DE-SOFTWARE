package model

import (
	"gorm.io/gorm"
)

type Inscriptos struct {
	gorm.Model
	CourseId uint
	UserId   uint

	User   User
	Course Course
}
