package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type SupplyDemand struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:供需ID" json:"id"`
	Title    string `gorm:"type:varchar(255);not null;comment:标题" json:"title"`
	Content  string `gorm:"type:text;comment:内容" json:"content"`
	TagWeigh string `gorm:"type:varchar(255);default:0;comment:标签" j重量son:"tag_weigh"`
	TagName  string `gorm:"type:varchar(255);default:'';comment:标签名称" json:"tag_name"`
	TagPrice string `gorm:"type:varchar(255);default:'';comment:标签价格" json:"tag_price"`

	CoverURL string         `gorm:"type:varchar(512);default:'';comment:封面图地址" json:"cover_url"`
	FilesURL datatypes.JSON `gorm:"type:json;comment:附件地址列表" json:"files_url"`
	Likes    int            `gorm:"type:int;default:0;comment:点赞数" json:"like"`

	UserID uint `gorm:"index;comment:用户id" json:"user_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}

type SupplyDemandComment struct {
	ID             int64  `gorm:"primaryKey;column:id" json:"id"`
	SupplyDemandID int64  `gorm:"column:supply_demand_id" json:"supply_demand_id"` // 外键：供需ID
	UserID         int64  `gorm:"column:user_id" json:"user_id"`                   // 外键：用户ID
	CommentContent string `gorm:"column:comment_content" json:"comment_content"`   // 评论内容
	LikeCount      int    `gorm:"column:like_count" json:"like_count"`             // 点赞数
	ReplyId        int64  `gorm:"column:reply_id" json:"reply_id"`

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
