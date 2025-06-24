package news

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service"
	svcDto "github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	ctrlDto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/admin/news"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type CategoryCtrl struct {
	dig.In
	Services service.INewsCategorySvc
}

func NewCategoryCtrl(srv service.INewsCategorySvc) Interface.INewsCategoryCtrl {
	return &CategoryCtrl{
		Services: srv,
	}
}

func (c *CategoryCtrl) CreateCategory(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.CategoryCreateDTO](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// DTO 转换
	dto := svcDto.CategoryCreateSvcDTO{
		Name:        apiCtx.Request.Name,
		Description: apiCtx.Request.Description,
	}

	code := c.Services.CreateCategory(dto)
	apiCtx.NoDataJSON(code)
}

func (c *CategoryCtrl) GetCategoryList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	code, vo := c.Services.ListCategories()
	apiCtx.WithDataJSON(code, vo)
}

func (c *CategoryCtrl) UpdateCategory(ctx *gin.Context) {

	apiCtx := controller.NewAPiContext[ctrlDto.CategoryUpdateDTO](ctx)
	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	id, _ := apiCtx.GetIdByPath()

	// DTO 转换
	dto := svcDto.CategoryUpdateSvcDTO{
		Name:        apiCtx.Request.Name,
		Description: apiCtx.Request.Description,
		SortOrder:   apiCtx.Request.SortOrder,
	}

	code := c.Services.UpdateCategory(id, dto)
	apiCtx.NoDataJSON(code)
}

func (c *CategoryCtrl) DeleteCategory(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id, _ := apiCtx.GetIdByPath()

	code := c.Services.DeleteCategory(id)

	apiCtx.NoDataJSON(code)
}
