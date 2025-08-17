package news

// NewsCreateDTO 创建新闻请求
type NewsCreateDTO struct {
	Title      string   `json:"title" binding:"required"`
	CategoryID uint     `json:"category_id" binding:"required"`
	Abstract   string   `json:"abstract"`
	Type       string   `json:"type"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content" binding:"required"`
	Status     string   `json:"status"`
	Cover      string   `json:"cover" binding:"required"` // 封面图
	Files      []string `json:"files"`                    // 普通多图
}

// NewsUpdateDTO 更新新闻请求
type NewsUpdateDTO struct {
	Title      string   `json:"title" binding:"required"`
	CategoryID uint     `json:"category_id" binding:"required"`
	Abstract   string   `json:"abstract"`
	Type       string   `json:"type"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content" binding:"required"`
	Status     string   `json:"status"`
	Cover      string   `json:"cover" binding:"required"` // 封面图
	Files      []string `json:"files"`                    // 普通多图
}

// NewsListFilterDTO 新闻列表筛选请求
type NewsListFilterDTO struct {
	Title  string `form:"title"`
	Author string `form:"author"`
	Status string `form:"status"`
	Page   int    `form:"page"`
	Count  int    `form:"count"`
}
