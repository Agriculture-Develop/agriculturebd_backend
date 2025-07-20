package entity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SupplyDemand struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	TagWeigh string `json:"tag_count"`
	TagName  string `json:"tag_name"`
	TagPrice string `json:"tag_price"`

	CoverURL string         `json:"cover_url"`
	FilesURL datatypes.JSON `json:"files_url"`
	Likes    int            `json:"like"`
	UserId   uint           `json:"user_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type SupplyDemandComment struct {
	ID             int64  `json:"id"`
	SupplyDemandID int64  `json:"supply_demand_id"`
	UserID         int64  `json:"user_id"`
	CommentContent string `json:"comment_content"`
	LikeCount      int    `json:"like_count"`
	ReplyId        int64  `json:"reply_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
