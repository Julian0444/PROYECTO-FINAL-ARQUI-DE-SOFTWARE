package categories

type CreateCategoryRequestDto struct {
	CategoryName string `json:"category_name"`
}

type CreateCategoryResponseDto struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
