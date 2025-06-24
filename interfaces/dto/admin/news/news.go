package news

// NewsCreateDTO 创建新闻请求
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
	Title  string `form:"title"`
	Author string `form:"author"`
	Status string `form:"status"`
	Page   int    `form:"page"`
	Count  int    `form:"count"`
}
