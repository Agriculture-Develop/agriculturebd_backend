package auth

import (
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	dto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/auth"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Ctrl struct {
	dig.In
}

func NewAuthCtrl() *Ctrl {
	return new(Ctrl)
}

// 密码登录
func (api *Ctrl) LoginByPassword(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.LoginByPwdDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(resp.CodeInvalidParams)
		return
	}

	//statusCode, token, userId := api.Services.LoginByPassword(ctx.Request.Phone, ctx.Request.Password)
	//ctx.WithDataJSON(statusCode, vo.LoginRespVo{Userid: userId, Token: token})
}

// 验证码登录
func (api *Ctrl) LoginByCode(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.LoginByCodeDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(resp.CodeInvalidParams)
		return
	}

	//statusCode, token, userId := api.Services.LoginByCode(ctx.Request.Phone, ctx.Request.AuthCode)
	//ctx.WithDataJSON(statusCode, vo.LoginRespVo{Userid: userId, Token: token})
}

// 注册
func (api *Ctrl) Register(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.RegisterDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(resp.CodeInvalidParams)
		return
	}

	//statusCode := api.Services.Register(ctx.Request)
	//ctx.NoDataJSON(statusCode)
}

// 发送手机验证码
func (api *Ctrl) SendPhoneCode(c *gin.Context) {
	ctx := controller.NewAPiContext[dto.SendCodeDTO](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(resp.CodeInvalidParams)
		return
	}

	//statusCode := api.Services.SendPhoneCode(ctx.Request.Phone)
	//ctx.NoDataJSON(statusCode)
}
