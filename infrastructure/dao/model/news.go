package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type News struct {
	ID       uint           `gorm:"primaryKey;autoIncrement;comment:新闻ID" json:"id"`
	Title    string         `gorm:"type:varchar(255);not null;index;comment:新闻标题" json:"title"`
	Abstract string         `gorm:"type:text;comment:新闻摘要" json:"abstract"`
	Type     string         `gorm:"type:varchar(50);not null;comment:类型" json:"type"`
	Keyword  datatypes.JSON `gorm:"type:json;comment:关键词列表" json:"keyword"`
	Source   string         `gorm:"type:varchar(100);default:'';comment:新闻来源" json:"source"`
	Content  string         `gorm:"type:longtext;comment:新闻内容" json:"content"`
	Status   string         `gorm:"type:varchar(20);default:'draft';index;comment:新闻状态" json:"status"`
	Comment  string         `gorm:"type:text;default:'';comment:审核批注" json:"comment"`
	FilesURL datatypes.JSON `gorm:"type:json;comment:新闻图片地址组" json:"files_url"`
	CoverURL string         `gorm:"type:varchar(512);default:'';comment:封面图地址" json:"cover_url"`

	// 添加外键约束
	UserID     uint `gorm:"index;comment:用户id" json:"user_id"`
	CategoryID uint `gorm:"index;comment:分类id" json:"category_id"`

	PublishedAt *time.Time     `gorm:"index;comment:发布时间" json:"published_at"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}
