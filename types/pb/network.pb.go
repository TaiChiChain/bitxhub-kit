// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: network.proto

package pb

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

type Status int32

const (
	Status_ERROR   Status = 0
	Status_SUCCESS Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ERROR",
		1: "SUCCESS",
	}
	Status_value = map[string]int32{
		"ERROR":   0,
		"SUCCESS": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_network_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_network_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0}
}

type Message_Type int32

const (
	Message_SYNC_STATE_REQUEST         Message_Type = 0
	Message_SYNC_STATE_RESPONSE        Message_Type = 1
	Message_SYNC_BLOCK_REQUEST         Message_Type = 2
	Message_SYNC_BLOCK_RESPONSE        Message_Type = 3
	Message_SYNC_CHAIN_DATA_REQUEST    Message_Type = 4
	Message_SYNC_CHAIN_DATA_RESPONSE   Message_Type = 5
	Message_FETCH_EPOCH_STATE_REQUEST  Message_Type = 6
	Message_FETCH_EPOCH_STATE_RESPONSE Message_Type = 7
)

// Enum value maps for Message_Type.
var (
	Message_Type_name = map[int32]string{
		0: "SYNC_STATE_REQUEST",
		1: "SYNC_STATE_RESPONSE",
		2: "SYNC_BLOCK_REQUEST",
		3: "SYNC_BLOCK_RESPONSE",
		4: "SYNC_CHAIN_DATA_REQUEST",
		5: "SYNC_CHAIN_DATA_RESPONSE",
		6: "FETCH_EPOCH_STATE_REQUEST",
		7: "FETCH_EPOCH_STATE_RESPONSE",
	}
	Message_Type_value = map[string]int32{
		"SYNC_STATE_REQUEST":         0,
		"SYNC_STATE_RESPONSE":        1,
		"SYNC_BLOCK_REQUEST":         2,
		"SYNC_BLOCK_RESPONSE":        3,
		"SYNC_CHAIN_DATA_REQUEST":    4,
		"SYNC_CHAIN_DATA_RESPONSE":   5,
		"FETCH_EPOCH_STATE_REQUEST":  6,
		"FETCH_EPOCH_STATE_RESPONSE": 7,
	}
)

func (x Message_Type) Enum() *Message_Type {
	p := new(Message_Type)
	*p = x
	return p
}

func (x Message_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_network_proto_enumTypes[1].Descriptor()
}

func (Message_Type) Type() protoreflect.EnumType {
	return &file_network_proto_enumTypes[1]
}

func (x Message_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Type.Descriptor instead.
func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string       `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Type    Message_Type `protobuf:"varint,2,opt,name=type,proto3,enum=pb.Message_Type" json:"type,omitempty"`
	Data    []byte       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Version uint64       `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetType() Message_Type {
	if x != nil {
		return x.Type
	}
	return Message_SYNC_STATE_REQUEST
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height  uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncStateRequest) Reset() {
	*x = SyncStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncStateRequest) ProtoMessage() {}

func (x *SyncStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncStateRequest.ProtoReflect.Descriptor instead.
func (*SyncStateRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{1}
}

func (x *SyncStateRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SyncStateRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckpointState *CheckpointState `protobuf:"bytes,1,opt,name=checkpoint_state,json=checkpointState,proto3" json:"checkpoint_state,omitempty"`
	Status          Status           `protobuf:"varint,2,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Error           string           `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Version         uint64           `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncStateResponse) Reset() {
	*x = SyncStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncStateResponse) ProtoMessage() {}

func (x *SyncStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncStateResponse.ProtoReflect.Descriptor instead.
func (*SyncStateResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{2}
}

func (x *SyncStateResponse) GetCheckpointState() *CheckpointState {
	if x != nil {
		return x.CheckpointState
	}
	return nil
}

func (x *SyncStateResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *SyncStateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SyncStateResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height  uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncBlockRequest) Reset() {
	*x = SyncBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBlockRequest) ProtoMessage() {}

func (x *SyncBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncBlockRequest.ProtoReflect.Descriptor instead.
func (*SyncBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{3}
}

func (x *SyncBlockRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SyncBlockRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block   []byte `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Status  Status `protobuf:"varint,2,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Version uint64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncBlockResponse) Reset() {
	*x = SyncBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBlockResponse) ProtoMessage() {}

func (x *SyncBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncBlockResponse.ProtoReflect.Descriptor instead.
func (*SyncBlockResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{4}
}

func (x *SyncBlockResponse) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *SyncBlockResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *SyncBlockResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SyncBlockResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncChainDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height  uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncChainDataRequest) Reset() {
	*x = SyncChainDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncChainDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncChainDataRequest) ProtoMessage() {}

func (x *SyncChainDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncChainDataRequest.ProtoReflect.Descriptor instead.
func (*SyncChainDataRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{5}
}

func (x *SyncChainDataRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SyncChainDataRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncChainDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block    []byte `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Receipts []byte `protobuf:"bytes,2,opt,name=receipts,proto3" json:"receipts,omitempty"`
	Status   Status `protobuf:"varint,3,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Error    string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Version  uint64 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncChainDataResponse) Reset() {
	*x = SyncChainDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncChainDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncChainDataResponse) ProtoMessage() {}

func (x *SyncChainDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncChainDataResponse.ProtoReflect.Descriptor instead.
func (*SyncChainDataResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{6}
}

func (x *SyncChainDataResponse) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *SyncChainDataResponse) GetReceipts() []byte {
	if x != nil {
		return x.Receipts
	}
	return nil
}

func (x *SyncChainDataResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *SyncChainDataResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SyncChainDataResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type FetchEpochStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (x *FetchEpochStateRequest) Reset() {
	*x = FetchEpochStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchEpochStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEpochStateRequest) ProtoMessage() {}

func (x *FetchEpochStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEpochStateRequest.ProtoReflect.Descriptor instead.
func (*FetchEpochStateRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{7}
}

func (x *FetchEpochStateRequest) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

type FetchEpochStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Status  Status `protobuf:"varint,2,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Version uint64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *FetchEpochStateResponse) Reset() {
	*x = FetchEpochStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchEpochStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEpochStateResponse) ProtoMessage() {}

func (x *FetchEpochStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEpochStateResponse.ProtoReflect.Descriptor instead.
func (*FetchEpochStateResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{8}
}

func (x *FetchEpochStateResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FetchEpochStateResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *FetchEpochStateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FetchEpochStateResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type CheckpointState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height       uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Digest       string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	LatestHeight uint64 `protobuf:"varint,3,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
}

func (x *CheckpointState) Reset() {
	*x = CheckpointState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointState) ProtoMessage() {}

func (x *CheckpointState) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointState.ProtoReflect.Descriptor instead.
func (*CheckpointState) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{9}
}

func (x *CheckpointState) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *CheckpointState) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CheckpointState) GetLatestHeight() uint64 {
	if x != nil {
		return x.LatestHeight
	}
	return 0
}

var File_network_proto protoreflect.FileDescriptor

var file_network_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xd6, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e,
	0x43, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x05, 0x12, 0x1d, 0x0a,
	0x19, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x50, 0x4f, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a,
	0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x50, 0x4f, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x07, 0x22, 0x44, 0x0a, 0x10,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x10,
	0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x48, 0x0a, 0x14, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x15,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x16, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x81, 0x01, 0x0a, 0x17,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x66, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2a, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_proto_rawDescOnce sync.Once
	file_network_proto_rawDescData = file_network_proto_rawDesc
)

func file_network_proto_rawDescGZIP() []byte {
	file_network_proto_rawDescOnce.Do(func() {
		file_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_proto_rawDescData)
	})
	return file_network_proto_rawDescData
}

var file_network_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_network_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_network_proto_goTypes = []interface{}{
	(Status)(0),                     // 0: pb.Status
	(Message_Type)(0),               // 1: pb.Message.Type
	(*Message)(nil),                 // 2: pb.Message
	(*SyncStateRequest)(nil),        // 3: pb.SyncStateRequest
	(*SyncStateResponse)(nil),       // 4: pb.SyncStateResponse
	(*SyncBlockRequest)(nil),        // 5: pb.SyncBlockRequest
	(*SyncBlockResponse)(nil),       // 6: pb.SyncBlockResponse
	(*SyncChainDataRequest)(nil),    // 7: pb.SyncChainDataRequest
	(*SyncChainDataResponse)(nil),   // 8: pb.SyncChainDataResponse
	(*FetchEpochStateRequest)(nil),  // 9: pb.FetchEpochStateRequest
	(*FetchEpochStateResponse)(nil), // 10: pb.FetchEpochStateResponse
	(*CheckpointState)(nil),         // 11: pb.CheckpointState
}
var file_network_proto_depIdxs = []int32{
	1,  // 0: pb.Message.type:type_name -> pb.Message.Type
	11, // 1: pb.SyncStateResponse.checkpoint_state:type_name -> pb.CheckpointState
	0,  // 2: pb.SyncStateResponse.status:type_name -> pb.Status
	0,  // 3: pb.SyncBlockResponse.status:type_name -> pb.Status
	0,  // 4: pb.SyncChainDataResponse.status:type_name -> pb.Status
	0,  // 5: pb.FetchEpochStateResponse.status:type_name -> pb.Status
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_network_proto_init() }
func file_network_proto_init() {
	if File_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncStateRequest); i {
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
		file_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncStateResponse); i {
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
		file_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncBlockRequest); i {
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
		file_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncBlockResponse); i {
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
		file_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncChainDataRequest); i {
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
		file_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncChainDataResponse); i {
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
		file_network_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchEpochStateRequest); i {
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
		file_network_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchEpochStateResponse); i {
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
		file_network_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointState); i {
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
			RawDescriptor: file_network_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_network_proto_goTypes,
		DependencyIndexes: file_network_proto_depIdxs,
		EnumInfos:         file_network_proto_enumTypes,
		MessageInfos:      file_network_proto_msgTypes,
	}.Build()
	File_network_proto = out.File
	file_network_proto_rawDesc = nil
	file_network_proto_goTypes = nil
	file_network_proto_depIdxs = nil
}
