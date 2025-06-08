package user

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
)

type Ctrl struct {
	dig.In
}

func NewUserCtrl() Interface.IUserCtrl {
	return new(Ctrl)
}

// 获取用户列表
func (api *Ctrl) GetUserList(c *gin.Context) {
	//ctx := controller.NewAPiContext[struct{}](c)

	//page, count, err := ctx.GetPageAndCount()
	//if err != nil {
	//	ctx.NoDataJSON(resp.CodeInvalidParams)
	//	return
	//}

	//statusCode, users := api.Services.GetUserList(page, count)
	//ctx.WithDataJSON(statusCode, resp.GetUserListResponse{UserList: users})
}

func (api *Ctrl) AddUser(c *gin.Context) {
	//ctx := controller.NewAPiCtrl[resp.AddUserRequestData](c)
	//
	//if err := ctx.BindJSON(); err != nil {
	//	ctx.NoDataJSON(apiCode.CodeInvalidParams)
	//	return
	//}
	//
	//statusCode := api.Services.AddUser(ctx.Request)
	//ctx.WithDataJSON(statusCode, nil)
}

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
	//ctx.WithDataJSON(statusCode, resp.ModifyUserInfoResponse{User: data})
}

func (api *Ctrl) ModifyPassword(c *gin.Context) {}

func (api *Ctrl) ModifyRole(c *gin.Context) {

}

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
