// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_manage_tunnel.proto

package opts

import (
	fmt "fmt"
	args "github.com/batchcorp/plumber-schemas/build/go/protos/args"
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

type GetTunnelOptions struct {
	// @gotags: kong:"help='ID of the tunnel to get (leave empty to get all)'"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the tunnel to get (leave empty to get all)'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTunnelOptions) Reset()         { *m = GetTunnelOptions{} }
func (m *GetTunnelOptions) String() string { return proto.CompactTextString(m) }
func (*GetTunnelOptions) ProtoMessage()    {}
func (*GetTunnelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d5e823fc1e6249, []int{0}
}

func (m *GetTunnelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTunnelOptions.Unmarshal(m, b)
}
func (m *GetTunnelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTunnelOptions.Marshal(b, m, deterministic)
}
func (m *GetTunnelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTunnelOptions.Merge(m, src)
}
func (m *GetTunnelOptions) XXX_Size() int {
	return xxx_messageInfo_GetTunnelOptions.Size(m)
}
func (m *GetTunnelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTunnelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GetTunnelOptions proto.InternalMessageInfo

func (m *GetTunnelOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTunnelOptions struct {
	// @gotags: kong:"help='Connection ID for the tunnel to use',required=true"
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" kong:"help='Connection ID for the tunnel to use',required=true"`
	// @gotags: kong:"help='Batch API token (create in settings -> security)',required=true"
	TunnelToken string `protobuf:"bytes,3,opt,name=tunnel_token,json=tunnelToken,proto3" json:"tunnel_token,omitempty" kong:"help='Batch API token (create in settings -> security)',required=true"`
	// @gotags: kong:"help='Name for the tunnel (auto-generated if left empty)'"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" kong:"help='Name for the tunnel (auto-generated if left empty)'"`
	// @gotags: kong:"help='Notes associated with the tunnel'"
	Notes string `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty" kong:"help='Notes associated with the tunnel'"`
	// @gotags: kong:"help='Tunnel API address',default='dproxy.batch.sh:443'"
	XTunnelAddress string `protobuf:"bytes,6,opt,name=_tunnel_address,json=TunnelAddress,proto3" json:"_tunnel_address,omitempty" kong:"help='Tunnel API address',default='dproxy.batch.sh:443'"`
	// @gotags: kong:"help='Tunnel API initial connection timeout',default=5"
	XTunnelTimeoutSeconds uint32 `protobuf:"varint,7,opt,name=_tunnel_timeout_seconds,json=TunnelTimeoutSeconds,proto3" json:"_tunnel_timeout_seconds,omitempty" kong:"help='Tunnel API initial connection timeout',default=5"`
	// @gotags: kong:"help='Use gRPC insecure mode when talking to Batch Tunnel API'"
	XTunnelInsecure bool `protobuf:"varint,8,opt,name=_tunnel_insecure,json=TunnelInsecure,proto3" json:"_tunnel_insecure,omitempty" kong:"help='Use gRPC insecure mode when talking to Batch Tunnel API'"`
	// @gotags: kong:"cmd,help='Apache Kafka'"
	Kafka *args.KafkaWriteArgs `protobuf:"bytes,100,opt,name=kafka,proto3" json:"kafka,omitempty" kong:"cmd,help='Apache Kafka'"`
	// @gotags: kong:"cmd,help='Apache ActiveMQ'"
	Activemq *args.ActiveMQWriteArgs `protobuf:"bytes,101,opt,name=activemq,proto3" json:"activemq,omitempty" kong:"cmd,help='Apache ActiveMQ'"`
	// @gotags: kong:"cmd,help='AWS Simple Queue System'"
	AwsSqs *args.AWSSQSWriteArgs `protobuf:"bytes,102,opt,name=aws_sqs,json=awsSqs,proto3" json:"aws_sqs,omitempty" kong:"cmd,help='AWS Simple Queue System'"`
	// @gotags: kong:"cmd,help='AWS Simple Notification System'"
	AwsSns *args.AWSSNSWriteArgs `protobuf:"bytes,103,opt,name=aws_sns,json=awsSns,proto3" json:"aws_sns,omitempty" kong:"cmd,help='AWS Simple Notification System'"`
	// @gotags: kong:"cmd,help='NATS'"
	Nats *args.NatsWriteArgs `protobuf:"bytes,104,opt,name=nats,proto3" json:"nats,omitempty" kong:"cmd,help='NATS'"`
	// @gotags: kong:"cmd,help='NATS Streaming'"
	NatsStreaming *args.NatsStreamingWriteArgs `protobuf:"bytes,105,opt,name=nats_streaming,json=natsStreaming,proto3" json:"nats_streaming,omitempty" kong:"cmd,help='NATS Streaming'"`
	// @gotags: kong:"cmd,help='NSQ'"
	Nsq *args.NSQWriteArgs `protobuf:"bytes,106,opt,name=nsq,proto3" json:"nsq,omitempty" kong:"cmd,help='NSQ'"`
	// @gotags: kong:"cmd,help='RabbitMQ'"
	Rabbit *args.RabbitWriteArgs `protobuf:"bytes,107,opt,name=rabbit,proto3" json:"rabbit,omitempty" kong:"cmd,help='RabbitMQ'"`
	// @gotags: kong:"cmd,help='MQTT'"
	Mqtt *args.MQTTWriteArgs `protobuf:"bytes,108,opt,name=mqtt,proto3" json:"mqtt,omitempty" kong:"cmd,help='MQTT'"`
	// @gotags: kong:"cmd,help='Azure Service Bus'"
	AzureServiceBus *args.AzureServiceBusWriteArgs `protobuf:"bytes,109,opt,name=azure_service_bus,json=azureServiceBus,proto3" json:"azure_service_bus,omitempty" kong:"cmd,help='Azure Service Bus'"`
	// @gotags: kong:"cmd,help='Azure Event Hub'"
	AzureEventHub *args.AzureEventHubWriteArgs `protobuf:"bytes,110,opt,name=azure_event_hub,json=azureEventHub,proto3" json:"azure_event_hub,omitempty" kong:"cmd,help='Azure Event Hub'"`
	// @gotags: kong:"cmd,help='Google Cloud Platform Pub/Sub'"
	GcpPubsub *args.GCPPubSubWriteArgs `protobuf:"bytes,111,opt,name=gcp_pubsub,json=gcpPubsub,proto3" json:"gcp_pubsub,omitempty" kong:"cmd,help='Google Cloud Platform Pub/Sub'"`
	// @gotags: kong:"cmd,help='KubeMQ Queue'"
	KubemqQueue *args.KubeMQQueueWriteArgs `protobuf:"bytes,112,opt,name=kubemq_queue,json=kubemqQueue,proto3" json:"kubemq_queue,omitempty" kong:"cmd,help='KubeMQ Queue'"`
	// @gotags: kong:"cmd,help='Redis PubSub'"
	RedisPubsub *args.RedisPubSubWriteArgs `protobuf:"bytes,113,opt,name=redis_pubsub,json=redisPubsub,proto3" json:"redis_pubsub,omitempty" kong:"cmd,help='Redis PubSub'"`
	// @gotags: kong:"cmd,help='Redis Streams'"
	RedisStreams *args.RedisStreamsWriteArgs `protobuf:"bytes,114,opt,name=redis_streams,json=redisStreams,proto3" json:"redis_streams,omitempty" kong:"cmd,help='Redis Streams'"`
	// @gotags: kong:"cmd,help='Apache Pulsar'"
	Pulsar *args.PulsarWriteArgs `protobuf:"bytes,115,opt,name=pulsar,proto3" json:"pulsar,omitempty" kong:"cmd,help='Apache Pulsar'"`
	// @gotags: kong:"cmd,help='RabbitMQ Streams'"
	RabbitStreams *args.RabbitStreamsWriteArgs `protobuf:"bytes,116,opt,name=rabbit_streams,json=rabbitStreams,proto3" json:"rabbit_streams,omitempty" kong:"cmd,help='RabbitMQ Streams'"`
	// @gotags: kong:"cmd,help='NATS JetStream'"
	NatsJetstream *args.NatsJetstreamWriteArgs `protobuf:"bytes,117,opt,name=nats_jetstream,json=natsJetstream,proto3" json:"nats_jetstream,omitempty" kong:"cmd,help='NATS JetStream'"`
	// @gotags: kong:"cmd,help='AWS Kinesis Streams'"
	AwsKinesis           *args.AWSKinesisWriteArgs `protobuf:"bytes,118,opt,name=aws_kinesis,json=awsKinesis,proto3" json:"aws_kinesis,omitempty" kong:"cmd,help='AWS Kinesis Streams'"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateTunnelOptions) Reset()         { *m = CreateTunnelOptions{} }
func (m *CreateTunnelOptions) String() string { return proto.CompactTextString(m) }
func (*CreateTunnelOptions) ProtoMessage()    {}
func (*CreateTunnelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d5e823fc1e6249, []int{1}
}

func (m *CreateTunnelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTunnelOptions.Unmarshal(m, b)
}
func (m *CreateTunnelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTunnelOptions.Marshal(b, m, deterministic)
}
func (m *CreateTunnelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTunnelOptions.Merge(m, src)
}
func (m *CreateTunnelOptions) XXX_Size() int {
	return xxx_messageInfo_CreateTunnelOptions.Size(m)
}
func (m *CreateTunnelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTunnelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTunnelOptions proto.InternalMessageInfo

func (m *CreateTunnelOptions) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *CreateTunnelOptions) GetTunnelToken() string {
	if m != nil {
		return m.TunnelToken
	}
	return ""
}

func (m *CreateTunnelOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTunnelOptions) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *CreateTunnelOptions) GetXTunnelAddress() string {
	if m != nil {
		return m.XTunnelAddress
	}
	return ""
}

func (m *CreateTunnelOptions) GetXTunnelTimeoutSeconds() uint32 {
	if m != nil {
		return m.XTunnelTimeoutSeconds
	}
	return 0
}

func (m *CreateTunnelOptions) GetXTunnelInsecure() bool {
	if m != nil {
		return m.XTunnelInsecure
	}
	return false
}

func (m *CreateTunnelOptions) GetKafka() *args.KafkaWriteArgs {
	if m != nil {
		return m.Kafka
	}
	return nil
}

func (m *CreateTunnelOptions) GetActivemq() *args.ActiveMQWriteArgs {
	if m != nil {
		return m.Activemq
	}
	return nil
}

func (m *CreateTunnelOptions) GetAwsSqs() *args.AWSSQSWriteArgs {
	if m != nil {
		return m.AwsSqs
	}
	return nil
}

func (m *CreateTunnelOptions) GetAwsSns() *args.AWSSNSWriteArgs {
	if m != nil {
		return m.AwsSns
	}
	return nil
}

func (m *CreateTunnelOptions) GetNats() *args.NatsWriteArgs {
	if m != nil {
		return m.Nats
	}
	return nil
}

func (m *CreateTunnelOptions) GetNatsStreaming() *args.NatsStreamingWriteArgs {
	if m != nil {
		return m.NatsStreaming
	}
	return nil
}

func (m *CreateTunnelOptions) GetNsq() *args.NSQWriteArgs {
	if m != nil {
		return m.Nsq
	}
	return nil
}

func (m *CreateTunnelOptions) GetRabbit() *args.RabbitWriteArgs {
	if m != nil {
		return m.Rabbit
	}
	return nil
}

func (m *CreateTunnelOptions) GetMqtt() *args.MQTTWriteArgs {
	if m != nil {
		return m.Mqtt
	}
	return nil
}

func (m *CreateTunnelOptions) GetAzureServiceBus() *args.AzureServiceBusWriteArgs {
	if m != nil {
		return m.AzureServiceBus
	}
	return nil
}

func (m *CreateTunnelOptions) GetAzureEventHub() *args.AzureEventHubWriteArgs {
	if m != nil {
		return m.AzureEventHub
	}
	return nil
}

func (m *CreateTunnelOptions) GetGcpPubsub() *args.GCPPubSubWriteArgs {
	if m != nil {
		return m.GcpPubsub
	}
	return nil
}

func (m *CreateTunnelOptions) GetKubemqQueue() *args.KubeMQQueueWriteArgs {
	if m != nil {
		return m.KubemqQueue
	}
	return nil
}

func (m *CreateTunnelOptions) GetRedisPubsub() *args.RedisPubSubWriteArgs {
	if m != nil {
		return m.RedisPubsub
	}
	return nil
}

func (m *CreateTunnelOptions) GetRedisStreams() *args.RedisStreamsWriteArgs {
	if m != nil {
		return m.RedisStreams
	}
	return nil
}

func (m *CreateTunnelOptions) GetPulsar() *args.PulsarWriteArgs {
	if m != nil {
		return m.Pulsar
	}
	return nil
}

func (m *CreateTunnelOptions) GetRabbitStreams() *args.RabbitStreamsWriteArgs {
	if m != nil {
		return m.RabbitStreams
	}
	return nil
}

func (m *CreateTunnelOptions) GetNatsJetstream() *args.NatsJetstreamWriteArgs {
	if m != nil {
		return m.NatsJetstream
	}
	return nil
}

func (m *CreateTunnelOptions) GetAwsKinesis() *args.AWSKinesisWriteArgs {
	if m != nil {
		return m.AwsKinesis
	}
	return nil
}

type DeleteTunnelOptions struct {
	// @gotags: kong:"help='ID of the tunnel to delete',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the tunnel to delete',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTunnelOptions) Reset()         { *m = DeleteTunnelOptions{} }
func (m *DeleteTunnelOptions) String() string { return proto.CompactTextString(m) }
func (*DeleteTunnelOptions) ProtoMessage()    {}
func (*DeleteTunnelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d5e823fc1e6249, []int{2}
}

func (m *DeleteTunnelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTunnelOptions.Unmarshal(m, b)
}
func (m *DeleteTunnelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTunnelOptions.Marshal(b, m, deterministic)
}
func (m *DeleteTunnelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTunnelOptions.Merge(m, src)
}
func (m *DeleteTunnelOptions) XXX_Size() int {
	return xxx_messageInfo_DeleteTunnelOptions.Size(m)
}
func (m *DeleteTunnelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTunnelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTunnelOptions proto.InternalMessageInfo

func (m *DeleteTunnelOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopTunnelOptions struct {
	// @gotags: kong:"help='ID of the tunnel to stop',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the tunnel to stop',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopTunnelOptions) Reset()         { *m = StopTunnelOptions{} }
func (m *StopTunnelOptions) String() string { return proto.CompactTextString(m) }
func (*StopTunnelOptions) ProtoMessage()    {}
func (*StopTunnelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d5e823fc1e6249, []int{3}
}

func (m *StopTunnelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopTunnelOptions.Unmarshal(m, b)
}
func (m *StopTunnelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopTunnelOptions.Marshal(b, m, deterministic)
}
func (m *StopTunnelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopTunnelOptions.Merge(m, src)
}
func (m *StopTunnelOptions) XXX_Size() int {
	return xxx_messageInfo_StopTunnelOptions.Size(m)
}
func (m *StopTunnelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StopTunnelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StopTunnelOptions proto.InternalMessageInfo

func (m *StopTunnelOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResumeTunnelOptions struct {
	// @gotags: kong:"help='ID of the tunnel to resume',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the tunnel to resume',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeTunnelOptions) Reset()         { *m = ResumeTunnelOptions{} }
func (m *ResumeTunnelOptions) String() string { return proto.CompactTextString(m) }
func (*ResumeTunnelOptions) ProtoMessage()    {}
func (*ResumeTunnelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d5e823fc1e6249, []int{4}
}

func (m *ResumeTunnelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeTunnelOptions.Unmarshal(m, b)
}
func (m *ResumeTunnelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeTunnelOptions.Marshal(b, m, deterministic)
}
func (m *ResumeTunnelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeTunnelOptions.Merge(m, src)
}
func (m *ResumeTunnelOptions) XXX_Size() int {
	return xxx_messageInfo_ResumeTunnelOptions.Size(m)
}
func (m *ResumeTunnelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeTunnelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeTunnelOptions proto.InternalMessageInfo

func (m *ResumeTunnelOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTunnelOptions)(nil), "protos.opts.GetTunnelOptions")
	proto.RegisterType((*CreateTunnelOptions)(nil), "protos.opts.CreateTunnelOptions")
	proto.RegisterType((*DeleteTunnelOptions)(nil), "protos.opts.DeleteTunnelOptions")
	proto.RegisterType((*StopTunnelOptions)(nil), "protos.opts.StopTunnelOptions")
	proto.RegisterType((*ResumeTunnelOptions)(nil), "protos.opts.ResumeTunnelOptions")
}

func init() { proto.RegisterFile("opts/ps_opts_manage_tunnel.proto", fileDescriptor_b9d5e823fc1e6249) }

var fileDescriptor_b9d5e823fc1e6249 = []byte{
	// 879 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0x5f, 0x6f, 0xdb, 0x36,
	0x10, 0x87, 0xdb, 0x26, 0x4d, 0xe9, 0x38, 0x69, 0x99, 0x62, 0x65, 0xd3, 0xad, 0x75, 0x9c, 0x65,
	0x30, 0x30, 0xcc, 0xc6, 0xfe, 0xf4, 0x61, 0x1b, 0x30, 0x20, 0x6d, 0x87, 0xac, 0x0d, 0xd2, 0xc5,
	0x96, 0x81, 0x02, 0x7b, 0x11, 0x28, 0xf9, 0xaa, 0xa8, 0xb6, 0x28, 0x99, 0x47, 0x26, 0xc0, 0x3e,
	0xde, 0x3e, 0xd9, 0x40, 0x52, 0xb2, 0x45, 0x5b, 0xc9, 0x93, 0xac, 0xdf, 0x3f, 0x91, 0xd4, 0xdd,
	0x59, 0xa4, 0x9b, 0x17, 0x0a, 0x87, 0x05, 0x86, 0xe6, 0x1a, 0x66, 0x5c, 0xf0, 0x04, 0x42, 0xa5,
	0x85, 0x80, 0xf9, 0xa0, 0x90, 0xb9, 0xca, 0x69, 0xdb, 0x5e, 0x70, 0x60, 0x04, 0x87, 0x2f, 0xb8,
	0x4c, 0xac, 0xdc, 0x5c, 0x43, 0x1e, 0xab, 0xf4, 0x1a, 0xb2, 0x85, 0x53, 0x1e, 0xbe, 0xf4, 0xc9,
	0x1b, 0x0c, 0x67, 0xa9, 0x00, 0x4c, 0xb1, 0xe4, 0x0f, 0x37, 0x78, 0x14, 0x77, 0x70, 0x8b, 0x8a,
	0xeb, 0xf9, 0xdc, 0xbf, 0x5a, 0x42, 0x08, 0xd7, 0x20, 0x54, 0x78, 0xa5, 0xa3, 0x52, 0xf3, 0x6d,
	0x83, 0x06, 0x41, 0x5e, 0xa7, 0x31, 0x84, 0x91, 0xae, 0x92, 0xbe, 0xf1, 0x54, 0x49, 0x5c, 0x84,
	0x85, 0x8e, 0x70, 0x19, 0xc2, 0x3c, 0x7a, 0xc6, 0x3f, 0xcf, 0x78, 0xc9, 0xbc, 0xf2, 0x19, 0x1d,
	0x41, 0xb6, 0x08, 0x17, 0x1a, 0x34, 0x94, 0x82, 0x67, 0x9e, 0x20, 0x5b, 0x28, 0xd5, 0x48, 0x08,
	0xae, 0xaa, 0xb5, 0x1c, 0x6d, 0x10, 0xe1, 0x17, 0x50, 0xa8, 0x24, 0xf0, 0xec, 0x76, 0x89, 0xe3,
	0x53, 0x91, 0x94, 0x92, 0xaf, 0x7c, 0x09, 0x56, 0xef, 0xe2, 0xb9, 0x87, 0x17, 0x7a, 0x8e, 0x5c,
	0x36, 0x52, 0x92, 0x47, 0x51, 0xaa, 0x1a, 0x1f, 0xe8, 0xa8, 0xf2, 0x91, 0xd8, 0x78, 0x12, 0x12,
	0xa6, 0x29, 0xfa, 0x87, 0xd8, 0x6d, 0x10, 0x78, 0x11, 0xbd, 0x1e, 0x79, 0x7c, 0x06, 0x6a, 0x62,
	0x8b, 0xec, 0xef, 0x42, 0xa5, 0xb9, 0x40, 0xba, 0x47, 0xee, 0xa5, 0x53, 0xd6, 0xea, 0xb6, 0xfa,
	0x8f, 0xc6, 0xf7, 0xd2, 0x69, 0xef, 0xbf, 0x36, 0x39, 0x78, 0x2b, 0x81, 0x2b, 0xf0, 0x75, 0xc7,
	0xa4, 0x13, 0xe7, 0x42, 0x40, 0x6c, 0x6e, 0xc3, 0xa5, 0x65, 0x77, 0x05, 0xbe, 0x9f, 0xd2, 0x23,
	0xb2, 0xeb, 0x4a, 0x38, 0x54, 0xf9, 0x0c, 0x04, 0xbb, 0x6f, 0x35, 0x6d, 0x87, 0x4d, 0x0c, 0x44,
	0x29, 0x79, 0x20, 0x78, 0x06, 0xec, 0x81, 0xa5, 0xec, 0x6f, 0xfa, 0x94, 0x6c, 0x89, 0x5c, 0x01,
	0xb2, 0x2d, 0x0b, 0xba, 0x1b, 0xfa, 0x1d, 0xd9, 0x2f, 0x1b, 0x22, 0xe4, 0xd3, 0xa9, 0x04, 0x44,
	0xb6, 0x6d, 0xf9, 0x8e, 0x5b, 0xd9, 0xa9, 0x03, 0xe9, 0x6b, 0xf2, 0xac, 0xd2, 0xa9, 0x34, 0x83,
	0x5c, 0xab, 0x10, 0x21, 0xce, 0xc5, 0x14, 0xd9, 0xc3, 0x6e, 0xab, 0xdf, 0x19, 0x3f, 0x75, 0xfa,
	0x89, 0x23, 0x03, 0xc7, 0xd1, 0x3e, 0x79, 0x5c, 0xd9, 0x52, 0x81, 0x10, 0x6b, 0x09, 0x6c, 0xa7,
	0xdb, 0xea, 0xef, 0x8c, 0xf7, 0x9c, 0xfe, 0x7d, 0x89, 0xd2, 0x1f, 0xc9, 0x96, 0x2d, 0x49, 0x36,
	0xed, 0xb6, 0xfa, 0xed, 0x9f, 0x5e, 0x0c, 0xca, 0xc6, 0x34, 0xe7, 0x3c, 0x38, 0x37, 0xcc, 0x27,
	0x99, 0x2a, 0x38, 0x95, 0x09, 0x8e, 0x9d, 0x92, 0xfe, 0x46, 0x76, 0xaa, 0x1e, 0x65, 0x60, 0x5d,
	0x2f, 0x3d, 0xd7, 0xa9, 0x25, 0x2f, 0x46, 0x2b, 0xe3, 0x52, 0x4f, 0x5f, 0x93, 0x87, 0x65, 0x1b,
	0xb2, 0xcf, 0xd6, 0xfa, 0xb5, 0x6f, 0xfd, 0x14, 0x04, 0xa3, 0x60, 0x65, 0xdc, 0xe6, 0x37, 0x18,
	0x2c, 0x70, 0x69, 0x13, 0xc8, 0x92, 0x5b, 0x6c, 0x1f, 0xd7, 0x6d, 0x02, 0xe9, 0xc0, 0xbc, 0x0f,
	0x85, 0xec, 0xca, 0x7a, 0x0e, 0x3d, 0xcf, 0x47, 0xae, 0x70, 0xe5, 0xb0, 0x3a, 0xfa, 0x81, 0xec,
	0xf9, 0xfd, 0xc0, 0x52, 0xeb, 0x3c, 0xde, 0x70, 0x06, 0x95, 0x62, 0x15, 0xd1, 0x11, 0x75, 0x9c,
	0x7e, 0x4f, 0xee, 0x0b, 0x5c, 0xb0, 0x2f, 0x36, 0xe0, 0xb9, 0x1f, 0x10, 0xd4, 0xce, 0xc6, 0xa8,
	0xe8, 0x2f, 0x64, 0xdb, 0xf5, 0x05, 0x9b, 0x35, 0x6c, 0x6f, 0x6c, 0xa9, 0xda, 0xf6, 0x9c, 0xd6,
	0x6c, 0xcf, 0xcc, 0x04, 0x36, 0x6f, 0xd8, 0xde, 0xc5, 0x68, 0x32, 0xa9, 0x6d, 0xcf, 0xe8, 0xe8,
	0x88, 0x3c, 0xd9, 0x98, 0x61, 0x2c, 0xb3, 0xe6, 0x13, 0xff, 0x3c, 0x8d, 0x2a, 0x70, 0xa2, 0x37,
	0xba, 0x76, 0x4c, 0xfb, 0xdc, 0x67, 0xe8, 0x39, 0xd9, 0x5f, 0x1b, 0x9d, 0x4c, 0x34, 0x1c, 0x99,
	0x0d, 0xfc, 0xd3, 0x48, 0xfe, 0xd2, 0x51, 0xed, 0xc8, 0x78, 0x1d, 0xa7, 0x7f, 0x10, 0xb2, 0x9a,
	0x9e, 0x2c, 0xb7, 0x39, 0xaf, 0xbc, 0x9c, 0xb3, 0xb7, 0x97, 0x97, 0x3a, 0x0a, 0xea, 0x19, 0x8f,
	0x92, 0xb8, 0xb8, 0xb4, 0x0e, 0xfa, 0x8e, 0xec, 0xd6, 0x87, 0x28, 0x2b, 0x6c, 0xc2, 0x91, 0x5f,
	0xd2, 0x3a, 0x82, 0x8b, 0xd1, 0xc8, 0xf0, 0xab, 0x8c, 0xb6, 0xb3, 0x59, 0xd4, 0xa4, 0xd4, 0x07,
	0x10, 0x5b, 0x34, 0xa4, 0x8c, 0x8d, 0x60, 0x7d, 0x25, 0x6d, 0x59, 0xa2, 0x66, 0x2d, 0x67, 0xa4,
	0xe3, 0x4d, 0x29, 0x26, 0x6d, 0x4c, 0x6f, 0x33, 0xc6, 0x95, 0x4c, 0xed, 0x90, 0xdd, 0xe3, 0x4b,
	0xd8, 0x94, 0x86, 0x1b, 0xb4, 0x0c, 0x1b, 0x4a, 0xe3, 0xd2, 0x52, 0xb5, 0xd2, 0x70, 0x5a, 0x53,
	0xc9, 0xfe, 0xa0, 0x65, 0xaa, 0xe1, 0xb5, 0xb8, 0xc2, 0xda, 0x58, 0x40, 0x47, 0xd6, 0xf1, 0x65,
	0x57, 0x2c, 0xff, 0x48, 0x98, 0xbe, 0xa5, 0x2b, 0x3e, 0x54, 0x8a, 0xb5, 0xae, 0x58, 0xe2, 0xf4,
	0x94, 0xb4, 0x6b, 0x7f, 0xe1, 0xec, 0xda, 0x06, 0x75, 0xd7, 0x9b, 0xf9, 0xdc, 0xd1, 0xab, 0x14,
	0xc2, 0x6f, 0xb0, 0x04, 0x7b, 0x27, 0xe4, 0xe0, 0x1d, 0xcc, 0x61, 0x7d, 0x86, 0xaf, 0xcf, 0xfa,
	0x63, 0xf2, 0x24, 0x50, 0x79, 0x71, 0xb7, 0xe8, 0x84, 0x1c, 0x8c, 0x01, 0x75, 0x76, 0x77, 0xd6,
	0x9b, 0xdf, 0xff, 0xf9, 0x35, 0x49, 0x95, 0xf9, 0x2e, 0x88, 0xf3, 0x6c, 0x18, 0x71, 0x15, 0x5f,
	0xc5, 0xb9, 0x2c, 0x86, 0xc5, 0x5c, 0x67, 0x11, 0xc8, 0x1f, 0x30, 0xbe, 0x82, 0x8c, 0xe3, 0x30,
	0xd2, 0xe9, 0x7c, 0x3a, 0x4c, 0xf2, 0xa1, 0xdb, 0xcf, 0xd0, 0x7c, 0xdd, 0x44, 0xdb, 0xf6, 0xe6,
	0xe7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xec, 0x7c, 0x0d, 0x27, 0x15, 0x09, 0x00, 0x00,
}