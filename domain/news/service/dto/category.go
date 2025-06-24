package dto

// CategoryCreateSvcDTO 创建分类服务层DTO
type CategoryCreateSvcDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CategoryUpdateSvcDTO 更新分类服务层DTO
type CategoryUpdateSvcDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
}
