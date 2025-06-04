package admin

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/admin/user"
	"github.com/gin-gonic/gin"
)

func Models(r *gin.RouterGroup) {

	// 用户管理接口
	var userApi Interface.IUserApi = user.NewApi()

	r.POST("/login", userApi.Login)

	users := r.Group("/user")
	{

		users.GET("", userApi.GetUserList)

		users.POST("", userApi.AddUser)

		users.PUT("", userApi.ModifyUserInfo)

		users.DELETE("", userApi.DeleteUser)
	}

}
