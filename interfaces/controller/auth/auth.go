package auth

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/service"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	dto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Ctrl struct {
	dig.In

	Services service.IAuthSvc
}

func NewAuthCtrl(srv service.IAuthSvc) Interface.IAuthCtrl {
	return &Ctrl{
		Services: srv,
	}
}

// 密码登录
func (api *Ctrl) LoginByPassword(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.LoginByPwdSCtrlDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(respCode.InvalidParams)
		return
	}

	statusCode, loginSvcVo := api.Services.LoginByPassword(ctx.Request.Phone, ctx.Request.Password)
	ctx.WithDataJSON(statusCode, loginSvcVo)
}

// 验证码登录
func (api *Ctrl) LoginByCode(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.LoginByCodeCtrlDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(respCode.InvalidParams)
		return
	}

	statusCode, loginSvcVo := api.Services.LoginByCode(ctx.Request.Phone, ctx.Request.AuthCode)
	ctx.WithDataJSON(statusCode, loginSvcVo)
}

// 注册
func (api *Ctrl) Register(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.RegisterCtrlDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(respCode.InvalidParams)
		return
	}

	statusCode := api.Services.Register(ctx.Request.Password, ctx.Request.Phone, ctx.Request.AuthCode)
	ctx.NoDataJSON(statusCode)
}

// 发送手机验证码
func (api *Ctrl) SendPhoneCode(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.SendCodeCtrlDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(respCode.InvalidParams)
		return
	}

	statusCode := api.Services.SendPhoneCode(ctx.Request.Phone)
	ctx.NoDataJSON(statusCode)
}

// UpdateUserPassword 更新用户密码
func (c *Ctrl) UpdateUserPassword(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[dto.UpdatePasswordCtrlDto](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code := c.Services.UpdatePassword(apiCtx.Request.Phone, apiCtx.Request.AuthCode, apiCtx.Request.NewPassword)
	apiCtx.NoDataJSON(code)
}
