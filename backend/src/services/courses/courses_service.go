package courses

import (
	"backend/src/clients"
	"backend/src/domain/dtos/courses"
	model "backend/src/models"
)

type CourseService struct {
	client clients.CourseClient
}

type ICourseService interface {
	CreateCourse(createCourseDto courses.CreateCoursesRequestDto) courses.CreateCoursesResponseDto
}

func NewCourseService(client *clients.CourseClient) ICourseService {
	return &CourseService{client: *client}
}

func (c *CourseService) CreateCourse(createCourseDto courses.CreateCoursesRequestDto) courses.CreateCoursesResponseDto {
	newCourse := model.Course{
		CourseName:        createCourseDto.CourseName,
		CourseDescription: createCourseDto.CourseDescription,
		CoursePrice:       createCourseDto.CoursePrice,
		CourseDuration:    createCourseDto.CourseDuration,
		CourseCapacity:    createCourseDto.CourseCapacity,
		CategoryID:        createCourseDto.CategoryID,
		CourseInitDate:    createCourseDto.CourseInitDate,
		CourseState:       createCourseDto.CourseState,
	}

	course := c.client.CreateCourse(newCourse)

	return courses.CreateCoursesResponseDto{
		CourseId:   uint(course.Id),
		CourseName: course.CourseName,
	}
}
