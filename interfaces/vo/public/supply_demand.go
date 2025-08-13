package public

// SupplyDemandDetailVO 供需详情
type SupplyDemandDetailVO struct {
	ID            uint     `json:"id"`
	UserId        uint     `json:"user_id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	CoverURL      string   `json:"cover_url"`
	FilesURL      []string `json:"files_url"`
	PublisherName string   `json:"publisher_name"`
	CreatedAt     string   `json:"created_at"`
	Like          string   `json:"like"`
	Tags          TagsVO   `json:"tags"`
}

// TagsVO 标签信息
type TagsVO struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Weight string `json:"weight"`
}

// SupplyDemandListVO 供需列表
type SupplyDemandListVO struct {
	Total int                  `json:"total"`
	List  []SupplyDemandItemVO `json:"list"`
}

// SupplyDemandItemVO 供需列表项
type SupplyDemandItemVO struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"userid"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagName   string `json:"tag_name"`
	TagWeigh  string `json:"tag_weigh"`
	TagPrice  string `json:"tag_price"`
	CoverURL  string `json:"cover_url"`
	Like      string `json:"like"`
}

// CommentDetailVO 评论详情
type CommentDetailVO struct {
	Avatar        string `json:"avatar"`
	ID            string `json:"id"`
	PublisherName string `json:"publisher_name"`
	Comment       string `json:"comment"`
	Role          string `json:"role"`
	Like          string `json:"like"`
	UserId        uint   `json:"user_id"`
	ReplyId       int64  `json:"reply_id"`
	CreatedAt     string `json:"created_at"`
}

// CommentListVO 评论列表
type CommentListVO struct {
	Total int               `json:"total"`
	List  []CommentDetailVO `json:"list"`
}
