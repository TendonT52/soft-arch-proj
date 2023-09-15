// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: v1/email.proto

package gen

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

type EmailData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL     string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *EmailData) Reset() {
	*x = EmailData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailData) ProtoMessage() {}

func (x *EmailData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailData.ProtoReflect.Descriptor instead.
func (*EmailData) Descriptor() ([]byte, []int) {
	return file_v1_email_proto_rawDescGZIP(), []int{0}
}

func (x *EmailData) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *EmailData) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmailData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_v1_email_proto protoreflect.FileDescriptor

var file_v1_email_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x61, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x69, 0x6b, 0x68, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x6e, 0x53, 0x6b, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_email_proto_rawDescOnce sync.Once
	file_v1_email_proto_rawDescData = file_v1_email_proto_rawDesc
)

func file_v1_email_proto_rawDescGZIP() []byte {
	file_v1_email_proto_rawDescOnce.Do(func() {
		file_v1_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_email_proto_rawDescData)
	})
	return file_v1_email_proto_rawDescData
}

var file_v1_email_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_email_proto_goTypes = []interface{}{
	(*EmailData)(nil), // 0: email.EmailData
}
var file_v1_email_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_email_proto_init() }
func file_v1_email_proto_init() {
	if File_v1_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailData); i {
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
			RawDescriptor: file_v1_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_email_proto_goTypes,
		DependencyIndexes: file_v1_email_proto_depIdxs,
		MessageInfos:      file_v1_email_proto_msgTypes,
	}.Build()
	File_v1_email_proto = out.File
	file_v1_email_proto_rawDesc = nil
	file_v1_email_proto_goTypes = nil
	file_v1_email_proto_depIdxs = nil
}