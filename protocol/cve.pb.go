// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: cve.proto

package CveSubscriptionSrv

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

type EmSubscriptionType int32

const (
	EmSubscriptionType_MoBile EmSubscriptionType = 0
	EmSubscriptionType_Mail   EmSubscriptionType = 1
)

// Enum value maps for EmSubscriptionType.
var (
	EmSubscriptionType_name = map[int32]string{
		0: "MoBile",
		1: "Mail",
	}
	EmSubscriptionType_value = map[string]int32{
		"MoBile": 0,
		"Mail":   1,
	}
)

func (x EmSubscriptionType) Enum() *EmSubscriptionType {
	p := new(EmSubscriptionType)
	*p = x
	return p
}

func (x EmSubscriptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmSubscriptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_cve_proto_enumTypes[0].Descriptor()
}

func (EmSubscriptionType) Type() protoreflect.EnumType {
	return &file_cve_proto_enumTypes[0]
}

func (x EmSubscriptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmSubscriptionType.Descriptor instead.
func (EmSubscriptionType) EnumDescriptor() ([]byte, []int) {
	return file_cve_proto_rawDescGZIP(), []int{0}
}

// The request
type HandleSubscriptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SType   EmSubscriptionType `protobuf:"varint,1,opt,name=sType,proto3,enum=CveSubscriptionSrv.EmSubscriptionType" json:"sType,omitempty"`
	Account string             `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *HandleSubscriptionReq) Reset() {
	*x = HandleSubscriptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleSubscriptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleSubscriptionReq) ProtoMessage() {}

func (x *HandleSubscriptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_cve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleSubscriptionReq.ProtoReflect.Descriptor instead.
func (*HandleSubscriptionReq) Descriptor() ([]byte, []int) {
	return file_cve_proto_rawDescGZIP(), []int{0}
}

func (x *HandleSubscriptionReq) GetSType() EmSubscriptionType {
	if x != nil {
		return x.SType
	}
	return EmSubscriptionType_MoBile
}

func (x *HandleSubscriptionReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// The response
type HandleSubscriptionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HandleSubscriptionRsp) Reset() {
	*x = HandleSubscriptionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleSubscriptionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleSubscriptionRsp) ProtoMessage() {}

func (x *HandleSubscriptionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleSubscriptionRsp.ProtoReflect.Descriptor instead.
func (*HandleSubscriptionRsp) Descriptor() ([]byte, []int) {
	return file_cve_proto_rawDescGZIP(), []int{1}
}

func (x *HandleSubscriptionRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cve_proto protoreflect.FileDescriptor

var file_cve_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x43, 0x76, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x22,
	0x6f, 0x0a, 0x15, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x43, 0x76, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x2e, 0x45, 0x6d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x31, 0x0a, 0x15, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x2a, 0x0a, 0x12, 0x45, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x42,
	0x69, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x32,
	0x82, 0x01, 0x0a, 0x12, 0x43, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x12, 0x6c, 0x0a, 0x12, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x43,
	0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72,
	0x76, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x43, 0x76, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x2e, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x3b, 0x43, 0x76, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cve_proto_rawDescOnce sync.Once
	file_cve_proto_rawDescData = file_cve_proto_rawDesc
)

func file_cve_proto_rawDescGZIP() []byte {
	file_cve_proto_rawDescOnce.Do(func() {
		file_cve_proto_rawDescData = protoimpl.X.CompressGZIP(file_cve_proto_rawDescData)
	})
	return file_cve_proto_rawDescData
}

var file_cve_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cve_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cve_proto_goTypes = []interface{}{
	(EmSubscriptionType)(0),       // 0: CveSubscriptionSrv.EmSubscriptionType
	(*HandleSubscriptionReq)(nil), // 1: CveSubscriptionSrv.HandleSubscriptionReq
	(*HandleSubscriptionRsp)(nil), // 2: CveSubscriptionSrv.HandleSubscriptionRsp
}
var file_cve_proto_depIdxs = []int32{
	0, // 0: CveSubscriptionSrv.HandleSubscriptionReq.sType:type_name -> CveSubscriptionSrv.EmSubscriptionType
	1, // 1: CveSubscriptionSrv.CveSubscriptionSrv.HandleSubscription:input_type -> CveSubscriptionSrv.HandleSubscriptionReq
	2, // 2: CveSubscriptionSrv.CveSubscriptionSrv.HandleSubscription:output_type -> CveSubscriptionSrv.HandleSubscriptionRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cve_proto_init() }
func file_cve_proto_init() {
	if File_cve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleSubscriptionReq); i {
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
		file_cve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleSubscriptionRsp); i {
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
			RawDescriptor: file_cve_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cve_proto_goTypes,
		DependencyIndexes: file_cve_proto_depIdxs,
		EnumInfos:         file_cve_proto_enumTypes,
		MessageInfos:      file_cve_proto_msgTypes,
	}.Build()
	File_cve_proto = out.File
	file_cve_proto_rawDesc = nil
	file_cve_proto_goTypes = nil
	file_cve_proto_depIdxs = nil
}
