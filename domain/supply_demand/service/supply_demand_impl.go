package service

import (
	"errors"
	userRepo "github.com/Agriculture-Develop/agriculturebd/domain/user/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/upload"
	"strconv"

	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/entity"
	demandRepo "github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/vo"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SupplyDemandSvc struct {
	dig.In
	DemandRepo demandRepo.ISupplyDemandRepo
	UserRepo   userRepo.IUserRepo
}

func NewSupplyDemandService(demandRepo demandRepo.ISupplyDemandRepo, userRepo userRepo.IUserRepo) ISupplyDemandService {
	return &SupplyDemandSvc{DemandRepo: demandRepo, UserRepo: userRepo}
}

// CreateSupplyDemand 创建供需
func (s *SupplyDemandSvc) CreateSupplyDemand(dto dto.SupplyDemandCreateSvcDTO) respCode.StatusCode {
	// 1. 参数校验
	if dto.Title == "" || dto.Content == "" {
		return respCode.InvalidParamsFormat
	}

	// 3. 创建供需实体
	supplyDemand := &entity.SupplyDemand{
		Title:    dto.Title,
		Content:  dto.Content,
		CoverURL: dto.CoverURL,
		FilesURL: dto.FilesURL,
		TagName:  dto.TagName,
		TagPrice: dto.TagPrice,
		TagWeigh: dto.TagWeigh,
		UserId:   dto.UserID,
	}

	// 4. 保存到数据库
	if err := s.DemandRepo.Create(supplyDemand); err != nil {
		zap.L().Error("CreateSupplyDemand fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

// GetSupplyDemandDetail 获取供需详情
func (s *SupplyDemandSvc) GetSupplyDemandDetail(id uint) (respCode.StatusCode, *vo.SupplyDemandDetailSvcVO) {
	// 1. 获取供需信息
	supplyDemand, err := s.DemandRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.SupplyDemandNotExist, nil
		}
		zap.L().Error("GetSupplyDemandById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	u, err := s.UserRepo.GetUserById(supplyDemand.UserId)
	if err != nil {
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 3. 构建返回VO
	supplyDemandVO := &vo.SupplyDemandDetailSvcVO{
		ID:            supplyDemand.ID,
		UserId:        supplyDemand.UserId,
		Title:         supplyDemand.Title,
		Content:       supplyDemand.Content,
		CoverURL:      supplyDemand.CoverURL,
		FilesURL:      supplyDemand.FilesURL,
		PublisherName: u.Nickname,
		CreatedAt:     supplyDemand.CreatedAt.Format("2006-01-02 15:04:05"),
		Like:          strconv.Itoa(supplyDemand.Likes),
		TagName:       supplyDemand.TagName,
		TagPrice:      supplyDemand.TagPrice,
		TagWeigh:      supplyDemand.TagWeigh,
	}

	return respCode.Success, supplyDemandVO
}

// ListSupplyDemand 获取供需列表
func (s *SupplyDemandSvc) ListSupplyDemand(filter dto.SupplyDemandListFilterSvcDTO) (respCode.StatusCode, *vo.SupplyDemandListSvcVO) {
	// 1. 设置默认分页参数
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Count <= 0 {
		filter.Count = 10
	}

	// 2. 构建筛选条件
	repoFilter := demandRepo.SupplyDemandListFilter{
		Title: filter.Title,
		Page:  filter.Page,
		Count: filter.Count,
	}

	// 3. 获取供需列表
	supplyDemandList, total, err := s.DemandRepo.List(repoFilter)
	if err != nil {
		zap.L().Error("ListSupplyDemand fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 4. 转换为VO
	supplyDemandVOs := make([]vo.SupplyDemandListItemSvcVO, 0, len(supplyDemandList))
	for _, item := range supplyDemandList {

		supplyDemandVOs = append(supplyDemandVOs, vo.SupplyDemandListItemSvcVO{
			Id:        item.ID,
			UserId:    item.UserId,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			Title:     item.Title,
			Content:   item.Content,
			TagName:   item.TagName,
			TagWeigh:  item.TagWeigh,
			TagPrice:  item.TagPrice,
			CoverURL:  item.CoverURL,
			Like:      strconv.Itoa(item.Likes),
		})
	}

	// 5. 构建返回结果
	result := &vo.SupplyDemandListSvcVO{
		Total: int(total),
		List:  supplyDemandVOs,
	}

	return respCode.Success, result
}

// DeleteSupplyDemand 删除供需
func (s *SupplyDemandSvc) DeleteSupplyDemand(userid, id uint) respCode.StatusCode {
	// 1. 检查供需是否存在
	demand, err := s.DemandRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.SupplyDemandNotExist
		}
		zap.L().Error("GetSupplyDemandById for delete fail", zap.Error(err))
		return respCode.ServerBusy
	}

	//2. 校验角色
	u, err := s.UserRepo.GetUserById(userid)
	if err != nil {
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	if u.Role < 1 && demand.UserId != userid {
		return respCode.Forbidden
	}

	//3. 删除文件
	if demand.CoverURL != "" {
		err := upload.DeleteFile(demand.CoverURL, "good")
		if err != nil {
			zap.L().Warn("Delete cover image fail", zap.Error(err))
			return respCode.ServerBusy
		}
	}

	for _, fileURL := range demand.FilesURL {
		err := upload.DeleteFile(fileURL, "good")
		if err != nil {
			zap.L().Warn("Delete file fail", zap.Error(err))
			return respCode.ServerBusy
		}
	}

	//4. 删除供需
	if err := s.DemandRepo.Delete(id); err != nil {
		zap.L().Error("DeleteSupplyDemand fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}
