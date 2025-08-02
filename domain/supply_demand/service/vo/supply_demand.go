package vo

// SupplyDemandDetailSvcVO 供需详情服务层VO
type SupplyDemandDetailSvcVO struct {
	ID            uint     `json:"id"`
	UserId        uint     `json:"user_id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	CoverURL      string   `json:"cover_url"`
	FilesURL      []string `json:"files_url"`
	PublisherName string   `json:"publisher_name"`
	CreatedAt     string   `json:"created_at"`
	Like          string   `json:"like"`
	TagName       string   `json:"tag_name"`
	TagPrice      string   `json:"tag_price"`
	TagWeigh      string   `json:"tag_weigh"`
}

// SupplyDemandListSvcVO 供需列表服务层VO
type SupplyDemandListSvcVO struct {
	Total int                         `json:"total"`
	List  []SupplyDemandListItemSvcVO `json:"list"`
}

// SupplyDemandListItemSvcVO 供需列表项服务层VO
type SupplyDemandListItemSvcVO struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagName   string `json:"tag_name"`
	TagWeigh  string `json:"tag_weigh"`
	TagPrice  string `json:"tag_price"`
	CoverURL  string `json:"cover_url"`
	Like      string `json:"like"`
}

// CommentDetailSvcVO 评论详情服务层VO
type CommentDetailSvcVO struct {
	Avatar        string `json:"avatar"`
	ID            string `json:"id"`
	PublisherName string `json:"publisher_name"`
	Comment       string `json:"comment"`
	Role          string `json:"role"`
	Like          string `json:"like"`
	UserId        uint   `json:"user_id"`
	CreatedAt     string `json:"created_at"`
}

// CommentListSvcVO 评论列表服务层VO
type CommentListSvcVO struct {
	Total int                  `json:"total"`
	List  []CommentDetailSvcVO `json:"list"`
}
