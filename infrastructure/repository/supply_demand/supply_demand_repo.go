package supply_demand

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SupplyDemandRepo struct {
	dig.In
	DB *gorm.DB
}

func NewSupplyDemandRepo(db *gorm.DB) repository.ISupplyDemandRepo {
	return &SupplyDemandRepo{DB: db}
}

// Create 创建供需
func (r *SupplyDemandRepo) Create(supplyDemand *entity.SupplyDemand) error {
	// 转换为DAO模型

	tag := model.TagInfo{
		Name:  supplyDemand.TagName,
		Price: supplyDemand.TagPrice,
		Weigh: supplyDemand.TagWeigh,
	}

	daoModel := &model.SupplyDemand{
		Title:    supplyDemand.Title,
		Content:  supplyDemand.Content,
		Tag:      tag,
		CoverURL: supplyDemand.CoverURL,
		FilesURL: supplyDemand.FilesURL,
		Likes:    supplyDemand.Likes,
		UserID:   supplyDemand.UserId,
	}

	if err := r.DB.Create(daoModel).Error; err != nil {
		zap.L().Error("Create supply_demand fail", zap.Error(err))
		return err
	}

	// 回填ID
	supplyDemand.ID = daoModel.ID
	return nil
}

// GetByID 根据ID获取供需详情
func (r *SupplyDemandRepo) GetByID(id uint) (*entity.SupplyDemand, error) {
	var daoModel model.SupplyDemand
	if err := r.DB.Where("id = ?", id).First(&daoModel).Error; err != nil {
		return nil, err
	}

	// 转换为实体
	e := &entity.SupplyDemand{
		ID:        daoModel.ID,
		Title:     daoModel.Title,
		Content:   daoModel.Content,
		TagName:   daoModel.Tag.Name,
		TagPrice:  daoModel.Tag.Price,
		TagWeigh:  daoModel.Tag.Weigh,
		CoverURL:  daoModel.CoverURL,
		FilesURL:  daoModel.FilesURL,
		Likes:     daoModel.Likes,
		UserId:    daoModel.UserID,
		CreatedAt: daoModel.CreatedAt,
		UpdatedAt: daoModel.UpdatedAt,
		DeletedAt: daoModel.DeletedAt,
	}

	return e, nil
}

// List 获取供需列表
func (r *SupplyDemandRepo) List(filter repository.SupplyDemandListFilter) ([]*entity.SupplyDemand, int64, error) {
	var daoModels []model.SupplyDemand
	var total int64

	fmt.Println("filter.Title: ", filter.Title)

	// 构建查询条件
	query := r.DB.Model(&model.SupplyDemand{})
	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (filter.Page - 1) * filter.Count
	if err := query.Offset(offset).Limit(filter.Count).Order("created_at DESC").Find(&daoModels).Error; err != nil {
		return nil, 0, err
	}

	// 转换为实体列表
	entities := make([]*entity.SupplyDemand, 0, len(daoModels))
	for _, daoModel := range daoModels {
		e := &entity.SupplyDemand{
			ID:       daoModel.ID,
			Title:    daoModel.Title,
			Content:  daoModel.Content,
			TagName:  daoModel.Tag.Name,
			TagPrice: daoModel.Tag.Price,
			TagWeigh: daoModel.Tag.Weigh,

			CoverURL:  daoModel.CoverURL,
			FilesURL:  daoModel.FilesURL,
			Likes:     daoModel.Likes,
			UserId:    daoModel.UserID,
			CreatedAt: daoModel.CreatedAt,
			UpdatedAt: daoModel.UpdatedAt,
			DeletedAt: daoModel.DeletedAt,
		}
		entities = append(entities, e)
	}

	return entities, total, nil
}

// Delete 删除供需
func (r *SupplyDemandRepo) Delete(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&model.SupplyDemand{}).Error
}
