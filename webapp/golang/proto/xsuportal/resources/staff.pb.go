// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: xsuportal/resources/staff.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Staff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GithubLogin string `protobuf:"bytes,2,opt,name=github_login,json=githubLogin,proto3" json:"github_login,omitempty"`
}

func (x *Staff) Reset() {
	*x = Staff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_resources_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Staff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Staff) ProtoMessage() {}

func (x *Staff) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_resources_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Staff.ProtoReflect.Descriptor instead.
func (*Staff) Descriptor() ([]byte, []int) {
	return file_xsuportal_resources_staff_proto_rawDescGZIP(), []int{0}
}

func (x *Staff) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Staff) GetGithubLogin() string {
	if x != nil {
		return x.GithubLogin
	}
	return ""
}

var File_xsuportal_resources_staff_proto protoreflect.FileDescriptor

var file_xsuportal_resources_staff_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73,
	0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78,
	0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xsuportal_resources_staff_proto_rawDescOnce sync.Once
	file_xsuportal_resources_staff_proto_rawDescData = file_xsuportal_resources_staff_proto_rawDesc
)

func file_xsuportal_resources_staff_proto_rawDescGZIP() []byte {
	file_xsuportal_resources_staff_proto_rawDescOnce.Do(func() {
		file_xsuportal_resources_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_xsuportal_resources_staff_proto_rawDescData)
	})
	return file_xsuportal_resources_staff_proto_rawDescData
}

var file_xsuportal_resources_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xsuportal_resources_staff_proto_goTypes = []interface{}{
	(*Staff)(nil), // 0: xsuportal.proto.resources.Staff
}
var file_xsuportal_resources_staff_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_xsuportal_resources_staff_proto_init() }
func file_xsuportal_resources_staff_proto_init() {
	if File_xsuportal_resources_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xsuportal_resources_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Staff); i {
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
			RawDescriptor: file_xsuportal_resources_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xsuportal_resources_staff_proto_goTypes,
		DependencyIndexes: file_xsuportal_resources_staff_proto_depIdxs,
		MessageInfos:      file_xsuportal_resources_staff_proto_msgTypes,
	}.Build()
	File_xsuportal_resources_staff_proto = out.File
	file_xsuportal_resources_staff_proto_rawDesc = nil
	file_xsuportal_resources_staff_proto_goTypes = nil
	file_xsuportal_resources_staff_proto_depIdxs = nil
}
