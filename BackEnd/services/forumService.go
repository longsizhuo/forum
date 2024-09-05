// Package services service/forumService.go
package services

import (
	"context"
	"fmt"
	"github.com/longsizhuo/forum/models"
	pb "github.com/longsizhuo/forum/proto/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	pb.UnimplementedForumServiceServer
	Db *gorm.DB
}

// CreateUser creates a new user
func (s *Server) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	defaultRoleID := 1

	dateOfBirth, err := time.Parse("2006/01/02", req.UserDateBirth)
	if err != nil {
		return nil, err
	}

	if req.UserPassword != req.UserRePassword {
		return nil, fmt.Errorf("passwords do not match")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Check if email is already registered (Assuming req contains Email field)
	var existingUser models.User
	if err := s.Db.Where("email = ?", req.UserEmail).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email is already registered")
	}

	user := models.User{
		UserName:     req.UserName,
		UserPassword: string(hashedPassword),
		UserSex:      req.UserSex,
		UserAge:      int(req.UserAge),
		RoleID:       defaultRoleID,
		UserEmail:    req.UserEmail,
		DateBirth:    dateOfBirth,
	}
	if err := s.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{UserId: int32(user.ID)}, nil
}

func (s *Server) CreateTopic(_ context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	topic := models.Topic{
		UID:     int(req.UserId),
		SID:     int(req.SectionId),
		Title:   req.Title,
		Content: req.Content,
	}
	if err := s.Db.Create(&topic).Error; err != nil {
		return nil, err
	}
	return &pb.CreateTopicResponse{TopicId: int32(topic.ID)}, nil
}

func (s *Server) CreateReply(_ context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyResponse, error) {
	reply := models.Reply{
		UID:     int(req.UserId),
		TID:     int(req.TopicId),
		Content: req.Content,
	}
	if err := s.Db.Create(&reply).Error; err != nil {
		return nil, err
	}
	return &pb.CreateReplyResponse{ReplyId: int32(reply.ID)}, nil
}

func (s *Server) GetTopics(_ context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	var topics []models.Topic
	if err := s.Db.Where("s_id = ?", req.SectionId).Find(&topics).Error; err != nil {
		return nil, err
	}

	var responseTopics []*pb.Topic
	for _, topic := range topics {
		responseTopics = append(responseTopics, &pb.Topic{
			TopicId:    int32(topic.ID),
			UserId:     int32(topic.UID),
			Title:      topic.Title,
			Content:    topic.Content,
			ReplyCount: int32(topic.ReplyCount),
		})
	}

	return &pb.GetTopicsResponse{Topics: responseTopics}, nil
}

func (s *Server) GetReplies(_ context.Context, req *pb.GetRepliesRequest) (*pb.GetRepliesResponse, error) {
	var replies []models.Reply
	if err := s.Db.Where("t_id = ?", req.TopicId).Find(&replies).Error; err != nil {
		return nil, err
	}

	var responseReplies []*pb.Reply
	for _, reply := range replies {
		responseReplies = append(responseReplies, &pb.Reply{
			ReplyId: int32(reply.ID),
			UserId:  int32(reply.UID),
			Content: reply.Content,
		})
	}

	return &pb.GetRepliesResponse{Replies: responseReplies}, nil
}
