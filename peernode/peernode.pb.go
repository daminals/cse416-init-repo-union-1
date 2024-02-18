// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: peernode.proto

package peernode

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FileLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link           string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Token          string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PaymentAddress string `protobuf:"bytes,3,opt,name=payment_address,json=paymentAddress,proto3" json:"payment_address,omitempty"`
}

func (x *FileLink) Reset() {
	*x = FileLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLink) ProtoMessage() {}

func (x *FileLink) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileLink.ProtoReflect.Descriptor instead.
func (*FileLink) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{2}
}

func (x *FileLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *FileLink) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FileLink) GetPaymentAddress() string {
	if x != nil {
		return x.PaymentAddress
	}
	return ""
}

type FileHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *FileHash) Reset() {
	*x = FileHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHash) ProtoMessage() {}

func (x *FileHash) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHash.ProtoReflect.Descriptor instead.
func (*FileHash) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{3}
}

func (x *FileHash) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{4}
}

func (x *FileRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *FileRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type FileRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*FileRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *FileRequestList) Reset() {
	*x = FileRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peernode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequestList) ProtoMessage() {}

func (x *FileRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_peernode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequestList.ProtoReflect.Descriptor instead.
func (*FileRequestList) Descriptor() ([]byte, []int) {
	return file_peernode_proto_rawDescGZIP(), []int{5}
}

func (x *FileRequestList) GetRequests() []*FileRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

var File_peernode_proto protoreflect.FileDescriptor

var file_peernode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x31, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x44, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x32, 0x4c,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x65,
	0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x52, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x32, 0x93, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x19, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x65, 0x65, 0x72,
	0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peernode_proto_rawDescOnce sync.Once
	file_peernode_proto_rawDescData = file_peernode_proto_rawDesc
)

func file_peernode_proto_rawDescGZIP() []byte {
	file_peernode_proto_rawDescOnce.Do(func() {
		file_peernode_proto_rawDescData = protoimpl.X.CompressGZIP(file_peernode_proto_rawDescData)
	})
	return file_peernode_proto_rawDescData
}

var file_peernode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_peernode_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),    // 0: peernode.HelloRequest
	(*HelloResponse)(nil),   // 1: peernode.HelloResponse
	(*FileLink)(nil),        // 2: peernode.FileLink
	(*FileHash)(nil),        // 3: peernode.FileHash
	(*FileRequest)(nil),     // 4: peernode.FileRequest
	(*FileRequestList)(nil), // 5: peernode.FileRequestList
	(*emptypb.Empty)(nil),   // 6: google.protobuf.Empty
}
var file_peernode_proto_depIdxs = []int32{
	4, // 0: peernode.FileRequestList.requests:type_name -> peernode.FileRequest
	0, // 1: peernode.HealthService.SayHello:input_type -> peernode.HelloRequest
	2, // 2: peernode.ConsumerService.ReceiveFileInfo:input_type -> peernode.FileLink
	3, // 3: peernode.MarketService.AddFileRequest:input_type -> peernode.FileHash
	3, // 4: peernode.MarketService.GetFileRequests:input_type -> peernode.FileHash
	1, // 5: peernode.HealthService.SayHello:output_type -> peernode.HelloResponse
	6, // 6: peernode.ConsumerService.ReceiveFileInfo:output_type -> google.protobuf.Empty
	6, // 7: peernode.MarketService.AddFileRequest:output_type -> google.protobuf.Empty
	5, // 8: peernode.MarketService.GetFileRequests:output_type -> peernode.FileRequestList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_peernode_proto_init() }
func file_peernode_proto_init() {
	if File_peernode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peernode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_peernode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_peernode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileLink); i {
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
		file_peernode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHash); i {
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
		file_peernode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_peernode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequestList); i {
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
			RawDescriptor: file_peernode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_peernode_proto_goTypes,
		DependencyIndexes: file_peernode_proto_depIdxs,
		MessageInfos:      file_peernode_proto_msgTypes,
	}.Build()
	File_peernode_proto = out.File
	file_peernode_proto_rawDesc = nil
	file_peernode_proto_goTypes = nil
	file_peernode_proto_depIdxs = nil
}
