// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: snapshot.proto

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

type SnapshotJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journals []*SnapshotJournalEntry `protobuf:"bytes,1,rep,name=journals,proto3" json:"journals,omitempty"`
}

func (x *SnapshotJournal) Reset() {
	*x = SnapshotJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotJournal) ProtoMessage() {}

func (x *SnapshotJournal) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotJournal.ProtoReflect.Descriptor instead.
func (*SnapshotJournal) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotJournal) GetJournals() []*SnapshotJournalEntry {
	if x != nil {
		return x.Journals
	}
	return nil
}

type SnapshotJournalEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        []byte            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrevAccount    []byte            `protobuf:"bytes,2,opt,name=prev_account,json=prevAccount,proto3" json:"prev_account,omitempty"`
	AccountChanged bool              `protobuf:"varint,3,opt,name=account_changed,json=accountChanged,proto3" json:"account_changed,omitempty"`
	PrevStates     map[string][]byte `protobuf:"bytes,4,rep,name=prev_states,json=prevStates,proto3" json:"prev_states,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SnapshotJournalEntry) Reset() {
	*x = SnapshotJournalEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotJournalEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotJournalEntry) ProtoMessage() {}

func (x *SnapshotJournalEntry) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotJournalEntry.ProtoReflect.Descriptor instead.
func (*SnapshotJournalEntry) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotJournalEntry) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *SnapshotJournalEntry) GetPrevAccount() []byte {
	if x != nil {
		return x.PrevAccount
	}
	return nil
}

func (x *SnapshotJournalEntry) GetAccountChanged() bool {
	if x != nil {
		return x.AccountChanged
	}
	return false
}

func (x *SnapshotJournalEntry) GetPrevStates() map[string][]byte {
	if x != nil {
		return x.PrevStates
	}
	return nil
}

var File_snapshot_proto protoreflect.FileDescriptor

var file_snapshot_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x47, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x22, 0x86, 0x02,
	0x0a, 0x14, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x0b,
	0x70, 0x72, 0x65, 0x76, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x76,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x65,
	0x76, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x76, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snapshot_proto_rawDescOnce sync.Once
	file_snapshot_proto_rawDescData = file_snapshot_proto_rawDesc
)

func file_snapshot_proto_rawDescGZIP() []byte {
	file_snapshot_proto_rawDescOnce.Do(func() {
		file_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_snapshot_proto_rawDescData)
	})
	return file_snapshot_proto_rawDescData
}

var file_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_snapshot_proto_goTypes = []interface{}{
	(*SnapshotJournal)(nil),      // 0: pb.SnapshotJournal
	(*SnapshotJournalEntry)(nil), // 1: pb.SnapshotJournalEntry
	nil,                          // 2: pb.SnapshotJournalEntry.PrevStatesEntry
}
var file_snapshot_proto_depIdxs = []int32{
	1, // 0: pb.SnapshotJournal.journals:type_name -> pb.SnapshotJournalEntry
	2, // 1: pb.SnapshotJournalEntry.prev_states:type_name -> pb.SnapshotJournalEntry.PrevStatesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_snapshot_proto_init() }
func file_snapshot_proto_init() {
	if File_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotJournal); i {
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
		file_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotJournalEntry); i {
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
			RawDescriptor: file_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_snapshot_proto_goTypes,
		DependencyIndexes: file_snapshot_proto_depIdxs,
		MessageInfos:      file_snapshot_proto_msgTypes,
	}.Build()
	File_snapshot_proto = out.File
	file_snapshot_proto_rawDesc = nil
	file_snapshot_proto_goTypes = nil
	file_snapshot_proto_depIdxs = nil
}
