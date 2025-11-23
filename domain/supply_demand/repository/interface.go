package repository

import "github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/entity"

type ISupplyDemandRepo interface {
	// 创建供需
	Create(supplyDemand *entity.SupplyDemand) error
	// 根据ID获取供需详情
	GetByID(id uint) (*entity.SupplyDemand, error)
	// 获取供需列表
	List(filter SupplyDemandListFilter) ([]*entity.SupplyDemand, int64, error)
	// 删除供需
	Delete(id uint) error
}

type SupplyDemandListFilter struct {
	Title     string
	Category  string
	UserIDs   []uint
	SortField string
	SortOrder string
	Page      int
	Count     int
}

const (
	SortFieldCreatedAt = "created_at"
	SortFieldPrice     = "price"

	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

type ISupplyDemandCommentRepo interface {
	// 创建评论
	Create(comment *entity.SupplyDemandComment) error
	// 根据ID获取评论详情
	GetByID(id int64) (*entity.SupplyDemandComment, error)
	// 获取评论列表
	List(supplyDemandID int64) ([]*entity.SupplyDemandComment, int64, error)
	// 删除评论
	Delete(id int64) error
	// 删除子评论
	DeleteByParentId(parentId int64) error
}
