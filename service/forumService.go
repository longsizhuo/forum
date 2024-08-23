package service

import (
	"context"
	"github.com/longsizhuo/forum/models"
	pb "github.com/longsizhuo/forum/proto"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedForumServiceServer
	db *gorm.DB
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := models.User{
		UserName:     req.UserName,
		UserPassword: req.UserPassword,
		UserSex:      req.UserSex,
		UserAge:      int(req.UserAge),
	}
	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{UserId: int32(user.UID)}, nil
}
