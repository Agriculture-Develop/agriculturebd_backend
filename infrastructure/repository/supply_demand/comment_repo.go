package supply_demand

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SupplyDemandCommentRepo struct {
	dig.In
	DB *gorm.DB
}

func NewSupplyDemandCommentRepo(db *gorm.DB) repository.ISupplyDemandCommentRepo {
	return &SupplyDemandCommentRepo{DB: db}
}

// Create 创建评论
func (r *SupplyDemandCommentRepo) Create(comment *entity.SupplyDemandComment) error {
	// 转换为DAO模型
	daoModel := &model.SupplyDemandComment{
		SupplyDemandID: comment.SupplyDemandID,
		UserID:         comment.UserID,
		CommentContent: comment.CommentContent,
		LikeCount:      comment.LikeCount,
		ReplyId:        comment.ReplyId,
	}

	if err := r.DB.Create(daoModel).Error; err != nil {
		zap.L().Error("Create supply_demand_comment fail", zap.Error(err))
		return err
	}

	// 回填ID
	comment.ID = daoModel.ID
	return nil
}

// GetByID 根据ID获取评论详情
func (r *SupplyDemandCommentRepo) GetByID(id int64) (*entity.SupplyDemandComment, error) {
	var daoModel model.SupplyDemandComment
	if err := r.DB.Where("id = ?", id).First(&daoModel).Error; err != nil {
		return nil, err
	}

	// 转换为实体
	entity := &entity.SupplyDemandComment{
		ID:             daoModel.ID,
		SupplyDemandID: daoModel.SupplyDemandID,
		UserID:         daoModel.UserID,
		CommentContent: daoModel.CommentContent,
		LikeCount:      daoModel.LikeCount,
		ReplyId:        daoModel.ReplyId,
		CreatedAt:      daoModel.CreatedAt,
		UpdatedAt:      daoModel.UpdatedAt,
		DeletedAt:      daoModel.DeletedAt,
	}

	return entity, nil
}

// List 获取评论列表
func (r *SupplyDemandCommentRepo) List(supplyDemandID int64) ([]*entity.SupplyDemandComment, int64, error) {
	var daoModels []model.SupplyDemandComment
	var total int64

	// 构建查询条件
	query := r.DB.Model(&model.SupplyDemandComment{}).Where("supply_demand_id = ?", supplyDemandID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询评论列表
	if err := query.Order("created_at DESC").Find(&daoModels).Error; err != nil {
		return nil, 0, err
	}

	// 转换为实体列表
	entities := make([]*entity.SupplyDemandComment, 0, len(daoModels))
	for _, daoModel := range daoModels {
		e := &entity.SupplyDemandComment{
			ID:             daoModel.ID,
			SupplyDemandID: daoModel.SupplyDemandID,
			UserID:         daoModel.UserID,
			CommentContent: daoModel.CommentContent,
			LikeCount:      daoModel.LikeCount,
			ReplyId:        daoModel.ReplyId,
			CreatedAt:      daoModel.CreatedAt,
			UpdatedAt:      daoModel.UpdatedAt,
			DeletedAt:      daoModel.DeletedAt,
		}
		entities = append(entities, e)
	}

	return entities, total, nil
}

// Delete 删除评论
func (r *SupplyDemandCommentRepo) Delete(id int64) error {
	return r.DB.Where("id = ?", id).Delete(&model.SupplyDemandComment{}).Error
}

// DeleteByParentId 根据父级ID删除
func (r *SupplyDemandCommentRepo) DeleteByParentId(parentId int64) error {
	return r.DB.Where("reply_id = ?", parentId).Delete(&model.SupplyDemandComment{}).Error
}
