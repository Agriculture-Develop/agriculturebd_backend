package model

import "time"

type NewsCategories struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;comment:分类ID" json:"id"`
	Name        string `gorm:"type:varchar(50);not null;uniqueIndex;comment:分类名称" json:"name"`
	Description string `gorm:"type:text;comment:分类描述" json:"description"`
	SortOrder   uint   `gorm:"default:0;comment:排序权重" json:"sort_order"`

	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}
