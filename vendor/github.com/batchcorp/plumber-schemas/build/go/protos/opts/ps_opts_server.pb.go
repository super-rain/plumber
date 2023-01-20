// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_server.proto

package opts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ServerOptions struct {
	// @gotags: kong:"default=plumber1,help='Unique ID that identifies this plumber node',env='PLUMBER_SERVER_NODE_ID',required"
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" kong:"default=plumber1,help='Unique ID that identifies this plumber node',env='PLUMBER_SERVER_NODE_ID',required"`
	// @gotags: kong:"default=aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa,help='ID of the plumber cluster (has to be the same across all nodes)',env='PLUMBER_SERVER_CLUSTER_ID',required"
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" kong:"default=aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa,help='ID of the plumber cluster (has to be the same across all nodes)',env='PLUMBER_SERVER_CLUSTER_ID',required"`
	// @gotags: kong:"help='Host:port that the gRPC server will bind to',env='PLUMBER_SERVER_GRPC_LISTEN_ADDRESS',default=0.0.0.0:9090"
	GrpcListenAddress string `protobuf:"bytes,3,opt,name=grpc_listen_address,json=grpcListenAddress,proto3" json:"grpc_listen_address,omitempty" kong:"help='Host:port that the gRPC server will bind to',env='PLUMBER_SERVER_GRPC_LISTEN_ADDRESS',default=0.0.0.0:9090"`
	// @gotags: kong:"default=batchcorp,help='All gRPC requests require this auth token to be set',env='PLUMBER_SERVER_AUTH_TOKEN',required"
	AuthToken string `protobuf:"bytes,4,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty" kong:"default=batchcorp,help='All gRPC requests require this auth token to be set',env='PLUMBER_SERVER_AUTH_TOKEN',required"`
	// @gotags: kong:"help='Comma separated list of NATS server URLs (can contain user:password if using auth; only used if --enable-cluster is true)',env='PLUMBER_SERVER_NATS_URL',default='nats://localhost:4222'"
	NatsUrl []string `protobuf:"bytes,5,rep,name=nats_url,json=natsUrl,proto3" json:"nats_url,omitempty" kong:"help='Comma separated list of NATS server URLs (can contain user:password if using auth; only used if --enable-cluster is true)',env='PLUMBER_SERVER_NATS_URL',default='nats://localhost:4222'"`
	// @gotags: kong:"help='Whether to use TLS (only used if --enable-cluster is true)',env='PLUMBER_SERVER_USE_TLS'"
	UseTls bool `protobuf:"varint,500,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Whether to use TLS (only used if --enable-cluster is true)',env='PLUMBER_SERVER_USE_TLS'"`
	// @gotags: kong:"help='TLS client cert file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CERT_FILE'"
	TlsCertFile string `protobuf:"bytes,6,opt,name=tls_cert_file,json=tlsCertFile,proto3" json:"tls_cert_file,omitempty" kong:"help='TLS client cert file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CERT_FILE'"`
	// @gotags: kong:"help='TLS client key file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_KEY_FILE'"
	TlsKeyFile string `protobuf:"bytes,7,opt,name=tls_key_file,json=tlsKeyFile,proto3" json:"tls_key_file,omitempty" kong:"help='TLS client key file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_KEY_FILE'"`
	// @gotags: kong:"help='TLS CA certificate file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CA_FILE'"
	TlsCaFile string `protobuf:"bytes,8,opt,name=tls_ca_file,json=tlsCaFile,proto3" json:"tls_ca_file,omitempty" kong:"help='TLS CA certificate file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CA_FILE'"`
	// @gotags: kong:"help='Skip server cert verification (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_SKIP_VERIFY',default=false"
	TlsSkipVerify bool `protobuf:"varint,9,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Skip server cert verification (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_SKIP_VERIFY',default=false"`
	// @gotags: kong:"help='Run plumber in cluster mode (will use NATS)',env='PLUMBER_SERVER_ENABLE_CLUSTER',default=false"
	EnableCluster bool `protobuf:"varint,10,opt,name=enable_cluster,json=enableCluster,proto3" json:"enable_cluster,omitempty" kong:"help='Run plumber in cluster mode (will use NATS)',env='PLUMBER_SERVER_ENABLE_CLUSTER',default=false"`
	// @gotags: kong:"help='Location to store time-series data for counters',default='./.tsdata'"
	StatsDatabasePath string `protobuf:"bytes,14,opt,name=stats_database_path,json=statsDatabasePath,proto3" json:"stats_database_path,omitempty" kong:"help='Location to store time-series data for counters',default='./.tsdata'"`
	// @gotags: kong:"help='How often to flush time-series data (in seconds) from memory to storage',default='60'"
	StatsFlushIntervalSeconds int32 `protobuf:"varint,15,opt,name=stats_flush_interval_seconds,json=statsFlushIntervalSeconds,proto3" json:"stats_flush_interval_seconds,omitempty" kong:"help='How often to flush time-series data (in seconds) from memory to storage',default='60'"`
	// @gotags: kong:"help='What address to bind the built-in HTTP server to',default='0.0.0.0:9191'"
	HttpListenAddress string `protobuf:"bytes,16,opt,name=http_listen_address,json=httpListenAddress,proto3" json:"http_listen_address,omitempty" kong:"help='What address to bind the built-in HTTP server to',default='0.0.0.0:9191'"`
	// @gotags: kong:"help='Allow plumber to be controlled from https://console.streamdal.com',env='PLUMBER_REMOTE_CONTROL_ENABLED',default=false"
	RemoteControlEnabled bool `protobuf:"varint,17,opt,name=remote_control_enabled,json=remoteControlEnabled,proto3" json:"remote_control_enabled,omitempty" kong:"help='Allow plumber to be controlled from https://console.streamdal.com',env='PLUMBER_REMOTE_CONTROL_ENABLED',default=false"`
	// @gotags: kong:"help='Address of Streamdal Plumber remote control service',env='PLUMBER_REMOTE_CONTROL_ADDRESS',default='foreman.streamdal.com:443'"
	RemoteControlAddress string `protobuf:"bytes,18,opt,name=remote_control_address,json=remoteControlAddress,proto3" json:"remote_control_address,omitempty" kong:"help='Address of Streamdal Plumber remote control service',env='PLUMBER_REMOTE_CONTROL_ADDRESS',default='foreman.streamdal.com:443'"`
	// @gotags: kong:"help='Streamdal API token, needed to access remote control service',env='PLUMBER_REMOTE_CONTROL_API_TOKEN'"
	RemoteControlApiToken string `protobuf:"bytes,19,opt,name=remote_control_api_token,json=remoteControlApiToken,proto3" json:"remote_control_api_token,omitempty" kong:"help='Streamdal API token, needed to access remote control service',env='PLUMBER_REMOTE_CONTROL_API_TOKEN'"`
	// @gotags: kong:"help='Connect to remote control server without TLS',default=false"
	RemoteControlDisableTls bool     `protobuf:"varint,20,opt,name=remote_control_disable_tls,json=remoteControlDisableTls,proto3" json:"remote_control_disable_tls,omitempty" kong:"help='Connect to remote control server without TLS',default=false"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ServerOptions) Reset()         { *m = ServerOptions{} }
func (m *ServerOptions) String() string { return proto.CompactTextString(m) }
func (*ServerOptions) ProtoMessage()    {}
func (*ServerOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_72fba0b7ae2941aa, []int{0}
}

func (m *ServerOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerOptions.Unmarshal(m, b)
}
func (m *ServerOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerOptions.Marshal(b, m, deterministic)
}
func (m *ServerOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerOptions.Merge(m, src)
}
func (m *ServerOptions) XXX_Size() int {
	return xxx_messageInfo_ServerOptions.Size(m)
}
func (m *ServerOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ServerOptions proto.InternalMessageInfo

func (m *ServerOptions) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ServerOptions) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ServerOptions) GetGrpcListenAddress() string {
	if m != nil {
		return m.GrpcListenAddress
	}
	return ""
}

func (m *ServerOptions) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *ServerOptions) GetNatsUrl() []string {
	if m != nil {
		return m.NatsUrl
	}
	return nil
}

func (m *ServerOptions) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *ServerOptions) GetTlsCertFile() string {
	if m != nil {
		return m.TlsCertFile
	}
	return ""
}

func (m *ServerOptions) GetTlsKeyFile() string {
	if m != nil {
		return m.TlsKeyFile
	}
	return ""
}

func (m *ServerOptions) GetTlsCaFile() string {
	if m != nil {
		return m.TlsCaFile
	}
	return ""
}

func (m *ServerOptions) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

func (m *ServerOptions) GetEnableCluster() bool {
	if m != nil {
		return m.EnableCluster
	}
	return false
}

func (m *ServerOptions) GetStatsDatabasePath() string {
	if m != nil {
		return m.StatsDatabasePath
	}
	return ""
}

func (m *ServerOptions) GetStatsFlushIntervalSeconds() int32 {
	if m != nil {
		return m.StatsFlushIntervalSeconds
	}
	return 0
}

func (m *ServerOptions) GetHttpListenAddress() string {
	if m != nil {
		return m.HttpListenAddress
	}
	return ""
}

func (m *ServerOptions) GetRemoteControlEnabled() bool {
	if m != nil {
		return m.RemoteControlEnabled
	}
	return false
}

func (m *ServerOptions) GetRemoteControlAddress() string {
	if m != nil {
		return m.RemoteControlAddress
	}
	return ""
}

func (m *ServerOptions) GetRemoteControlApiToken() string {
	if m != nil {
		return m.RemoteControlApiToken
	}
	return ""
}

func (m *ServerOptions) GetRemoteControlDisableTls() bool {
	if m != nil {
		return m.RemoteControlDisableTls
	}
	return false
}

func init() {
	proto.RegisterType((*ServerOptions)(nil), "protos.opts.ServerOptions")
}

func init() { proto.RegisterFile("opts/ps_opts_server.proto", fileDescriptor_72fba0b7ae2941aa) }

var fileDescriptor_72fba0b7ae2941aa = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0x95, 0xb7, 0x6d, 0xfe, 0x38, 0x6f, 0xd3, 0x76, 0x5b, 0xa8, 0x83, 0x00, 0x45, 0x95,
	0x40, 0xb9, 0x90, 0x3d, 0x80, 0x84, 0x50, 0x0f, 0x08, 0x52, 0x2a, 0x05, 0x90, 0x40, 0x49, 0xe0,
	0xc0, 0xc5, 0xf2, 0xee, 0x4e, 0xb2, 0x56, 0x9c, 0xb5, 0xe5, 0x99, 0x8d, 0x94, 0x8f, 0xc0, 0xf7,
	0xe5, 0x03, 0x20, 0xdb, 0xe9, 0x21, 0x11, 0x27, 0xcb, 0xf3, 0x7b, 0x9e, 0x19, 0x7b, 0x3c, 0x66,
	0x7d, 0x63, 0x09, 0x53, 0x8b, 0xc2, 0xaf, 0x02, 0xc1, 0x6d, 0xc0, 0x8d, 0xac, 0x33, 0x64, 0x92,
	0x6e, 0x58, 0x70, 0xe4, 0xc9, 0xcd, 0xef, 0x26, 0x3b, 0x9d, 0x05, 0xfa, 0xcd, 0x92, 0x32, 0x15,
	0x26, 0xd7, 0xac, 0x55, 0x99, 0x02, 0x84, 0x2a, 0x78, 0x63, 0xd0, 0x18, 0x76, 0xa6, 0x4d, 0xbf,
	0x9d, 0x14, 0xc9, 0x33, 0xc6, 0x72, 0x5d, 0x23, 0x81, 0xf3, 0xec, 0xbf, 0xc0, 0x3a, 0xbb, 0xc8,
	0xa4, 0x48, 0x46, 0xec, 0x72, 0xe9, 0x6c, 0x2e, 0xb4, 0x42, 0x82, 0x4a, 0xc8, 0xa2, 0x70, 0x80,
	0xc8, 0x8f, 0x82, 0xee, 0xc2, 0xa3, 0xaf, 0x81, 0x7c, 0x88, 0xc0, 0xa7, 0x93, 0x35, 0x95, 0x82,
	0xcc, 0x0a, 0x2a, 0x7e, 0x1c, 0xd3, 0xf9, 0xc8, 0xdc, 0x07, 0x92, 0x3e, 0x6b, 0x57, 0x92, 0x50,
	0xd4, 0x4e, 0xf3, 0x93, 0xc1, 0xd1, 0xb0, 0x33, 0x6d, 0xf9, 0xfd, 0x0f, 0xa7, 0x13, 0xce, 0x5a,
	0x35, 0x82, 0x20, 0x8d, 0xfc, 0x8f, 0x4f, 0xdf, 0x9e, 0x36, 0x6b, 0x84, 0xb9, 0xc6, 0xe4, 0x86,
	0x9d, 0x92, 0x46, 0x91, 0x83, 0x23, 0xb1, 0x50, 0x1a, 0x78, 0x33, 0xa4, 0xed, 0x92, 0xc6, 0x31,
	0x38, 0xba, 0x57, 0x1a, 0x92, 0x01, 0xfb, 0xdf, 0x6b, 0x56, 0xb0, 0x8d, 0x92, 0x56, 0x90, 0x30,
	0xd2, 0xf8, 0x05, 0xb6, 0x41, 0xf1, 0x9c, 0x75, 0x43, 0x16, 0x19, 0x05, 0xed, 0x78, 0x34, 0x9f,
	0x43, 0x06, 0xfe, 0x92, 0x9d, 0x79, 0x8e, 0x2b, 0x65, 0xc5, 0x06, 0x9c, 0x5a, 0x6c, 0x79, 0x27,
	0x1c, 0xc3, 0x17, 0x9f, 0xad, 0x94, 0xfd, 0x19, 0x82, 0xc9, 0x0b, 0xd6, 0x83, 0x4a, 0x66, 0x1a,
	0xc4, 0xae, 0x4b, 0x9c, 0x45, 0x59, 0x8c, 0x8e, 0x63, 0xd0, 0x37, 0x0e, 0xc9, 0x5f, 0xb5, 0x90,
	0x24, 0x33, 0x89, 0x20, 0xac, 0xa4, 0x92, 0xf7, 0x62, 0xe3, 0x02, 0xba, 0xdb, 0x91, 0xef, 0x92,
	0xca, 0xe4, 0x3d, 0x7b, 0x1a, 0xf5, 0x0b, 0x5d, 0x63, 0x29, 0x54, 0x45, 0xe0, 0x36, 0x52, 0x0b,
	0x84, 0xdc, 0x54, 0x05, 0xf2, 0xb3, 0x41, 0x63, 0x78, 0x32, 0xed, 0x07, 0xcd, 0xbd, 0x97, 0x4c,
	0x76, 0x8a, 0x59, 0x14, 0xf8, 0x82, 0x25, 0x91, 0x3d, 0x7c, 0xa9, 0xf3, 0x58, 0xd0, 0xa3, 0xfd,
	0x97, 0x7a, 0xc3, 0x1e, 0x3b, 0x58, 0x1b, 0x02, 0x91, 0x9b, 0x8a, 0x9c, 0xd1, 0x22, 0x5e, 0xa0,
	0xe0, 0x17, 0xe1, 0x3e, 0x57, 0x91, 0x8e, 0x23, 0xfc, 0x14, 0xd9, 0x3f, 0x5c, 0x0f, 0x85, 0x92,
	0x50, 0x68, 0xdf, 0xf5, 0x50, 0xeb, 0x2d, 0xe3, 0x87, 0x2e, 0xab, 0x76, 0x33, 0x72, 0x19, 0x7c,
	0x8f, 0xf6, 0x7d, 0x56, 0xc5, 0x79, 0xb9, 0x65, 0x4f, 0x0e, 0x8c, 0x85, 0xc2, 0xd0, 0x7c, 0x3f,
	0x27, 0x57, 0xe1, 0xa0, 0xd7, 0x7b, 0xd6, 0xbb, 0xc8, 0xe7, 0x1a, 0x3f, 0x1f, 0xb7, 0xbb, 0xe7,
	0xbd, 0x8f, 0xb7, 0xbf, 0xde, 0x2d, 0x15, 0x95, 0x75, 0x36, 0xca, 0xcd, 0x3a, 0xcd, 0x24, 0xe5,
	0x65, 0x6e, 0x9c, 0x4d, 0xad, 0xae, 0xd7, 0x19, 0xb8, 0x57, 0x98, 0x97, 0xb0, 0x96, 0x98, 0x66,
	0xb5, 0xd2, 0x45, 0xba, 0x34, 0x69, 0xfc, 0x48, 0xa9, 0xff, 0x48, 0x59, 0x33, 0x6c, 0x5e, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x42, 0xc4, 0xc9, 0x79, 0x03, 0x00, 0x00,
}
