// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/protos/app.proto

package protos

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

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Editor      string `protobuf:"bytes,2,opt,name=editor,proto3" json:"editor,omitempty"`
	ReferenceId string `protobuf:"bytes,3,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Data        string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Metadata    string `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Label       string `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_api_protos_app_proto_rawDescGZIP(), []int{0}
}

func (x *Snapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Snapshot) GetEditor() string {
	if x != nil {
		return x.Editor
	}
	return ""
}

func (x *Snapshot) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Snapshot) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Snapshot) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Snapshot) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type SnapshotCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Previous  *Snapshot `protobuf:"bytes,1,opt,name=previous,proto3" json:"previous,omitempty"`
	Current   *Snapshot `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	EventName string    `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Label     string    `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *SnapshotCreatedEvent) Reset() {
	*x = SnapshotCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotCreatedEvent) ProtoMessage() {}

func (x *SnapshotCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotCreatedEvent.ProtoReflect.Descriptor instead.
func (*SnapshotCreatedEvent) Descriptor() ([]byte, []int) {
	return file_api_protos_app_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotCreatedEvent) GetPrevious() *Snapshot {
	if x != nil {
		return x.Previous
	}
	return nil
}

func (x *SnapshotCreatedEvent) GetCurrent() *Snapshot {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *SnapshotCreatedEvent) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *SnapshotCreatedEvent) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_api_protos_app_proto protoreflect.FileDescriptor

var file_api_protos_app_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x70, 0x22, 0x9b, 0x01, 0x0a, 0x08,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x9f, 0x01, 0x0a, 0x14, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12, 0x27, 0x0a,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_app_proto_rawDescOnce sync.Once
	file_api_protos_app_proto_rawDescData = file_api_protos_app_proto_rawDesc
)

func file_api_protos_app_proto_rawDescGZIP() []byte {
	file_api_protos_app_proto_rawDescOnce.Do(func() {
		file_api_protos_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_app_proto_rawDescData)
	})
	return file_api_protos_app_proto_rawDescData
}

var file_api_protos_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_protos_app_proto_goTypes = []interface{}{
	(*Snapshot)(nil),             // 0: app.Snapshot
	(*SnapshotCreatedEvent)(nil), // 1: app.SnapshotCreatedEvent
}
var file_api_protos_app_proto_depIdxs = []int32{
	0, // 0: app.SnapshotCreatedEvent.previous:type_name -> app.Snapshot
	0, // 1: app.SnapshotCreatedEvent.current:type_name -> app.Snapshot
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_protos_app_proto_init() }
func file_api_protos_app_proto_init() {
	if File_api_protos_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_api_protos_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotCreatedEvent); i {
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
			RawDescriptor: file_api_protos_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protos_app_proto_goTypes,
		DependencyIndexes: file_api_protos_app_proto_depIdxs,
		MessageInfos:      file_api_protos_app_proto_msgTypes,
	}.Build()
	File_api_protos_app_proto = out.File
	file_api_protos_app_proto_rawDesc = nil
	file_api_protos_app_proto_goTypes = nil
	file_api_protos_app_proto_depIdxs = nil
}
