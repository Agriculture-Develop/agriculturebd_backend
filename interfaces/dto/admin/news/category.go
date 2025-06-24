package news

// CategoryCreateDTO 创建分类请求
type CategoryCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CategoryUpdateDTO 更新分类请求
type CategoryUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
}
