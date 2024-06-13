// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: iam/service_accounts/v1/access_keys.proto

package serviceaccountsv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccessKeyCredentialsServiceClient is the client API for AccessKeyCredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessKeyCredentialsServiceClient interface {
	// Add добавляет ключ доступа сервисного аккаунта
	Add(ctx context.Context, in *AddAccessKeyRequest, opts ...grpc.CallOption) (*AddAccessKeyResponse, error)
	// Get возвращает информацию о ключе доступа по id для сервисного аккаунта
	Get(ctx context.Context, in *GetAccessKeyRequest, opts ...grpc.CallOption) (*GetAccessKeyResponse, error)
	// List возвращает список ключей доступа для сервисного аккаунта
	List(ctx context.Context, in *ListAccessKeyRequest, opts ...grpc.CallOption) (*ListAccessKeyResponse, error)
	// Delete удаляет ключ доступа по id ключа для сервисного аккаунта
	Delete(ctx context.Context, in *DeleteAccessKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type accessKeyCredentialsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessKeyCredentialsServiceClient(cc grpc.ClientConnInterface) AccessKeyCredentialsServiceClient {
	return &accessKeyCredentialsServiceClient{cc}
}

func (c *accessKeyCredentialsServiceClient) Add(ctx context.Context, in *AddAccessKeyRequest, opts ...grpc.CallOption) (*AddAccessKeyResponse, error) {
	out := new(AddAccessKeyResponse)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.service_accounts.AccessKeyCredentialsService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyCredentialsServiceClient) Get(ctx context.Context, in *GetAccessKeyRequest, opts ...grpc.CallOption) (*GetAccessKeyResponse, error) {
	out := new(GetAccessKeyResponse)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.service_accounts.AccessKeyCredentialsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyCredentialsServiceClient) List(ctx context.Context, in *ListAccessKeyRequest, opts ...grpc.CallOption) (*ListAccessKeyResponse, error) {
	out := new(ListAccessKeyResponse)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.service_accounts.AccessKeyCredentialsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyCredentialsServiceClient) Delete(ctx context.Context, in *DeleteAccessKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.service_accounts.AccessKeyCredentialsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
