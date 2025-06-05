package user

import (
	"go.uber.org/dig"

	//"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/apiCode"
	//"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/apiModel"
	"github.com/gin-gonic/gin"
)

type Ctrl struct {
	dig.In
}

func NewUserCtrl() *Ctrl {
	return new(Ctrl)
}

// 管理员登录
func (api *Ctrl) Login(c *gin.Context) {
	// 待实现
	c.JSON(200, gin.H{"code": 200, "data": "success"})
}

// 获取用户列表
func (api *Ctrl) GetUserList(c *gin.Context) {
	//ctx := controller.NewAPiCtrl[struct{}](c)
	//
	//page, count, err := ctx.GetPageAndCount()
	//if err != nil {
	//	ctx.NoDataJSON(apiCode.CodeInvalidParams)
	//	return
	//}
	//
	//statusCode, users := api.Services.GetUserList(page, count)
	//ctx.WithDataJSON(statusCode, apiModel.GetUserListResponse{UserList: users})
}

// 添加用户
func (api *Ctrl) AddUser(c *gin.Context) {
	//ctx := controller.NewAPiCtrl[apiModel.AddUserRequestData](c)
	//
	//if err := ctx.BindJSON(); err != nil {
	//	ctx.NoDataJSON(apiCode.CodeInvalidParams)
	//	return
	//}
	//
	//statusCode := api.Services.AddUser(ctx.Request)
	//ctx.WithDataJSON(statusCode, nil)
}

// 修改用户信息
func (api *Ctrl) ModifyUserInfo(c *gin.Context) {
	//ctx := controller.NewAPiCtrl[struct{}](c)
	//
	//userId := ctx.GetUserID()
	//rawData, err := c.GetRawData()
	//if err != nil || len(rawData) == 0 {
	//	ctx.NoDataJSON(apiCode.CodeInvalidParams)
	//	return
	//}
	//
	//statusCode, data := api.Services.ModifyInfo(rawData, userId)
	//ctx.WithDataJSON(statusCode, apiModel.ModifyUserInfoResponse{User: data})
}

// 删除用户
func (api *Ctrl) DeleteUser(c *gin.Context) {
	//ctx := controller.NewAPiCtrl[struct{}](c)
	//
	//idStr := c.Query("id")
	//id, err := strconv.ParseInt(idStr, 10, 64)
	//if err != nil || id <= 0 {
	//	ctx.NoDataJSON(apiCode.CodeInvalidParams)
	//	return
	//}
	//
	//statusCode := api.Services.DeleteUser(id)
	//ctx.NoDataJSON(statusCode)
}
