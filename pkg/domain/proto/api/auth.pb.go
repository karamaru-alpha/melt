// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/auth.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	model "github.com/karamaru-alpha/melt/pkg/domain/proto/model"
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

type AuthSignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthSignupRequest) Reset() {
	*x = AuthSignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSignupRequest) ProtoMessage() {}

func (x *AuthSignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSignupRequest.ProtoReflect.Descriptor instead.
func (*AuthSignupRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthSignupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthSignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AuthSignupResponse) Reset() {
	*x = AuthSignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSignupResponse) ProtoMessage() {}

func (x *AuthSignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSignupResponse.ProtoReflect.Descriptor instead.
func (*AuthSignupResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthSignupResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthSignupResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AuthRefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AuthRefreshTokenRequest) Reset() {
	*x = AuthRefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshTokenRequest) ProtoMessage() {}

func (x *AuthRefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthRefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AuthRefreshTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *AuthRefreshTokenResponse) Reset() {
	*x = AuthRefreshTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshTokenResponse) ProtoMessage() {}

func (x *AuthRefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthRefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthRefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_api_auth_proto protoreflect.FileDescriptor

var file_api_auth_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a,
	0x11, 0x41, 0x75, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x5a, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a,
	0x17, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0xc0, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3b, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x72, 0x61, 0x6d, 0x61, 0x72, 0x75, 0x2d, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x6c, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_proto_rawDescOnce sync.Once
	file_api_auth_proto_rawDescData = file_api_auth_proto_rawDesc
)

func file_api_auth_proto_rawDescGZIP() []byte {
	file_api_auth_proto_rawDescOnce.Do(func() {
		file_api_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_proto_rawDescData)
	})
	return file_api_auth_proto_rawDescData
}

var file_api_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_auth_proto_goTypes = []interface{}{
	(*AuthSignupRequest)(nil),        // 0: api.AuthSignupRequest
	(*AuthSignupResponse)(nil),       // 1: api.AuthSignupResponse
	(*AuthRefreshTokenRequest)(nil),  // 2: api.AuthRefreshTokenRequest
	(*AuthRefreshTokenResponse)(nil), // 3: api.AuthRefreshTokenResponse
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
	(*model.User)(nil),               // 5: model.User
}
var file_api_auth_proto_depIdxs = []int32{
	0, // 0: api.Auth.Signup:input_type -> api.AuthSignupRequest
	2, // 1: api.Auth.RefreshToken:input_type -> api.AuthRefreshTokenRequest
	4, // 2: api.Auth.Get:input_type -> google.protobuf.Empty
	1, // 3: api.Auth.Signup:output_type -> api.AuthSignupResponse
	3, // 4: api.Auth.RefreshToken:output_type -> api.AuthRefreshTokenResponse
	5, // 5: api.Auth.Get:output_type -> model.User
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_auth_proto_init() }
func file_api_auth_proto_init() {
	if File_api_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthSignupRequest); i {
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
		file_api_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthSignupResponse); i {
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
		file_api_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRefreshTokenRequest); i {
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
		file_api_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRefreshTokenResponse); i {
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
			RawDescriptor: file_api_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_auth_proto_goTypes,
		DependencyIndexes: file_api_auth_proto_depIdxs,
		MessageInfos:      file_api_auth_proto_msgTypes,
	}.Build()
	File_api_auth_proto = out.File
	file_api_auth_proto_rawDesc = nil
	file_api_auth_proto_goTypes = nil
	file_api_auth_proto_depIdxs = nil
}
