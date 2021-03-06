// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: avg/avgpb/avg.proto

package avgpb

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

type Calculate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Calculate) Reset() {
	*x = Calculate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avg_avgpb_avg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calculate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calculate) ProtoMessage() {}

func (x *Calculate) ProtoReflect() protoreflect.Message {
	mi := &file_avg_avgpb_avg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calculate.ProtoReflect.Descriptor instead.
func (*Calculate) Descriptor() ([]byte, []int) {
	return file_avg_avgpb_avg_proto_rawDescGZIP(), []int{0}
}

func (x *Calculate) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calc *Calculate `protobuf:"bytes,1,opt,name=calc,proto3" json:"calc,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avg_avgpb_avg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avg_avgpb_avg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_avg_avgpb_avg_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateRequest) GetCalc() *Calculate {
	if x != nil {
		return x.Calc
	}
	return nil
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res float64 `protobuf:"fixed64,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avg_avgpb_avg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avg_avgpb_avg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_avg_avgpb_avg_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateResponse) GetRes() float64 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_avg_avgpb_avg_proto protoreflect.FileDescriptor

var file_avg_avgpb_avg_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x76, 0x67, 0x2f, 0x61, 0x76, 0x67, 0x70, 0x62, 0x2f, 0x61, 0x76, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x76, 0x67, 0x70, 0x62, 0x22, 0x23, 0x0a, 0x09,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x76, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x25, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x72,
	0x65, 0x73, 0x32, 0x51, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x42, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x61, 0x76, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x76, 0x67, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x61, 0x76, 0x67, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avg_avgpb_avg_proto_rawDescOnce sync.Once
	file_avg_avgpb_avg_proto_rawDescData = file_avg_avgpb_avg_proto_rawDesc
)

func file_avg_avgpb_avg_proto_rawDescGZIP() []byte {
	file_avg_avgpb_avg_proto_rawDescOnce.Do(func() {
		file_avg_avgpb_avg_proto_rawDescData = protoimpl.X.CompressGZIP(file_avg_avgpb_avg_proto_rawDescData)
	})
	return file_avg_avgpb_avg_proto_rawDescData
}

var file_avg_avgpb_avg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_avg_avgpb_avg_proto_goTypes = []interface{}{
	(*Calculate)(nil),         // 0: avgpb.Calculate
	(*CalculateRequest)(nil),  // 1: avgpb.CalculateRequest
	(*CalculateResponse)(nil), // 2: avgpb.CalculateResponse
}
var file_avg_avgpb_avg_proto_depIdxs = []int32{
	0, // 0: avgpb.CalculateRequest.calc:type_name -> avgpb.Calculate
	1, // 1: avgpb.CalcService.Calculate:input_type -> avgpb.CalculateRequest
	2, // 2: avgpb.CalcService.Calculate:output_type -> avgpb.CalculateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_avg_avgpb_avg_proto_init() }
func file_avg_avgpb_avg_proto_init() {
	if File_avg_avgpb_avg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avg_avgpb_avg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calculate); i {
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
		file_avg_avgpb_avg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_avg_avgpb_avg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
			RawDescriptor: file_avg_avgpb_avg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avg_avgpb_avg_proto_goTypes,
		DependencyIndexes: file_avg_avgpb_avg_proto_depIdxs,
		MessageInfos:      file_avg_avgpb_avg_proto_msgTypes,
	}.Build()
	File_avg_avgpb_avg_proto = out.File
	file_avg_avgpb_avg_proto_rawDesc = nil
	file_avg_avgpb_avg_proto_goTypes = nil
	file_avg_avgpb_avg_proto_depIdxs = nil
}
