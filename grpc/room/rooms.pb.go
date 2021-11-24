// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: proto/rooms.proto

package room

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	shared "github.com/dimitarsi/go-chat-app/grpc/shared"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JoinResponse_JoinStatus int32

const (
	JoinResponse_Allowed JoinResponse_JoinStatus = 0
	JoinResponse_Denied  JoinResponse_JoinStatus = 1
)

// Enum value maps for JoinResponse_JoinStatus.
var (
	JoinResponse_JoinStatus_name = map[int32]string{
		0: "Allowed",
		1: "Denied",
	}
	JoinResponse_JoinStatus_value = map[string]int32{
		"Allowed": 0,
		"Denied":  1,
	}
)

func (x JoinResponse_JoinStatus) Enum() *JoinResponse_JoinStatus {
	p := new(JoinResponse_JoinStatus)
	*p = x
	return p
}

func (x JoinResponse_JoinStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JoinResponse_JoinStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_rooms_proto_enumTypes[0].Descriptor()
}

func (JoinResponse_JoinStatus) Type() protoreflect.EnumType {
	return &file_proto_rooms_proto_enumTypes[0]
}

func (x JoinResponse_JoinStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JoinResponse_JoinStatus.Descriptor instead.
func (JoinResponse_JoinStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{3, 0}
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomName string `protobuf:"bytes,1,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

type NewRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *shared.Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *NewRoomResponse) Reset() {
	*x = NewRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRoomResponse) ProtoMessage() {}

func (x *NewRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRoomResponse.ProtoReflect.Descriptor instead.
func (*NewRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{1}
}

func (x *NewRoomResponse) GetRoom() *shared.Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status JoinResponse_JoinStatus `protobuf:"varint,1,opt,name=status,proto3,enum=rooms.JoinResponse_JoinStatus" json:"status,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{3}
}

func (x *JoinResponse) GetStatus() JoinResponse_JoinStatus {
	if x != nil {
		return x.Status
	}
	return JoinResponse_Allowed
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{5}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term string `protobuf:"bytes,1,opt,name=Term,proto3" json:"Term,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{6}
}

func (x *SearchRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type RoomsFoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomsFound []*shared.Room `protobuf:"bytes,1,rep,name=RoomsFound,proto3" json:"RoomsFound,omitempty"`
}

func (x *RoomsFoundResponse) Reset() {
	*x = RoomsFoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsFoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsFoundResponse) ProtoMessage() {}

func (x *RoomsFoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsFoundResponse.ProtoReflect.Descriptor instead.
func (*RoomsFoundResponse) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{7}
}

func (x *RoomsFoundResponse) GetRoomsFound() []*shared.Room {
	if x != nil {
		return x.RoomsFound
	}
	return nil
}

type RoomDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *RoomDetailsRequest) Reset() {
	*x = RoomDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDetailsRequest) ProtoMessage() {}

func (x *RoomDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDetailsRequest.ProtoReflect.Descriptor instead.
func (*RoomDetailsRequest) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{8}
}

func (x *RoomDetailsRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type RoomDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participants *RoomDetailsResponse_Participants `protobuf:"bytes,1,opt,name=participants,proto3" json:"participants,omitempty"`
}

func (x *RoomDetailsResponse) Reset() {
	*x = RoomDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDetailsResponse) ProtoMessage() {}

func (x *RoomDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDetailsResponse.ProtoReflect.Descriptor instead.
func (*RoomDetailsResponse) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{9}
}

func (x *RoomDetailsResponse) GetParticipants() *RoomDetailsResponse_Participants {
	if x != nil {
		return x.Participants
	}
	return nil
}

type RoomDetailsResponse_Participants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*shared.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Count int32          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RoomDetailsResponse_Participants) Reset() {
	*x = RoomDetailsResponse_Participants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rooms_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDetailsResponse_Participants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDetailsResponse_Participants) ProtoMessage() {}

func (x *RoomDetailsResponse_Participants) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rooms_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDetailsResponse_Participants.ProtoReflect.Descriptor instead.
func (*RoomDetailsResponse_Participants) Descriptor() ([]byte, []int) {
	return file_proto_rooms_proto_rawDescGZIP(), []int{9, 0}
}

func (x *RoomDetailsResponse_Participants) GetUsers() []*shared.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *RoomDetailsResponse_Participants) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_rooms_proto protoreflect.FileDescriptor

var file_proto_rooms_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x33, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x0c, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x25, 0x0a, 0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x10, 0x01, 0x22, 0x26, 0x0a, 0x0c, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x42, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x2c, 0x0a, 0x12,
	0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x52,
	0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x1a,
	0x48, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12,
	0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb9, 0x02, 0x0a, 0x05, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x12, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x2e, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x6f,
	0x6f, 0x6d, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rooms_proto_rawDescOnce sync.Once
	file_proto_rooms_proto_rawDescData = file_proto_rooms_proto_rawDesc
)

func file_proto_rooms_proto_rawDescGZIP() []byte {
	file_proto_rooms_proto_rawDescOnce.Do(func() {
		file_proto_rooms_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rooms_proto_rawDescData)
	})
	return file_proto_rooms_proto_rawDescData
}

var file_proto_rooms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_rooms_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_rooms_proto_goTypes = []interface{}{
	(JoinResponse_JoinStatus)(0),             // 0: rooms.JoinResponse.JoinStatus
	(*CreateRoomRequest)(nil),                // 1: rooms.CreateRoomRequest
	(*NewRoomResponse)(nil),                  // 2: rooms.NewRoomResponse
	(*JoinRequest)(nil),                      // 3: rooms.JoinRequest
	(*JoinResponse)(nil),                     // 4: rooms.JoinResponse
	(*LeaveRequest)(nil),                     // 5: rooms.LeaveRequest
	(*LeaveResponse)(nil),                    // 6: rooms.LeaveResponse
	(*SearchRequest)(nil),                    // 7: rooms.SearchRequest
	(*RoomsFoundResponse)(nil),               // 8: rooms.RoomsFoundResponse
	(*RoomDetailsRequest)(nil),               // 9: rooms.RoomDetailsRequest
	(*RoomDetailsResponse)(nil),              // 10: rooms.RoomDetailsResponse
	(*RoomDetailsResponse_Participants)(nil), // 11: rooms.RoomDetailsResponse.Participants
	(*shared.Room)(nil),                      // 12: shared.Room
	(*shared.User)(nil),                      // 13: shared.User
}
var file_proto_rooms_proto_depIdxs = []int32{
	12, // 0: rooms.NewRoomResponse.room:type_name -> shared.Room
	0,  // 1: rooms.JoinResponse.status:type_name -> rooms.JoinResponse.JoinStatus
	12, // 2: rooms.RoomsFoundResponse.RoomsFound:type_name -> shared.Room
	11, // 3: rooms.RoomDetailsResponse.participants:type_name -> rooms.RoomDetailsResponse.Participants
	13, // 4: rooms.RoomDetailsResponse.Participants.users:type_name -> shared.User
	1,  // 5: rooms.Rooms.CreateRoom:input_type -> rooms.CreateRoomRequest
	3,  // 6: rooms.Rooms.JoinRoom:input_type -> rooms.JoinRequest
	5,  // 7: rooms.Rooms.LeaveRoom:input_type -> rooms.LeaveRequest
	7,  // 8: rooms.Rooms.FindRoom:input_type -> rooms.SearchRequest
	9,  // 9: rooms.Rooms.RoomDetails:input_type -> rooms.RoomDetailsRequest
	2,  // 10: rooms.Rooms.CreateRoom:output_type -> rooms.NewRoomResponse
	4,  // 11: rooms.Rooms.JoinRoom:output_type -> rooms.JoinResponse
	6,  // 12: rooms.Rooms.LeaveRoom:output_type -> rooms.LeaveResponse
	8,  // 13: rooms.Rooms.FindRoom:output_type -> rooms.RoomsFoundResponse
	10, // 14: rooms.Rooms.RoomDetails:output_type -> rooms.RoomDetailsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_rooms_proto_init() }
func file_proto_rooms_proto_init() {
	if File_proto_rooms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rooms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_proto_rooms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRoomResponse); i {
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
		file_proto_rooms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_proto_rooms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_proto_rooms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_proto_rooms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveResponse); i {
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
		file_proto_rooms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_rooms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsFoundResponse); i {
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
		file_proto_rooms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDetailsRequest); i {
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
		file_proto_rooms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDetailsResponse); i {
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
		file_proto_rooms_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDetailsResponse_Participants); i {
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
			RawDescriptor: file_proto_rooms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rooms_proto_goTypes,
		DependencyIndexes: file_proto_rooms_proto_depIdxs,
		EnumInfos:         file_proto_rooms_proto_enumTypes,
		MessageInfos:      file_proto_rooms_proto_msgTypes,
	}.Build()
	File_proto_rooms_proto = out.File
	file_proto_rooms_proto_rawDesc = nil
	file_proto_rooms_proto_goTypes = nil
	file_proto_rooms_proto_depIdxs = nil
}