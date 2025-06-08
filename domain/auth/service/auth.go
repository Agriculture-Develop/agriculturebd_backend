package service

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/constant"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/service/vo"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/random"
	"go.uber.org/dig"
	"golang.org/x/crypto/bcrypt"
)

type IAuthSvc interface {
	LoginByPassword(username, password string) (respCode.StatusCode, vo.LoginSvcVo)
	LoginByCode(phone, code string) (respCode.StatusCode, vo.LoginSvcVo)
	Register(password, phone, code string) respCode.StatusCode
	SendPhoneCode(phone string) respCode.StatusCode
}

type Svc struct {
	dig.In
	Repo repository.IAuthRepo
}

func NewAuthSvc(r repository.IAuthRepo) IAuthSvc {
	return &Svc{Repo: r}
}

func (a *Svc) LoginByPassword(phone, password string) (respCode.StatusCode, vo.LoginSvcVo) {
	// 1. 获取用户信息
	user, err := a.Repo.GetUserByPhone(phone)
	if err != nil {
		return respCode.UserNotExist, vo.LoginSvcVo{}
	}

	// 2. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return respCode.InvalidPassword, vo.LoginSvcVo{}
	}

	// 3. 生成token
	token, err := a.Repo.GenerateToken(user.ID)
	if err != nil {
		return respCode.ServerBusy, vo.LoginSvcVo{}
	}

	return respCode.Success, vo.LoginSvcVo{
		Id:    user.ID,
		Token: token,
	}
}

func (a *Svc) LoginByCode(phone, code string) (respCode.StatusCode, vo.LoginSvcVo) {
	// 1. 验证验证码
	if !a.Repo.VerifyPhoneCode(phone, code) {
		return respCode.InvalidCaptcha, vo.LoginSvcVo{}
	}

	// 2. 获取用户信息
	user, err := a.Repo.GetUserByPhone(phone)
	if err != nil {
		return respCode.UserNotExist, vo.LoginSvcVo{}
	}

	// 3. 生成token
	token, err := a.Repo.GenerateToken(user.ID)
	if err != nil {
		return respCode.ServerBusy, vo.LoginSvcVo{}
	}

	return respCode.Success, vo.LoginSvcVo{
		Id:    user.ID,
		Token: token,
	}
}

func (a *Svc) Register(password, phone, code string) respCode.StatusCode {
	// 1. 验证验证码
	if !a.Repo.VerifyPhoneCode(phone, code) {
		return respCode.InvalidCaptcha
	}

	// 2. 检查用户是否已存在
	if user, _ := a.Repo.GetUserByPhone(phone); user != nil {
		return respCode.UserExist
	}

	// 3. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return respCode.ServerBusy
	}

	// 4. 创建用户
	user := &entity.User{
		Phone:    phone,
		Password: string(hashedPassword),
	}
	if err = a.Repo.CreateUser(user); err != nil {
		return respCode.ServerBusy
	}

	return respCode.Success
}

func (a *Svc) SendPhoneCode(phone string) respCode.StatusCode {
	// 生成6位随机验证码
	code := random.GetRandomNum(constant.CaptchaLens)

	// 保存验证码
	if err := a.Repo.SavePhoneCode(phone, code); err != nil {
		return respCode.ServerBusy
	}

	// TODO: 调用短信服务发送验证码

	return respCode.Success
}
