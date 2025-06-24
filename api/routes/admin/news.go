package admin

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"github.com/gin-gonic/gin"
)

func NewsModels(r *gin.RouterGroup, ctrl Interface.INewsCtrl, catCtrl Interface.INewsCategoryCtrl) {
	newsGroup := r.Group("/news", middleware.Auth())
	{
		newsGroup.POST("", ctrl.CreateNews)
		newsGroup.GET("list", ctrl.GetNewsList)
		newsGroup.GET("/:id", ctrl.GetNewsDetail)
		newsGroup.PUT("/:id", ctrl.UpdateNews)
		newsGroup.PUT("status/:id", ctrl.UpdateNewsStatus)
		newsGroup.DELETE("/:id", ctrl.DeleteNews)

		// 分类相关
		newsGroup.POST("categories", catCtrl.CreateCategory)
		newsGroup.GET("categories/list", catCtrl.GetCategoryList)
		newsGroup.PUT("categories/:id", catCtrl.UpdateCategory)
		newsGroup.DELETE("categories/:id", catCtrl.DeleteCategory)
	}
}
