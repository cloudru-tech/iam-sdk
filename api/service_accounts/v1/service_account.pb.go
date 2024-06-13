// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: iam/service_accounts/v1/service_account.proto

package serviceaccountsv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceAccounts - список сервисных аккаунтов
type ServiceAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ServiceAccounts - список сервисных аккаунтов
	ServiceAccounts []*ServiceAccount `protobuf:"bytes,1,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty"`
}

func (x *ServiceAccounts) Reset() {
	*x = ServiceAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccounts) ProtoMessage() {}

func (x *ServiceAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccounts.ProtoReflect.Descriptor instead.
func (*ServiceAccounts) Descriptor() ([]byte, []int) {
	return file_iam_service_accounts_v1_service_account_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceAccounts) GetServiceAccounts() []*ServiceAccount {
	if x != nil {
		return x.ServiceAccounts
	}
	return nil
}

type ServiceAccountID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UserID - идентификатор сервисного аккаунта
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServiceAccountID) Reset() {
	*x = ServiceAccountID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccountID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountID) ProtoMessage() {}

func (x *ServiceAccountID) ProtoReflect() protoreflect.Message {
	mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountID.ProtoReflect.Descriptor instead.
func (*ServiceAccountID) Descriptor() ([]byte, []int) {
	return file_iam_service_accounts_v1_service_account_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceAccountID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ServiceAccount - информация о сервисном аккаунте
type ServiceAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UserID - идентификатор сервисного аккаунта
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// NamespaceID - область видимости
	NamespaceId string `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	// Name - название сервисного аккаунта.
	// Уникально для namespace_id.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Email - email сервисного аккаунта.
	// Только для чтения, виртуальный.
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Description - описание
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// UseRefreshTokens - выпускать refresh_token или нет (по умолчанию false)
	UseRefreshTokens bool `protobuf:"varint,6,opt,name=use_refresh_tokens,json=useRefreshTokens,proto3" json:"use_refresh_tokens,omitempty"`
	// Enabled - включен или выключен сервисный аккаунт
	Enabled bool `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// CreatedAt - дата создания записи
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UpdatedAt - дата обновления записи
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ServiceAccount) Reset() {
	*x = ServiceAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccount) ProtoMessage() {}

func (x *ServiceAccount) ProtoReflect() protoreflect.Message {
	mi := &file_iam_service_accounts_v1_service_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccount.ProtoReflect.Descriptor instead.
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return file_iam_service_accounts_v1_service_account_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceAccount) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *ServiceAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ServiceAccount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceAccount) GetUseRefreshTokens() bool {
	if x != nil {
		return x.UseRefreshTokens
	}
	return false
}

func (x *ServiceAccount) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ServiceAccount) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ServiceAccount) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_iam_service_accounts_v1_service_account_proto protoreflect.FileDescriptor

var file_iam_service_accounts_v1_service_account_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x28, 0x73, 0x62, 0x65, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x70, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x63, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x62, 0x65, 0x72, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x70, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcd, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x75, 0x73, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x6b, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x2e, 0x73, 0x62,
	0x65, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x73, 0x76, 0x70, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_service_accounts_v1_service_account_proto_rawDescOnce sync.Once
	file_iam_service_accounts_v1_service_account_proto_rawDescData = file_iam_service_accounts_v1_service_account_proto_rawDesc
)

func file_iam_service_accounts_v1_service_account_proto_rawDescGZIP() []byte {
	file_iam_service_accounts_v1_service_account_proto_rawDescOnce.Do(func() {
		file_iam_service_accounts_v1_service_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_service_accounts_v1_service_account_proto_rawDescData)
	})
	return file_iam_service_accounts_v1_service_account_proto_rawDescData
}

var file_iam_service_accounts_v1_service_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iam_service_accounts_v1_service_account_proto_goTypes = []interface{}{
	(*ServiceAccounts)(nil),       // 0: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccounts
	(*ServiceAccountID)(nil),      // 1: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccountID
	(*ServiceAccount)(nil),        // 2: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccount
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_iam_service_accounts_v1_service_account_proto_depIdxs = []int32{
	2, // 0: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccounts.service_accounts:type_name -> sbercloud.cp.iam.api.v1.service_accounts.ServiceAccount
	3, // 1: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccount.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: sbercloud.cp.iam.api.v1.service_accounts.ServiceAccount.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_iam_service_accounts_v1_service_account_proto_init() }
func file_iam_service_accounts_v1_service_account_proto_init() {
	if File_iam_service_accounts_v1_service_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_service_accounts_v1_service_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccounts); i {
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
		file_iam_service_accounts_v1_service_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountID); i {
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
		file_iam_service_accounts_v1_service_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccount); i {
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
			RawDescriptor: file_iam_service_accounts_v1_service_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iam_service_accounts_v1_service_account_proto_goTypes,
		DependencyIndexes: file_iam_service_accounts_v1_service_account_proto_depIdxs,
		MessageInfos:      file_iam_service_accounts_v1_service_account_proto_msgTypes,
	}.Build()
	File_iam_service_accounts_v1_service_account_proto = out.File
	file_iam_service_accounts_v1_service_account_proto_rawDesc = nil
	file_iam_service_accounts_v1_service_account_proto_goTypes = nil
	file_iam_service_accounts_v1_service_account_proto_depIdxs = nil
}
