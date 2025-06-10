package Interface

import "github.com/gin-gonic/gin"

type IUserCtrl interface {
	UpdateUserInfo(c *gin.Context)
	GetUserList(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type IAuthCtrl interface {
	LoginByPassword(c *gin.Context)
	LoginByCode(c *gin.Context)
	Register(c *gin.Context)
	SendPhoneCode(c *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
}
