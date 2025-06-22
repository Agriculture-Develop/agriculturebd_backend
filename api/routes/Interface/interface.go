package Interface

import "github.com/gin-gonic/gin"

type IUserCtrl interface {
	UpdateUserInfo(c *gin.Context)
	GetUserList(c *gin.Context)
	DeleteUser(c *gin.Context)
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
