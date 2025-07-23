package news

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	upload "github.com/Agriculture-Develop/agriculturebd/domain/common/service"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service"
	svcDto "github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	dto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin"
	ctrlDto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin/news"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
)

type Ctrl struct {
	dig.In
	Services  service.INewsSvc
	UploadSvc upload.IUploadSvc
}

func NewCtrl(srv service.INewsSvc, upload upload.IUploadSvc) Interface.INewsCtrl {
	return &Ctrl{
		Services:  srv,
		UploadSvc: upload,
	}
}

// 提交新闻
func (c *Ctrl) CreateNews(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsCreateDTO](ctx)
	if err := apiCtx.BindForm(); err != nil {

		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// 文件校验与上传
	coverUrl, err := c.UploadSvc.UploadFile(apiCtx.Request.Cover, "news")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
		return
	}

	filesUrl, err := c.UploadSvc.UploadFiles(apiCtx.Request.Files, "news")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
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
		Type:       apiCtx.Request.Type,
		CoverURL:   coverUrl,
		FilesURL:   filesUrl,
		Status:     apiCtx.Request.Status,
		UserID:     apiCtx.GetUserIdByToken(),
	}

	code := c.Services.CreateNews(dto)
	apiCtx.NoDataJSON(code)
}

// 修改新闻信息
func (c *Ctrl) UpdateNews(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsUpdateDTO](ctx)
	if err := apiCtx.BindForm(); err != nil {
		log.Println(err)
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	id, _ := apiCtx.GetIdByPath()

	// 文件校验与上传
	coverUrl, err := c.UploadSvc.UploadFile(apiCtx.Request.Cover, "news")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
		return
	}

	filesUrl, err := c.UploadSvc.UploadFiles(apiCtx.Request.Files, "news")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
		return
	}

	// DTO 转换
	dto := svcDto.NewsUpdateSvcDTO{
		Title:      apiCtx.Request.Title,
		CategoryID: apiCtx.Request.CategoryID,
		Abstract:   apiCtx.Request.Abstract,
		Keyword:    apiCtx.Request.Keyword,
		Source:     apiCtx.Request.Source,
		Content:    apiCtx.Request.Content,
		Type:       apiCtx.Request.Type,
		CoverURL:   coverUrl,
		FilesURL:   filesUrl,
		Status:     apiCtx.Request.Status,
		UserID:     apiCtx.GetUserIdByToken(),
	}

	code := c.Services.UpdateNews(id, dto)
	apiCtx.NoDataJSON(code)
}

// 获取新闻列表
func (c *Ctrl) GetNewsList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.NewsListFilterDTO](ctx)
	if err := apiCtx.BindQuery(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		fmt.Println(err)
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

// 删除新闻
func (c *Ctrl) DeleteNews(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id, _ := apiCtx.GetIdByPath()
	code := c.Services.DeleteNews(id)
	apiCtx.NoDataJSON(code)
}
