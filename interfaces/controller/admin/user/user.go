package user

import (
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	//"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/apiCode"
	//"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/apiModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Api struct{}

func NewApi() *Api {
	return new(Api)
}

// 管理员登录
func (api *Api) Login(c *gin.Context) {
	// 待实现
}

// 获取用户列表
func (api *Api) GetUserList(c *gin.Context) {
	ctx := controller.NewAPiCtrl[struct{}](c)

	page, count, err := ctx.GetPageAndCount()
	if err != nil {
		ctx.NoDataJSON(apiCode.CodeInvalidParams)
		return
	}

	statusCode, users := api.Services.GetUserList(page, count)
	ctx.WithDataJSON(statusCode, apiModel.GetUserListResponse{UserList: users})
}

// 添加用户
func (api *Api) AddUser(c *gin.Context) {
	ctx := controller.NewAPiCtrl[apiModel.AddUserRequestData](c)

	if err := ctx.BindJSON(); err != nil {
		ctx.NoDataJSON(apiCode.CodeInvalidParams)
		return
	}

	statusCode := api.Services.AddUser(ctx.Request)
	ctx.WithDataJSON(statusCode, nil)
}

// 修改用户信息
func (api *Api) ModifyUserInfo(c *gin.Context) {
	ctx := controller.NewAPiCtrl[struct{}](c)

	userId := ctx.GetUserID()
	rawData, err := c.GetRawData()
	if err != nil || len(rawData) == 0 {
		ctx.NoDataJSON(apiCode.CodeInvalidParams)
		return
	}

	statusCode, data := api.Services.ModifyInfo(rawData, userId)
	ctx.WithDataJSON(statusCode, apiModel.ModifyUserInfoResponse{User: data})
}

// 删除用户
func (api *Api) DeleteUser(c *gin.Context) {
	ctx := controller.NewAPiCtrl[struct{}](c)

	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		ctx.NoDataJSON(apiCode.CodeInvalidParams)
		return
	}

	statusCode := api.Services.DeleteUser(id)
	ctx.NoDataJSON(statusCode)
}
