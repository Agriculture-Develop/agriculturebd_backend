package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         int64  `gorm:"primarykey;NOT NULL;comment:用户ID" json:"id"`    // 传给前端时应当转为字符串避免精度丢失
	Password   string `json:"-"`                                             // 用户密码
	AvatarPath string `json:"avatar_path"`                                   // 用户头像的Url
	Nickname   string `gorm:"type:varchar(32);comment:用户昵称" json:"nickname"` // 昵称
	Role       int    `gorm:"comment:用户身份，具体参照文档" json:"role"`               // 用户身份，用户的权限控制
	Biography  string `json:"biography"`                                     // 个性签名
	TrueName   string `json:"true_name"`
	GradeTitle string `json:"grade_title"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
