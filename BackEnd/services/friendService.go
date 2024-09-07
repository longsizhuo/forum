package services

import (
	"context"
	"fmt"
	"github.com/longsizhuo/forum/models"
	pb "github.com/longsizhuo/forum/proto/friend"
	"gorm.io/gorm"
)

type FriendService struct {
	pb.UnimplementedFriendServiceServer
	Db *gorm.DB
}

func (s *FriendService) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.AddFriendResponse, error) {
	var fromUser, toUser models.User

	// 检查 fromId 用户是否存在
	if err := s.Db.Where("id = ?", req.FromId).First(&fromUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d does not exist", req.FromId)
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	// 检查 toId 用户是否存在
	if err := s.Db.Where("id = ?", req.ToId).First(&toUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d does not exist", req.ToId)
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	// 调用 CheckFriendship 方法，检查是否已经是好友
	checkFriendReq := &pb.CheckFriendshipRequest{
		FromId: req.FromId,
		ToId:   req.ToId,
	}

	checkFriendRes, err := s.CheckFriendship(ctx, checkFriendReq)
	if err != nil {
		return nil, fmt.Errorf("failed to check friendship: %v", err)
	}

	// 如果已经是好友，返回提示信息
	if checkFriendRes.IsFriend {
		return nil, fmt.Errorf("user %d and user %d are already friends", req.FromId, req.ToId)
	}

	// 添加好友，使用 GORM 的 Many-to-Many 关联关系
	if err := s.Db.Model(&fromUser).Association("Friends").Append(&toUser); err != nil {
		return nil, fmt.Errorf("failed to add friend: %v", err)
	}

	// 返回成功响应
	return &pb.AddFriendResponse{
		FromId:   req.FromId,
		ToId:     req.ToId,
		IsFriend: true,
	}, nil
}

func (s *FriendService) CheckFriendship(_ context.Context, req *pb.CheckFriendshipRequest) (*pb.CheckFriendshipResponse, error) {
	var user models.User

	// 从数据库中获取 fromId 对应的用户记录，并预加载好友列表
	if err := s.Db.Where("id = ?", req.FromId).Preload("Friends").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d does not exist", req.FromId)
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	// 检查好友列表中是否包含 toId
	isFriend := false
	for _, friend := range user.Friends {
		if int32(friend.ID) == req.ToId {
			isFriend = true
			break
		}
	}

	// 返回检查结果
	return &pb.CheckFriendshipResponse{
		FromId:   req.FromId,
		ToId:     req.ToId,
		IsFriend: isFriend,
	}, nil
}

func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	var fromUser, toUser models.User

	// 检查 fromId 用户是否存在
	if err := s.Db.Where("id = ?", req.FromId).First(&fromUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d does not exist", req.FromId)
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	// 检查 toId 用户是否存在
	if err := s.Db.Where("id = ?", req.ToId).First(&toUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d does not exist", req.ToId)
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	// 调用 CheckFriendship 方法，检查是否已经是好友
	checkFriendReq := &pb.CheckFriendshipRequest{
		FromId: req.FromId,
		ToId:   req.ToId,
	}

	checkFriendRes, err := s.CheckFriendship(ctx, checkFriendReq)
	if err != nil {
		return nil, fmt.Errorf("failed to check friendship: %v", err)
	}

	// 如果还不是好友，返回提示信息
	if checkFriendRes.IsFriend == false {
		return nil, fmt.Errorf("user %d and user %d are not friends", req.FromId, req.ToId)
	}

	// 删除好友，使用 GORM 的 Many-to-Many 关联关系
	if err := s.Db.Model(&fromUser).Association("Friends").Delete(&toUser); err != nil {
		return nil, fmt.Errorf("failed to delete friend: %v", err)
	}

	// 返回成功响应
	return &pb.DeleteFriendResponse{
		FromId:   req.FromId,
		ToId:     req.ToId,
		IsFriend: false,
	}, nil
}
