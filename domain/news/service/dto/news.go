package dto

// NewsCreateSvcDTO 创建新闻服务层DTO
type NewsCreateSvcDTO struct {
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Abstract   string   `json:"abstract"`
	Types      string   `json:"types"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
	Status     string   `json:"status"`
	UserID     uint     `json:"user_id"`
}

// NewsUpdateSvcDTO 更新新闻服务层DTO
type NewsUpdateSvcDTO struct {
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Abstract   string   `json:"abstract"`
	Type       string   `json:"types"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
	Status     string   `json:"status"`
	UserID     uint     `json:"user_id"`
}

// NewsListFilterSvcDTO 新闻列表筛选服务层DTO
type NewsListFilterSvcDTO struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Status   string `json:"status"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
