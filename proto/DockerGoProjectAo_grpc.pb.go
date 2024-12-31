// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: DockerGoProjectAo.proto

package dockerGoProjectProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

import (
	"github.com/qwe826344858/dockerGoProject/CommonLogic"
	grpcClient "github.com/qwe826344858/dockerGoProject/GRpcCommon"
	"fmt"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DockerGoProjectAo_GetItemInfo_FullMethodName = "/dockerGoProjectProto.DockerGoProjectAo/GetItemInfo"
)

// DockerGoProjectAoClient is the client API for DockerGoProjectAo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerGoProjectAoClient interface {
	// 获取商品信息
	GetItemInfo(ctx context.Context, in *GetItemInfoReq, opts ...grpc.CallOption) (*GetItemInfoResp, error)
}

type dockerGoProjectAoClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerGoProjectAoClient(cc grpc.ClientConnInterface) DockerGoProjectAoClient {
	return &dockerGoProjectAoClient{cc}
}

func (c *dockerGoProjectAoClient) GetItemInfo(ctx context.Context, in *GetItemInfoReq, opts ...grpc.CallOption) (*GetItemInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemInfoResp)
	err := c.cc.Invoke(ctx, DockerGoProjectAo_GetItemInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerGoProjectAoServer is the server API for DockerGoProjectAo service.
// All implementations must embed UnimplementedDockerGoProjectAoServer
// for forward compatibility.
type DockerGoProjectAoServer interface {
	// 获取商品信息
	GetItemInfo(context.Context, *GetItemInfoReq) (*GetItemInfoResp, error)
	mustEmbedUnimplementedDockerGoProjectAoServer()
}

// UnimplementedDockerGoProjectAoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDockerGoProjectAoServer struct{}

func (UnimplementedDockerGoProjectAoServer) GetItemInfo(context.Context, *GetItemInfoReq) (*GetItemInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemInfo not implemented")
}
func (UnimplementedDockerGoProjectAoServer) mustEmbedUnimplementedDockerGoProjectAoServer() {}
func (UnimplementedDockerGoProjectAoServer) testEmbeddedByValue()                           {}

// UnsafeDockerGoProjectAoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerGoProjectAoServer will
// result in compilation errors.
type UnsafeDockerGoProjectAoServer interface {
	mustEmbedUnimplementedDockerGoProjectAoServer()
}

func RegisterDockerGoProjectAoServer(s grpc.ServiceRegistrar, srv DockerGoProjectAoServer) {
	// If the following call pancis, it indicates UnimplementedDockerGoProjectAoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DockerGoProjectAo_ServiceDesc, srv)
}

func _DockerGoProjectAo_GetItemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerGoProjectAoServer).GetItemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerGoProjectAo_GetItemInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerGoProjectAoServer).GetItemInfo(ctx, req.(*GetItemInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerGoProjectAo_ServiceDesc is the grpc.ServiceDesc for DockerGoProjectAo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerGoProjectAo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dockerGoProjectProto.DockerGoProjectAo",
	HandlerType: (*DockerGoProjectAoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItemInfo",
			Handler:    _DockerGoProjectAo_GetItemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DockerGoProjectAo.proto",
}

func GetDockerGoProjectAoClient()(f *grpcClient.GRpcFactory,Client DockerGoProjectAoClient,err error){
	var ServiceName CommonLogic.ServiceName= "DockerGoProjectAo"
	f = grpcClient.NewGRpcFactory()
	// 注册客户端
	f.RegisterClient(ServiceName, func(conn *grpc.ClientConn) grpcClient.AoClient {return NewDockerGoProjectAoClient(conn)})
	client,err := f.GetClient(ServiceName)
	if err != nil {
		return
	}

	// 类型断言为具体的客户端类型
	Client, ok := client.(DockerGoProjectAoClient)
	if !ok {
		err = fmt.Errorf("client is not of type DockerProjectAoClient")
		return
	}
	return
}
