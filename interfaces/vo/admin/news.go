package vo

import "time"

type NewsDetailVO struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	CategoryID uint      `json:"category_id"`
	Category   string    `json:"category"`
	Abstract   string    `json:"abstract"`
	Keyword    []string  `json:"keyword"`
	Source     string    `json:"source"`
	Content    string    `json:"content"`
	CoverURL   string    `json:"cover_url"`
	FilesURL   []string  `json:"files_url"`
	Status     string    `json:"status"`
	Author     string    `json:"author"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type NewsListVO struct {
	Total int            `json:"total"`
	List  []NewsDetailVO `json:"list"`
}

type CategoryVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
}

type CategoryListVO struct {
	List []CategoryVO `json:"list"`
}
