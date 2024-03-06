// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
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
	0x65, 0x61, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_jmt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jmt_proto_goTypes = []interface{}{
	(*Node)(nil),         // 0: pb.Node
	(*InternalNode)(nil), // 1: pb.InternalNode
	(*LeafNode)(nil),     // 2: pb.LeafNode
	(*Child)(nil),        // 3: pb.Child
}
var file_jmt_proto_depIdxs = []int32{
	3, // 0: pb.InternalNode.children:type_name -> pb.Child
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
