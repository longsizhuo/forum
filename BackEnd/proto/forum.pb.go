// proto/forum.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.20.3
// source: proto/forum.proto

package forum

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName       string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPassword   string `protobuf:"bytes,2,opt,name=userPassword,proto3" json:"userPassword,omitempty"`
	UserRePassword string `protobuf:"bytes,3,opt,name=userRePassword,proto3" json:"userRePassword,omitempty"`
	UserSex        string `protobuf:"bytes,4,opt,name=userSex,proto3" json:"userSex,omitempty"`
	UserAge        int32  `protobuf:"varint,5,opt,name=userAge,proto3" json:"userAge,omitempty"`
	UserEmail      string `protobuf:"bytes,6,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserDateBirth  string `protobuf:"bytes,7,opt,name=userDateBirth,proto3" json:"userDateBirth,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateUserRequest) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *CreateUserRequest) GetUserRePassword() string {
	if x != nil {
		return x.UserRePassword
	}
	return ""
}

func (x *CreateUserRequest) GetUserSex() string {
	if x != nil {
		return x.UserSex
	}
	return ""
}

func (x *CreateUserRequest) GetUserAge() int32 {
	if x != nil {
		return x.UserAge
	}
	return 0
}

func (x *CreateUserRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *CreateUserRequest) GetUserDateBirth() string {
	if x != nil {
		return x.UserDateBirth
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SectionId int32  `protobuf:"varint,2,opt,name=sectionId,proto3" json:"sectionId,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateTopicRequest) Reset() {
	*x = CreateTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicRequest) ProtoMessage() {}

func (x *CreateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicRequest.ProtoReflect.Descriptor instead.
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTopicRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTopicRequest) GetSectionId() int32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

func (x *CreateTopicRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTopicRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateTopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId int32 `protobuf:"varint,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
}

func (x *CreateTopicResponse) Reset() {
	*x = CreateTopicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicResponse) ProtoMessage() {}

func (x *CreateTopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicResponse.ProtoReflect.Descriptor instead.
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopicResponse) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

type CreateReplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TopicId int32  `protobuf:"varint,2,opt,name=topicId,proto3" json:"topicId,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateReplyRequest) Reset() {
	*x = CreateReplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReplyRequest) ProtoMessage() {}

func (x *CreateReplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReplyRequest.ProtoReflect.Descriptor instead.
func (*CreateReplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{4}
}

func (x *CreateReplyRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateReplyRequest) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *CreateReplyRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateReplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyId int32 `protobuf:"varint,1,opt,name=replyId,proto3" json:"replyId,omitempty"`
}

func (x *CreateReplyResponse) Reset() {
	*x = CreateReplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReplyResponse) ProtoMessage() {}

func (x *CreateReplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReplyResponse.ProtoReflect.Descriptor instead.
func (*CreateReplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{5}
}

func (x *CreateReplyResponse) GetReplyId() int32 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

type GetTopicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId int32 `protobuf:"varint,1,opt,name=sectionId,proto3" json:"sectionId,omitempty"`
}

func (x *GetTopicsRequest) Reset() {
	*x = GetTopicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicsRequest) ProtoMessage() {}

func (x *GetTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicsRequest.ProtoReflect.Descriptor instead.
func (*GetTopicsRequest) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{6}
}

func (x *GetTopicsRequest) GetSectionId() int32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

type GetTopicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *GetTopicsResponse) Reset() {
	*x = GetTopicsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicsResponse) ProtoMessage() {}

func (x *GetTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicsResponse.ProtoReflect.Descriptor instead.
func (*GetTopicsResponse) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{7}
}

func (x *GetTopicsResponse) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId    int32  `protobuf:"varint,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
	UserId     int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ReplyCount int32  `protobuf:"varint,5,opt,name=replyCount,proto3" json:"replyCount,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{8}
}

func (x *Topic) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *Topic) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Topic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Topic) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Topic) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

type GetRepliesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId int32 `protobuf:"varint,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
}

func (x *GetRepliesRequest) Reset() {
	*x = GetRepliesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepliesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepliesRequest) ProtoMessage() {}

func (x *GetRepliesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepliesRequest.ProtoReflect.Descriptor instead.
func (*GetRepliesRequest) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{9}
}

func (x *GetRepliesRequest) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

type GetRepliesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replies []*Reply `protobuf:"bytes,1,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *GetRepliesResponse) Reset() {
	*x = GetRepliesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepliesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepliesResponse) ProtoMessage() {}

func (x *GetRepliesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepliesResponse.ProtoReflect.Descriptor instead.
func (*GetRepliesResponse) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{10}
}

func (x *GetRepliesResponse) GetReplies() []*Reply {
	if x != nil {
		return x.Replies
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyId int32  `protobuf:"varint,1,opt,name=replyId,proto3" json:"replyId,omitempty"`
	UserId  int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forum_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forum_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_forum_proto_rawDescGZIP(), []int{11}
}

func (x *Reply) GetReplyId() int32 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

func (x *Reply) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Reply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_proto_forum_proto protoreflect.FileDescriptor

var file_proto_forum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7a,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x30,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x05,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xe0, 0x02, 0x0a, 0x0c, 0x46, 0x6f,
	0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_forum_proto_rawDescOnce sync.Once
	file_proto_forum_proto_rawDescData = file_proto_forum_proto_rawDesc
)

func file_proto_forum_proto_rawDescGZIP() []byte {
	file_proto_forum_proto_rawDescOnce.Do(func() {
		file_proto_forum_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_forum_proto_rawDescData)
	})
	return file_proto_forum_proto_rawDescData
}

var file_proto_forum_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_forum_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),   // 0: proto.CreateUserRequest
	(*CreateUserResponse)(nil),  // 1: proto.CreateUserResponse
	(*CreateTopicRequest)(nil),  // 2: proto.CreateTopicRequest
	(*CreateTopicResponse)(nil), // 3: proto.CreateTopicResponse
	(*CreateReplyRequest)(nil),  // 4: proto.CreateReplyRequest
	(*CreateReplyResponse)(nil), // 5: proto.CreateReplyResponse
	(*GetTopicsRequest)(nil),    // 6: proto.GetTopicsRequest
	(*GetTopicsResponse)(nil),   // 7: proto.GetTopicsResponse
	(*Topic)(nil),               // 8: proto.Topic
	(*GetRepliesRequest)(nil),   // 9: proto.GetRepliesRequest
	(*GetRepliesResponse)(nil),  // 10: proto.GetRepliesResponse
	(*Reply)(nil),               // 11: proto.Reply
}
var file_proto_forum_proto_depIdxs = []int32{
	8,  // 0: proto.GetTopicsResponse.topics:type_name -> proto.Topic
	11, // 1: proto.GetRepliesResponse.replies:type_name -> proto.Reply
	0,  // 2: proto.ForumService.CreateUser:input_type -> proto.CreateUserRequest
	2,  // 3: proto.ForumService.CreateTopic:input_type -> proto.CreateTopicRequest
	4,  // 4: proto.ForumService.CreateReply:input_type -> proto.CreateReplyRequest
	6,  // 5: proto.ForumService.GetTopics:input_type -> proto.GetTopicsRequest
	9,  // 6: proto.ForumService.GetReplies:input_type -> proto.GetRepliesRequest
	1,  // 7: proto.ForumService.CreateUser:output_type -> proto.CreateUserResponse
	3,  // 8: proto.ForumService.CreateTopic:output_type -> proto.CreateTopicResponse
	5,  // 9: proto.ForumService.CreateReply:output_type -> proto.CreateReplyResponse
	7,  // 10: proto.ForumService.GetTopics:output_type -> proto.GetTopicsResponse
	10, // 11: proto.ForumService.GetReplies:output_type -> proto.GetRepliesResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_forum_proto_init() }
func file_proto_forum_proto_init() {
	if File_proto_forum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_forum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopicRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopicResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReplyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopicsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopicsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepliesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepliesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_forum_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_forum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_forum_proto_goTypes,
		DependencyIndexes: file_proto_forum_proto_depIdxs,
		MessageInfos:      file_proto_forum_proto_msgTypes,
	}.Build()
	File_proto_forum_proto = out.File
	file_proto_forum_proto_rawDesc = nil
	file_proto_forum_proto_goTypes = nil
	file_proto_forum_proto_depIdxs = nil
}