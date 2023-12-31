// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: internal/grpc/proto/grpc_server.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetLongUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLongUrlRequest) Reset() {
	*x = GetLongUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLongUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLongUrlRequest) ProtoMessage() {}

func (x *GetLongUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLongUrlRequest.ProtoReflect.Descriptor instead.
func (*GetLongUrlRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{0}
}

func (x *GetLongUrlRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *GetShortUrlRequest) Reset() {
	*x = GetShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortUrlRequest) ProtoMessage() {}

func (x *GetShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortUrlRequest.ProtoReflect.Descriptor instead.
func (*GetShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetShortUrlRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type GetShortUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetShortUrlResponse) Reset() {
	*x = GetShortUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortUrlResponse) ProtoMessage() {}

func (x *GetShortUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortUrlResponse.ProtoReflect.Descriptor instead.
func (*GetShortUrlResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetShortUrlResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type UrlData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *UrlData) Reset() {
	*x = UrlData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlData) ProtoMessage() {}

func (x *UrlData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlData.ProtoReflect.Descriptor instead.
func (*UrlData) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{3}
}

func (x *UrlData) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *UrlData) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type GetUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*UrlData `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetUrlsResponse) Reset() {
	*x = GetUrlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlsResponse) ProtoMessage() {}

func (x *GetUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlsResponse.ProtoReflect.Descriptor instead.
func (*GetUrlsResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetUrlsResponse) GetUrls() []*UrlData {
	if x != nil {
		return x.Urls
	}
	return nil
}

type DeleteUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *DeleteUrlsRequest) Reset() {
	*x = DeleteUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlsRequest) ProtoMessage() {}

func (x *DeleteUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_grpc_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUrlsRequest.ProtoReflect.Descriptor instead.
func (*DeleteUrlsRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_grpc_server_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUrlsRequest) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_internal_grpc_proto_grpc_server_proto protoreflect.FileDescriptor

var file_internal_grpc_proto_grpc_server_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c,
	0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x49, 0x0a, 0x07, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22,
	0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x32, 0xc1,
	0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x50, 0x69, 0x6e, 0x67,
	0x44, 0x62, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x75, 0x72, 0x74, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_proto_grpc_server_proto_rawDescOnce sync.Once
	file_internal_grpc_proto_grpc_server_proto_rawDescData = file_internal_grpc_proto_grpc_server_proto_rawDesc
)

func file_internal_grpc_proto_grpc_server_proto_rawDescGZIP() []byte {
	file_internal_grpc_proto_grpc_server_proto_rawDescOnce.Do(func() {
		file_internal_grpc_proto_grpc_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_proto_grpc_server_proto_rawDescData)
	})
	return file_internal_grpc_proto_grpc_server_proto_rawDescData
}

var file_internal_grpc_proto_grpc_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_grpc_proto_grpc_server_proto_goTypes = []interface{}{
	(*GetLongUrlRequest)(nil),   // 0: grpc.GetLongUrlRequest
	(*GetShortUrlRequest)(nil),  // 1: grpc.GetShortUrlRequest
	(*GetShortUrlResponse)(nil), // 2: grpc.GetShortUrlResponse
	(*UrlData)(nil),             // 3: grpc.UrlData
	(*GetUrlsResponse)(nil),     // 4: grpc.GetUrlsResponse
	(*DeleteUrlsRequest)(nil),   // 5: grpc.DeleteUrlsRequest
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_internal_grpc_proto_grpc_server_proto_depIdxs = []int32{
	3, // 0: grpc.GetUrlsResponse.urls:type_name -> grpc.UrlData
	0, // 1: grpc.Shortener.GetLongUrl:input_type -> grpc.GetLongUrlRequest
	1, // 2: grpc.Shortener.GetShortUrl:input_type -> grpc.GetShortUrlRequest
	6, // 3: grpc.Shortener.GetUrls:input_type -> google.protobuf.Empty
	5, // 4: grpc.Shortener.DeleteUrls:input_type -> grpc.DeleteUrlsRequest
	6, // 5: grpc.Shortener.PingDb:input_type -> google.protobuf.Empty
	6, // 6: grpc.Shortener.GetLongUrl:output_type -> google.protobuf.Empty
	2, // 7: grpc.Shortener.GetShortUrl:output_type -> grpc.GetShortUrlResponse
	4, // 8: grpc.Shortener.GetUrls:output_type -> grpc.GetUrlsResponse
	6, // 9: grpc.Shortener.DeleteUrls:output_type -> google.protobuf.Empty
	6, // 10: grpc.Shortener.PingDb:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_grpc_proto_grpc_server_proto_init() }
func file_internal_grpc_proto_grpc_server_proto_init() {
	if File_internal_grpc_proto_grpc_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_proto_grpc_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLongUrlRequest); i {
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
		file_internal_grpc_proto_grpc_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortUrlRequest); i {
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
		file_internal_grpc_proto_grpc_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortUrlResponse); i {
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
		file_internal_grpc_proto_grpc_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlData); i {
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
		file_internal_grpc_proto_grpc_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlsResponse); i {
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
		file_internal_grpc_proto_grpc_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUrlsRequest); i {
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
			RawDescriptor: file_internal_grpc_proto_grpc_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_proto_grpc_server_proto_goTypes,
		DependencyIndexes: file_internal_grpc_proto_grpc_server_proto_depIdxs,
		MessageInfos:      file_internal_grpc_proto_grpc_server_proto_msgTypes,
	}.Build()
	File_internal_grpc_proto_grpc_server_proto = out.File
	file_internal_grpc_proto_grpc_server_proto_rawDesc = nil
	file_internal_grpc_proto_grpc_server_proto_goTypes = nil
	file_internal_grpc_proto_grpc_server_proto_depIdxs = nil
}
