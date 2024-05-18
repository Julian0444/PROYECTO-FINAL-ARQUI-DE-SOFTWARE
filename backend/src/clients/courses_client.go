package clients

import (
	model "backend/src/models"

	"gorm.io/gorm"
)

type CourseClient struct {
	db *gorm.DB
}

func NewCoursesClient(db *gorm.DB) *CourseClient {
	return &CourseClient{db: db}
}

func (c *CourseClient) CreateCourse(course model.Course) model.Course {
	result := c.db.Create(&course)
	if result.Error != nil {
		return model.Course{}
	}
	return course
}
