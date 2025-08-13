package model

import (
	"gorm.io/gorm"
	"time"
)

type SupplyDemand struct {
	BaseModel

	Title    string   `gorm:"type:varchar(255);not null;comment:标题" json:"title"`
	Content  string   `gorm:"type:text;comment:内容" json:"content"`
	CoverURL string   `gorm:"type:varchar(512);default:'';comment:封面图地址" json:"cover_url"`
	FilesURL []string `gorm:"serializer:json;type:json;comment:附件地址列表" json:"files_url"`
	Likes    int      `gorm:"type:int;default:0;comment:点赞数" json:"likes"`

	Tag    TagInfo `gorm:"embedded;embeddedPrefix:tag_" json:"tag"`
	UserID uint    `gorm:"index;comment:用户id" json:"user_id"`
}

type TagInfo struct {
	Weigh string `gorm:"type:varchar(255);default:0;comment:标签" json:"weigh"`
	Name  string `gorm:"type:varchar(255);default:'';comment:标签名称" json:"name"`
	Price string `gorm:"type:varchar(255);default:'';comment:标签价格" json:"price"`
}

type SupplyDemandComment struct {
	ID             int64  `gorm:"primaryKey;column:id" json:"id"`
	SupplyDemandID int64  `gorm:"column:supply_demand_id" json:"supply_demand_id"` // 外键：供需ID
	UserID         int64  `gorm:"column:user_id" json:"user_id"`                   // 外键：用户ID
	CommentContent string `gorm:"column:comment_content" json:"comment_content"`   // 评论内容
	LikeCount      int    `gorm:"column:like_count" json:"like_count"`             // 点赞数
	ReplyId        int64  `gorm:"index;column:reply_id" json:"reply_id"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}

func (SupplyDemand) TableName() string {
	return "supply_demand"
}

func (SupplyDemandComment) TableName() string {
	return "supply_demand_comment"
}
