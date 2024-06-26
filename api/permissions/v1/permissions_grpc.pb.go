// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: iam/permissions/v1/permissions.proto

package permissionsv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PermissionsServiceClient is the client API for PermissionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionsServiceClient interface {
	// Add добавляет новое разрешение
	Add(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete удаляет разрешение
	Delete(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type permissionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsServiceClient(cc grpc.ClientConnInterface) PermissionsServiceClient {
	return &permissionsServiceClient{cc}
}

func (c *permissionsServiceClient) Add(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.permissions.PermissionsService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) Delete(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sbercloud.cp.iam.api.v1.permissions.PermissionsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
