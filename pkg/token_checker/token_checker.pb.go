// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.1
// source: proto/token_checker.proto

package token_checker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenSafetyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TokenMint     string                 `protobuf:"bytes,1,opt,name=tokenMint,proto3" json:"tokenMint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenSafetyRequest) Reset() {
	*x = TokenSafetyRequest{}
	mi := &file_proto_token_checker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenSafetyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenSafetyRequest) ProtoMessage() {}

func (x *TokenSafetyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_checker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenSafetyRequest.ProtoReflect.Descriptor instead.
func (*TokenSafetyRequest) Descriptor() ([]byte, []int) {
	return file_proto_token_checker_proto_rawDescGZIP(), []int{0}
}

func (x *TokenSafetyRequest) GetTokenMint() string {
	if x != nil {
		return x.TokenMint
	}
	return ""
}

type TokenSafetyResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	IsNonStandardProgram bool                   `protobuf:"varint,1,opt,name=isNonStandardProgram,proto3" json:"isNonStandardProgram,omitempty"`
	HasMintAuthority     bool                   `protobuf:"varint,2,opt,name=hasMintAuthority,proto3" json:"hasMintAuthority,omitempty"`
	HasFreezeAuthority   bool                   `protobuf:"varint,3,opt,name=hasFreezeAuthority,proto3" json:"hasFreezeAuthority,omitempty"`
	Reason               string                 `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *TokenSafetyResponse) Reset() {
	*x = TokenSafetyResponse{}
	mi := &file_proto_token_checker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenSafetyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenSafetyResponse) ProtoMessage() {}

func (x *TokenSafetyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_checker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenSafetyResponse.ProtoReflect.Descriptor instead.
func (*TokenSafetyResponse) Descriptor() ([]byte, []int) {
	return file_proto_token_checker_proto_rawDescGZIP(), []int{1}
}

func (x *TokenSafetyResponse) GetIsNonStandardProgram() bool {
	if x != nil {
		return x.IsNonStandardProgram
	}
	return false
}

func (x *TokenSafetyResponse) GetHasMintAuthority() bool {
	if x != nil {
		return x.HasMintAuthority
	}
	return false
}

func (x *TokenSafetyResponse) GetHasFreezeAuthority() bool {
	if x != nil {
		return x.HasFreezeAuthority
	}
	return false
}

func (x *TokenSafetyResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_proto_token_checker_proto protoreflect.FileDescriptor

const file_proto_token_checker_proto_rawDesc = "" +
	"\n" +
	"\x19proto/token_checker.proto\x12\ftokenchecker\"2\n" +
	"\x12TokenSafetyRequest\x12\x1c\n" +
	"\ttokenMint\x18\x01 \x01(\tR\ttokenMint\"\xbd\x01\n" +
	"\x13TokenSafetyResponse\x122\n" +
	"\x14isNonStandardProgram\x18\x01 \x01(\bR\x14isNonStandardProgram\x12*\n" +
	"\x10hasMintAuthority\x18\x02 \x01(\bR\x10hasMintAuthority\x12.\n" +
	"\x12hasFreezeAuthority\x18\x03 \x01(\bR\x12hasFreezeAuthority\x12\x16\n" +
	"\x06reason\x18\x04 \x01(\tR\x06reason2g\n" +
	"\fTokenChecker\x12W\n" +
	"\x10CheckTokenSafety\x12 .tokenchecker.TokenSafetyRequest\x1a!.tokenchecker.TokenSafetyResponseB\x15Z\x13./pkg/token_checkerb\x06proto3"

var (
	file_proto_token_checker_proto_rawDescOnce sync.Once
	file_proto_token_checker_proto_rawDescData []byte
)

func file_proto_token_checker_proto_rawDescGZIP() []byte {
	file_proto_token_checker_proto_rawDescOnce.Do(func() {
		file_proto_token_checker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_token_checker_proto_rawDesc), len(file_proto_token_checker_proto_rawDesc)))
	})
	return file_proto_token_checker_proto_rawDescData
}

var file_proto_token_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_token_checker_proto_goTypes = []any{
	(*TokenSafetyRequest)(nil),  // 0: tokenchecker.TokenSafetyRequest
	(*TokenSafetyResponse)(nil), // 1: tokenchecker.TokenSafetyResponse
}
var file_proto_token_checker_proto_depIdxs = []int32{
	0, // 0: tokenchecker.TokenChecker.CheckTokenSafety:input_type -> tokenchecker.TokenSafetyRequest
	1, // 1: tokenchecker.TokenChecker.CheckTokenSafety:output_type -> tokenchecker.TokenSafetyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_token_checker_proto_init() }
func file_proto_token_checker_proto_init() {
	if File_proto_token_checker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_token_checker_proto_rawDesc), len(file_proto_token_checker_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_token_checker_proto_goTypes,
		DependencyIndexes: file_proto_token_checker_proto_depIdxs,
		MessageInfos:      file_proto_token_checker_proto_msgTypes,
	}.Build()
	File_proto_token_checker_proto = out.File
	file_proto_token_checker_proto_goTypes = nil
	file_proto_token_checker_proto_depIdxs = nil
}
