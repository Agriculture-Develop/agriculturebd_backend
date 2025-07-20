package public

import "mime/multipart"

// SupplyDemandCreateCtrlDTO 创建供需请求
type SupplyDemandCreateCtrlDTO struct {
	Title    string                  `form:"title" binding:"required"`
	Content  string                  `form:"content" binding:"required"`
	Cover    *multipart.FileHeader   `form:"cover" binding:"required"`
	Files    []*multipart.FileHeader `form:"files"`
	TagName  string                  `form:"tag_name"`
	TagPrice string                  `form:"tag_price"`
	TagWeigh string                  `form:"tag_weigh"`
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
