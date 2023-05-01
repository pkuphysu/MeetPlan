// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: model/common.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "meetplan/biz/model/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryPageParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   int32 `protobuf:"varint,1,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty" form:"page_no" query:"page_no"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size" query:"page_size"`
}

func (x *QueryPageParam) Reset() {
	*x = QueryPageParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPageParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPageParam) ProtoMessage() {}

func (x *QueryPageParam) ProtoReflect() protoreflect.Message {
	mi := &file_model_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPageParam.ProtoReflect.Descriptor instead.
func (*QueryPageParam) Descriptor() ([]byte, []int) {
	return file_model_common_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPageParam) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *QueryPageParam) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo     int32 `protobuf:"varint,1,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty" form:"page_no" query:"page_no"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size" query:"page_size"`
	TotalCount int64 `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" query:"total_count" form:"total_count" json:"total_count"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_model_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_model_common_proto_rawDescGZIP(), []int{1}
}

func (x *Pagination) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// 通用返回结构，code为0表示成功，其他值表示失败，message为失败原因
// 所有返回结构都应包含这两个字段
// 本结构不应该被直接使用
type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" form:"message" query:"message"` // any data = 3;
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_model_common_proto_rawDescGZIP(), []int{2}
}

func (x *BaseResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_model_common_proto protoreflect.FileDescriptor

var file_model_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x10, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a,
	0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7b, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x16,
	0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x14, 0x5a, 0x12, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x62, 0x69, 0x7a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_common_proto_rawDescOnce sync.Once
	file_model_common_proto_rawDescData = file_model_common_proto_rawDesc
)

func file_model_common_proto_rawDescGZIP() []byte {
	file_model_common_proto_rawDescOnce.Do(func() {
		file_model_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_common_proto_rawDescData)
	})
	return file_model_common_proto_rawDescData
}

var file_model_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_common_proto_goTypes = []interface{}{
	(*QueryPageParam)(nil), // 0: model.QueryPageParam
	(*Pagination)(nil),     // 1: model.Pagination
	(*BaseResponse)(nil),   // 2: model.BaseResponse
}
var file_model_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_common_proto_init() }
func file_model_common_proto_init() {
	if File_model_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPageParam); i {
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
		file_model_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_model_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
			RawDescriptor: file_model_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_common_proto_goTypes,
		DependencyIndexes: file_model_common_proto_depIdxs,
		MessageInfos:      file_model_common_proto_msgTypes,
	}.Build()
	File_model_common_proto = out.File
	file_model_common_proto_rawDesc = nil
	file_model_common_proto_goTypes = nil
	file_model_common_proto_depIdxs = nil
}
