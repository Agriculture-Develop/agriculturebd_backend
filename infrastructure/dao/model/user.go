package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户表
type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	Password   string `gorm:"type:varchar(255);not null;comment:用户密码" json:"-"`
	AvatarPath string `gorm:"type:varchar(255);default:'';comment:头像地址" json:"avatar_path"`
	Nickname   string `gorm:"type:varchar(32);not null;comment:用户昵称" json:"nickname"`
	Role       int    `gorm:"type:tinyint;default:0;not null;comment:用户角色" json:"role"`
	Status     int    `gorm:"type:tinyint(1);default:0;comment:用户状态(0启用 1禁用)" json:"status"`
	Phone      string `gorm:"type:varchar(20);not null;unique;comment:手机号" json:"phone"`

	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}
