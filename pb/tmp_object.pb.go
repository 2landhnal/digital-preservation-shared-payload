// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc1
// source: tmp_object.proto

package pb

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

type UpdateTmpObjectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ObjectId      string                 `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTmpObjectRequest) Reset() {
	*x = UpdateTmpObjectRequest{}
	mi := &file_tmp_object_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTmpObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTmpObjectRequest) ProtoMessage() {}

func (x *UpdateTmpObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tmp_object_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTmpObjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateTmpObjectRequest) Descriptor() ([]byte, []int) {
	return file_tmp_object_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTmpObjectRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *UpdateTmpObjectRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateTmpObjectResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTmpObjectResponse) Reset() {
	*x = UpdateTmpObjectResponse{}
	mi := &file_tmp_object_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTmpObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTmpObjectResponse) ProtoMessage() {}

func (x *UpdateTmpObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tmp_object_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTmpObjectResponse.ProtoReflect.Descriptor instead.
func (*UpdateTmpObjectResponse) Descriptor() ([]byte, []int) {
	return file_tmp_object_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTmpObjectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateTmpObjectRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ObjectId         string                 `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	UploaderEmail    string                 `protobuf:"bytes,2,opt,name=uploaderEmail,proto3" json:"uploaderEmail,omitempty"`
	FolderId         int64                  `protobuf:"varint,3,opt,name=folderId,proto3" json:"folderId,omitempty"`
	OriginalFilename string                 `protobuf:"bytes,4,opt,name=originalFilename,proto3" json:"originalFilename,omitempty"`
	StorageFilename  string                 `protobuf:"bytes,5,opt,name=storageFilename,proto3" json:"storageFilename,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateTmpObjectRequest) Reset() {
	*x = CreateTmpObjectRequest{}
	mi := &file_tmp_object_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTmpObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTmpObjectRequest) ProtoMessage() {}

func (x *CreateTmpObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tmp_object_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTmpObjectRequest.ProtoReflect.Descriptor instead.
func (*CreateTmpObjectRequest) Descriptor() ([]byte, []int) {
	return file_tmp_object_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTmpObjectRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *CreateTmpObjectRequest) GetUploaderEmail() string {
	if x != nil {
		return x.UploaderEmail
	}
	return ""
}

func (x *CreateTmpObjectRequest) GetFolderId() int64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *CreateTmpObjectRequest) GetOriginalFilename() string {
	if x != nil {
		return x.OriginalFilename
	}
	return ""
}

func (x *CreateTmpObjectRequest) GetStorageFilename() string {
	if x != nil {
		return x.StorageFilename
	}
	return ""
}

type CreateTmpObjectResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ObjectId         string                 `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	UploaderEmail    string                 `protobuf:"bytes,2,opt,name=uploaderEmail,proto3" json:"uploaderEmail,omitempty"`
	FolderId         int64                  `protobuf:"varint,3,opt,name=folderId,proto3" json:"folderId,omitempty"`
	OriginalFilename string                 `protobuf:"bytes,4,opt,name=originalFilename,proto3" json:"originalFilename,omitempty"`
	StorageFilename  string                 `protobuf:"bytes,5,opt,name=storageFilename,proto3" json:"storageFilename,omitempty"`
	Status           string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Id               int64                  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateTmpObjectResponse) Reset() {
	*x = CreateTmpObjectResponse{}
	mi := &file_tmp_object_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTmpObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTmpObjectResponse) ProtoMessage() {}

func (x *CreateTmpObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tmp_object_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTmpObjectResponse.ProtoReflect.Descriptor instead.
func (*CreateTmpObjectResponse) Descriptor() ([]byte, []int) {
	return file_tmp_object_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTmpObjectResponse) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *CreateTmpObjectResponse) GetUploaderEmail() string {
	if x != nil {
		return x.UploaderEmail
	}
	return ""
}

func (x *CreateTmpObjectResponse) GetFolderId() int64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *CreateTmpObjectResponse) GetOriginalFilename() string {
	if x != nil {
		return x.OriginalFilename
	}
	return ""
}

func (x *CreateTmpObjectResponse) GetStorageFilename() string {
	if x != nil {
		return x.StorageFilename
	}
	return ""
}

func (x *CreateTmpObjectResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateTmpObjectResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateTmpObjectResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CreateTmpObjectResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_tmp_object_proto protoreflect.FileDescriptor

const file_tmp_object_proto_rawDesc = "" +
	"\n" +
	"\x10tmp_object.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"L\n" +
	"\x16UpdateTmpObjectRequest\x12\x1a\n" +
	"\bobjectId\x18\x01 \x01(\tR\bobjectId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"3\n" +
	"\x17UpdateTmpObjectResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\xcc\x01\n" +
	"\x16CreateTmpObjectRequest\x12\x1a\n" +
	"\bobjectId\x18\x01 \x01(\tR\bobjectId\x12$\n" +
	"\ruploaderEmail\x18\x02 \x01(\tR\ruploaderEmail\x12\x1a\n" +
	"\bfolderId\x18\x03 \x01(\x03R\bfolderId\x12*\n" +
	"\x10originalFilename\x18\x04 \x01(\tR\x10originalFilename\x12(\n" +
	"\x0fstorageFilename\x18\x05 \x01(\tR\x0fstorageFilename\"\xe9\x02\n" +
	"\x17CreateTmpObjectResponse\x12\x1a\n" +
	"\bobjectId\x18\x01 \x01(\tR\bobjectId\x12$\n" +
	"\ruploaderEmail\x18\x02 \x01(\tR\ruploaderEmail\x12\x1a\n" +
	"\bfolderId\x18\x03 \x01(\x03R\bfolderId\x12*\n" +
	"\x10originalFilename\x18\x04 \x01(\tR\x10originalFilename\x12(\n" +
	"\x0fstorageFilename\x18\x05 \x01(\tR\x0fstorageFilename\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x128\n" +
	"\tcreatedAt\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x128\n" +
	"\tupdatedAt\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x0e\n" +
	"\x02id\x18\t \x01(\x03R\x02idB=Z;github.com/2landhnal/digital-preservation-shared-payload/pbb\x06proto3"

var (
	file_tmp_object_proto_rawDescOnce sync.Once
	file_tmp_object_proto_rawDescData []byte
)

func file_tmp_object_proto_rawDescGZIP() []byte {
	file_tmp_object_proto_rawDescOnce.Do(func() {
		file_tmp_object_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tmp_object_proto_rawDesc), len(file_tmp_object_proto_rawDesc)))
	})
	return file_tmp_object_proto_rawDescData
}

var file_tmp_object_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tmp_object_proto_goTypes = []any{
	(*UpdateTmpObjectRequest)(nil),  // 0: pb.UpdateTmpObjectRequest
	(*UpdateTmpObjectResponse)(nil), // 1: pb.UpdateTmpObjectResponse
	(*CreateTmpObjectRequest)(nil),  // 2: pb.CreateTmpObjectRequest
	(*CreateTmpObjectResponse)(nil), // 3: pb.CreateTmpObjectResponse
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_tmp_object_proto_depIdxs = []int32{
	4, // 0: pb.CreateTmpObjectResponse.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: pb.CreateTmpObjectResponse.updatedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tmp_object_proto_init() }
func file_tmp_object_proto_init() {
	if File_tmp_object_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tmp_object_proto_rawDesc), len(file_tmp_object_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tmp_object_proto_goTypes,
		DependencyIndexes: file_tmp_object_proto_depIdxs,
		MessageInfos:      file_tmp_object_proto_msgTypes,
	}.Build()
	File_tmp_object_proto = out.File
	file_tmp_object_proto_goTypes = nil
	file_tmp_object_proto_depIdxs = nil
}
