// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: math/v1alpha1/math.proto

package v1alpha1

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

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content float32 `protobuf:"fixed32,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1alpha1_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1alpha1_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_math_v1alpha1_math_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetContent() float32 {
	if x != nil {
		return x.Content
	}
	return 0
}

type GetSquareOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Number `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetSquareOfRequest) Reset() {
	*x = GetSquareOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1alpha1_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSquareOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSquareOfRequest) ProtoMessage() {}

func (x *GetSquareOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1alpha1_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSquareOfRequest.ProtoReflect.Descriptor instead.
func (*GetSquareOfRequest) Descriptor() ([]byte, []int) {
	return file_math_v1alpha1_math_proto_rawDescGZIP(), []int{1}
}

func (x *GetSquareOfRequest) GetReq() *Number {
	if x != nil {
		return x.Req
	}
	return nil
}

type GetSquareOfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res *Number `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *GetSquareOfResponse) Reset() {
	*x = GetSquareOfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1alpha1_math_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSquareOfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSquareOfResponse) ProtoMessage() {}

func (x *GetSquareOfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1alpha1_math_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSquareOfResponse.ProtoReflect.Descriptor instead.
func (*GetSquareOfResponse) Descriptor() ([]byte, []int) {
	return file_math_v1alpha1_math_proto_rawDescGZIP(), []int{2}
}

func (x *GetSquareOfResponse) GetRes() *Number {
	if x != nil {
		return x.Res
	}
	return nil
}

type StreamSquareOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Number `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *StreamSquareOfRequest) Reset() {
	*x = StreamSquareOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1alpha1_math_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamSquareOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSquareOfRequest) ProtoMessage() {}

func (x *StreamSquareOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1alpha1_math_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSquareOfRequest.ProtoReflect.Descriptor instead.
func (*StreamSquareOfRequest) Descriptor() ([]byte, []int) {
	return file_math_v1alpha1_math_proto_rawDescGZIP(), []int{3}
}

func (x *StreamSquareOfRequest) GetReq() *Number {
	if x != nil {
		return x.Req
	}
	return nil
}

type StreamSquareOfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res *Number `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *StreamSquareOfResponse) Reset() {
	*x = StreamSquareOfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1alpha1_math_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamSquareOfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSquareOfResponse) ProtoMessage() {}

func (x *StreamSquareOfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1alpha1_math_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSquareOfResponse.ProtoReflect.Descriptor instead.
func (*StreamSquareOfResponse) Descriptor() ([]byte, []int) {
	return file_math_v1alpha1_math_proto_rawDescGZIP(), []int{4}
}

func (x *StreamSquareOfResponse) GetRes() *Number {
	if x != nil {
		return x.Res
	}
	return nil
}

var File_math_v1alpha1_math_proto protoreflect.FileDescriptor

var file_math_v1alpha1_math_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x22, 0x0a, 0x06, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x3e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x15,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x41,
	0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x03, 0x72, 0x65,
	0x73, 0x32, 0xc8, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x4f, 0x66, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x12, 0x24, 0x2e, 0x6d, 0x61,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x65, 0x6c, 0x6f, 0x6f, 0x70, 0x6d, 0x65, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_math_v1alpha1_math_proto_rawDescOnce sync.Once
	file_math_v1alpha1_math_proto_rawDescData = file_math_v1alpha1_math_proto_rawDesc
)

func file_math_v1alpha1_math_proto_rawDescGZIP() []byte {
	file_math_v1alpha1_math_proto_rawDescOnce.Do(func() {
		file_math_v1alpha1_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_math_v1alpha1_math_proto_rawDescData)
	})
	return file_math_v1alpha1_math_proto_rawDescData
}

var file_math_v1alpha1_math_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_math_v1alpha1_math_proto_goTypes = []interface{}{
	(*Number)(nil),                 // 0: math.v1alpha1.Number
	(*GetSquareOfRequest)(nil),     // 1: math.v1alpha1.GetSquareOfRequest
	(*GetSquareOfResponse)(nil),    // 2: math.v1alpha1.GetSquareOfResponse
	(*StreamSquareOfRequest)(nil),  // 3: math.v1alpha1.StreamSquareOfRequest
	(*StreamSquareOfResponse)(nil), // 4: math.v1alpha1.StreamSquareOfResponse
}
var file_math_v1alpha1_math_proto_depIdxs = []int32{
	0, // 0: math.v1alpha1.GetSquareOfRequest.req:type_name -> math.v1alpha1.Number
	0, // 1: math.v1alpha1.GetSquareOfResponse.res:type_name -> math.v1alpha1.Number
	0, // 2: math.v1alpha1.StreamSquareOfRequest.req:type_name -> math.v1alpha1.Number
	0, // 3: math.v1alpha1.StreamSquareOfResponse.res:type_name -> math.v1alpha1.Number
	1, // 4: math.v1alpha1.MathOperation.GetSquareOf:input_type -> math.v1alpha1.GetSquareOfRequest
	3, // 5: math.v1alpha1.MathOperation.StreamSquareOf:input_type -> math.v1alpha1.StreamSquareOfRequest
	2, // 6: math.v1alpha1.MathOperation.GetSquareOf:output_type -> math.v1alpha1.GetSquareOfResponse
	4, // 7: math.v1alpha1.MathOperation.StreamSquareOf:output_type -> math.v1alpha1.StreamSquareOfResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_math_v1alpha1_math_proto_init() }
func file_math_v1alpha1_math_proto_init() {
	if File_math_v1alpha1_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_math_v1alpha1_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_math_v1alpha1_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSquareOfRequest); i {
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
		file_math_v1alpha1_math_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSquareOfResponse); i {
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
		file_math_v1alpha1_math_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamSquareOfRequest); i {
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
		file_math_v1alpha1_math_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamSquareOfResponse); i {
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
			RawDescriptor: file_math_v1alpha1_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_math_v1alpha1_math_proto_goTypes,
		DependencyIndexes: file_math_v1alpha1_math_proto_depIdxs,
		MessageInfos:      file_math_v1alpha1_math_proto_msgTypes,
	}.Build()
	File_math_v1alpha1_math_proto = out.File
	file_math_v1alpha1_math_proto_rawDesc = nil
	file_math_v1alpha1_math_proto_goTypes = nil
	file_math_v1alpha1_math_proto_depIdxs = nil
}