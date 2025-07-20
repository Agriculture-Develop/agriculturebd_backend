package admin

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"github.com/gin-gonic/gin"
)

func UserModels(r *gin.RouterGroup, ctrl Interface.IUserCtrl) {
	// 用户管理
	userGroup := r.Group("/user", middleware.Auth(), middleware.WithSuperAdmin())
	{
		userGroup.GET("list", ctrl.GetUserList)           // 获取用户列表
		userGroup.PUT("/:id", ctrl.UpdateUserInfoByAdmin) // 更新用户信息
		userGroup.DELETE("/:id", ctrl.DeleteUser)         // 删除用户
	}
}
