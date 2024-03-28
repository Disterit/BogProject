// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: adder.proto

package api

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

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageRequestTelegram `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{0}
}

func (x *MessageResponse) GetMessages() []*MessageRequestTelegram {
	if x != nil {
		return x.Messages
	}
	return nil
}

type MessageRequestTelegram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageRequestTelegram) Reset() {
	*x = MessageRequestTelegram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequestTelegram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequestTelegram) ProtoMessage() {}

func (x *MessageRequestTelegram) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequestTelegram.ProtoReflect.Descriptor instead.
func (*MessageRequestTelegram) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRequestTelegram) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MessageRequestTelegram) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_adder_proto protoreflect.FileDescriptor

var file_adder_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x55, 0x0a, 0x0b, 0x54, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x69, 0x73, 0x74, 0x65, 0x72, 0x69, 0x74, 0x2f, 0x42, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adder_proto_rawDescOnce sync.Once
	file_adder_proto_rawDescData = file_adder_proto_rawDesc
)

func file_adder_proto_rawDescGZIP() []byte {
	file_adder_proto_rawDescOnce.Do(func() {
		file_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_adder_proto_rawDescData)
	})
	return file_adder_proto_rawDescData
}

var file_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_adder_proto_goTypes = []interface{}{
	(*MessageResponse)(nil),        // 0: proto.MessageResponse
	(*MessageRequestTelegram)(nil), // 1: proto.MessageRequestTelegram
}
var file_adder_proto_depIdxs = []int32{
	1, // 0: proto.MessageResponse.messages:type_name -> proto.MessageRequestTelegram
	1, // 1: proto.TelegramBot.GetMessages:input_type -> proto.MessageRequestTelegram
	0, // 2: proto.TelegramBot.GetMessages:output_type -> proto.MessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_adder_proto_init() }
func file_adder_proto_init() {
	if File_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
		file_adder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequestTelegram); i {
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
			RawDescriptor: file_adder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adder_proto_goTypes,
		DependencyIndexes: file_adder_proto_depIdxs,
		MessageInfos:      file_adder_proto_msgTypes,
	}.Build()
	File_adder_proto = out.File
	file_adder_proto_rawDesc = nil
	file_adder_proto_goTypes = nil
	file_adder_proto_depIdxs = nil
}
