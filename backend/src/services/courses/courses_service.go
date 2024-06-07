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
	GetCourseById(id int) model.Course
	GetAllCourses(terms string) []model.Course
	CreateCourse(createCourseDto courses.CreateCoursesRequestDto) courses.CreateCoursesResponseDto
	UpdateCourse(updateCourseDto courses.UpdateCoursesRequestDto) courses.UpdateCoursesResponseDto
}

func NewCourseService(client *clients.CourseClient) ICourseService {
	return &CourseService{client: *client}
}

func (c *CourseService) GetCourseById(id int) model.Course {
	return c.client.GetCourseById(id)
}

func (c *CourseService) GetAllCourses(terms string) []model.Course {
	return c.client.GetAllCourses(terms)
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

func (c *CourseService) UpdateCourse(updateCourseDto courses.UpdateCoursesRequestDto) courses.UpdateCoursesResponseDto {

	checkCourse := c.client.GetCourseById(int(updateCourseDto.CourseId))

	if checkCourse.Id == 0 {
		panic("Course not found")
	}

	course := model.Course{
		Id:                updateCourseDto.CourseId,
		CourseName:        updateCourseDto.CourseName,
		CourseDescription: updateCourseDto.CourseDescription,
		CoursePrice:       updateCourseDto.CoursePrice,
		CourseDuration:    updateCourseDto.CourseDuration,
		CourseCapacity:    updateCourseDto.CourseCapacity,
		CategoryID:        updateCourseDto.CategoryID,
		CourseInitDate:    updateCourseDto.CourseInitDate,
		CourseState:       updateCourseDto.CourseState,
	}

	course = c.client.UpdateCourse(course)

	return courses.UpdateCoursesResponseDto{
		CourseId:   uint(course.Id),
		CourseName: course.CourseName,
	}
}
