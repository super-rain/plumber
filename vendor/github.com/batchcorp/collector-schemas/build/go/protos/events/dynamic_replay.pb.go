// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dynamic_replay.proto

package events

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

type DynamicReplay_Type int32

const (
	DynamicReplay_UNSET DynamicReplay_Type = 0
	// Sent by plumber to dProxy
	DynamicReplay_AUTH_REQUEST DynamicReplay_Type = 1
	// Sent by dProxy to plumber
	DynamicReplay_AUTH_RESPONSE DynamicReplay_Type = 2
	// Sent by dProxy to plumber
	// Contains an events.Outbound message with a replay payload
	DynamicReplay_OUTBOUND_MESSAGE DynamicReplay_Type = 3
	// Sent by dProxy to plumber
	// Replicates an ISB replay event message for display in plumber logs
	DynamicReplay_REPLAY_EVENT DynamicReplay_Type = 4
)

var DynamicReplay_Type_name = map[int32]string{
	0: "UNSET",
	1: "AUTH_REQUEST",
	2: "AUTH_RESPONSE",
	3: "OUTBOUND_MESSAGE",
	4: "REPLAY_EVENT",
}

var DynamicReplay_Type_value = map[string]int32{
	"UNSET":            0,
	"AUTH_REQUEST":     1,
	"AUTH_RESPONSE":    2,
	"OUTBOUND_MESSAGE": 3,
	"REPLAY_EVENT":     4,
}

func (x DynamicReplay_Type) String() string {
	return proto.EnumName(DynamicReplay_Type_name, int32(x))
}

func (DynamicReplay_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{0, 0}
}

type ReplayEvent_Type int32

const (
	ReplayEvent_UNSET         ReplayEvent_Type = 0
	ReplayEvent_CREATE_REPLAY ReplayEvent_Type = 1
	ReplayEvent_PAUSE_REPLAY  ReplayEvent_Type = 2
	ReplayEvent_RESUME_REPLAY ReplayEvent_Type = 3
	ReplayEvent_ABORT_REPLAY  ReplayEvent_Type = 4
	ReplayEvent_FINISH_REPLAY ReplayEvent_Type = 5
)

var ReplayEvent_Type_name = map[int32]string{
	0: "UNSET",
	1: "CREATE_REPLAY",
	2: "PAUSE_REPLAY",
	3: "RESUME_REPLAY",
	4: "ABORT_REPLAY",
	5: "FINISH_REPLAY",
}

var ReplayEvent_Type_value = map[string]int32{
	"UNSET":         0,
	"CREATE_REPLAY": 1,
	"PAUSE_REPLAY":  2,
	"RESUME_REPLAY": 3,
	"ABORT_REPLAY":  4,
	"FINISH_REPLAY": 5,
}

func (x ReplayEvent_Type) String() string {
	return proto.EnumName(ReplayEvent_Type_name, int32(x))
}

func (ReplayEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{3, 0}
}

// DynamicReplay is an envelope message for dynamic replay communication from dproxy to plumber
type DynamicReplay struct {
	Type     DynamicReplay_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.DynamicReplay_Type" json:"type,omitempty"`
	ReplayId string             `protobuf:"bytes,2,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*DynamicReplay_AuthRequest
	//	*DynamicReplay_AuthResponse
	//	*DynamicReplay_OutboundMessage
	//	*DynamicReplay_ReplayMessage
	Payload              isDynamicReplay_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DynamicReplay) Reset()         { *m = DynamicReplay{} }
func (m *DynamicReplay) String() string { return proto.CompactTextString(m) }
func (*DynamicReplay) ProtoMessage()    {}
func (*DynamicReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{0}
}

func (m *DynamicReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicReplay.Unmarshal(m, b)
}
func (m *DynamicReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicReplay.Marshal(b, m, deterministic)
}
func (m *DynamicReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicReplay.Merge(m, src)
}
func (m *DynamicReplay) XXX_Size() int {
	return xxx_messageInfo_DynamicReplay.Size(m)
}
func (m *DynamicReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicReplay.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicReplay proto.InternalMessageInfo

func (m *DynamicReplay) GetType() DynamicReplay_Type {
	if m != nil {
		return m.Type
	}
	return DynamicReplay_UNSET
}

func (m *DynamicReplay) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

type isDynamicReplay_Payload interface {
	isDynamicReplay_Payload()
}

type DynamicReplay_AuthRequest struct {
	AuthRequest *AuthRequest `protobuf:"bytes,100,opt,name=auth_request,json=authRequest,proto3,oneof"`
}

type DynamicReplay_AuthResponse struct {
	AuthResponse *AuthResponse `protobuf:"bytes,101,opt,name=auth_response,json=authResponse,proto3,oneof"`
}

type DynamicReplay_OutboundMessage struct {
	OutboundMessage *Outbound `protobuf:"bytes,102,opt,name=outbound_message,json=outboundMessage,proto3,oneof"`
}

type DynamicReplay_ReplayMessage struct {
	ReplayMessage *ReplayEvent `protobuf:"bytes,103,opt,name=replay_message,json=replayMessage,proto3,oneof"`
}

func (*DynamicReplay_AuthRequest) isDynamicReplay_Payload() {}

func (*DynamicReplay_AuthResponse) isDynamicReplay_Payload() {}

func (*DynamicReplay_OutboundMessage) isDynamicReplay_Payload() {}

func (*DynamicReplay_ReplayMessage) isDynamicReplay_Payload() {}

func (m *DynamicReplay) GetPayload() isDynamicReplay_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DynamicReplay) GetAuthRequest() *AuthRequest {
	if x, ok := m.GetPayload().(*DynamicReplay_AuthRequest); ok {
		return x.AuthRequest
	}
	return nil
}

func (m *DynamicReplay) GetAuthResponse() *AuthResponse {
	if x, ok := m.GetPayload().(*DynamicReplay_AuthResponse); ok {
		return x.AuthResponse
	}
	return nil
}

func (m *DynamicReplay) GetOutboundMessage() *Outbound {
	if x, ok := m.GetPayload().(*DynamicReplay_OutboundMessage); ok {
		return x.OutboundMessage
	}
	return nil
}

func (m *DynamicReplay) GetReplayMessage() *ReplayEvent {
	if x, ok := m.GetPayload().(*DynamicReplay_ReplayMessage); ok {
		return x.ReplayMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DynamicReplay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DynamicReplay_AuthRequest)(nil),
		(*DynamicReplay_AuthResponse)(nil),
		(*DynamicReplay_OutboundMessage)(nil),
		(*DynamicReplay_ReplayMessage)(nil),
	}
}

type AuthRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	MessageBus           string   `protobuf:"bytes,2,opt,name=message_bus,json=messageBus,proto3" json:"message_bus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{1}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthRequest) GetMessageBus() string {
	if m != nil {
		return m.MessageBus
	}
	return ""
}

type AuthResponse struct {
	// Whether or not the connection is authorized
	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	// Error message if any
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReplayEvent struct {
	Type                 ReplayEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.ReplayEvent_Type" json:"type,omitempty"`
	ReplayId             string           `protobuf:"bytes,2,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplayEvent) Reset()         { *m = ReplayEvent{} }
func (m *ReplayEvent) String() string { return proto.CompactTextString(m) }
func (*ReplayEvent) ProtoMessage()    {}
func (*ReplayEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d74edecdc1efa3, []int{3}
}

func (m *ReplayEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplayEvent.Unmarshal(m, b)
}
func (m *ReplayEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplayEvent.Marshal(b, m, deterministic)
}
func (m *ReplayEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplayEvent.Merge(m, src)
}
func (m *ReplayEvent) XXX_Size() int {
	return xxx_messageInfo_ReplayEvent.Size(m)
}
func (m *ReplayEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplayEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReplayEvent proto.InternalMessageInfo

func (m *ReplayEvent) GetType() ReplayEvent_Type {
	if m != nil {
		return m.Type
	}
	return ReplayEvent_UNSET
}

func (m *ReplayEvent) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.DynamicReplay_Type", DynamicReplay_Type_name, DynamicReplay_Type_value)
	proto.RegisterEnum("events.ReplayEvent_Type", ReplayEvent_Type_name, ReplayEvent_Type_value)
	proto.RegisterType((*DynamicReplay)(nil), "events.DynamicReplay")
	proto.RegisterType((*AuthRequest)(nil), "events.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "events.AuthResponse")
	proto.RegisterType((*ReplayEvent)(nil), "events.ReplayEvent")
}

func init() { proto.RegisterFile("dynamic_replay.proto", fileDescriptor_04d74edecdc1efa3) }

var fileDescriptor_04d74edecdc1efa3 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5d, 0x8f, 0x93, 0x4e,
	0x18, 0xc5, 0x4b, 0xb7, 0xdd, 0xdd, 0x3e, 0x7d, 0xf9, 0xd3, 0xf9, 0xf7, 0x82, 0xac, 0x89, 0x36,
	0x5c, 0xf5, 0x42, 0x21, 0x59, 0x6f, 0x4c, 0x74, 0x2f, 0xa8, 0x1d, 0xa5, 0x89, 0x7d, 0x71, 0x00,
	0x13, 0xbd, 0x90, 0xf0, 0x32, 0x16, 0x62, 0xcb, 0x20, 0x03, 0x9b, 0xd4, 0xaf, 0xe6, 0x17, 0xf0,
	0x63, 0x19, 0x18, 0x58, 0xdb, 0xe8, 0x85, 0x97, 0xe7, 0x3c, 0xe7, 0x4c, 0xfb, 0xfc, 0x66, 0x80,
	0x49, 0x78, 0x4c, 0xbc, 0x43, 0x1c, 0xb8, 0x19, 0x4d, 0xf7, 0xde, 0x51, 0x4b, 0x33, 0x96, 0x33,
	0x74, 0x49, 0xef, 0x69, 0x92, 0xf3, 0x9b, 0x11, 0x2b, 0x72, 0x9f, 0x15, 0x49, 0x28, 0x7c, 0xf5,
	0xc7, 0x05, 0x0c, 0x17, 0xa2, 0x40, 0xaa, 0x3c, 0xd2, 0xa0, 0x93, 0x1f, 0x53, 0xaa, 0x48, 0x53,
	0x69, 0x36, 0xba, 0xbd, 0xd1, 0x44, 0x51, 0x3b, 0x0b, 0x69, 0xf6, 0x31, 0xa5, 0xa4, 0xca, 0xa1,
	0x47, 0xd0, 0x13, 0xbf, 0xe4, 0xc6, 0xa1, 0xd2, 0x9e, 0x4a, 0xb3, 0x1e, 0xb9, 0x16, 0xc6, 0x32,
	0x44, 0x2f, 0x60, 0xe0, 0x15, 0x79, 0xe4, 0x66, 0xf4, 0x5b, 0x41, 0x79, 0xae, 0x84, 0x53, 0x69,
	0xd6, 0xbf, 0xfd, 0xbf, 0x39, 0xd4, 0x28, 0xf2, 0x88, 0x88, 0x91, 0xd9, 0x22, 0x7d, 0xef, 0xb7,
	0x44, 0x2f, 0x61, 0x58, 0x37, 0x79, 0xca, 0x12, 0x4e, 0x15, 0x5a, 0x55, 0x27, 0xe7, 0x55, 0x31,
	0x33, 0x5b, 0x64, 0xe0, 0x9d, 0x68, 0x74, 0x07, 0x72, 0xb3, 0xa7, 0x7b, 0xa0, 0x9c, 0x7b, 0x3b,
	0xaa, 0x7c, 0xa9, 0xfa, 0x72, 0xd3, 0xdf, 0xd4, 0x73, 0xb3, 0x45, 0xfe, 0x6b, 0xb2, 0x2b, 0x11,
	0x45, 0xaf, 0x60, 0x54, 0xaf, 0xd4, 0x94, 0x77, 0xe7, 0xff, 0x5b, 0x50, 0xc0, 0xa5, 0x30, 0x5b,
	0x64, 0x28, 0xc2, 0x75, 0x5b, 0xfd, 0x0c, 0x9d, 0x12, 0x0f, 0xea, 0x41, 0xd7, 0x59, 0x5b, 0xd8,
	0x96, 0x5b, 0x48, 0x86, 0x81, 0xe1, 0xd8, 0xa6, 0x4b, 0xf0, 0x7b, 0x07, 0x5b, 0xb6, 0x2c, 0xa1,
	0x31, 0x0c, 0x6b, 0xc7, 0xda, 0x6e, 0xd6, 0x16, 0x96, 0xdb, 0x68, 0x02, 0xf2, 0xc6, 0xb1, 0xe7,
	0x1b, 0x67, 0xbd, 0x70, 0x57, 0xd8, 0xb2, 0x8c, 0xb7, 0x58, 0xbe, 0x28, 0xab, 0x04, 0x6f, 0xdf,
	0x19, 0x1f, 0x5d, 0xfc, 0x01, 0xaf, 0x6d, 0xb9, 0x33, 0xef, 0xc1, 0x55, 0xea, 0x1d, 0xf7, 0xcc,
	0x0b, 0xd5, 0x05, 0xf4, 0x4f, 0x10, 0xa2, 0x09, 0x74, 0x73, 0xf6, 0x95, 0x26, 0xd5, 0xdd, 0xf5,
	0x88, 0x10, 0xe8, 0x09, 0xf4, 0xeb, 0x35, 0x5c, 0xbf, 0xe0, 0xf5, 0x15, 0x41, 0x6d, 0xcd, 0x0b,
	0xae, 0x9a, 0x30, 0x38, 0xa5, 0x89, 0x1e, 0x03, 0x94, 0x34, 0x59, 0x16, 0x7f, 0xa7, 0x61, 0x75,
	0xd6, 0x35, 0x39, 0x71, 0x90, 0x02, 0x57, 0x0d, 0x17, 0x71, 0x58, 0x23, 0xd5, 0x9f, 0x12, 0xf4,
	0x4f, 0xd8, 0xa0, 0xa7, 0x67, 0x6f, 0x49, 0xf9, 0x0b, 0xbe, 0x7f, 0x7d, 0x49, 0x6a, 0xf2, 0x27,
	0xd5, 0x31, 0x0c, 0x5f, 0x13, 0x6c, 0xd8, 0xd8, 0x15, 0x84, 0x64, 0xa9, 0xa4, 0xb5, 0x35, 0x1c,
	0xeb, 0xc1, 0x69, 0x97, 0x21, 0x82, 0x2d, 0x67, 0xf5, 0x60, 0x55, 0x48, 0x8d, 0xf9, 0x86, 0xd8,
	0x8d, 0xd3, 0x29, 0x43, 0x6f, 0x96, 0xeb, 0xa5, 0x65, 0x36, 0x56, 0x77, 0x6e, 0xc3, 0x98, 0x47,
	0x9a, 0xef, 0xe5, 0x41, 0xa4, 0x71, 0x9a, 0xdd, 0xc7, 0x01, 0xe5, 0x5b, 0xe9, 0xd3, 0xdd, 0x2e,
	0xce, 0xa3, 0xc2, 0xd7, 0x02, 0x76, 0xd0, 0xab, 0x61, 0xc0, 0xb2, 0x54, 0x0f, 0xd8, 0x7e, 0x4f,
	0x83, 0x9c, 0x65, 0xcf, 0x78, 0x10, 0xd1, 0x83, 0xc7, 0x75, 0xbf, 0x88, 0xf7, 0xa1, 0xbe, 0x63,
	0x7a, 0xf5, 0x99, 0x71, 0x5d, 0x6c, 0xee, 0x5f, 0x56, 0xf2, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0xc7, 0x16, 0x93, 0xa5, 0x03, 0x00, 0x00,
}
