package vo

type LoginSvcVo struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
