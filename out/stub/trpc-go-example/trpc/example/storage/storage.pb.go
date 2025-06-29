// This is a Protocol Buffers file used by trpc-go to define the StorageService API.
// It defines two RPC methods: Store (to store key-value pairs) and Fetch (to retrieve values).
// The messages define the request and response formats for these methods.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/storage.proto

package storage

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // The key to store (e.g., a string identifier)
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // The value to store for that key
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	mi := &file_proto_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StoreRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StoreRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type StoreResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // A simple response message, like "ok" or an error status
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	mi := &file_proto_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StoreResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FetchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // The key to retrieve the value for
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	mi := &file_proto_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{2}
}

func (x *FetchRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"` // The value associated with the key, if found
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	mi := &file_proto_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{3}
}

func (x *FetchResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_storage_proto protoreflect.FileDescriptor

const file_proto_storage_proto_rawDesc = "" +
	"\n" +
	"\x13proto/storage.proto\x12\x14trpc.example.storage\"6\n" +
	"\fStoreRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\"'\n" +
	"\rStoreResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\" \n" +
	"\fFetchRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"%\n" +
	"\rFetchResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value2\xb4\x01\n" +
	"\x0eStorageService\x12P\n" +
	"\x05Store\x12\".trpc.example.storage.StoreRequest\x1a#.trpc.example.storage.StoreResponse\x12P\n" +
	"\x05Fetch\x12\".trpc.example.storage.FetchRequest\x1a#.trpc.example.storage.FetchResponseB&Z$trpc-go-example/trpc/example/storageb\x06proto3"

var (
	file_proto_storage_proto_rawDescOnce sync.Once
	file_proto_storage_proto_rawDescData []byte
)

func file_proto_storage_proto_rawDescGZIP() []byte {
	file_proto_storage_proto_rawDescOnce.Do(func() {
		file_proto_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_storage_proto_rawDesc), len(file_proto_storage_proto_rawDesc)))
	})
	return file_proto_storage_proto_rawDescData
}

var file_proto_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_storage_proto_goTypes = []any{
	(*StoreRequest)(nil),  // 0: trpc.example.storage.StoreRequest
	(*StoreResponse)(nil), // 1: trpc.example.storage.StoreResponse
	(*FetchRequest)(nil),  // 2: trpc.example.storage.FetchRequest
	(*FetchResponse)(nil), // 3: trpc.example.storage.FetchResponse
}
var file_proto_storage_proto_depIdxs = []int32{
	0, // 0: trpc.example.storage.StorageService.Store:input_type -> trpc.example.storage.StoreRequest
	2, // 1: trpc.example.storage.StorageService.Fetch:input_type -> trpc.example.storage.FetchRequest
	1, // 2: trpc.example.storage.StorageService.Store:output_type -> trpc.example.storage.StoreResponse
	3, // 3: trpc.example.storage.StorageService.Fetch:output_type -> trpc.example.storage.FetchResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_storage_proto_init() }
func file_proto_storage_proto_init() {
	if File_proto_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_storage_proto_rawDesc), len(file_proto_storage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_storage_proto_goTypes,
		DependencyIndexes: file_proto_storage_proto_depIdxs,
		MessageInfos:      file_proto_storage_proto_msgTypes,
	}.Build()
	File_proto_storage_proto = out.File
	file_proto_storage_proto_goTypes = nil
	file_proto_storage_proto_depIdxs = nil
}
