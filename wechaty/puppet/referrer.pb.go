// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: wechaty/puppet/referrer.proto

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

//*
// "Referrer" and "Referral" refers to different things. "Referrer" is something or somebody who refers. "Referral" is the act of referring.
//  - https://english.stackexchange.com/questions/33135/referrer-versus-referral-versus-referer
type Referrer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactId string `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	RoomId    string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *Referrer) Reset() {
	*x = Referrer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_referrer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Referrer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Referrer) ProtoMessage() {}

func (x *Referrer) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_referrer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Referrer.ProtoReflect.Descriptor instead.
func (*Referrer) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_referrer_proto_rawDescGZIP(), []int{0}
}

func (x *Referrer) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *Referrer) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

var File_wechaty_puppet_referrer_proto protoreflect.FileDescriptor

var file_wechaty_puppet_referrer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x22,
	0x42, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x42, 0x67, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75,
	0x70, 0x70, 0x65, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0xaa,
	0x02, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechaty_puppet_referrer_proto_rawDescOnce sync.Once
	file_wechaty_puppet_referrer_proto_rawDescData = file_wechaty_puppet_referrer_proto_rawDesc
)

func file_wechaty_puppet_referrer_proto_rawDescGZIP() []byte {
	file_wechaty_puppet_referrer_proto_rawDescOnce.Do(func() {
		file_wechaty_puppet_referrer_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechaty_puppet_referrer_proto_rawDescData)
	})
	return file_wechaty_puppet_referrer_proto_rawDescData
}

var file_wechaty_puppet_referrer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wechaty_puppet_referrer_proto_goTypes = []interface{}{
	(*Referrer)(nil), // 0: wechaty.puppet.Referrer
}
var file_wechaty_puppet_referrer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wechaty_puppet_referrer_proto_init() }
func file_wechaty_puppet_referrer_proto_init() {
	if File_wechaty_puppet_referrer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechaty_puppet_referrer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Referrer); i {
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
			RawDescriptor: file_wechaty_puppet_referrer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wechaty_puppet_referrer_proto_goTypes,
		DependencyIndexes: file_wechaty_puppet_referrer_proto_depIdxs,
		MessageInfos:      file_wechaty_puppet_referrer_proto_msgTypes,
	}.Build()
	File_wechaty_puppet_referrer_proto = out.File
	file_wechaty_puppet_referrer_proto_rawDesc = nil
	file_wechaty_puppet_referrer_proto_goTypes = nil
	file_wechaty_puppet_referrer_proto_depIdxs = nil
}
