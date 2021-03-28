// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

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

type Account_OnboardingState int32

const (
	Account_PENDING       Account_OnboardingState = 0
	Account_WELCOME       Account_OnboardingState = 1
	Account_DESTINATION   Account_OnboardingState = 2
	Account_INITIAL_EVENT Account_OnboardingState = 3
	Account_CLOSED        Account_OnboardingState = 4
)

var Account_OnboardingState_name = map[int32]string{
	0: "PENDING",
	1: "WELCOME",
	2: "DESTINATION",
	3: "INITIAL_EVENT",
	4: "CLOSED",
}

var Account_OnboardingState_value = map[string]int32{
	"PENDING":       0,
	"WELCOME":       1,
	"DESTINATION":   2,
	"INITIAL_EVENT": 3,
	"CLOSED":        4,
}

func (x Account_OnboardingState) String() string {
	return proto.EnumName(Account_OnboardingState_name, int32(x))
}

func (Account_OnboardingState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0, 0}
}

type Account_OnboardingStateStatus int32

const (
	Account_STATUS_UNSET Account_OnboardingStateStatus = 0
	Account_COMPLETE     Account_OnboardingStateStatus = 1
	Account_PROCESSING   Account_OnboardingStateStatus = 2
	Account_ERROR        Account_OnboardingStateStatus = 3
	Account_SKIPPED      Account_OnboardingStateStatus = 4
)

var Account_OnboardingStateStatus_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "COMPLETE",
	2: "PROCESSING",
	3: "ERROR",
	4: "SKIPPED",
}

var Account_OnboardingStateStatus_value = map[string]int32{
	"STATUS_UNSET": 0,
	"COMPLETE":     1,
	"PROCESSING":   2,
	"ERROR":        3,
	"SKIPPED":      4,
}

func (x Account_OnboardingStateStatus) String() string {
	return proto.EnumName(Account_OnboardingStateStatus_name, int32(x))
}

func (Account_OnboardingStateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0, 1}
}

type Account struct {
	Id                    string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName              string                        `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email                 string                        `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	OnboardingState       Account_OnboardingState       `protobuf:"varint,4,opt,name=onboarding_state,json=onboardingState,proto3,enum=events.Account_OnboardingState" json:"onboarding_state,omitempty"`
	OnboardingStateStatus Account_OnboardingStateStatus `protobuf:"varint,5,opt,name=onboarding_state_status,json=onboardingStateStatus,proto3,enum=events.Account_OnboardingStateStatus" json:"onboarding_state_status,omitempty"`
	BillingPlanId         string                        `protobuf:"bytes,6,opt,name=billing_plan_id,json=billingPlanId,proto3" json:"billing_plan_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetOnboardingState() Account_OnboardingState {
	if m != nil {
		return m.OnboardingState
	}
	return Account_PENDING
}

func (m *Account) GetOnboardingStateStatus() Account_OnboardingStateStatus {
	if m != nil {
		return m.OnboardingStateStatus
	}
	return Account_STATUS_UNSET
}

func (m *Account) GetBillingPlanId() string {
	if m != nil {
		return m.BillingPlanId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.Account_OnboardingState", Account_OnboardingState_name, Account_OnboardingState_value)
	proto.RegisterEnum("events.Account_OnboardingStateStatus", Account_OnboardingStateStatus_name, Account_OnboardingStateStatus_value)
	proto.RegisterType((*Account)(nil), "events.Account")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0x1b, 0xa7, 0x49, 0x9b, 0xaf, 0x4d, 0xa2, 0x89, 0x95, 0x19, 0x76, 0xb1, 0x12, 0xd8,
	0xe8, 0xc5, 0xb0, 0x61, 0x7b, 0x02, 0xcf, 0x11, 0x43, 0x9b, 0x2b, 0x1b, 0x4b, 0xdd, 0x60, 0x63,
	0x18, 0xf9, 0xcf, 0x12, 0x81, 0x6c, 0x85, 0xd8, 0xde, 0xdb, 0xee, 0x5d, 0x86, 0xec, 0x32, 0x58,
	0x09, 0xec, 0x46, 0xf0, 0x9d, 0xef, 0xf0, 0x3b, 0x47, 0x42, 0xb0, 0x94, 0x45, 0x61, 0xfa, 0xa6,
	0xf3, 0x0e, 0x47, 0xd3, 0x19, 0x3c, 0xaf, 0x7e, 0x55, 0x4d, 0xd7, 0x6e, 0x7e, 0x4f, 0xe1, 0x22,
	0x18, 0x37, 0x78, 0x05, 0x8e, 0x2a, 0xdd, 0xc9, 0xed, 0xe4, 0x6e, 0x91, 0x3a, 0xaa, 0xc4, 0x2f,
	0x61, 0xf1, 0xb3, 0xd7, 0x3a, 0x6b, 0x64, 0x5d, 0xb9, 0xce, 0x20, 0x5f, 0x5a, 0x81, 0xc9, 0xba,
	0xc2, 0xcf, 0x61, 0x56, 0xd5, 0x52, 0x69, 0x77, 0x3a, 0x2c, 0xc6, 0x01, 0x7f, 0x02, 0x64, 0x9a,
	0xdc, 0xc8, 0x63, 0xa9, 0x9a, 0x5d, 0xd6, 0x76, 0xb2, 0xab, 0xdc, 0xf3, 0xdb, 0xc9, 0xdd, 0xea,
	0xdd, 0x2b, 0x6f, 0x4c, 0xf4, 0x1e, 0xd3, 0xbc, 0xf8, 0xaf, 0x8f, 0x5b, 0x5b, 0xba, 0x36, 0xff,
	0x0a, 0xf8, 0x07, 0xbc, 0x78, 0xca, 0x1a, 0xce, 0xbe, 0x75, 0x67, 0x03, 0xf2, 0xf5, 0x7f, 0x90,
	0x7c, 0x30, 0xa7, 0x37, 0xe6, 0x94, 0x8c, 0xdf, 0xc0, 0x3a, 0x57, 0x5a, 0x5b, 0xf6, 0x41, 0xcb,
	0x26, 0x53, 0xa5, 0x3b, 0x1f, 0xae, 0xb2, 0x7c, 0x94, 0x13, 0x2d, 0x1b, 0x5a, 0x6e, 0xbe, 0xc3,
	0xfa, 0x09, 0x17, 0x5f, 0xc1, 0x45, 0x42, 0xd8, 0x96, 0xb2, 0x8f, 0xe8, 0xcc, 0x0e, 0x5f, 0x49,
	0x14, 0xc6, 0xf7, 0x04, 0x4d, 0xf0, 0x1a, 0xae, 0xb6, 0x84, 0x0b, 0xca, 0x02, 0x41, 0x63, 0x86,
	0x1c, 0xfc, 0x0c, 0x96, 0x94, 0x51, 0x41, 0x83, 0x28, 0x23, 0x5f, 0x08, 0x13, 0x68, 0x8a, 0x01,
	0xe6, 0x61, 0x14, 0x73, 0xb2, 0x45, 0xe7, 0x9b, 0x0c, 0x6e, 0x4e, 0x96, 0xc6, 0x08, 0xae, 0xb9,
	0x08, 0xc4, 0x03, 0xcf, 0x1e, 0x18, 0x27, 0x02, 0x9d, 0xe1, 0x6b, 0xb8, 0x0c, 0xe3, 0xfb, 0x24,
	0x22, 0xc2, 0x06, 0xad, 0x00, 0x92, 0x34, 0x0e, 0x09, 0xe7, 0xb6, 0x85, 0x83, 0x17, 0x30, 0x23,
	0x69, 0x1a, 0xa7, 0x68, 0x6a, 0x0b, 0xf1, 0xcf, 0x34, 0x49, 0x6c, 0xc0, 0x07, 0xef, 0xdb, 0xdb,
	0x9d, 0xea, 0xf6, 0x7d, 0xee, 0x15, 0xa6, 0xf6, 0x73, 0xd9, 0x15, 0xfb, 0xc2, 0x1c, 0x0f, 0x7e,
	0x5b, 0xec, 0xab, 0x5a, 0xb6, 0x7e, 0xde, 0x2b, 0x5d, 0xfa, 0x3b, 0xe3, 0x8f, 0x4f, 0x99, 0xcf,
	0x87, 0xef, 0xf1, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x2e, 0xdc, 0xfc, 0x2f, 0x02,
	0x00, 0x00,
}