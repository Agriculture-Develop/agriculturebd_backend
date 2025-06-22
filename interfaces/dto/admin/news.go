package dto

type NewsCreateDTO struct {
	Title      string   `json:"title" binding:"required"`
	CategoryID uint     `json:"category_id" binding:"required"`
	Abstract   string   `json:"abstract"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content" binding:"required"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
	Status     string   `json:"status"`
}

type NewsUpdateDTO struct {
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Abstract   string   `json:"abstract"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
	Status     string   `json:"status"`
}

type NewsStatusUpdateDTO struct {
	Status string `json:"status" binding:"required"`
}

type NewsListFilterDTO struct {
	Title  string `form:"title"`
	Author string `form:"author"`
	Status string `form:"status"`
	Page   int    `form:"page"`
	Count  int    `form:"count"`
}

type CategoryCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
}

type CategoryUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
}
