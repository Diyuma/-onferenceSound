// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: service.proto

package protosound

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

type ChatServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []float32 `protobuf:"fixed32,3,rep,packed,name=data,proto3" json:"data,omitempty"` // fixed32?
	Rate    int64     `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	SoundId uint64    `protobuf:"varint,4,opt,name=soundId,proto3" json:"soundId,omitempty"`
	OnlyOne bool      `protobuf:"varint,5,opt,name=onlyOne,proto3" json:"onlyOne,omitempty"`
}

func (x *ChatServerMessage) Reset() {
	*x = ChatServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatServerMessage) ProtoMessage() {}

func (x *ChatServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatServerMessage.ProtoReflect.Descriptor instead.
func (*ChatServerMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChatServerMessage) GetData() []float32 {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChatServerMessage) GetRate() int64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ChatServerMessage) GetSoundId() uint64 {
	if x != nil {
		return x.SoundId
	}
	return 0
}

func (x *ChatServerMessage) GetOnlyOne() bool {
	if x != nil {
		return x.OnlyOne
	}
	return false
}

type ChatClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []float32 `protobuf:"fixed32,4,rep,packed,name=data,proto3" json:"data,omitempty"`
	Rate       int64     `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	UserId     uint32    `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	ConfId     uint64    `protobuf:"varint,7,opt,name=confId,proto3" json:"confId,omitempty"`
	TimeStamp  uint64    `protobuf:"varint,8,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	MessageInd uint32    `protobuf:"varint,9,opt,name=messageInd,proto3" json:"messageInd,omitempty"`
}

func (x *ChatClientMessage) Reset() {
	*x = ChatClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatClientMessage) ProtoMessage() {}

func (x *ChatClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatClientMessage.ProtoReflect.Descriptor instead.
func (*ChatClientMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChatClientMessage) GetData() []float32 {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChatClientMessage) GetRate() int64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ChatClientMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatClientMessage) GetConfId() uint64 {
	if x != nil {
		return x.ConfId
	}
	return 0
}

func (x *ChatClientMessage) GetTimeStamp() uint64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *ChatClientMessage) GetMessageInd() uint32 {
	if x != nil {
		return x.MessageInd
	}
	return 0
}

type ClientResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate       int64  `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
	SoundId    uint64 `protobuf:"varint,2,opt,name=soundId,proto3" json:"soundId,omitempty"`
	MessageInd uint32 `protobuf:"varint,3,opt,name=messageInd,proto3" json:"messageInd,omitempty"`
}

func (x *ClientResponseMessage) Reset() {
	*x = ClientResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponseMessage) ProtoMessage() {}

func (x *ClientResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponseMessage.ProtoReflect.Descriptor instead.
func (*ClientResponseMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClientResponseMessage) GetRate() int64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ClientResponseMessage) GetSoundId() uint64 {
	if x != nil {
		return x.SoundId
	}
	return 0
}

func (x *ClientResponseMessage) GetMessageInd() uint32 {
	if x != nil {
		return x.MessageInd
	}
	return 0
}

type ClientInfoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfId uint64 `protobuf:"varint,1,opt,name=confId,proto3" json:"confId,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ClientInfoMessage) Reset() {
	*x = ClientInfoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoMessage) ProtoMessage() {}

func (x *ClientInfoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoMessage.ProtoReflect.Descriptor instead.
func (*ClientInfoMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClientInfoMessage) GetConfId() uint64 {
	if x != nil {
		return x.ConfId
	}
	return 0
}

func (x *ClientInfoMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ClientUserInitResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ClientUserInitResponseMessage) Reset() {
	*x = ClientUserInitResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientUserInitResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientUserInitResponseMessage) ProtoMessage() {}

func (x *ClientUserInitResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientUserInitResponseMessage.ProtoReflect.Descriptor instead.
func (*ClientUserInitResponseMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *ClientUserInitResponseMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ClientConfInitResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfId uint64 `protobuf:"varint,1,opt,name=confId,proto3" json:"confId,omitempty"`
}

func (x *ClientConfInitResponseMessage) Reset() {
	*x = ClientConfInitResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfInitResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfInitResponseMessage) ProtoMessage() {}

func (x *ClientConfInitResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfInitResponseMessage.ProtoReflect.Descriptor instead.
func (*ClientConfInitResponseMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *ClientConfInitResponseMessage) GetConfId() uint64 {
	if x != nil {
		return x.ConfId
	}
	return 0
}

// not to use basic google.Empty now, may be later change
type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x75, 0x0a, 0x11, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6e, 0x6c, 0x79, 0x4f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x6c, 0x79, 0x4f, 0x6e, 0x65, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x22, 0xbb, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x4a, 0x04,
	0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06,
	0x22, 0x65, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x1d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x49, 0x64, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x92,
	0x03, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x09, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x50, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x4f, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_proto_goTypes = []interface{}{
	(*ChatServerMessage)(nil),             // 0: protosound.ChatServerMessage
	(*ChatClientMessage)(nil),             // 1: protosound.ChatClientMessage
	(*ClientResponseMessage)(nil),         // 2: protosound.ClientResponseMessage
	(*ClientInfoMessage)(nil),             // 3: protosound.ClientInfoMessage
	(*ClientUserInitResponseMessage)(nil), // 4: protosound.ClientUserInitResponseMessage
	(*ClientConfInitResponseMessage)(nil), // 5: protosound.ClientConfInitResponseMessage
	(*EmptyMessage)(nil),                  // 6: protosound.EmptyMessage
}
var file_service_proto_depIdxs = []int32{
	3, // 0: protosound.SoundService.GetSound:input_type -> protosound.ClientInfoMessage
	1, // 1: protosound.SoundService.SendSound:input_type -> protosound.ChatClientMessage
	3, // 2: protosound.SoundService.PingServer:input_type -> protosound.ClientInfoMessage
	6, // 3: protosound.SoundService.InitUser:input_type -> protosound.EmptyMessage
	6, // 4: protosound.SoundService.InitConf:input_type -> protosound.EmptyMessage
	0, // 5: protosound.SoundService.GetSound:output_type -> protosound.ChatServerMessage
	2, // 6: protosound.SoundService.SendSound:output_type -> protosound.ClientResponseMessage
	6, // 7: protosound.SoundService.PingServer:output_type -> protosound.EmptyMessage
	4, // 8: protosound.SoundService.InitUser:output_type -> protosound.ClientUserInitResponseMessage
	5, // 9: protosound.SoundService.InitConf:output_type -> protosound.ClientConfInitResponseMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatServerMessage); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatClientMessage); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponseMessage); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfoMessage); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientUserInitResponseMessage); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfInitResponseMessage); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}