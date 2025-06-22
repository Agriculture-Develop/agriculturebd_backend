package news

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service"
	svcDto "github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	dto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin"
	ctrlDto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin/news"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Ctrl struct {
	dig.In
	Services service.INewsSvc
}

func NewCtrl(srv service.INewsSvc) Interface.INewsCtrl {
	return &Ctrl{
		Services: srv,
	}
}

// 提交新闻
func (c *Ctrl) CreateNews(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsCreateDTO](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// DTO 转换
	dto := svcDto.NewsCreateSvcDTO{
		Title:      apiCtx.Request.Title,
		CategoryID: apiCtx.Request.CategoryID,
		Abstract:   apiCtx.Request.Abstract,
		Keyword:    apiCtx.Request.Keyword,
		Source:     apiCtx.Request.Source,
		Content:    apiCtx.Request.Content,
		CoverURL:   apiCtx.Request.CoverURL,
		FilesURL:   apiCtx.Request.FilesURL,
		Status:     apiCtx.Request.Status,
		UserID:     apiCtx.GetUserIdByToken(),
	}

	code := c.Services.CreateNews(dto)
	apiCtx.NoDataJSON(code)
}

// 获取新闻列表
func (c *Ctrl) GetNewsList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsListFilterDTO](ctx)
	if err := apiCtx.BindQuery(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// DTO 转换
	filter := svcDto.NewsListFilterSvcDTO{
		Title:    apiCtx.Request.Title,
		Author:   apiCtx.Request.Author,
		Status:   apiCtx.Request.Status,
		Page:     apiCtx.Request.Page,
		PageSize: apiCtx.Request.Count,
	}

	code, vo := c.Services.ListNews(filter)
	apiCtx.WithDataJSON(code, vo)
}

// 获取新闻详情
func (c *Ctrl) GetNewsDetail(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id, _ := apiCtx.GetIdByPath()

	code, vo := c.Services.GetNewsDetail(id)

	apiCtx.WithDataJSON(code, vo)
}

// 修改新闻信息
func (c *Ctrl) UpdateNews(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsUpdateDTO](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	id, _ := apiCtx.GetIdByPath()

	// DTO 转换
	dto := svcDto.NewsUpdateSvcDTO{
		Title:      apiCtx.Request.Title,
		CategoryID: apiCtx.Request.CategoryID,
		Abstract:   apiCtx.Request.Abstract,
		Keyword:    apiCtx.Request.Keyword,
		Source:     apiCtx.Request.Source,
		Content:    apiCtx.Request.Content,
		CoverURL:   apiCtx.Request.CoverURL,
		FilesURL:   apiCtx.Request.FilesURL,
		Status:     apiCtx.Request.Status,
	}

	code := c.Services.UpdateNews(id, dto)
	apiCtx.NoDataJSON(code)
}

// 修改新闻状态
func (c *Ctrl) UpdateNewsStatus(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[dto.NewsStatusUpdateDTO](ctx)

	id, _ := apiCtx.GetIdByPath()

	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}
	code := c.Services.UpdateNewsStatus(id, apiCtx.Request.Status)
	apiCtx.NoDataJSON(code)
}
