package Interface

import "github.com/gin-gonic/gin"

type IUserCtrl interface {
	UpdateUserInfoByAdmin(c *gin.Context)
	UpdateUserInfoByUser(c *gin.Context)
	GetUserList(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserDetail(c *gin.Context)
	UpdateUserAvatar(c *gin.Context)
}

type IAuthCtrl interface {
	LoginByPassword(c *gin.Context)
	LoginByCode(c *gin.Context)
	Register(c *gin.Context)
	SendPhoneCode(c *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
}

type INewsCtrl interface {
	// 提交新闻
	CreateNews(ctx *gin.Context)
	// 获取新闻列表
	GetNewsList(ctx *gin.Context)
	// 获取新闻详情
	GetNewsDetail(ctx *gin.Context)
	// 修改新闻信息
	UpdateNews(ctx *gin.Context)
	// 修改新闻状态
	UpdateNewsStatus(ctx *gin.Context)
	// 删除新闻
	DeleteNews(ctx *gin.Context)
}

type INewsCategoryCtrl interface {
	// 创建新闻分类
	CreateCategory(ctx *gin.Context)
	// 获取新闻分类列表
	GetCategoryList(ctx *gin.Context)
	// 修改新闻分类
	UpdateCategory(ctx *gin.Context)
	// 删除新闻分类
	DeleteCategory(ctx *gin.Context)
}

// ISupplyDemandCtrl 供需控制器接口
type ISupplyDemandCtrl interface {
	// 获取供需列表
	GetSupplyDemandList(ctx *gin.Context)
	// 获取供需详情
	GetSupplyDemandDetail(ctx *gin.Context)
	// 创建供需
	CreateSupplyDemand(ctx *gin.Context)
	// 删除供需
	DeleteSupplyDemand(ctx *gin.Context)
	// 获取评论列表
	GetCommentList(ctx *gin.Context)
	// 获取评论详情
	GetCommentDetail(ctx *gin.Context)
	// 创建评论
	CreateComment(ctx *gin.Context)
	// 删除评论
	DeleteComment(ctx *gin.Context)
}
