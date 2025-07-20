package service

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/vo"
)

type ISupplyDemandService interface {
	// 创建供需
	CreateSupplyDemand(dto dto.SupplyDemandCreateSvcDTO) respCode.StatusCode
	// 获取供需详情
	GetSupplyDemandDetail(id uint) (respCode.StatusCode, *vo.SupplyDemandDetailSvcVO)
	// 获取供需列表
	ListSupplyDemand(filter dto.SupplyDemandListFilterSvcDTO) (respCode.StatusCode, *vo.SupplyDemandListSvcVO)
	// 删除供需
	DeleteSupplyDemand(userid, id uint) respCode.StatusCode
}

type ISupplyDemandCommentService interface {
	// 创建评论
	CreateComment(dto dto.CommentCreateSvcDTO) respCode.StatusCode
	// 获取评论详情
	GetCommentDetail(id int64) (respCode.StatusCode, *vo.CommentDetailSvcVO)
	// 获取评论列表
	ListComments(supplyDemandID int64) (respCode.StatusCode, *vo.CommentListSvcVO)
	// 删除评论
	DeleteComment(userid, id uint) respCode.StatusCode
}
