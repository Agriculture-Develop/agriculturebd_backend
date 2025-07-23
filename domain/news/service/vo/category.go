package vo

// CategorySvcVO 分类服务层VO
type CategorySvcVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   uint   `json:"sort_order"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CategoryListSvcVO 分类列表服务层VO
type CategoryListSvcVO struct {
	List []CategorySvcVO `json:"list"`
}
