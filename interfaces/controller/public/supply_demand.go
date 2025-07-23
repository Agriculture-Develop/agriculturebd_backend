package public

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	upload "github.com/Agriculture-Develop/agriculturebd/domain/common/service"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service"
	svcDto "github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	ctrlDto "github.com/Agriculture-Develop/agriculturebd/interfaces/dto/public"
	ctrlVo "github.com/Agriculture-Develop/agriculturebd/interfaces/vo/public"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type SupplyDemandCtrl struct {
	dig.In
	Services   service.ISupplyDemandService
	CommentSvc service.ISupplyDemandCommentService
	UploadSvc  upload.IUploadSvc
}

func NewSupplyDemandCtrl(svc service.ISupplyDemandService, commentSvc service.ISupplyDemandCommentService, upload upload.IUploadSvc) Interface.ISupplyDemandCtrl {
	return &SupplyDemandCtrl{
		Services:   svc,
		CommentSvc: commentSvc,
		UploadSvc:  upload,
	}
}

// GetSupplyDemandList 获取供需列表
func (c *SupplyDemandCtrl) GetSupplyDemandList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.SupplyDemandListFilterCtrlDTO](ctx)
	if err := apiCtx.BindQuery(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// DTO 转换
	filter := svcDto.SupplyDemandListFilterSvcDTO{
		Title: apiCtx.Request.Title,
		Page:  apiCtx.Request.Page,
		Count: apiCtx.Request.Count,
	}

	code, vo := c.Services.ListSupplyDemand(filter)
	if code != respCode.Success {
		apiCtx.NoDataJSON(code)
		return
	}

	// VO 转换
	res := &ctrlVo.SupplyDemandListVO{
		Total: vo.Total,
		List:  make([]ctrlVo.SupplyDemandItemVO, 0, len(vo.List)),
	}

	for _, item := range vo.List {
		res.List = append(res.List, ctrlVo.SupplyDemandItemVO{
			Id:        item.Id,
			UserId:    item.UserId,
			CreatedAt: item.CreatedAt,
			Title:     item.Title,
			Content:   item.Content,
			TagName:   item.TagName,
			TagWeigh:  item.TagWeigh,
			TagPrice:  item.TagPrice,
			CoverURL:  item.CoverURL,
			Like:      item.Like,
		})
	}

	apiCtx.WithDataJSON(code, res)
}

// GetSupplyDemandDetail 获取供需详情
func (c *SupplyDemandCtrl) GetSupplyDemandDetail(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code, vo := c.Services.GetSupplyDemandDetail(id)
	if code != respCode.Success {
		apiCtx.NoDataJSON(code)
		return
	}

	// VO 转换
	res := &ctrlVo.SupplyDemandDetailVO{
		ID:            vo.ID,
		UserId:        vo.UserId,
		Title:         vo.Title,
		Content:       vo.Content,
		CoverURL:      vo.CoverURL,
		FilesURL:      vo.FilesURL,
		PublisherName: vo.PublisherName,
		CreatedAt:     vo.CreatedAt,
		Like:          vo.Like,
		Tags: ctrlVo.TagsVO{
			Name:   vo.TagName,
			Price:  vo.TagPrice,
			Weight: vo.TagWeigh,
		},
	}

	apiCtx.WithDataJSON(code, res)
}

// CreateSupplyDemand 创建供需
func (c *SupplyDemandCtrl) CreateSupplyDemand(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.SupplyDemandCreateCtrlDTO](ctx)

	if err := apiCtx.BindForm(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// 文件上传处理
	coverURL, err := c.UploadSvc.UploadFile(apiCtx.Request.Cover, "good")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
		return
	}

	filesURL, err := c.UploadSvc.UploadFiles(apiCtx.Request.Files, "good")
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParams, err.Error())
		return
	}

	// DTO 转换
	dto := svcDto.SupplyDemandCreateSvcDTO{
		Title:    apiCtx.Request.Title,
		Content:  apiCtx.Request.Content,
		CoverURL: coverURL,
		FilesURL: filesURL,
		TagName:  apiCtx.Request.TagName,
		TagPrice: apiCtx.Request.TagPrice,
		TagWeigh: apiCtx.Request.TagWeigh,
		UserID:   apiCtx.GetUserIdByToken(),
	}

	code := c.Services.CreateSupplyDemand(dto)
	apiCtx.NoDataJSON(code)
}

// DeleteSupplyDemand 删除供需
func (c *SupplyDemandCtrl) DeleteSupplyDemand(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	id, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	userId := apiCtx.GetUserIdByToken()

	code := c.Services.DeleteSupplyDemand(userId, id)
	apiCtx.NoDataJSON(code)
}

// GetCommentList 获取评论列表
func (c *SupplyDemandCtrl) GetCommentList(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	supplyDemandID, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code, vo := c.CommentSvc.ListComments(int64(supplyDemandID))
	if code != respCode.Success {
		apiCtx := controller.NewAPiContext[struct{}](ctx)
		apiCtx.NoDataJSON(code)
		return
	}

	// VO 转换
	res := &ctrlVo.CommentListVO{
		Total: vo.Total,
		List:  make([]ctrlVo.CommentDetailVO, 0, len(vo.List)),
	}

	for _, item := range vo.List {
		res.List = append(res.List, ctrlVo.CommentDetailVO{
			Avatar:        item.Avatar,
			ID:            item.ID,
			PublisherName: item.PublisherName,
			Comment:       item.Comment,
			Role:          item.Role,
			Like:          item.Like,
			CreatedAt:     item.CreatedAt,
		})
	}

	apiCtx.WithDataJSON(code, res)
}

// GetCommentDetail 获取评论详情
func (c *SupplyDemandCtrl) GetCommentDetail(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[struct{}](ctx)
	commentID, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code, vo := c.CommentSvc.GetCommentDetail(int64(commentID))
	if code != respCode.Success {
		apiCtx := controller.NewAPiContext[struct{}](ctx)
		apiCtx.NoDataJSON(code)
		return
	}

	// VO 转换
	res := &ctrlVo.CommentDetailVO{
		Avatar:        vo.Avatar,
		ID:            vo.ID,
		PublisherName: vo.PublisherName,
		Comment:       vo.Comment,
		Role:          vo.Role,
		Like:          vo.Like,
		CreatedAt:     vo.CreatedAt,
	}

	apiCtx.WithDataJSON(code, res)
}

// CreateComment 创建评论
func (c *SupplyDemandCtrl) CreateComment(ctx *gin.Context) {
	apiCtx := controller.NewAPiContext[ctrlDto.CommentCreateCtrlDTO](ctx)
	supplyDemandID, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	if err := apiCtx.BindJSON(); err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	// DTO 转换
	dto := svcDto.CommentCreateSvcDTO{
		SupplyDemandID: int64(supplyDemandID),
		CommentContent: apiCtx.Request.Comment,
		UserID:         apiCtx.GetUserIdByToken(),
	}

	code := c.CommentSvc.CreateComment(dto)
	apiCtx.NoDataJSON(code)
}

// DeleteComment 删除评论
func (c *SupplyDemandCtrl) DeleteComment(ctx *gin.Context) {

	apiCtx := controller.NewAPiContext[ctrlDto.CommentCreateCtrlDTO](ctx)

	userId := apiCtx.GetUserIdByToken()
	commentID, err := apiCtx.GetIdByPath()
	if err != nil {
		apiCtx.NoDataJSON(respCode.InvalidParamsFormat)
		return
	}

	code := c.CommentSvc.DeleteComment(userId, commentID)
	apiCtx.NoDataJSON(code)
}
