// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xc5, 0x03, 0x08, 0xae, 0x14, 0x81, 0x05, 0x62, 0x0d, 0x03, 0xc6, 0x07, 0xa0, 0x41,
	0x30, 0xf1, 0x86, 0x04, 0x2b, 0xda, 0x84, 0x84, 0x04, 0x6a, 0x99, 0x40, 0xbc, 0x39, 0xe9, 0x69,
	0xad, 0xe4, 0xc6, 0xc1, 0x76, 0x40, 0xf9, 0x0a, 0x7c, 0x6a, 0x94, 0xd8, 0x37, 0xdb, 0x71, 0xc2,
	0x53, 0xeb, 0xff, 0xff, 0xee, 0xd7, 0xf3, 0xdd, 0x35, 0x01, 0x28, 0xb8, 0xc6, 0x65, 0xad, 0xa4,
	0x91, 0xec, 0x56, 0xff, 0xa1, 0xb3, 0x79, 0x29, 0xab, 0x0a, 0x4b, 0x63, 0xe5, 0x0c, 0x14, 0xf2,
	0xad, 0xfb, 0x3e, 0xfb, 0xa3, 0xf6, 0x06, 0xe9, 0xa0, 0x50, 0xf0, 0xd6, 0x1e, 0x5e, 0xff, 0x9d,
	0xc1, 0xfc, 0xab, 0x68, 0x0e, 0x05, 0xaa, 0x0d, 0xaa, 0xdf, 0xa8, 0xd8, 0x0f, 0x78, 0x70, 0x81,
	0xe6, 0x83, 0x10, 0x2b, 0x8b, 0xdb, 0xcb, 0x4a, 0xb3, 0x13, 0x1b, 0xae, 0x97, 0x89, 0xb5, 0xc6,
	0x5f, 0x0d, 0x6a, 0x93, 0xbd, 0xf8, 0x4f, 0x84, 0xae, 0x65, 0xa5, 0x91, 0x7d, 0x86, 0xf9, 0x05,
	0x1a, 0xef, 0xb0, 0xe3, 0x20, 0xc7, 0xcb, 0x44, 0x7c, 0x3a, 0xe1, 0x3a, 0xda, 0x25, 0xdc, 0x5f,
	0x29, 0xe4, 0x06, 0x03, 0xe0, 0x73, 0x4a, 0x19, 0x3a, 0xc4, 0x3c, 0x99, 0x0e, 0x70, 0xd8, 0x2f,
	0x70, 0xef, 0x1b, 0xea, 0xb0, 0xca, 0xeb, 0x3a, 0x62, 0x9d, 0x90, 0xcf, 0xa6, 0x6c, 0x5f, 0xe7,
	0x65, 0xbd, 0x9d, 0xa8, 0x73, 0xe8, 0x24, 0x75, 0xa6, 0x01, 0x1e, 0xfb, 0x11, 0x05, 0x8e, 0x63,
	0x87, 0x4e, 0x82, 0x4d, 0x03, 0x1c, 0x76, 0x05, 0x60, 0x5b, 0xb3, 0x46, 0xbe, 0x65, 0x8b, 0xb8,
	0x5d, 0x9d, 0x46, 0xa8, 0x6c, 0xcc, 0x72, 0x90, 0x33, 0xb8, 0xb3, 0x31, 0x5c, 0x99, 0x9e, 0x71,
	0x44, 0x81, 0xd7, 0x12, 0x21, 0x16, 0x23, 0x8e, 0x25, 0xbc, 0xba, 0xc1, 0xce, 0x61, 0x66, 0x37,
	0xa9, 0xd3, 0x35, 0xcb, 0xe2, 0xf5, 0xea, 0x45, 0xe2, 0x3c, 0x19, 0xf5, 0x5c, 0x2d, 0xef, 0xe0,
	0xf6, 0xc6, 0xc8, 0xba, 0x2f, 0xe5, 0xb1, 0xff, 0x41, 0xab, 0x10, 0xe1, 0x28, 0x35, 0x7c, 0x3f,
	0xd6, 0xa8, 0x9b, 0xc3, 0xa0, 0x1f, 0x5e, 0x4b, 0xfa, 0x11, 0x5a, 0x1e, 0x62, 0x1b, 0x1e, 0x43,
	0xbc, 0x96, 0x40, 0x42, 0xcb, 0x41, 0x4e, 0xe1, 0xe6, 0xf7, 0xee, 0x5f, 0xcc, 0x1e, 0x52, 0x50,
	0x7f, 0xa4, 0xd4, 0x47, 0x03, 0xd5, 0x65, 0x9d, 0xc3, 0x8c, 0x06, 0x24, 0x78, 0xcb, 0x92, 0xa9,
	0x09, 0xde, 0x26, 0x6d, 0x8c, 0x3c, 0xcf, 0xb1, 0xab, 0x38, 0xe0, 0x04, 0x62, 0xc2, 0x89, 0x3c,
	0xcf, 0xa1, 0x06, 0x45, 0x9c, 0x40, 0x4c, 0x38, 0x91, 0xe7, 0x38, 0xef, 0xbb, 0x15, 0xeb, 0x66,
	0xd5, 0x51, 0x06, 0xe3, 0x0b, 0x18, 0x8b, 0x11, 0xc7, 0x11, 0x3e, 0xc1, 0x5d, 0xda, 0x17, 0xc1,
	0x5b, 0xcd, 0x92, 0x2d, 0xea, 0x54, 0xe2, 0x1c, 0x8f, 0x9b, 0xfe, 0x52, 0x34, 0xb0, 0xe8, 0x52,
	0x81, 0x98, 0x5c, 0x2a, 0xf2, 0x2c, 0xe7, 0xec, 0xed, 0xcf, 0xd3, 0xab, 0xbd, 0xd9, 0x35, 0xc5,
	0xb2, 0x94, 0x87, 0xbc, 0xe0, 0xa6, 0xdc, 0x95, 0x52, 0xd5, 0x79, 0x6d, 0x1f, 0xd0, 0x2f, 0x75,
	0xb9, 0xc3, 0x03, 0xd7, 0x79, 0xd1, 0xec, 0xc5, 0x36, 0xbf, 0x92, 0xb9, 0x65, 0x15, 0xf6, 0x0d,
	0xf0, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x35, 0x2d, 0x01, 0x16, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlumberServerClient is the client API for PlumberServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlumberServerClient interface {
	// List configured/known connections
	GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	// Start reading data from a connection
	CreateRead(ctx context.Context, in *CreateReadRequest, opts ...grpc.CallOption) (*CreateReadResponse, error)
	// Streams messages received off of a read
	StartRead(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (PlumberServer_StartReadClient, error)
	// List all reads that have been created
	GetAllReads(ctx context.Context, in *GetAllReadsRequest, opts ...grpc.CallOption) (*GetAllReadsResponse, error)
	// Stop reading data from a connection
	StopRead(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*StopReadResponse, error)
	// Resume reading data from an existing read
	ResumeRead(ctx context.Context, in *ResumeReadRequest, opts ...grpc.CallOption) (*ResumeReadResponse, error)
	// Resume reading data from an existing read
	DeleteRead(ctx context.Context, in *DeleteReadRequest, opts ...grpc.CallOption) (*DeleteReadResponse, error)
	// Write data to a connection
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	// Create a data relay from plumber server to the Batch platform
	CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error)
	// Update a relay (such as API token) - relay will be interrupted!
	UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error)
	ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error)
	StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error)
	GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error)
	// Delete an existing relay
	DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error)
}

type plumberServerClient struct {
	cc *grpc.ClientConn
}

func NewPlumberServerClient(cc *grpc.ClientConn) PlumberServerClient {
	return &plumberServerClient{cc}
}

func (c *plumberServerClient) GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error) {
	out := new(GetAllConnectionsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error) {
	out := new(UpdateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateRead(ctx context.Context, in *CreateReadRequest, opts ...grpc.CallOption) (*CreateReadResponse, error) {
	out := new(CreateReadResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StartRead(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (PlumberServer_StartReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlumberServer_serviceDesc.Streams[0], "/protos.PlumberServer/StartRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &plumberServerStartReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlumberServer_StartReadClient interface {
	Recv() (*StartReadResponse, error)
	grpc.ClientStream
}

type plumberServerStartReadClient struct {
	grpc.ClientStream
}

func (x *plumberServerStartReadClient) Recv() (*StartReadResponse, error) {
	m := new(StartReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plumberServerClient) GetAllReads(ctx context.Context, in *GetAllReadsRequest, opts ...grpc.CallOption) (*GetAllReadsResponse, error) {
	out := new(GetAllReadsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllReads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopRead(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*StopReadResponse, error) {
	out := new(StopReadResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeRead(ctx context.Context, in *ResumeReadRequest, opts ...grpc.CallOption) (*ResumeReadResponse, error) {
	out := new(ResumeReadResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteRead(ctx context.Context, in *DeleteReadRequest, opts ...grpc.CallOption) (*DeleteReadResponse, error) {
	out := new(DeleteReadResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error) {
	out := new(CreateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error) {
	out := new(UpdateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error) {
	out := new(ResumeRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error) {
	out := new(StopRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error) {
	out := new(GetAllRelaysResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllRelays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error) {
	out := new(DeleteRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlumberServerServer is the server API for PlumberServer service.
type PlumberServerServer interface {
	// List configured/known connections
	GetAllConnections(context.Context, *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	// Start reading data from a connection
	CreateRead(context.Context, *CreateReadRequest) (*CreateReadResponse, error)
	// Streams messages received off of a read
	StartRead(*StartReadRequest, PlumberServer_StartReadServer) error
	// List all reads that have been created
	GetAllReads(context.Context, *GetAllReadsRequest) (*GetAllReadsResponse, error)
	// Stop reading data from a connection
	StopRead(context.Context, *StopReadRequest) (*StopReadResponse, error)
	// Resume reading data from an existing read
	ResumeRead(context.Context, *ResumeReadRequest) (*ResumeReadResponse, error)
	// Resume reading data from an existing read
	DeleteRead(context.Context, *DeleteReadRequest) (*DeleteReadResponse, error)
	// Write data to a connection
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	// Create a data relay from plumber server to the Batch platform
	CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error)
	// Update a relay (such as API token) - relay will be interrupted!
	UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error)
	ResumeRelay(context.Context, *ResumeRelayRequest) (*ResumeRelayResponse, error)
	StopRelay(context.Context, *StopRelayRequest) (*StopRelayResponse, error)
	GetAllRelays(context.Context, *GetAllRelaysRequest) (*GetAllRelaysResponse, error)
	// Delete an existing relay
	DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error)
}

// UnimplementedPlumberServerServer can be embedded to have forward compatible implementations.
type UnimplementedPlumberServerServer struct {
}

func (*UnimplementedPlumberServerServer) GetAllConnections(ctx context.Context, req *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConnections not implemented")
}
func (*UnimplementedPlumberServerServer) GetConnection(ctx context.Context, req *GetConnectionRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (*UnimplementedPlumberServerServer) CreateConnection(ctx context.Context, req *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) TestConnection(ctx context.Context, req *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateConnection(ctx context.Context, req *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteConnection(ctx context.Context, req *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (*UnimplementedPlumberServerServer) CreateRead(ctx context.Context, req *CreateReadRequest) (*CreateReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRead not implemented")
}
func (*UnimplementedPlumberServerServer) StartRead(req *StartReadRequest, srv PlumberServer_StartReadServer) error {
	return status.Errorf(codes.Unimplemented, "method StartRead not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllReads(ctx context.Context, req *GetAllReadsRequest) (*GetAllReadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReads not implemented")
}
func (*UnimplementedPlumberServerServer) StopRead(ctx context.Context, req *StopReadRequest) (*StopReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRead not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeRead(ctx context.Context, req *ResumeReadRequest) (*ResumeReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeRead not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteRead(ctx context.Context, req *DeleteReadRequest) (*DeleteReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRead not implemented")
}
func (*UnimplementedPlumberServerServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedPlumberServerServer) CreateRelay(ctx context.Context, req *CreateRelayRequest) (*CreateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateRelay(ctx context.Context, req *UpdateRelayRequest) (*UpdateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeRelay(ctx context.Context, req *ResumeRelayRequest) (*ResumeRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeRelay not implemented")
}
func (*UnimplementedPlumberServerServer) StopRelay(ctx context.Context, req *StopRelayRequest) (*StopRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRelay not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllRelays(ctx context.Context, req *GetAllRelaysRequest) (*GetAllRelaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRelays not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteRelay(ctx context.Context, req *DeleteRelayRequest) (*DeleteRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelay not implemented")
}

func RegisterPlumberServerServer(s *grpc.Server, srv PlumberServerServer) {
	s.RegisterService(&_PlumberServer_serviceDesc, srv)
}

func _PlumberServer_GetAllConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllConnections(ctx, req.(*GetAllConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateRead(ctx, req.(*CreateReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StartRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlumberServerServer).StartRead(m, &plumberServerStartReadServer{stream})
}

type PlumberServer_StartReadServer interface {
	Send(*StartReadResponse) error
	grpc.ServerStream
}

type plumberServerStartReadServer struct {
	grpc.ServerStream
}

func (x *plumberServerStartReadServer) Send(m *StartReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PlumberServer_GetAllReads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllReads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllReads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllReads(ctx, req.(*GetAllReadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRead(ctx, req.(*StopReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeRead(ctx, req.(*ResumeReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteRead(ctx, req.(*DeleteReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateRelay(ctx, req.(*CreateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateRelay(ctx, req.(*UpdateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeRelay(ctx, req.(*ResumeRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRelay(ctx, req.(*StopRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetAllRelays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRelaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllRelays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllRelays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllRelays(ctx, req.(*GetAllRelaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteRelay(ctx, req.(*DeleteRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlumberServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PlumberServer",
	HandlerType: (*PlumberServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllConnections",
			Handler:    _PlumberServer_GetAllConnections_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _PlumberServer_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _PlumberServer_CreateConnection_Handler,
		},
		{
			MethodName: "TestConnection",
			Handler:    _PlumberServer_TestConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _PlumberServer_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _PlumberServer_DeleteConnection_Handler,
		},
		{
			MethodName: "CreateRead",
			Handler:    _PlumberServer_CreateRead_Handler,
		},
		{
			MethodName: "GetAllReads",
			Handler:    _PlumberServer_GetAllReads_Handler,
		},
		{
			MethodName: "StopRead",
			Handler:    _PlumberServer_StopRead_Handler,
		},
		{
			MethodName: "ResumeRead",
			Handler:    _PlumberServer_ResumeRead_Handler,
		},
		{
			MethodName: "DeleteRead",
			Handler:    _PlumberServer_DeleteRead_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _PlumberServer_Write_Handler,
		},
		{
			MethodName: "CreateRelay",
			Handler:    _PlumberServer_CreateRelay_Handler,
		},
		{
			MethodName: "UpdateRelay",
			Handler:    _PlumberServer_UpdateRelay_Handler,
		},
		{
			MethodName: "ResumeRelay",
			Handler:    _PlumberServer_ResumeRelay_Handler,
		},
		{
			MethodName: "StopRelay",
			Handler:    _PlumberServer_StopRelay_Handler,
		},
		{
			MethodName: "GetAllRelays",
			Handler:    _PlumberServer_GetAllRelays_Handler,
		},
		{
			MethodName: "DeleteRelay",
			Handler:    _PlumberServer_DeleteRelay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartRead",
			Handler:       _PlumberServer_StartRead_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "base.proto",
}