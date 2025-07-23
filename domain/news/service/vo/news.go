package vo

// NewsDetailSvcVO 新闻详情服务层VO
type NewsDetailSvcVO struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Category   string   `json:"category"`
	Abstract   string   `json:"abstract"`
	Keyword    []string `json:"keyword"`
	Source     string   `json:"source"`
	Content    string   `json:"content"`
	CoverURL   string   `json:"cover_url"`
	FilesURL   []string `json:"files_url"`
	Status     string   `json:"status"`
	Author     string   `json:"author"`
	Type       string   `json:"type"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

// NewsListSvcVO 新闻列表服务层VO
type NewsListSvcVO struct {
	Total int               `json:"total"`
	List  []NewsDetailSvcVO `json:"list"`
}
