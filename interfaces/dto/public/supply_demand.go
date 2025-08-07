package public

// SupplyDemandCreateCtrlDTO 创建供需请求
type SupplyDemandCreateCtrlDTO struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Cover    string   `json:"cover" binding:"required"`
	Files    []string `json:"files"`
	TagName  string   `json:"tag_name"`
	TagPrice string   `json:"tag_price"`
	TagWeigh string   `json:"tag_weigh"`
}

// SupplyDemandListFilterCtrlDTO 供需列表筛选请求
type SupplyDemandListFilterCtrlDTO struct {
	Title string `form:"title"`
	Page  int    `form:"page"`
	Count int    `form:"count"`
}

// CommentCreateCtrlDTO 创建评论请求
type CommentCreateCtrlDTO struct {
	Comment string `json:"comment" binding:"required"`
}
