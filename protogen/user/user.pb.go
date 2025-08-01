// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/user/type/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,proto3" json:"full_name,omitempty"`
	Whatsapp      string                 `protobuf:"bytes,4,opt,name=whatsapp,proto3" json:"whatsapp,omitempty"`
	PhotoProfile  string                 `protobuf:"bytes,5,opt,name=photo_profile,proto3" json:"photo_profile,omitempty"`
	Password      string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,7,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	Role          string                 `protobuf:"bytes,10,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_user_type_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_type_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_user_type_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetWhatsapp() string {
	if x != nil {
		return x.Whatsapp
	}
	return ""
}

func (x *User) GetPhotoProfile() string {
	if x != nil {
		return x.PhotoProfile
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type FindUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *User                  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	mi := &file_proto_user_type_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_type_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_type_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserResponse) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_user_type_user_proto protoreflect.FileDescriptor

const file_proto_user_type_user_proto_rawDesc = "" +
	"\n" +
	"\x1aproto/user/type/user.proto\x12\tuser.type\x1a%proto/google/protobuf/timestamp.proto\"\xe4\x02\n" +
	"\x04User\x12\x18\n" +
	"\auser_id\x18\x01 \x01(\tR\auser_id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1c\n" +
	"\tfull_name\x18\x03 \x01(\tR\tfull_name\x12\x1a\n" +
	"\bwhatsapp\x18\x04 \x01(\tR\bwhatsapp\x12$\n" +
	"\rphoto_profile\x18\x05 \x01(\tR\rphoto_profile\x12\x1a\n" +
	"\bpassword\x18\x06 \x01(\tR\bpassword\x12$\n" +
	"\rrefresh_token\x18\a \x01(\tR\rrefresh_token\x12:\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"created_at\x12:\n" +
	"\n" +
	"updated_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updated_at\x12\x12\n" +
	"\x04role\x18\n" +
	" \x01(\tR\x04role\"7\n" +
	"\x10FindUserResponse\x12#\n" +
	"\x04data\x18\x01 \x01(\v2\x0f.user.type.UserR\x04dataB2Z0github.com/faujiahmat/cleuse-proto/protogen/userb\x06proto3"

var (
	file_proto_user_type_user_proto_rawDescOnce sync.Once
	file_proto_user_type_user_proto_rawDescData []byte
)

func file_proto_user_type_user_proto_rawDescGZIP() []byte {
	file_proto_user_type_user_proto_rawDescOnce.Do(func() {
		file_proto_user_type_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_type_user_proto_rawDesc), len(file_proto_user_type_user_proto_rawDesc)))
	})
	return file_proto_user_type_user_proto_rawDescData
}

var file_proto_user_type_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_type_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: user.type.User
	(*FindUserResponse)(nil),      // 1: user.type.FindUserResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_user_type_user_proto_depIdxs = []int32{
	2, // 0: user.type.User.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: user.type.User.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: user.type.FindUserResponse.data:type_name -> user.type.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_user_type_user_proto_init() }
func file_proto_user_type_user_proto_init() {
	if File_proto_user_type_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_type_user_proto_rawDesc), len(file_proto_user_type_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_type_user_proto_goTypes,
		DependencyIndexes: file_proto_user_type_user_proto_depIdxs,
		MessageInfos:      file_proto_user_type_user_proto_msgTypes,
	}.Build()
	File_proto_user_type_user_proto = out.File
	file_proto_user_type_user_proto_goTypes = nil
	file_proto_user_type_user_proto_depIdxs = nil
}
