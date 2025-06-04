package Interface

import "github.com/gin-gonic/gin"

type IUserApi interface {
	Login(c *gin.Context)

	AddUser(c *gin.Context)

	ModifyUserInfo(c *gin.Context)

	GetUserList(c *gin.Context)

	DeleteUser(c *gin.Context)
}
