package user

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/service"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	userDto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin"
	userVo "github.com/Agriculture-Develop/agriculturebd/interfaces/vo/admin"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Ctrl struct {
	dig.In
	Services service.IUserSvc
}

func NewUserCtrl(srv service.IUserSvc) Interface.IUserCtrl {
	return &Ctrl{Services: srv}
}

// GetUserList 获取用户列表
func (c *Ctrl) GetUserList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	page, count, err := apiCtx.GetPageAndCount()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code, users := c.Services.GetUserList(page, count)
	if code != respCode.Success {
		apiCtx.NoDataJSON(code)
		return
	}

	// 转换为接口层VO
	userVos := make([]userVo.UserCtrlVo, 0, len(users))
	for _, user := range users {
		userVos = append(userVos, userVo.UserCtrlVo{
			ID:         user.ID,
			Phone:      user.Phone,
			Nickname:   user.Nickname,
			AvatarPath: user.AvatarPath,
			Role:       user.Role,
			Status:     user.Status,
			CreatedAt:  user.CreatedAt,
		})
	}

	apiCtx.WithDataJSON(code, userVo.UserListCtrlVo{
		List:  userVos,
		Total: int64(len(users)),
	})
}

// UpdateUserInfo 更新用户信息
func (c *Ctrl) UpdateUserInfoByAdmin(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[userDto.UpdateUserInfoCtrlDto](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	userId, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code := c.Services.UpdateUserInfo(userId, apiCtx.Request.Nickname, apiCtx.Request.Role, apiCtx.Request.Status)
	apiCtx.NoDataJSON(code)
}

// DeleteUser 删除用户
func (c *Ctrl) DeleteUser(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)

	userId, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code := c.Services.DeleteUser(userId)
	apiCtx.NoDataJSON(code)
}

// 公用模块

func (c *Ctrl) GetUserDetail(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id := apiCtx.GetUserIdByToken()

	code, user := c.Services.GetUserDetail(id)

	apiCtx.WithDataJSON(code, user)

}

func (c *Ctrl) UpdateUserAvatar(ctx *gin.Context) {

	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id := apiCtx.GetUserIdByToken()

	// 从表单中获取头像
	file, err := ctx.FormFile("avatar")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code := c.Services.UpdateUserAvatar(id, file)

	apiCtx.NoDataJSON(code)
}

// UpdateUserInfo 更新用户信息
func (c *Ctrl) UpdateUserInfoByUser(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[userDto.UpdateUserInfoCtrlDtoByUser](ctx)
	if err := apiCtx.BindForm(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}
	userId := apiCtx.GetUserIdByToken()

	code := c.Services.UpdateUserInfoByUser(userId, apiCtx.Request.Nickname, apiCtx.Request.Role, apiCtx.Request.Avatar)
	apiCtx.NoDataJSON(code)
}
