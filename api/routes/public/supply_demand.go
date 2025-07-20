package public

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"github.com/gin-gonic/gin"
)

func SupplyDemandModels(r *gin.RouterGroup, ctrl Interface.ISupplyDemandCtrl) {
	goodGroup := r.Group("/good", middleware.Auth())
	{
		// 供需相关
		goodGroup.GET("list", ctrl.GetSupplyDemandList)   // 获取供需列表
		goodGroup.GET("/:id", ctrl.GetSupplyDemandDetail) // 获取供需详情
		goodGroup.POST("", ctrl.CreateSupplyDemand)       // 创建供需
		goodGroup.DELETE("/:id", ctrl.DeleteSupplyDemand)

		// 评论相关
		goodGroup.GET("/:id/comment", ctrl.GetCommentList)   // 获取评论列表
		goodGroup.GET("/comment/:id", ctrl.GetCommentDetail) // 获取评论详情
		goodGroup.POST("/:id/comment", ctrl.CreateComment)   // 创建评论
		goodGroup.DELETE("/comment/:id", ctrl.DeleteComment) // 删除评论
	}
}
