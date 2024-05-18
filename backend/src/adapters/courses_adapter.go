package adapters

import (
	"backend/src/clients"
	controllers "backend/src/controllers/courses"
	"backend/src/services/courses"

	"gorm.io/gorm"
)

func CoursesAdapter(db *gorm.DB) *controllers.CoursesControllers {
	client := clients.NewCoursesClient(db)
	service := courses.NewCourseService(client)
	return controllers.NewCoursesControllers(service)
}
