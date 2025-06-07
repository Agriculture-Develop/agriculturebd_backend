package dto

type LoginByPwdDTO struct {
	Phone    string `json:"phone" binding:"required,phone"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}

type LoginByCodeDTO struct {
	Phone    string `json:"phone" binding:"required,phone"`
	AuthCode string `json:"auth_code" binding:"required,len=6,numeric"`
}

type RegisterDTO struct {
	Phone    string `json:"phone" binding:"required,phone"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	AuthCode string `json:"auth_code" binding:"required,len=6,numeric"`
}

type SendCodeDTO struct {
	Phone string `json:"phone" binding:"required,phone"`
}
