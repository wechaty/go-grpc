// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: wechaty/puppet/file-box.proto

package puppet

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

type FileBoxChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*FileBoxChunk_Data
	//	*FileBoxChunk_Name
	Payload isFileBoxChunk_Payload `protobuf_oneof:"payload"`
}

func (x *FileBoxChunk) Reset() {
	*x = FileBoxChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_file_box_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileBoxChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileBoxChunk) ProtoMessage() {}

func (x *FileBoxChunk) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_file_box_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileBoxChunk.ProtoReflect.Descriptor instead.
func (*FileBoxChunk) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_file_box_proto_rawDescGZIP(), []int{0}
}

func (m *FileBoxChunk) GetPayload() isFileBoxChunk_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *FileBoxChunk) GetData() []byte {
	if x, ok := x.GetPayload().(*FileBoxChunk_Data); ok {
		return x.Data
	}
	return nil
}

func (x *FileBoxChunk) GetName() string {
	if x, ok := x.GetPayload().(*FileBoxChunk_Name); ok {
		return x.Name
	}
	return ""
}

type isFileBoxChunk_Payload interface {
	isFileBoxChunk_Payload()
}

type FileBoxChunk_Data struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type FileBoxChunk_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*FileBoxChunk_Data) isFileBoxChunk_Payload() {}

func (*FileBoxChunk_Name) isFileBoxChunk_Payload() {}

var File_wechaty_puppet_file_box_proto protoreflect.FileDescriptor

var file_wechaty_puppet_file_box_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x22,
	0x45, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x6f, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x48, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70,
	0x70, 0x65, 0x74, 0xaa, 0x02, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechaty_puppet_file_box_proto_rawDescOnce sync.Once
	file_wechaty_puppet_file_box_proto_rawDescData = file_wechaty_puppet_file_box_proto_rawDesc
)

func file_wechaty_puppet_file_box_proto_rawDescGZIP() []byte {
	file_wechaty_puppet_file_box_proto_rawDescOnce.Do(func() {
		file_wechaty_puppet_file_box_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechaty_puppet_file_box_proto_rawDescData)
	})
	return file_wechaty_puppet_file_box_proto_rawDescData
}

var file_wechaty_puppet_file_box_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wechaty_puppet_file_box_proto_goTypes = []interface{}{
	(*FileBoxChunk)(nil), // 0: wechaty.puppet.FileBoxChunk
}
var file_wechaty_puppet_file_box_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wechaty_puppet_file_box_proto_init() }
func file_wechaty_puppet_file_box_proto_init() {
	if File_wechaty_puppet_file_box_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechaty_puppet_file_box_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileBoxChunk); i {
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
	file_wechaty_puppet_file_box_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FileBoxChunk_Data)(nil),
		(*FileBoxChunk_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wechaty_puppet_file_box_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wechaty_puppet_file_box_proto_goTypes,
		DependencyIndexes: file_wechaty_puppet_file_box_proto_depIdxs,
		MessageInfos:      file_wechaty_puppet_file_box_proto_msgTypes,
	}.Build()
	File_wechaty_puppet_file_box_proto = out.File
	file_wechaty_puppet_file_box_proto_rawDesc = nil
	file_wechaty_puppet_file_box_proto_goTypes = nil
	file_wechaty_puppet_file_box_proto_depIdxs = nil
}
