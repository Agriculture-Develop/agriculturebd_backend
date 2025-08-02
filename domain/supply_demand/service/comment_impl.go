package service

import (
	"errors"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/valobj"
	userRepo "github.com/Agriculture-Develop/agriculturebd/domain/user/repository"
	"strconv"

	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/supply_demand/service/vo"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SupplyDemandCommentSvc struct {
	dig.In
	Repo     repository.ISupplyDemandCommentRepo
	userRepo userRepo.IUserRepo
}

func NewSupplyDemandCommentService(repo repository.ISupplyDemandCommentRepo, userRepo userRepo.IUserRepo) ISupplyDemandCommentService {
	return &SupplyDemandCommentSvc{Repo: repo, userRepo: userRepo}
}

// CreateComment 创建评论
func (s *SupplyDemandCommentSvc) CreateComment(dto dto.CommentCreateSvcDTO) respCode.StatusCode {
	// 1. 参数校验
	if dto.CommentContent == "" {
		return respCode.InvalidParamsFormat
	}

	// 2. 创建评论实体
	comment := &entity.SupplyDemandComment{
		SupplyDemandID: dto.SupplyDemandID,
		UserID:         int64(dto.UserID),
		CommentContent: dto.CommentContent,
		LikeCount:      0,
		ReplyId:        -1, // 默认回复ID为0
	}

	// 3. 保存到数据库
	if err := s.Repo.Create(comment); err != nil {
		zap.L().Error("CreateComment fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

// GetCommentDetail 获取评论详情
func (s *SupplyDemandCommentSvc) GetCommentDetail(id int64) (respCode.StatusCode, *vo.CommentDetailSvcVO) {
	// 1. 获取评论信息
	comment, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.CommentNotExist, nil
		}
		zap.L().Error("GetCommentById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 获取用户信息
	user, err := s.userRepo.GetUserById(uint(comment.UserID))
	if err != nil {
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 2. 构建返回VO
	commentVO := &vo.CommentDetailSvcVO{
		Avatar:        user.AvatarPath,
		ID:            strconv.FormatInt(comment.ID, 10),
		PublisherName: user.Nickname,
		Comment:       comment.CommentContent,
		Role:          valobj.UserRole(user.Role).Desc(),
		Like:          strconv.Itoa(comment.LikeCount),
		CreatedAt:     comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return respCode.Success, commentVO
}

// ListComments 获取评论列表
func (s *SupplyDemandCommentSvc) ListComments(supplyDemandID int64) (respCode.StatusCode, *vo.CommentListSvcVO) {
	// 1. 获取评论列表
	commentList, total, err := s.Repo.List(supplyDemandID)
	if err != nil {
		zap.L().Error("ListComments fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 2. 转换为VO
	commentVOs := make([]vo.CommentDetailSvcVO, 0, len(commentList))
	for _, item := range commentList {

		// 获取用户信息
		user, err := s.userRepo.GetUserById(uint(item.UserID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			zap.L().Error("GetUserById fail", zap.Error(err))
			return respCode.ServerBusy, nil
		}

		commentVOs = append(commentVOs, vo.CommentDetailSvcVO{
			Avatar:        user.AvatarPath,
			ID:            strconv.FormatInt(item.ID, 10),
			PublisherName: user.Nickname,
			Comment:       item.CommentContent,
			Role:          valobj.UserRole(user.Role).Desc(),
			Like:          strconv.Itoa(item.LikeCount),
			UserId:        uint(item.UserID),
			CreatedAt:     item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 3. 构建返回结果
	result := &vo.CommentListSvcVO{
		Total: int(total),
		List:  commentVOs,
	}

	return respCode.Success, result
}

// DeleteComment 删除评论
func (s *SupplyDemandCommentSvc) DeleteComment(userid, id uint) respCode.StatusCode {
	// 1. 检查评论是否存在
	comment, err := s.Repo.GetByID(int64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.CommentNotExist
		}
		zap.L().Error("GetCommentById for delete fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 检查用户是否是该评论的作者
	user, err := s.userRepo.GetUserById(uint(comment.UserID))
	if err != nil {
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}
	if user.Role < 1 && uint(comment.UserID) != userid {
		return respCode.Forbidden
	}

	// 2. 删除评论
	if err := s.Repo.Delete(int64(id)); err != nil {
		zap.L().Error("DeleteComment fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}
