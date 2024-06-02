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

func (c *CourseClient) GetCourseById(id int) model.Course {
	var course model.Course
	c.db.Where("id = ?", id).First(&course).Preload("Category")
	return course
}

func (c *CourseClient) GetAllCourses() []model.Course {
	var courses []model.Course
	c.db.Find(&courses)
	return courses

}

func (c *CourseClient) CreateCourse(course model.Course) model.Course {
	result := c.db.Create(&course)
	if result.Error != nil {
		return model.Course{}
	}
	return course
}

func (c *CourseClient) UpdateCourse(course model.Course) model.Course {
	result := c.db.Save(&course)
	if result.Error != nil {
		return model.Course{}
	}
	return course
}
