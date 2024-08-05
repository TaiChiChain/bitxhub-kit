// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: jmt.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leaf    bool   `protobuf:"varint,1,opt,name=leaf,proto3" json:"leaf,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetLeaf() bool {
	if x != nil {
		return x.Leaf
	}
	return false
}

func (x *Node) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type InternalNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children []*Child `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *InternalNode) Reset() {
	*x = InternalNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalNode) ProtoMessage() {}

func (x *InternalNode) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalNode.ProtoReflect.Descriptor instead.
func (*InternalNode) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{1}
}

func (x *InternalNode) GetChildren() []*Child {
	if x != nil {
		return x.Children
	}
	return nil
}

type LeafNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Hash  []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *LeafNode) Reset() {
	*x = LeafNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeafNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeafNode) ProtoMessage() {}

func (x *LeafNode) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeafNode.ProtoReflect.Descriptor instead.
func (*LeafNode) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{2}
}

func (x *LeafNode) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *LeafNode) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *LeafNode) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Leaf    bool   `protobuf:"varint,3,opt,name=leaf,proto3" json:"leaf,omitempty"`
}

func (x *Child) Reset() {
	*x = Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Child) ProtoMessage() {}

func (x *Child) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Child.ProtoReflect.Descriptor instead.
func (*Child) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{3}
}

func (x *Child) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Child) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Child) GetLeaf() bool {
	if x != nil {
		return x.Leaf
	}
	return false
}

type StateJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrieJournal []*TrieJournal    `protobuf:"bytes,1,rep,name=trie_journal,json=trieJournal,proto3" json:"trie_journal,omitempty"`
	CodeJournal map[string][]byte `protobuf:"bytes,2,rep,name=code_journal,json=codeJournal,proto3" json:"code_journal,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SnapJournal *SnapJournal      `protobuf:"bytes,3,opt,name=snap_journal,json=snapJournal,proto3" json:"snap_journal,omitempty"`
	RootHash    []byte            `protobuf:"bytes,4,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
}

func (x *StateJournal) Reset() {
	*x = StateJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateJournal) ProtoMessage() {}

func (x *StateJournal) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateJournal.ProtoReflect.Descriptor instead.
func (*StateJournal) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{4}
}

func (x *StateJournal) GetTrieJournal() []*TrieJournal {
	if x != nil {
		return x.TrieJournal
	}
	return nil
}

func (x *StateJournal) GetCodeJournal() map[string][]byte {
	if x != nil {
		return x.CodeJournal
	}
	return nil
}

func (x *StateJournal) GetSnapJournal() *SnapJournal {
	if x != nil {
		return x.SnapJournal
	}
	return nil
}

func (x *StateJournal) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

type TrieJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirtySet    map[string][]byte `protobuf:"bytes,1,rep,name=dirty_set,json=dirtySet,proto3" json:"dirty_set,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PruneSet    map[string][]byte `protobuf:"bytes,2,rep,name=prune_set,json=pruneSet,proto3" json:"prune_set,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type        uint64            `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	RootHash    []byte            `protobuf:"bytes,4,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	RootNodeKey []byte            `protobuf:"bytes,5,opt,name=root_node_key,json=rootNodeKey,proto3" json:"root_node_key,omitempty"`
}

func (x *TrieJournal) Reset() {
	*x = TrieJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrieJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrieJournal) ProtoMessage() {}

func (x *TrieJournal) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrieJournal.ProtoReflect.Descriptor instead.
func (*TrieJournal) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{5}
}

func (x *TrieJournal) GetDirtySet() map[string][]byte {
	if x != nil {
		return x.DirtySet
	}
	return nil
}

func (x *TrieJournal) GetPruneSet() map[string][]byte {
	if x != nil {
		return x.PruneSet
	}
	return nil
}

func (x *TrieJournal) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TrieJournal) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

func (x *TrieJournal) GetRootNodeKey() []byte {
	if x != nil {
		return x.RootNodeKey
	}
	return nil
}

type SnapJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destruct map[string][]byte       `protobuf:"bytes,1,rep,name=destruct,proto3" json:"destruct,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Account  map[string][]byte       `protobuf:"bytes,2,rep,name=account,proto3" json:"account,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Storage  map[string]*SnapStorage `protobuf:"bytes,3,rep,name=storage,proto3" json:"storage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SnapJournal) Reset() {
	*x = SnapJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapJournal) ProtoMessage() {}

func (x *SnapJournal) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapJournal.ProtoReflect.Descriptor instead.
func (*SnapJournal) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{6}
}

func (x *SnapJournal) GetDestruct() map[string][]byte {
	if x != nil {
		return x.Destruct
	}
	return nil
}

func (x *SnapJournal) GetAccount() map[string][]byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *SnapJournal) GetStorage() map[string]*SnapStorage {
	if x != nil {
		return x.Storage
	}
	return nil
}

type SnapStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapStorage map[string][]byte `protobuf:"bytes,1,rep,name=snap_storage,json=snapStorage,proto3" json:"snap_storage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SnapStorage) Reset() {
	*x = SnapStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapStorage) ProtoMessage() {}

func (x *SnapStorage) ProtoReflect() protoreflect.Message {
	mi := &file_jmt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapStorage.ProtoReflect.Descriptor instead.
func (*SnapStorage) Descriptor() ([]byte, []int) {
	return file_jmt_proto_rawDescGZIP(), []int{7}
}

func (x *SnapStorage) GetSnapStorage() map[string][]byte {
	if x != nil {
		return x.SnapStorage
	}
	return nil
}

var File_jmt_proto protoreflect.FileDescriptor

var file_jmt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x34, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x08,
	0x4c, 0x65, 0x61, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x49, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x65, 0x61, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x22,
	0x99, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x12, 0x32, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x65, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x69, 0x65,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x65, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63,
	0x6f, 0x64, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x0c, 0x73, 0x6e,
	0x61, 0x70, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x52, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x3e, 0x0a, 0x10, 0x43,
	0x6f, 0x64, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd4, 0x02, 0x0a, 0x0b,
	0x54, 0x72, 0x69, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x69, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x44, 0x69, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x64,
	0x69, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x72, 0x75, 0x6e, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x72, 0x69, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x75, 0x6e, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x6f, 0x6f,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x74,
	0x79, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x53, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xfe, 0x02, 0x0a, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x36, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x3b, 0x0a,
	0x0d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x6e, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x6e, 0x61,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jmt_proto_rawDescOnce sync.Once
	file_jmt_proto_rawDescData = file_jmt_proto_rawDesc
)

func file_jmt_proto_rawDescGZIP() []byte {
	file_jmt_proto_rawDescOnce.Do(func() {
		file_jmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_jmt_proto_rawDescData)
	})
	return file_jmt_proto_rawDescData
}

var file_jmt_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_jmt_proto_goTypes = []interface{}{
	(*Node)(nil),         // 0: pb.Node
	(*InternalNode)(nil), // 1: pb.InternalNode
	(*LeafNode)(nil),     // 2: pb.LeafNode
	(*Child)(nil),        // 3: pb.Child
	(*StateJournal)(nil), // 4: pb.StateJournal
	(*TrieJournal)(nil),  // 5: pb.TrieJournal
	(*SnapJournal)(nil),  // 6: pb.SnapJournal
	(*SnapStorage)(nil),  // 7: pb.SnapStorage
	nil,                  // 8: pb.StateJournal.CodeJournalEntry
	nil,                  // 9: pb.TrieJournal.DirtySetEntry
	nil,                  // 10: pb.TrieJournal.PruneSetEntry
	nil,                  // 11: pb.SnapJournal.DestructEntry
	nil,                  // 12: pb.SnapJournal.AccountEntry
	nil,                  // 13: pb.SnapJournal.StorageEntry
	nil,                  // 14: pb.SnapStorage.SnapStorageEntry
}
var file_jmt_proto_depIdxs = []int32{
	3,  // 0: pb.InternalNode.children:type_name -> pb.Child
	5,  // 1: pb.StateJournal.trie_journal:type_name -> pb.TrieJournal
	8,  // 2: pb.StateJournal.code_journal:type_name -> pb.StateJournal.CodeJournalEntry
	6,  // 3: pb.StateJournal.snap_journal:type_name -> pb.SnapJournal
	9,  // 4: pb.TrieJournal.dirty_set:type_name -> pb.TrieJournal.DirtySetEntry
	10, // 5: pb.TrieJournal.prune_set:type_name -> pb.TrieJournal.PruneSetEntry
	11, // 6: pb.SnapJournal.destruct:type_name -> pb.SnapJournal.DestructEntry
	12, // 7: pb.SnapJournal.account:type_name -> pb.SnapJournal.AccountEntry
	13, // 8: pb.SnapJournal.storage:type_name -> pb.SnapJournal.StorageEntry
	14, // 9: pb.SnapStorage.snap_storage:type_name -> pb.SnapStorage.SnapStorageEntry
	7,  // 10: pb.SnapJournal.StorageEntry.value:type_name -> pb.SnapStorage
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_jmt_proto_init() }
func file_jmt_proto_init() {
	if File_jmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_jmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalNode); i {
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
		file_jmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeafNode); i {
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
		file_jmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Child); i {
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
		file_jmt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateJournal); i {
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
		file_jmt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrieJournal); i {
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
		file_jmt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapJournal); i {
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
		file_jmt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapStorage); i {
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
			RawDescriptor: file_jmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jmt_proto_goTypes,
		DependencyIndexes: file_jmt_proto_depIdxs,
		MessageInfos:      file_jmt_proto_msgTypes,
	}.Build()
	File_jmt_proto = out.File
	file_jmt_proto_rawDesc = nil
	file_jmt_proto_goTypes = nil
	file_jmt_proto_depIdxs = nil
}
