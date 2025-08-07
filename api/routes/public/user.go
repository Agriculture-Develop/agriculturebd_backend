package public

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"github.com/gin-gonic/gin"
)

func UserModels(r *gin.RouterGroup, ctrl Interface.IUserCtrl) {
	userGroup := r.Group("/user", middleware.Auth())
	{
		userGroup.GET("", ctrl.GetUserDetail)
		userGroup.PUT("", ctrl.UpdateUserInfoByUser)
	}

	fileGroup := r.Group("/files", middleware.Auth())
	{
		fileGroup.POST("/:types", ctrl.UploadFile)
		fileGroup.DELETE("/:types/:name", ctrl.DeleteFile)
	}
}
