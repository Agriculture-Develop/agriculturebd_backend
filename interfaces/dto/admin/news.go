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
