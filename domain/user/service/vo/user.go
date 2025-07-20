package vo

type UserSvcVo struct {
	ID         uint   `json:"id"`
	Phone      string `json:"phone"`
	Nickname   string `json:"nickname"`
	AvatarPath string `json:"avatar_path"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}
