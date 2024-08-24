// service/forumService.go
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

// CreateUser creates a new user
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

func (s *server) CreateTopic(ctx context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	topic := models.Topic{
		UID:     int(req.UserId),
		SID:     int(req.SectionId),
		Title:   req.Title,
		Content: req.Content,
	}
	if err := s.db.Create(&topic).Error; err != nil {
		return nil, err
	}
	return &pb.CreateTopicResponse{TopicId: int32(topic.TID)}, nil
}

func (s *server) CreateReply(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyResponse, error) {
	reply := models.Reply{
		UID:     int(req.UserId),
		TID:     int(req.TopicId),
		Content: req.Content,
	}
	if err := s.db.Create(&reply).Error; err != nil {
		return nil, err
	}
	return &pb.CreateReplyResponse{ReplyId: int32(reply.RID)}, nil
}

func (s *server) GetTopics(ctx context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	var topics []models.Topic
	if err := s.db.Where("s_id = ?", req.SectionId).Find(&topics).Error; err != nil {
		return nil, err
	}

	var responseTopics []*pb.Topic
	for _, topic := range topics {
		responseTopics = append(responseTopics, &pb.Topic{
			TopicId:    int32(topic.TID),
			UserId:     int32(topic.UID),
			Title:      topic.Title,
			Content:    topic.Content,
			ReplyCount: int32(topic.ReplyCount),
		})
	}

	return &pb.GetTopicsResponse{Topics: responseTopics}, nil
}

func (s *server) GetReplies(ctx context.Context, req *pb.GetRepliesRequest) (*pb.GetRepliesResponse, error) {
	var replies []models.Reply
	if err := s.db.Where("t_id = ?", req.TopicId).Find(&replies).Error; err != nil {
		return nil, err
	}

	var responseReplies []*pb.Reply
	for _, reply := range replies {
		responseReplies = append(responseReplies, &pb.Reply{
			ReplyId: int32(reply.RID),
			UserId:  int32(reply.UID),
			Content: reply.Content,
		})
	}

	return &pb.GetRepliesResponse{Replies: responseReplies}, nil
}
