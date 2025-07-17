package news

import "mime/multipart"

// NewsCreateDTO 创建新闻请求
type NewsCreateDTO struct {
	Title      string                  `form:"title" binding:"required"`
	CategoryID uint                    `form:"category_id" binding:"required"`
	Abstract   string                  `form:"abstract"`
	Type       string                  `form:"type"`
	Keyword    []string                `form:"keyword"`
	Source     string                  `form:"source"`
	Content    string                  `form:"content" binding:"required"`
	Status     string                  `form:"status"`
	Cover      *multipart.FileHeader   `form:"cover" binding:"required"` // 封面图
	Files      []*multipart.FileHeader `form:"files"`                    // 普通多图
}

// NewsUpdateDTO 更新新闻请求
type NewsUpdateDTO struct {
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Abstract   string   `json:"abstract"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
}

// NewsListFilterDTO 新闻列表筛选请求
type NewsListFilterDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
	Page   int    `json:"page"`
	Count  int    `json:"count"`
}
