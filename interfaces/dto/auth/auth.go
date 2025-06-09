package dto

type LoginByPwdSCtrlDTO struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByCodeCtrlDTO struct {
	Phone    string `json:"phone" binding:"required"`
	AuthCode string `json:"auth_code" binding:"required,len=6,numeric"`
}

type RegisterCtrlDTO struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	AuthCode string `json:"auth_code" binding:"required,len=6,numeric"`
}

type SendCodeCtrlDTO struct {
	Phone string `json:"phone" binding:"required"`
}
