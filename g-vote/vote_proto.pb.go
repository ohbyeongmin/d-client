// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: vote_proto.proto

package g_vote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId     uint64   `protobuf:"varint,1,opt,name=Board_id,json=BoardId,proto3" json:"Board_id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Vote_Items  []string `protobuf:"bytes,4,rep,name=Vote_Items,json=VoteItems,proto3" json:"Vote_Items,omitempty"`
	Deadline    string   `protobuf:"bytes,5,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
}

func (x *CreateQuery) Reset() {
	*x = CreateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuery) ProtoMessage() {}

func (x *CreateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuery.ProtoReflect.Descriptor instead.
func (*CreateQuery) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{0}
}

func (x *CreateQuery) GetBoardId() uint64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *CreateQuery) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateQuery) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateQuery) GetVote_Items() []string {
	if x != nil {
		return x.Vote_Items
	}
	return nil
}

func (x *CreateQuery) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteId string `protobuf:"bytes,1,opt,name=Vote_id,json=VoteId,proto3" json:"Vote_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetVoteId() string {
	if x != nil {
		return x.VoteId
	}
	return ""
}

type SearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId uint64 `protobuf:"varint,1,opt,name=Board_id,json=BoardId,proto3" json:"Board_id,omitempty"`
	VoteId  string `protobuf:"bytes,2,opt,name=Vote_id,json=VoteId,proto3" json:"Vote_id,omitempty"`
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuery.ProtoReflect.Descriptor instead.
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{2}
}

func (x *SearchQuery) GetBoardId() uint64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *SearchQuery) GetVoteId() string {
	if x != nil {
		return x.VoteId
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string           `protobuf:"bytes,1,opt,name=User_id,json=UserId,proto3" json:"User_id,omitempty"`
	Title       string           `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	VoteItems   map[int32]string `protobuf:"bytes,4,rep,name=Vote_items,json=VoteItems,proto3" json:"Vote_items,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Deadline    string           `protobuf:"bytes,5,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	Voted       bool             `protobuf:"varint,6,opt,name=Voted,proto3" json:"Voted,omitempty"`
	VoteNum     map[int32]int32  `protobuf:"bytes,7,rep,name=Vote_num,json=VoteNum,proto3" json:"Vote_num,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{3}
}

func (x *SearchResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SearchResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SearchResponse) GetVoteItems() map[int32]string {
	if x != nil {
		return x.VoteItems
	}
	return nil
}

func (x *SearchResponse) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *SearchResponse) GetVoted() bool {
	if x != nil {
		return x.Voted
	}
	return false
}

func (x *SearchResponse) GetVoteNum() map[int32]int32 {
	if x != nil {
		return x.VoteNum
	}
	return nil
}

type SelectQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId    uint64 `protobuf:"varint,1,opt,name=Board_id,json=BoardId,proto3" json:"Board_id,omitempty"`
	VoteId     string `protobuf:"bytes,2,opt,name=Vote_id,json=VoteId,proto3" json:"Vote_id,omitempty"`
	VoteItemId int32  `protobuf:"varint,3,opt,name=Vote_item_id,json=VoteItemId,proto3" json:"Vote_item_id,omitempty"`
}

func (x *SelectQuery) Reset() {
	*x = SelectQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectQuery) ProtoMessage() {}

func (x *SelectQuery) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectQuery.ProtoReflect.Descriptor instead.
func (*SelectQuery) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{4}
}

func (x *SelectQuery) GetBoardId() uint64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *SelectQuery) GetVoteId() string {
	if x != nil {
		return x.VoteId
	}
	return ""
}

func (x *SelectQuery) GetVoteItemId() int32 {
	if x != nil {
		return x.VoteItemId
	}
	return 0
}

type SelectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SelectResponse) Reset() {
	*x = SelectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectResponse) ProtoMessage() {}

func (x *SelectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectResponse.ProtoReflect.Descriptor instead.
func (*SelectResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{5}
}

func (x *SelectResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *SelectResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MyVotesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MyVotesQuery) Reset() {
	*x = MyVotesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyVotesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyVotesQuery) ProtoMessage() {}

func (x *MyVotesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyVotesQuery.ProtoReflect.Descriptor instead.
func (*MyVotesQuery) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{6}
}

type MyVotesField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId int64  `protobuf:"varint,1,opt,name=Board_id,json=BoardId,proto3" json:"Board_id,omitempty"`
	VoteId  string `protobuf:"bytes,2,opt,name=Vote_id,json=VoteId,proto3" json:"Vote_id,omitempty"`
}

func (x *MyVotesField) Reset() {
	*x = MyVotesField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyVotesField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyVotesField) ProtoMessage() {}

func (x *MyVotesField) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyVotesField.ProtoReflect.Descriptor instead.
func (*MyVotesField) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{7}
}

func (x *MyVotesField) GetBoardId() int64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *MyVotesField) GetVoteId() string {
	if x != nil {
		return x.VoteId
	}
	return ""
}

type MyVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyVotes []*MyVotesField `protobuf:"bytes,1,rep,name=My_votes,json=MyVotes,proto3" json:"My_votes,omitempty"`
}

func (x *MyVotesResponse) Reset() {
	*x = MyVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyVotesResponse) ProtoMessage() {}

func (x *MyVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyVotesResponse.ProtoReflect.Descriptor instead.
func (*MyVotesResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_proto_rawDescGZIP(), []int{8}
}

func (x *MyVotesResponse) GetMyVotes() []*MyVotesField {
	if x != nil {
		return x.MyVotes
	}
	return nil
}

var File_vote_proto_proto protoreflect.FileDescriptor

var file_vote_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56, 0x6f, 0x74,
	0x65, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x56, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x93, 0x03, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x56, 0x6f,
	0x74, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x6f, 0x74,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x4e,
	0x75, 0x6d, 0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3a, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x0b,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x4f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x0c, 0x4d, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0f, 0x4d, 0x79, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x4d, 0x79, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x4d, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x07, 0x4d, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xf1, 0x01,
	0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x5f, 0x76, 0x6f,
	0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x5f, 0x76,
	0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x4d, 0x79, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x4d, 0x79, 0x56, 0x6f,
	0x74, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x2e, 0x4d, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x67, 0x2d, 0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vote_proto_proto_rawDescOnce sync.Once
	file_vote_proto_proto_rawDescData = file_vote_proto_proto_rawDesc
)

func file_vote_proto_proto_rawDescGZIP() []byte {
	file_vote_proto_proto_rawDescOnce.Do(func() {
		file_vote_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_vote_proto_proto_rawDescData)
	})
	return file_vote_proto_proto_rawDescData
}

var file_vote_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_vote_proto_proto_goTypes = []interface{}{
	(*CreateQuery)(nil),     // 0: g_vote.CreateQuery
	(*CreateResponse)(nil),  // 1: g_vote.CreateResponse
	(*SearchQuery)(nil),     // 2: g_vote.SearchQuery
	(*SearchResponse)(nil),  // 3: g_vote.SearchResponse
	(*SelectQuery)(nil),     // 4: g_vote.SelectQuery
	(*SelectResponse)(nil),  // 5: g_vote.SelectResponse
	(*MyVotesQuery)(nil),    // 6: g_vote.MyVotesQuery
	(*MyVotesField)(nil),    // 7: g_vote.MyVotesField
	(*MyVotesResponse)(nil), // 8: g_vote.MyVotesResponse
	nil,                     // 9: g_vote.SearchResponse.VoteItemsEntry
	nil,                     // 10: g_vote.SearchResponse.VoteNumEntry
}
var file_vote_proto_proto_depIdxs = []int32{
	9,  // 0: g_vote.SearchResponse.Vote_items:type_name -> g_vote.SearchResponse.VoteItemsEntry
	10, // 1: g_vote.SearchResponse.Vote_num:type_name -> g_vote.SearchResponse.VoteNumEntry
	7,  // 2: g_vote.MyVotesResponse.My_votes:type_name -> g_vote.MyVotesField
	0,  // 3: g_vote.Vote.CreateVote:input_type -> g_vote.CreateQuery
	2,  // 4: g_vote.Vote.SearchVote:input_type -> g_vote.SearchQuery
	4,  // 5: g_vote.Vote.SelectVote:input_type -> g_vote.SelectQuery
	6,  // 6: g_vote.Vote.MyVotes:input_type -> g_vote.MyVotesQuery
	1,  // 7: g_vote.Vote.CreateVote:output_type -> g_vote.CreateResponse
	3,  // 8: g_vote.Vote.SearchVote:output_type -> g_vote.SearchResponse
	5,  // 9: g_vote.Vote.SelectVote:output_type -> g_vote.SelectResponse
	8,  // 10: g_vote.Vote.MyVotes:output_type -> g_vote.MyVotesResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_vote_proto_proto_init() }
func file_vote_proto_proto_init() {
	if File_vote_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vote_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuery); i {
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
		file_vote_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_vote_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQuery); i {
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
		file_vote_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_vote_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectQuery); i {
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
		file_vote_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectResponse); i {
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
		file_vote_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyVotesQuery); i {
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
		file_vote_proto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyVotesField); i {
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
		file_vote_proto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyVotesResponse); i {
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
			RawDescriptor: file_vote_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vote_proto_proto_goTypes,
		DependencyIndexes: file_vote_proto_proto_depIdxs,
		MessageInfos:      file_vote_proto_proto_msgTypes,
	}.Build()
	File_vote_proto_proto = out.File
	file_vote_proto_proto_rawDesc = nil
	file_vote_proto_proto_goTypes = nil
	file_vote_proto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VoteClient is the client API for Vote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VoteClient interface {
	CreateVote(ctx context.Context, in *CreateQuery, opts ...grpc.CallOption) (*CreateResponse, error)
	SearchVote(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResponse, error)
	SelectVote(ctx context.Context, in *SelectQuery, opts ...grpc.CallOption) (*SelectResponse, error)
	MyVotes(ctx context.Context, in *MyVotesQuery, opts ...grpc.CallOption) (*MyVotesResponse, error)
}

type voteClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteClient(cc grpc.ClientConnInterface) VoteClient {
	return &voteClient{cc}
}

func (c *voteClient) CreateVote(ctx context.Context, in *CreateQuery, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/g_vote.Vote/CreateVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteClient) SearchVote(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/g_vote.Vote/SearchVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteClient) SelectVote(ctx context.Context, in *SelectQuery, opts ...grpc.CallOption) (*SelectResponse, error) {
	out := new(SelectResponse)
	err := c.cc.Invoke(ctx, "/g_vote.Vote/SelectVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteClient) MyVotes(ctx context.Context, in *MyVotesQuery, opts ...grpc.CallOption) (*MyVotesResponse, error) {
	out := new(MyVotesResponse)
	err := c.cc.Invoke(ctx, "/g_vote.Vote/MyVotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServer is the server API for Vote service.
type VoteServer interface {
	CreateVote(context.Context, *CreateQuery) (*CreateResponse, error)
	SearchVote(context.Context, *SearchQuery) (*SearchResponse, error)
	SelectVote(context.Context, *SelectQuery) (*SelectResponse, error)
	MyVotes(context.Context, *MyVotesQuery) (*MyVotesResponse, error)
}

// UnimplementedVoteServer can be embedded to have forward compatible implementations.
type UnimplementedVoteServer struct {
}

func (*UnimplementedVoteServer) CreateVote(context.Context, *CreateQuery) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVote not implemented")
}
func (*UnimplementedVoteServer) SearchVote(context.Context, *SearchQuery) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVote not implemented")
}
func (*UnimplementedVoteServer) SelectVote(context.Context, *SelectQuery) (*SelectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectVote not implemented")
}
func (*UnimplementedVoteServer) MyVotes(context.Context, *MyVotesQuery) (*MyVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyVotes not implemented")
}

func RegisterVoteServer(s *grpc.Server, srv VoteServer) {
	s.RegisterService(&_Vote_serviceDesc, srv)
}

func _Vote_CreateVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).CreateVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g_vote.Vote/CreateVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).CreateVote(ctx, req.(*CreateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vote_SearchVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).SearchVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g_vote.Vote/SearchVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).SearchVote(ctx, req.(*SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vote_SelectVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).SelectVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g_vote.Vote/SelectVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).SelectVote(ctx, req.(*SelectQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vote_MyVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyVotesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).MyVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g_vote.Vote/MyVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).MyVotes(ctx, req.(*MyVotesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vote_serviceDesc = grpc.ServiceDesc{
	ServiceName: "g_vote.Vote",
	HandlerType: (*VoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVote",
			Handler:    _Vote_CreateVote_Handler,
		},
		{
			MethodName: "SearchVote",
			Handler:    _Vote_SearchVote_Handler,
		},
		{
			MethodName: "SelectVote",
			Handler:    _Vote_SelectVote_Handler,
		},
		{
			MethodName: "MyVotes",
			Handler:    _Vote_MyVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vote_proto.proto",
}
