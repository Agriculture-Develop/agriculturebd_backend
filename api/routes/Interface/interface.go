package Interface

import "github.com/gin-gonic/gin"

type IUserCtrl interface {
	AddUser(c *gin.Context)
	ModifyUserInfo(c *gin.Context)
	GetUserList(c *gin.Context)
	DeleteUser(c *gin.Context)
	ModifyRole(c *gin.Context)
	ModifyPassword(c *gin.Context)
}

type IAuthCtrl interface {
	LoginByPassword(c *gin.Context)
	LoginByCode(c *gin.Context)
	Register(c *gin.Context)
	SendPhoneCode(c *gin.Context)
}
