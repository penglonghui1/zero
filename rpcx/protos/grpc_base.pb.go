// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: grpc_base.proto

package pb

import (
	any "github.com/golang/protobuf/ptypes/any"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Total   int32    `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Data    *any.Any `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_base_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Response) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Platform      string `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	ClientVersion string `protobuf:"bytes,3,opt,name=ClientVersion,proto3" json:"ClientVersion,omitempty"`
	Token         string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	ClientIP      string `protobuf:"bytes,5,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`
	RequestID     string `protobuf:"bytes,6,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
}

func (x *UserContext) Reset() {
	*x = UserContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserContext) ProtoMessage() {}

func (x *UserContext) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserContext.ProtoReflect.Descriptor instead.
func (*UserContext) Descriptor() ([]byte, []int) {
	return file_grpc_base_proto_rawDescGZIP(), []int{1}
}

func (x *UserContext) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserContext) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *UserContext) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *UserContext) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserContext) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *UserContext) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

type NopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NopResponse) Reset() {
	*x = NopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NopResponse) ProtoMessage() {}

func (x *NopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NopResponse.ProtoReflect.Descriptor instead.
func (*NopResponse) Descriptor() ([]byte, []int) {
	return file_grpc_base_proto_rawDescGZIP(), []int{2}
}

type Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=PageNumber,proto3" json:"PageNumber,omitempty"`
	PageRecord int32 `protobuf:"varint,2,opt,name=PageRecord,proto3" json:"PageRecord,omitempty"`
}

func (x *Paging) Reset() {
	*x = Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_grpc_base_proto_rawDescGZIP(), []int{3}
}

func (x *Paging) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Paging) GetPageRecord() int32 {
	if x != nil {
		return x.PageRecord
	}
	return 0
}

var File_grpc_base_proto protoreflect.FileDescriptor

var file_grpc_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x64, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44,
	0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x48, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_base_proto_rawDescOnce sync.Once
	file_grpc_base_proto_rawDescData = file_grpc_base_proto_rawDesc
)

func file_grpc_base_proto_rawDescGZIP() []byte {
	file_grpc_base_proto_rawDescOnce.Do(func() {
		file_grpc_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_base_proto_rawDescData)
	})
	return file_grpc_base_proto_rawDescData
}

var file_grpc_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_base_proto_goTypes = []interface{}{
	(*Response)(nil),    // 0: pb.Response
	(*UserContext)(nil), // 1: pb.UserContext
	(*NopResponse)(nil), // 2: pb.NopResponse
	(*Paging)(nil),      // 3: pb.Paging
	(*any.Any)(nil),     // 4: google.protobuf.Any
}
var file_grpc_base_proto_depIdxs = []int32{
	4, // 0: pb.Response.Data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_base_proto_init() }
func file_grpc_base_proto_init() {
	if File_grpc_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpc_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserContext); i {
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
		file_grpc_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NopResponse); i {
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
		file_grpc_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paging); i {
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
			RawDescriptor: file_grpc_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_base_proto_goTypes,
		DependencyIndexes: file_grpc_base_proto_depIdxs,
		MessageInfos:      file_grpc_base_proto_msgTypes,
	}.Build()
	File_grpc_base_proto = out.File
	file_grpc_base_proto_rawDesc = nil
	file_grpc_base_proto_goTypes = nil
	file_grpc_base_proto_depIdxs = nil
}
