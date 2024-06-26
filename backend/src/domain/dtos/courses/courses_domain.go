package courses

type CreateCoursesRequestDto struct {
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"description"`
	CoursePrice       float64 `json:"price"`
	CourseDuration    int     `json:"duration"`
	CourseCapacity    int     `json:"capacity"`
	CategoryID        uint    `json:"category_id"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
}

type CreateCoursesResponseDto struct {
	CourseName string `json:"course_name"`
	CourseId   uint   `json:"course_id"`
}

type GetOneCourseResponseDto struct {
	CategoryID        uint    `json:"category_id"`
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"description"`
	CoursePrice       float64 `json:"price"`
	CourseDuration    int     `json:"duration"`
	CourseCapacity    int     `json:"capacity"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
}

type UpdateCoursesRequestDto struct {
	CourseId          uint    `json:"id"`
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"description"`
	CoursePrice       float64 `json:"price"`
	CourseDuration    int     `json:"duration"`
	CourseCapacity    int     `json:"capacity"`
	CategoryID        uint    `json:"category_id"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
}

type UpdateCoursesResponseDto struct {
	CourseId   uint   `json:"course_id"`
	CourseName string `json:"course_name"`
}
