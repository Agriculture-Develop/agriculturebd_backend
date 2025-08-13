package dto

// SupplyDemandCreateSvcDTO 创建供需服务层DTO
type SupplyDemandCreateSvcDTO struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	CoverURL string   `json:"cover_url"`
	FilesURL []string `json:"files_url"`
	TagName  string   `json:"tag_name"`
	TagPrice string   `json:"tag_price"`
	TagWeigh string   `json:"tag_weigh"`
	UserID   uint     `json:"user_id"`
}

// SupplyDemandListFilterSvcDTO 供需列表筛选服务层DTO
type SupplyDemandListFilterSvcDTO struct {
	Title string `json:"title"`
	Page  int    `json:"page"`
	Count int    `json:"count"`
}

// CommentCreateSvcDTO 创建评论服务层DTO
type CommentCreateSvcDTO struct {
	SupplyDemandID int64  `json:"supply_demand_id"`
	CommentContent string `json:"comment_content"`
	UserID         uint   `json:"user_id"`
	ReplyID        int64  `json:"reply_id"`
}
