// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: services.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Document_UploadFile_FullMethodName   = "/pb.Document/UploadFile"
	Document_DownloadFile_FullMethodName = "/pb.Document/DownloadFile"
)

// DocumentClient is the client API for Document service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// >> Document rpc...
type DocumentClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, UploadStatus], error)
	DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error)
}

type documentClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentClient(cc grpc.ClientConnInterface) DocumentClient {
	return &documentClient{cc}
}

func (c *documentClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, UploadStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Document_ServiceDesc.Streams[0], Document_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileChunk, UploadStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Document_UploadFileClient = grpc.ClientStreamingClient[FileChunk, UploadStatus]

func (c *documentClient) DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Document_ServiceDesc.Streams[1], Document_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileRequest, FileChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Document_DownloadFileClient = grpc.ServerStreamingClient[FileChunk]

// DocumentServer is the server API for Document service.
// All implementations must embed UnimplementedDocumentServer
// for forward compatibility.
//
// >> Document rpc...
type DocumentServer interface {
	UploadFile(grpc.ClientStreamingServer[FileChunk, UploadStatus]) error
	DownloadFile(*FileRequest, grpc.ServerStreamingServer[FileChunk]) error
	mustEmbedUnimplementedDocumentServer()
}

// UnimplementedDocumentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDocumentServer struct{}

func (UnimplementedDocumentServer) UploadFile(grpc.ClientStreamingServer[FileChunk, UploadStatus]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedDocumentServer) DownloadFile(*FileRequest, grpc.ServerStreamingServer[FileChunk]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedDocumentServer) mustEmbedUnimplementedDocumentServer() {}
func (UnimplementedDocumentServer) testEmbeddedByValue()                  {}

// UnsafeDocumentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServer will
// result in compilation errors.
type UnsafeDocumentServer interface {
	mustEmbedUnimplementedDocumentServer()
}

func RegisterDocumentServer(s grpc.ServiceRegistrar, srv DocumentServer) {
	// If the following call pancis, it indicates UnimplementedDocumentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Document_ServiceDesc, srv)
}

func _Document_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DocumentServer).UploadFile(&grpc.GenericServerStream[FileChunk, UploadStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Document_UploadFileServer = grpc.ClientStreamingServer[FileChunk, UploadStatus]

func _Document_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServer).DownloadFile(m, &grpc.GenericServerStream[FileRequest, FileChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Document_DownloadFileServer = grpc.ServerStreamingServer[FileChunk]

// Document_ServiceDesc is the grpc.ServiceDesc for Document service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Document_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Document",
	HandlerType: (*DocumentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Document_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _Document_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}
