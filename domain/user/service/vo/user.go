package vo

import "time"

type UserSvcVo struct {
	ID         uint      `json:"id"`
	Phone      string    `json:"phone"`
	Nickname   string    `json:"nickname"`
	AvatarPath string    `json:"avatar_path"`
	Role       int       `json:"role"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
