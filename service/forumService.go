package service

import pb "github.com/longsizhuo/forum/proto/github.com/longsizhuo/forum/proto/forum_grpc.pb.go"

type server struct {
	pb.UnimplementedForumServiceServer
}
