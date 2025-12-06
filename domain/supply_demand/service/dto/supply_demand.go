package dto

// SupplyDemandCreateSvcDTO 创建供需服务层DTO
type SupplyDemandCreateSvcDTO struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	CoverURL string   `json:"cover_url"`
	FilesURL []string `json:"files_url"`
	TagName  string   `json:"tag_name"`
	TagPrice string   `json:"tag_price"`
	TagWeigh string   `json:"tag_weigh"`
	UserID   uint     `json:"user_id"`
}

// SupplyDemandListFilterSvcDTO 供需列表筛选服务层DTO
type SupplyDemandListFilterSvcDTO struct {
	Title     string `json:"title"`
	Category  string `json:"category"`
	UserRole  string `json:"user_role"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	Count     int    `json:"count"`
}

const (
	SortFieldCreatedAt = "created_at"
	SortFieldPrice     = "price"

	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

// CommentCreateSvcDTO 创建评论服务层DTO
type CommentCreateSvcDTO struct {
	SupplyDemandID int64  `json:"supply_demand_id"`
	CommentContent string `json:"comment_content"`
	UserID         uint   `json:"user_id"`
	ReplyID        int64  `json:"reply_id"`
}
