package admin

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/gin-gonic/gin"
)

func Models(r *gin.RouterGroup, userCtrl Interface.IUserCtrl) {

	users := r.Group("/user")
	{

		users.GET("", userCtrl.GetUserList)

		users.POST("", userCtrl.AddUser)

		users.PUT("", userCtrl.ModifyUserInfo)

		users.DELETE("", userCtrl.DeleteUser)
	}

}
