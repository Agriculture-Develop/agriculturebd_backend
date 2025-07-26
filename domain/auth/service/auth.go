package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/constant"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/valobj"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/service/vo"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizcode"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizerr"

	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/random"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthSvc interface {
	LoginByPassword(ctx context.Context, username, password string) (vo.LoginSvcVo, *bizerr.BizErr)
	LoginByCode(phone, code string) (respCode.StatusCode, vo.LoginSvcVo)
	Register(password, phone, code string) respCode.StatusCode
	SendPhoneCode(phone string) respCode.StatusCode
	UpdatePassword(phone string, authCode string, newPassword string) respCode.StatusCode
}

type Svc struct {
	dig.In
	Repo     repository.IAuthRepo
	SmsUtils repository.ISMSUtils
	biz      bizerr.Biz
}

func NewAuthSvc(r repository.IAuthRepo, sms repository.ISMSUtils) IAuthSvc {
	return &Svc{Repo: r, SmsUtils: sms, biz: bizerr.NewBiz("auth")}
}

func (a *Svc) LoginByPassword(ctx context.Context, phone, password string) (vo.LoginSvcVo, *bizerr.BizErr) {
	// 0. 验证参数
	if !entity.CheckPhone(phone) || !entity.CheckPassword(password) {
		return vo.LoginSvcVo{}, a.biz.CodeErr(bizcode.InvalidParams).WithExtraMsg(errors.New("手机号或者格式错误"))
	}

	// 1. 获取用户信息
	user, err := a.Repo.GetUserByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || user.ID == 0 {
			return vo.LoginSvcVo{}, a.biz.CodeErr(bizcode.RecordNotFound)
		} else {
			return vo.LoginSvcVo{}, a.biz.CodeErr(bizcode.RecordNotFound).Log(ctx, "CheckPassword fail")
		}
	}

	// 2. 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return vo.LoginSvcVo{}, a.biz.CodeErr(bizcode.BadRequest, fmt.Errorf("用户密码错误"))
	}

	// 3. 生成token
	token, err := a.Repo.GenerateToken(user.ID, user.Role.Int())
	if err != nil {
		return vo.LoginSvcVo{}, a.biz.CodeErr(bizcode.ServerBusy).Log(ctx, "GenerateToken fail")
	}

	return vo.LoginSvcVo{
		Id:    user.ID,
		Token: token,
		Role:  user.Role.Desc(),
	}, nil
}

func (a *Svc) LoginByCode(phone, code string) (respCode.StatusCode, vo.LoginSvcVo) {
	// 0. 参数校验
	if !entity.CheckPhone(phone) {
		return respCode.InvalidParamsFormat, vo.LoginSvcVo{}
	}

	// 1. 验证验证码
	if !a.Repo.VerifyPhoneCode(phone, code) {
		return respCode.InvalidCaptcha, vo.LoginSvcVo{}
	}

	// 2. 获取用户信息
	user, err := a.Repo.GetUserByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || user.ID == 0 {
			return respCode.UserNotExist, vo.LoginSvcVo{}
		} else {
			zap.L().Error("CheckPassword fail", zap.Error(err))
			return respCode.ServerBusy, vo.LoginSvcVo{}
		}
	}

	// 3. 生成token
	token, err := a.Repo.GenerateToken(user.ID, user.Role.Int())
	if err != nil {
		zap.L().Error("GenerateToken fail", zap.Error(err))
		return respCode.ServerBusy, vo.LoginSvcVo{}
	}

	return respCode.Success, vo.LoginSvcVo{
		Id:    user.ID,
		Token: token,
		Role:  user.Role.Desc(),
	}
}

func (a *Svc) Register(password, phone, code string) respCode.StatusCode {
	// 0. 参数校验
	if !entity.CheckPhone(phone) || !entity.CheckPassword(password) {
		return respCode.InvalidParamsFormat
	}

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
	user := entity.NewUser(phone, string(hashedPassword), valobj.RoleUser)

	if err = a.Repo.CreateUser(user); err != nil {
		return respCode.ServerBusy
	}

	return respCode.Success
}

func (a *Svc) SendPhoneCode(phone string) respCode.StatusCode {
	// 0. 参数校验
	if !entity.CheckPhone(phone) {
		return respCode.InvalidParamsFormat
	}

	// 1. 生成6位随机验证码
	code := random.GetRandomNum(constant.CaptchaLens)

	// 2. 调用短信服务发送验证码
	err := a.SmsUtils.SendCaptcha(phone, code)
	if err != nil {
		zap.L().Error("SendCaptcha fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 3. 保存验证码
	if err := a.Repo.SavePhoneCode(phone, code); err != nil {
		return respCode.ServerBusy
	}

	return respCode.Success
}

func (a *Svc) UpdatePassword(phone string, authCode string, newPassword string) respCode.StatusCode {
	// 0. 参数校验
	if !entity.CheckPassword(newPassword) {
		return respCode.InvalidParamsFormat
	}

	// 1. 获取用户信息
	user, err := a.Repo.GetUserByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist
		}
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 验证码验证
	if !a.Repo.VerifyPhoneCode(phone, authCode) {
		return respCode.InvalidCaptcha
	}

	// 3. 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("GeneratePassword fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 4. 更新密码
	if err := a.Repo.UpdateNewPassword(user.ID, string(hashedPassword)); err != nil {
		zap.L().Error("UpdateUser fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}
