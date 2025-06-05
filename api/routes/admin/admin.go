package admin

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin/Interface"
	"github.com/gin-gonic/gin"
)

func Models(r *gin.RouterGroup, userCtrl Interface.IUserCtrl) {

	r.POST("/login", userCtrl.Login)

	users := r.Group("/user")
	{

		users.GET("", userCtrl.GetUserList)

		users.POST("", userCtrl.AddUser)

		users.PUT("", userCtrl.ModifyUserInfo)

		users.DELETE("", userCtrl.DeleteUser)
	}

}
